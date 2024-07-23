package searchstore

import (
	"log"

	"github.com/blevesearch/bleve/v2"
	bleveIdx "github.com/blevesearch/bleve_index_api"
)

type Store struct {
	index bleve.Index
}

func NewStore(path string) (*Store, error) {
	index, err := bleve.Open(path)
	if err != nil {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(path, mapping)
		if err != nil {
			log.Println(err)

			return nil, err
		}
	}

	return &Store{index: index}, nil
}

func (s *Store) CreateIndex() error {
	users := []*User{
		{
			EmpNum:       "S11601",
			FullName:     "黒澤 拓磨",
			FullNameKana: "ｸﾛｻﾜ　ﾀｸﾏ",
			Email:        "kurosawa_takuma@cyberagenet.co.jp",
		},
		{
			EmpNum:       "S11602",
			FullName:     "黒澤 あきら",
			FullNameKana: "ｸﾛｻﾜ　ｱｷﾗ",
			Email:        "kurosawa_akira@cyberagenet.co.jp",
		},
		{
			EmpNum:       "S11603",
			FullName:     "黒澤 拓実",
			FullNameKana: "ｸﾛｻﾜ　ﾀｸﾐ",
			Email:        "kurosawa_takumi@cyberagenet.co.jp",
		},
	}
	for _, u := range users {
		if err := s.index.Index(u.EmpNum, u); err != nil {
			return err
		}
	}

	return nil
}

func (s *Store) Search(keyword string) ([]*User, error) {
	query := bleve.NewMatchQuery(keyword)
	search := bleve.NewSearchRequest(query)
	searchResults, err := s.index.Search(search)
	if err != nil {
		return nil, err
	}

	result := make([]*User, 0, len(searchResults.Hits))
	for _, hit := range searchResults.Hits {
		user := &User{
			EmpNum: hit.ID,
			Score:  hit.Score,
		}
		doc, err := s.index.Document(hit.ID)
		if err != nil {
			return nil, err
		}
		doc.VisitFields(func(field bleveIdx.Field) {
			if field.Name() == "FullName" {
				user.FullName = string(field.Value())
			}
			if field.Name() == "FullNameKana" {
				user.FullNameKana = string(field.Value())
			}
			if field.Name() == "Email" {
				user.Email = string(field.Value())
			}
		})
		result = append(result, user)
	}

	return result, nil
}
