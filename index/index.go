package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/olivere/elastic"
	"os"
)

const (
	PostIndex = "post"
	UserIndex = "user"
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
  This functions handles the initialization of ElasticSearch service.
 */
func main() {

	// get credentials from local environment
	getEnvVars()
	esUrl := os.Getenv("ELASTIC_SEARCH_URL")
	password := os.Getenv("ELASTIC_SEARCH_PASSWORD")

	fmt.Println(esUrl)
	fmt.Println(password)

	// Connect to the elasticsearch server
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(esUrl),
		elastic.SetBasicAuth("elastic", password))
	if err != nil {
		panic(err)
	}

	// check if the post collections exists. If not, initialize a new one.
	exists, err := client.IndexExists(PostIndex).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		mapping := `{
                        "mappings": {
                                "properties": {
                                        "user":     { "type": "keyword" },
                                        "message":  { "type": "text" },
                                        "url":      { "type": "keyword", "index": false },
                                        "type":     { "type": "keyword", "index": false }
                                }
                        }
                }`
		_, err := client.CreateIndex(PostIndex).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

	// check if user collection exists. If not, create a new one.
	exists, err = client.IndexExists(UserIndex).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		mapping := `{
                        "mappings": {
                                "properties": {
                                        "username": {"type": "keyword"},
                                        "password": {"type": "keyword", "index": false},
                                        "age":      {"type": "long", "index": false},
                                        "gender":   {"type": "keyword", "index": false}
                                }
                        }
                }`
		_, err = client.CreateIndex(UserIndex).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Index are created")
}
