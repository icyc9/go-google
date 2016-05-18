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
  google_response, err := client.GetSearchResults("golang")
  assert.Equal(t, err, nil)

  for result := range google_response.Results {
    current_result := google_response.Results[result]
    assert.NotNil(t, current_result.Description)
    assert.NotNil(t, current_result.Link)
  }
}