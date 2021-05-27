package gotradier

// EndpointType is an enum of the supported Tradier API endpoints
//
// (The streaming endpoint is not supported)
//
// https://documentation.tradier.com/brokerage-api/overview/endpoints
type EndpointType string

// EndpointType Enum
const (
	EndpointTypeAPI     EndpointType = "https://api.tradier.com/v1"
	EndpointTypeSandbox EndpointType = "https://sandbox.tradier.com/v1"
)

type symbolQuotes struct {
	Quotes           []Quote  `xml:"quote"`
	UnmatchedSymbols []string `xml:"unmatched_symbols"`
}

// Quotes using the symbol string as keys
type Quotes map[string]Quote

// Quote data for a particular symbol
type Quote struct {
	Symbol           string    `xml:"symbol"`
	Description      string    `xml:"description"`
	Exchange         string    `xml:"exch"`
	Type             QuoteType `xml:"type"`
	Last             float64   `xml:"last"`
	Change           float64   `xml:"change"`
	Volume           float64   `xml:"volume"`
	Open             float64   `xml:"open"`
	High             float64   `xml:"high"`
	Low              float64   `xml:"low"`
	Close            float64   `xml:"close"`
	Bid              float64   `xml:"bid"`
	Ask              float64   `xml:"ask"`
	ChangePercentage float64   `xml:"change_percentage"`
	AverageVolume    int       `xml:"average_volume"`
	LastVolume       int       `xml:"last_volume"`
	TradeDate        int64     `xml:"trade_date"`
	PreviousClose    float64   `xml:"prevclose"`
	Week52High       float64   `xml:"week_52_high"`
	Week52Low        float64   `xml:"week_52_low"`
	BidSize          int       `xml:"bidsize"`
	BidExchange      string    `xml:"bidexch"`
	BidDate          int64     `xml:"bid_date"`
	AskSize          int       `xml:"asksize"`
	AskExchange      string    `xml:"askexch"`
	AskDate          int64     `xml:"ask_date"`
	RootSymbols      string    `xml:"root_symbols"`
}

// QuoteType is an enum of the supported Tradier API quote types
type QuoteType string

const (
	// QuoteTypeOption indicates the Quote is for an option
	QuoteTypeOption QuoteType = "option"
	// QuoteTypeStock indicates the Quote is for an stock
	QuoteTypeStock QuoteType = "stock"
	// QuoteTypeETF indicates the Quote is for an ETF
	QuoteTypeETF QuoteType = "etf"
)

type optionExpirations struct {
	Expirations Expirations `xml:"date"`
}

// Expirations is a slice of expiration date strings in format YYYY-MM-DD
//
// Golang time format is 2006-01-02
type Expirations []string

type optionChain struct {
	Options Options `xml:"option"`
}

// Options is a slice of options data
type Options []Option

// Option contract data for a particular symbol
type Option struct {
	Symbol           string     `xml:"symbol"`
	Description      string     `xml:"description"`
	Exchange         string     `xml:"exch"`
	Type             QuoteType  `xml:"type"`
	Last             float64    `xml:"last"`
	Change           float64    `xml:"change"`
	Volume           float64    `xml:"volume"`
	Open             float64    `xml:"open"`
	High             float64    `xml:"high"`
	Low              float64    `xml:"low"`
	Close            float64    `xml:"close"`
	Bid              float64    `xml:"bid"`
	Ask              float64    `xml:"ask"`
	UnderlyingSymbol string     `xml:"underlying"`
	Strike           float64    `xml:"strike"`
	Greeks           Greeks     `xml:"greeks"`
	ChangePercentage float64    `xml:"change_percentage"`
	AverageVolume    int        `xml:"average_volume"`
	LastVolume       int        `xml:"last_volume"`
	TradeDate        int64      `xml:"trade_date"`
	PreviousClose    float64    `xml:"prevclose"`
	Week52High       float64    `xml:"week_52_high"`
	Week52Low        float64    `xml:"week_52_low"`
	BidSize          int        `xml:"bidsize"`
	BidExchange      string     `xml:"bidexch"`
	BidDate          int64      `xml:"bid_date"`
	AskSize          int        `xml:"asksize"`
	AskExchange      string     `xml:"askexch"`
	AskDate          int64      `xml:"ask_date"`
	OpenInterest     float64    `xml:"open_interest"`
	ContractSize     float64    `xml:"contract_size"`
	ExpirationDate   string     `xml:"expiration_date"`
	ExpirationType   string     `xml:"expiration_type"` // TODO convert to enum
	OptionType       OptionType `xml:"option_type"`
	RootSymbol       string     `xml:"root_symbol"`
}

// Greeks of an option contract
type Greeks struct {
	Delta     float64 `xml:"delta"`
	Gamma     float64 `xml:"gamma"`
	Theta     float64 `xml:"theta"`
	Vega      float64 `xml:"vega"`
	Rho       float64 `xml:"rho"`
	Phi       float64 `xml:"phi"`
	BidIV     float64 `xml:"bid_iv"`
	MidIV     float64 `xml:"mid_iv"`
	AskIV     float64 `xml:"ask_iv"`
	SMVVol    float64 `xml:"smv_vol"`
	UpdatedAt string  `xml:"updated_at"`
}

//OptionType is an enum of types of option contracts
type OptionType string

const (
	// OptionTypePut indicates the Option is a PUT
	OptionTypePut OptionType = "put"
	// OptionTypeCall indicates the Option is a CALL
	OptionTypeCall OptionType = "call"
)

type fault struct {
	Fault  string      `xml:"faultstring"`
	Detail faultDetail `xml:"detail"`
}

type faultDetail struct {
	ErrorCode string `xml:"errorcode"`
}
