package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Book struct {
	Title  string `yaml:"title"`
	Author string `yaml:"author"`
	Pages  []Page `yaml:"pages"`
}

type Page struct {
	PageNumber int      `yaml:"pageNumber"`
	Text       string   `yaml:"text"`
	Choices    []Choice `yaml:"choices"`
}

type Choice struct {
	PageNumber int    `yaml:"pageNumber"`
	Text       string `yaml:"text"`
}

func main() {
	book, err := getBook()
	if err != nil {
		fmt.Println("There was an error ", err)
	}

	readPage(book, 1)
}

func readPage(book *Book, pageNumber int) (nextPage int) {
	fmt.Println("Page ", pageNumber)
	for _, page := range book.Pages {
		if page.PageNumber == pageNumber {
			println(page.Text)
			for _, choice := range page.Choices {
				println(choice.Text)
			}
		}
	}

	return 0
}

func getBook() (*Book, error) {
	book := &Book{}
	source, err := ioutil.ReadFile("sample.yaml")
	if err != nil {
		return book, err
	}

	err = yaml.Unmarshal([]byte(source), &book)
	if err != nil {
		return book, err
	}

	return book, nil
}
