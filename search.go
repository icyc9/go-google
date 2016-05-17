package googlesearch

type SearchResult struct {
  Description string
  Link string
}

type SearchResultList struct {
  List []SearchResult
}

func (self *SearchResult) GetDescription() string {
  return self.Description
}

func (self *SearchResult) GetLink() string {
  return self.Link
}
