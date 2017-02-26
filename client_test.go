package googlesearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetProxy(t *testing.T) {
	client := NewClient()
	client.SetProxy("http://proxy:8000")
	assert.Equal(t, client.proxyURL.String(), "http://proxy:8000")
}

func TestSearch(t *testing.T) {
	client := NewClient()
	googleResponse, err := client.Search("golang")
	checkAnswer(t, googleResponse, err)
	googleResponse, err = client.SearchPage("golang", 1) // second page of answers.
	checkAnswer(t, googleResponse, err)
}

func checkAnswer(t *testing.T, googleResponse List, err error) {
	assert.Equal(t, err, nil)
	for _, result := range googleResponse {
		assert.NotNil(t, result.Name)
		assert.NotNil(t, result.Desc)
		assert.NotNil(t, result.Link)
	}
}
