package main

import (
	"fmt"
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
	defer store.Close()

	if err := store.CreateIndex(); err != nil {
		log.Println(err)

		return
	}

	func() {
		keyword := "s11601"
		result, err := store.Search(searchstore.QueryTypeTerm, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "kurosawa"
		result, err := store.Search(searchstore.QueryTypeWildcard, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "kurosawa_takuma"
		result, err := store.Search(searchstore.QueryTypeWildcard, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "ｸﾛｻﾜ"
		result, err := store.Search(searchstore.QueryTypeString, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- string query keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- string query keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "ｸﾛｻﾜ"
		result, err := store.Search(searchstore.QueryTypeMatchPhraseQuery, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- match phrase query keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- match phrase query keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "ｸﾛｻﾜ ﾀｸﾏ"
		result, err := store.Search(searchstore.QueryTypeString, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- string query keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- string query keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "ｸﾛｻﾜ ﾀｸﾏ"
		result, err := store.Search(searchstore.QueryTypeMatchPhraseQuery, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- match phrase query keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- match phrase query keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "拓磨"
		result, err := store.Search(searchstore.QueryTypeString, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- string query keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- string query keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "拓磨"
		result, err := store.Search(searchstore.QueryTypeMatchPhraseQuery, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- match phrase query keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- match phrase query keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "黒澤　拓磨"
		result, err := store.Search(searchstore.QueryTypeString, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- string query keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- string query keyword: %s ----------\n\n", keyword)
	}()

	func() {
		keyword := "黒澤　拓磨"
		result, err := store.Search(searchstore.QueryTypeMatchPhraseQuery, keyword)
		if err != nil {
			log.Println(err)

			return
		}

		fmt.Printf("-------- match phrase query keyword: %s ----------\n", keyword)
		for i, user := range result {
			log.Printf("user%d: %#v", i+1, user)
		}
		fmt.Printf("-------- match phrase query keyword: %s ----------\n\n", keyword)
	}()
}
