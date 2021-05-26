package gotradier

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const contentType string = "application/xml"

type Client struct {
	Endpoint EndpointType
	// Tradier API Access Token
	// https://developer.tradier.com/getting_started
	Token string
}

// GetQuotes for provided symbols
func (client Client) GetQuotes(symbols []string, greeks bool) (Quotes, error) {
	path := "/markets/quotes"
	params := url.Values{
		"symbols": []string{strings.Join(symbols, ",")},
		"greeks":  []string{fmt.Sprint(greeks)},
	}
	url := fmt.Sprintf("%s%s?%s", client.Endpoint, path, params.Encode())
	headers := map[string]string{"accept": contentType, "Authorization": fmt.Sprintf("Bearer %s", client.Token)}
	response, data, err := request(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	} else if response.StatusCode != http.StatusOK {
		return nil, getFault(data, response.StatusCode)
	}
	quotes := &symbolQuotes{}
	if err := xml.Unmarshal(data, quotes); err != nil {
		if string(data) == `</quotes>` {
			return Quotes{}, nil
		}
		return nil, err
	}
	quotesMap := Quotes{}
	for _, quote := range quotes.Quotes {
		quotesMap[quote.Symbol] = quote
	}
	return quotesMap, nil
}

// GetOptionExpirations for provided underlying symbol
func (client Client) GetOptionExpirations(symbol string) (Expirations, error) {
	path := "/markets/options/expirations"
	params := url.Values{
		"symbol":          []string{symbol},
		"includeAllRoots": []string{"true"},
		"strikes":         []string{"false"},
	}
	url := fmt.Sprintf("%s%s?%s", client.Endpoint, path, params.Encode())
	headers := map[string]string{"accept": contentType, "Authorization": fmt.Sprintf("Bearer %s", client.Token)}
	response, data, err := request(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	} else if response.StatusCode != http.StatusOK {
		return nil, getFault(data, response.StatusCode)
	}
	expirations := &optionExpirations{}
	if err := xml.Unmarshal(data, expirations); err != nil {
		if string(data) == `</expirations>` {
			return Expirations{}, nil
		}
		return nil, err
	}
	return expirations.Expirations, nil
}

//GetOptionChain at provided expiration for provided symbol
func (client Client) GetOptionChain(symbol string, expiration string, greeks bool) (Options, error) {
	path := "/markets/options/chains"
	params := url.Values{
		"symbol":     []string{symbol},
		"expiration": []string{expiration},
		"greeks":     []string{fmt.Sprint(greeks)},
	}
	url := fmt.Sprintf("%s%s?%s", client.Endpoint, path, params.Encode())
	headers := map[string]string{"accept": contentType, "Authorization": fmt.Sprintf("Bearer %s", client.Token)}
	response, data, err := request(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	} else if response.StatusCode != http.StatusOK {
		return nil, getFault(data, response.StatusCode)
	}
	optionChain := &optionChain{}
	if err := xml.Unmarshal(data, optionChain); err != nil {
		if string(data) == `</options>` {
			return Options{}, nil
		}
		return nil, err
	}
	return optionChain.Options, nil
}

func getFault(data []byte, status int) error {
	fault := &fault{}
	if err := xml.Unmarshal(data, fault); err != nil {
		return fmt.Errorf("%d - %s", status, string(data))
	}
	return fmt.Errorf("%d (%s) - %s", status, fault.Detail.ErrorCode, fault.Fault)
}

func request(method string, url string, headers map[string]string, body io.Reader) (*http.Response, []byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	for key, val := range headers {
		request.Header.Set(key, val)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	return response, data, nil
}
