package googlesearch

import (
	"github.com/PuerkitoBio/goquery"

	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// A Client provides a simple google search client.
//
type Client struct {
	*Config
	Header     http.Header
	HTTPClient *http.Client
	transport  *http.Transport
	proxyURL   *url.URL
}

// NewClient creates a new googlesearch Client with package Defaults.
//
func NewClient() *Client {
	client := &Client{
		Config:     &Defaults,
		Header:     http.Header{},
		HTTPClient: &http.Client{},
		transport:  &http.Transport{},
	}
	return client
}

// SetProxy sets a custom proxy for the client.
//
func (c *Client) SetProxy(proxy string) (*Client, error) {
	pURL, err := url.Parse(proxy)
	if err != nil {
		return nil, err
	}
	c.proxyURL = pURL
	return c, nil
}

// SetHeader sets a single header values for the client.
//
func (c *Client) SetHeader(header, value string) *Client {
	c.Header.Set(header, value)
	return c
}

//
//-------------------------------------------------------------------[ QUERY ]--

// FirstLink gets the first answer link for the query.
//
func (c *Client) FirstLink(query string) (link string, e error) {
	list, e := c.Search(query)
	if len(list) == 0 {
		return "", e
	}
	return list[0].Link, e
}

// Search asks a google server what he knows about your query string.
//
//
func (c *Client) Search(query string) (List, error) {
	return c.SearchPage(query, 0)
}

// SearchPage asks a google server what he knows about your query string.
// The page number is a step of n x NbResults answers.
//
//
func (c *Client) SearchPage(query string, start int) (List, error) {
	uri := c.FormatURL(query, c.Config.NbResults, c.Config.NbResults*start)
	resp, e := c.Download(uri)
	if e != nil {
		return nil, e
	}
	return c.Parse(resp)
}

//
//----------------------------------------------------------------[ INTERNAL ]--

// FormatURL formats the search URL.
//
func (c *Client) FormatURL(query string, count, start int) string {
	values := make(url.Values)
	values.Set("q", url.QueryEscape(query))
	values.Set("safe", c.Safe)
	values.Set("hl", c.Lang)
	values.Set("num", strconv.Itoa(count))
	values.Set("start", strconv.Itoa(start))
	values.Set("ie", c.CodesetInput)
	values.Set("oe", c.CodesetOutput)
	values.Set("filter", "0")

	return fmt.Sprintf(c.SearchURL, c.TLD) + "?" + values.Encode()
}

// Download gets the search html page from the server.
//
func (c *Client) Download(uri string) (*http.Response, error) {
	req, e := http.NewRequest("GET", uri, nil)
	if e != nil {
		return nil, e
	}
	req.Header = c.Header

	if c.proxyURL != nil {
		c.transport.Proxy = http.ProxyURL(c.proxyURL)
	}

	c.HTTPClient.Transport = c.transport
	return c.HTTPClient.Do(req)
}

// Parse finds links in the search html page.
//
func (c *Client) Parse(resp *http.Response) (List, error) {
	document, e := goquery.NewDocumentFromResponse(resp)
	if e != nil {
		return nil, e
	}

	var results []SearchResult
	document.Find("div.g").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Find(".r").Find("a").Attr("href")
		if !exists {
			return
		}

		link = strings.TrimPrefix(link, "/url?q=")
		idx := strings.Index(link, "&sa=")
		if idx > 0 {
			link = link[:idx]
		} else {
			println("parse search results: suffix not found", link)
		}

		results = append(results, SearchResult{
			Name: s.Find(".r").Find("a").Text(),
			Desc: s.Find(".s").Find(".st").Text(),
			Link: link,
		})
	})

	return results, e
}
