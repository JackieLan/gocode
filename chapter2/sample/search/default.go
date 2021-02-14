package search

import (
	"log"
)

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	log.Printf("search %s in default engine\n", searchTerm)
	return nil, nil
}
