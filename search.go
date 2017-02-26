// Package googlesearch launch search queries on google.
package googlesearch

// Defaults provides default search settings.
//
var Defaults = Config{
	TLD:           "com",   // google.COM
	Lang:          "en",    // in ENglish
	Safe:          "off",   // and NSFW.
	CodesetInput:  "utf-8", // input charset format.
	CodesetOutput: "utf-8", // output charset format.
	NbResults:     10,      // Number of answers per query.

	SearchURL: "https://www.google.%s/search", // Base url.
}

// Config defines a search settings.
//
type Config struct {
	TLD           string
	Lang          string
	Safe          string
	SearchURL     string
	CodesetInput  string
	CodesetOutput string
	NbResults     int
}

// List represents a list of SearchResult.
//
type List []SearchResult

// SearchResult defines a single search result website.
//
type SearchResult struct {
	Name string
	Desc string
	Link string
}

// GetName gets the result description.
//
func (res *SearchResult) GetName() string {
	return res.Name
}

// GetDescription gets the result description.
//
func (res *SearchResult) GetDescription() string {
	return res.Desc
}

// GetLink gets the result link location.
//
func (res *SearchResult) GetLink() string {
	return res.Link
}

// Resulter represents the result get data interface.
//
type Resulter interface {
	GetName() string
	GetDescription() string
	GetLink() string
}
