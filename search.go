package googlesearch

type SearchResult struct {
  Description string
  Link string
}

type GoogleResults struct {
  Results []SearchResult
}

func (self *SearchResult) GetDescription() string {
  return self.Description
}

func (self *SearchResult) GetLink() string {
  return self.Link
}
