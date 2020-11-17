package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/olivere/elastic"
	"os"

)

/**
  This function handles the loading of credential.env credential variables.
*/

func getEnvVars() {

	// get path of working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// load env variable to the program
	err = godotenv.Load(cwd + "/credential.env")
	if err != nil {
		panic(err)
	}
}

/**
  This function search for a specific query from a specific index
 */

func readFromElasticSearch(query elastic.Query, Index string) (*elastic.SearchResult, error) {
	getEnvVars()
	esUrl := os.Getenv("ELASTIC_SEARCH_URL")
	username := os.Getenv("ELASTIC_SEARCH_USERNAME")
	password := os.Getenv("ELASTIC_SEARCH_PASSWORD")
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(esUrl),
		elastic.SetBasicAuth(username, password))

	if err != nil {
		return nil, err
	}

	searchResult, err := client.Search().Index(Index).Query(query).Pretty(true).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}
