package main

/*
	Реализовать паттерн «адаптер» на любом примере.
*/

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func main() {
	hub := NewArticleHub(NewJsonToArticleAdapter())
	hub.PrintRandomArticle()
}

type Article struct {
	By     string `json:"by"`
	ID     int    `json:"id"`
	Parent int    `json:"parent"`
	Text   string `json:"text"`
	Time   int    `json:"time"`
	Type   string `json:"type"`
}

// Represents hub of articles, allows to read a random article
type ArticleHub struct {
	provider ArticleProvider
}

func NewArticleHub(provider ArticleProvider) *ArticleHub {
	return &ArticleHub{provider: provider}
}

// Prints random article, article is given by `hub.provider`
func (hub *ArticleHub) PrintRandomArticle() {
	article, err := hub.provider.GetRandomArticle()
	if err != nil {
		fmt.Printf("Could not get article, see: %v", err)
	}
	fmt.Printf("Author: %s\nID: %d\nParent: %d\nText: %s\n", article.By, article.ID, article.Parent, article.Text)
}

type ArticleProvider interface {
	GetRandomArticle() (*Article, error)
}

// Use this type to adapt ArticleService to ArticleProvider interface
type JsonToArticleAdapter struct {
	articleService *ArticleService
}

func NewJsonToArticleAdapter() *JsonToArticleAdapter {
	return &JsonToArticleAdapter{NewArticleService()}
}

func (jAdapter *JsonToArticleAdapter) GetRandomArticle() (*Article, error) {
	jsonString, err := jAdapter.articleService.GetArticleJSON(13000 + rand.Intn(1000))
	if err != nil {
		return nil, err
	}
	var article Article
	err = json.Unmarshal(jsonString, &article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

// This object is used to provide access to articles (which are available within web-api)
type ArticleService struct {
	s *http.Client
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		s: http.DefaultClient,
	}
}

// Gets article with specific id from hacker-news as json
func (ag *ArticleService) GetArticleJSON(id int) ([]byte, error) {
	resp, err := ag.s.Get(fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id))
	if err != nil {
		return nil, err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
