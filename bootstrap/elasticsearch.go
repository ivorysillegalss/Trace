package bootstrap

import (
	"go-quickstart/infrastructure/elasticsearch"
	"log"
)

func NewEsEngine(env *Env) elasticsearch.Client {
	client, err := elasticsearch.NewElasticSearchClient(env.ElasticSearchUrl)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewSearchEngine(es elasticsearch.Client) *SearchEngine {
	return &SearchEngine{EsEngine: es}
}
