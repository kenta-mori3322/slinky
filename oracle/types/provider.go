package types

import (
	"fmt"
)

// DefaultTickerDecimals is the number of decimal places every single price
// is scaled to before being sent for aggregation.
const DefaultTickerDecimals = 18

type (
	// ProviderTicker is the interface for the ticker that provider's utilize/return.
	ProviderTicker interface {
		fmt.Stringer

		// Provider returns the provider for the ticker.
		Provider() string
		// OffChainTicker returns the off-chain representation for the ticker.
		OffChainTicker() string
		// Decimals returns the number of decimals for the ticker.
		Decimals() uint64
		// JSON returns additional JSON data for the ticker.
		JSON() string
	}

	// DefaultProviderTicker is a basic implementation of the ProviderTicker interface.
	// Provider's that utilize this implementation should be able to easily configure
	// custom json data for their tickers.
	DefaultProviderTicker struct {
		provider string
		offChain string
		decimals uint64
		json     string
	}

	// ProviderTickers is helper struct to manage a list of provider tickers.
	ProviderTickers struct {
		cache map[string]ProviderTicker
	}
)

// NewProviderTicker returns a new provider ticker.
func NewProviderTicker(
	provider, offChain, json string,
	decimals uint64,
) ProviderTicker {
	return DefaultProviderTicker{
		provider: provider,
		offChain: offChain,
		json:     json,
		decimals: decimals,
	}
}

// Provider returns the provider for the ticker.
func (t DefaultProviderTicker) Provider() string {
	return t.provider
}

// OffChainTicker returns the off-chain representation for the ticker.
func (t DefaultProviderTicker) OffChainTicker() string {
	return t.offChain
}

// Decimals returns the number of decimals for the ticker.
func (t DefaultProviderTicker) Decimals() uint64 {
	return t.decimals
}

// JSON returns additional JSON data for the ticker.
func (t DefaultProviderTicker) JSON() string {
	return t.json
}

// String returns the string representation of the provider ticker.
func (t DefaultProviderTicker) String() string {
	return fmt.Sprintf(
		"provider: %s, off-chain-ticker: %s, decimals: %d, json: %s",
		t.provider,
		t.offChain,
		t.decimals,
		t.json,
	)
}

// NewProviderTickers returns a new list of provider tickers.
func NewProviderTickers(tickers ...ProviderTicker) ProviderTickers {
	cache := make(map[string]ProviderTicker)
	for _, ticker := range tickers {
		cache[ticker.OffChainTicker()] = ticker
	}
	return ProviderTickers{
		cache: cache,
	}
}

// FromOffChain returns the provider ticker from the off-chain ticker.
func (t ProviderTickers) FromOffChain(offChain string) (ProviderTicker, bool) {
	ticker, ok := t.cache[offChain]
	return ticker, ok
}

// Add adds a provider ticker to the list of provider tickers.
func (t *ProviderTickers) Add(ticker ProviderTicker) {
	t.cache[ticker.OffChainTicker()] = ticker
}

// Reset resets the provider tickers.
func (t *ProviderTickers) Reset() {
	t.cache = make(map[string]ProviderTicker)
}
