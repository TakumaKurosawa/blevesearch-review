package main

import (
	"log"

	"github.com/TakumaKurosawa/blevesearch-review/pkg/searchstore"
)

func main() {
	const path = "bleve.user"
	store, err := searchstore.NewStore(path)
	if err != nil {
		log.Println(err)

		return
	}

	if err := store.CreateIndex(); err != nil {
		log.Println(err)

		return
	}

	result, err := store.Search("拓磨")
	if err != nil {
		log.Println(err)

		return
	}

	for i, user := range result {
		log.Printf("user%d: %#v", i+1, user)
	}
}
