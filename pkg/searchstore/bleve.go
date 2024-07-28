package searchstore

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/web"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	bleveIdx "github.com/blevesearch/bleve_index_api"
	"github.com/ikawaha/bleveplugin/analysis/lang/ja"
)

type Store struct {
	index bleve.Index
}

func (s *Store) Close() {
	s.index.Close()
}

func NewStore(path string) (*Store, error) {
	idxMapping := bleve.NewIndexMapping()
	dm := bleve.NewDocumentMapping()
	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = keyword.Name
	webFieldMapping := bleve.NewTextFieldMapping()
	webFieldMapping.Analyzer = web.Name
	jaTextFieldMapping := bleve.NewTextFieldMapping()
	jaTextFieldMapping.Analyzer = "ja"
	dm.AddFieldMappingsAt("EmpNum", keywordFieldMapping)
	dm.AddFieldMappingsAt("FullName", jaTextFieldMapping)
	dm.AddFieldMappingsAt("FullNameKana", jaTextFieldMapping)
	dm.AddFieldMappingsAt("Email", webFieldMapping)
	idxMapping.AddDocumentMapping("user", dm)

	if err := idxMapping.AddCustomTokenizer("ja_tokenizer", map[string]any{
		"type":      ja.Name,
		"dict":      ja.DictIPA,
		"base_form": true,
		"stop_tags": true,
	}); err != nil {
		return nil, fmt.Errorf("failed to create ja tokenizer: %w", err)
	}
	if err := idxMapping.AddCustomAnalyzer("ja", map[string]any{
		"type":      custom.Name,
		"tokenizer": "ja_tokenizer",
		"token_filters": []string{
			ja.StopWordsName,
			lowercase.Name,
		},
	}); err != nil {
		return nil, fmt.Errorf("failed to create ja analyzer: %w", err)
	}

	index, err := bleve.NewMemOnly(idxMapping)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &Store{index: index}, nil
}

func (s *Store) CreateIndex() error {
	users := []*User{
		{
			EmpNum:       "S11601",
			FullName:     "黒澤　拓磨",
			FullNameKana: "ｸﾛｻﾜ　ﾀｸﾏ",
			Email:        "kurosawa_takuma@cyberagenet.co.jp",
		},
		{
			EmpNum:       "S11602",
			FullName:     "黒澤　あきら",
			FullNameKana: "ｸﾛｻﾜ　ｱｷﾗ",
			Email:        "kurosawa_akira@cyberagenet.co.jp",
		},
		{
			EmpNum:       "S11603",
			FullName:     "黒澤　拓実",
			FullNameKana: "ｸﾛｻﾜ　ﾀｸﾐ",
			Email:        "kurosawa_takumi@cyberagenet.co.jp",
		},
	}
	for _, u := range users {
		if err := s.index.Index(u.EmpNum, u); err != nil {
			return err
		}
	}

	docCount, err := s.index.DocCount()
	if err != nil {
		log.Println(err)

		return err
	}
	fmt.Printf("doc count: %d\n", docCount)

	return nil
}

type QueryType int

const (
	QueryTypeTerm QueryType = iota
	QueryTypeString
	QueryTypeWildcard
	QueryTypeMatchPhraseQuery
)

func (s *Store) Search(queryType QueryType, keyword string) ([]*User, error) {
	var search *bleve.SearchRequest
	switch queryType {
	case QueryTypeTerm:
		query := bleve.NewTermQuery(keyword)
		search = bleve.NewSearchRequest(query)
	case QueryTypeString:
		query := bleve.NewQueryStringQuery(keyword)
		search = bleve.NewSearchRequest(query)
	case QueryTypeWildcard:
		query := bleve.NewWildcardQuery("*" + keyword + "*")
		search = bleve.NewSearchRequest(query)
	case QueryTypeMatchPhraseQuery:
		query := bleve.NewMatchPhraseQuery(keyword)
		search = bleve.NewSearchRequest(query)
	default:
		return nil, fmt.Errorf("unsupported query type: %d", queryType)
	}

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
