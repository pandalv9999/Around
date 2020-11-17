package main

import (
	"github.com/olivere/elastic"
	"reflect"
)

const (
	PostIndex = "post"
)

type Post struct {
	User string `json:"user"`
	Message string `json:"message"`
	Url string `json:"url"`
	Type string `json:"type"`
}

func searchPostByUser(user string)([]Post, error) {
	query := elastic.NewTermQuery("user", user)
	searchResult, err := readFromElasticSearch(query, PostIndex)
	if err != nil {
		return nil, err
	}
	return getPostFromSearchResult(searchResult), nil
}

func searchPostsByKeywords(keywords string)([]Post, error) {
	query := elastic.NewMatchQuery("message", keywords)
	query.Operator("AND")
	if keywords == "" {
		query.ZeroTermsQuery("all")
	}
	searcResult, err := readFromElasticSearch(query, PostIndex)
	if err != nil {
		return nil, err
	}
	return getPostFromSearchResult(searcResult), nil
}

func getPostFromSearchResult(searchResult *elastic.SearchResult) []Post {
	var ptype Post
	var posts []Post

	for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
		if p, ok := item.(Post); ok {
			posts = append(posts, p)
		}
	}

	return posts
}
