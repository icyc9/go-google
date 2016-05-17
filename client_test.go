package googlesearch

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSetProxy(t *testing.T) {
  client := NewClient()
  client.SetProxy("http://proxy:8000")
  assert.Equal(t, client.proxyURL.String(), "http://proxy:8000")
}

func TestGetSearchResults(t *testing.T) {
  client := NewClient()
  results, err := client.GetSearchResults("golang")
  assert.Equal(t, err, nil)

  for result := range results.List {
    current_result := results.List[result]
    assert.NotNil(t, current_result.Description)
    assert.NotNil(t, current_result.Link)
  }
}