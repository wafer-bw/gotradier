package gotradier

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetQuotes(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockResponse := `<quotes>
			<quote><symbol>AAPL</symbol></quote>
			<unmatched_symbols><symbol>AMDDDDD</symbol></unmatched_symbols>
		</quotes>`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(mockResponse))
			require.NoError(t, err)
		}))
		defer mockServer.Close()
		client := &Client{Endpoint: EndpointType(mockServer.URL), Token: "abc"}
		response, err := client.GetQuotes([]string{"AAPL"}, false)
		require.NoError(t, err)
		require.Equal(t, 1, len(response), response)
	})
	t.Run("success/no quotes", func(t *testing.T) {
		mockResponse := `<quotes>
			<unmatched_symbols><symbol>AMDDDDD</symbol></unmatched_symbols>
		</quotes>`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(mockResponse))
			require.NoError(t, err)
		}))
		defer mockServer.Close()
		client := &Client{Endpoint: EndpointType(mockServer.URL), Token: "abc"}
		response, err := client.GetQuotes([]string{"AAPL"}, false)
		require.NoError(t, err)
		require.Equal(t, 0, len(response), response)
	})
	t.Run("success/no unmatched symbols", func(t *testing.T) {
		mockResponse := `<quotes>
			<quote><symbol>AAPL</symbol></quote>
		</quotes>`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(mockResponse))
			require.NoError(t, err)
		}))
		defer mockServer.Close()
		client := &Client{Endpoint: EndpointType(mockServer.URL), Token: "abc"}
		response, err := client.GetQuotes([]string{"AAPL"}, false)
		require.NoError(t, err)
		require.Equal(t, 1, len(response), response)
	})
	t.Run("success/no unmatched symbols or quotes", func(t *testing.T) {
		mockResponse := `</quotes>`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(mockResponse))
			require.NoError(t, err)
		}))
		defer mockServer.Close()
		client := &Client{Endpoint: EndpointType(mockServer.URL), Token: "abc"}
		response, err := client.GetQuotes([]string{"AAPL"}, false)
		require.NoError(t, err)
		require.Equal(t, 0, len(response), response)
	})
}

func TestGetOptionExpirations(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockResponse := `<expirations><date>2021-05-28</date></expirations>`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(mockResponse))
			require.NoError(t, err)
		}))
		defer mockServer.Close()
		client := &Client{Endpoint: EndpointType(mockServer.URL), Token: "abc"}
		response, err := client.GetOptionExpirations("AAPL")
		require.NoError(t, err)
		require.Equal(t, 1, len(response), response)
	})
	t.Run("success/no expiration dates", func(t *testing.T) {
		mockResponse := `</expirations>`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(mockResponse))
			require.NoError(t, err)
		}))
		defer mockServer.Close()
		client := &Client{Endpoint: EndpointType(mockServer.URL), Token: "abc"}
		response, err := client.GetOptionExpirations("AAPL")
		require.NoError(t, err)
		require.Equal(t, 0, len(response), response)
	})
}

func TestGetOptionChain(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockResponse := `<options><option><symbol>AAPL210618P00003000</symbol><greeks><delta>0.0</delta></greeks></option></options>`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(mockResponse))
			require.NoError(t, err)
		}))
		defer mockServer.Close()
		client := &Client{Endpoint: EndpointType(mockServer.URL), Token: "abc"}
		response, err := client.GetOptionChain("AAPL", "2023-01-20", true)
		require.NoError(t, err)
		require.Equal(t, 1, len(response), response)
	})
	t.Run("success/no option chains", func(t *testing.T) {
		mockResponse := `</options>`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(mockResponse))
			require.NoError(t, err)
		}))
		defer mockServer.Close()
		client := &Client{Endpoint: EndpointType(mockServer.URL), Token: "abc"}
		response, err := client.GetOptionChain("AAPL", "2023-01-20", true)
		require.NoError(t, err)
		require.Equal(t, 0, len(response), response)
	})
}

func TestGetFault(t *testing.T) {
	t.Run("valid fault", func(t *testing.T) {
		mockResponse := `<?xml version='1.0' encoding='UTF-8'?><fault><faultstring>Invalid API call as no apiproduct match found</faultstring><detail><errorcode>keymanagement.service.InvalidAPICallAsNoApiProductMatchFound</errorcode></detail></fault>`
		err := getFault([]byte(mockResponse), http.StatusUnauthorized)
		require.Error(t, err)
		require.Equal(t, "401 (keymanagement.service.InvalidAPICallAsNoApiProductMatchFound) - Invalid API call as no apiproduct match found", err.Error())
	})
	t.Run("no fault object", func(t *testing.T) {
		mockResponse := `Invalid Access Token`
		err := getFault([]byte(mockResponse), http.StatusUnauthorized)
		require.Error(t, err)
		require.Equal(t, "401 - Invalid Access Token", err.Error())
	})
}
