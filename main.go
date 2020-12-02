package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Book struct {

	Title  string `yaml:"title"`
	Author string `yaml:"author"`
	Pages  []page `yaml:"pages"`
}

type page struct {
	PageNumber int      `yaml:"pageNumber"`
	Text       string   `yaml:"text"`
	Choices    []string `yaml:"choices"`
}

const (
	firstPageNumber = 1
	separator       = "~~~~~~~~~~~~~~~~~~~~~~~~~~~"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must pass a single argument, the filename of the story (ex: story.yaml)")
		os.Exit(1)
	}

	book, err := getBook()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	readBook(book)
}

func readBook(book *Book) {
	printBookDetails(book)

	p, _ := getPage(book, firstPageNumber)

	for {
		printPage(p)

		if isTerminalPage(p) {
			break
		}

		p = chooseNextPage(book)
	}
}

func getPage(book *Book, pageNumber int) (page, error) {
	var matchingPage page
	for _, p := range book.Pages {
		if p.PageNumber == pageNumber {
			matchingPage = p

			return matchingPage, nil
		}
	}

	return matchingPage, fmt.Errorf("there is no page %v in this book", pageNumber)
}

func chooseNextPage(book *Book) page {
	for {
		fmt.Print("\n", "Choose a page: ")

		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")

		pageNumber, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Printf("\n\"%v\" is not a valid input.\n\n", userInput)
			fmt.Println(separator)

			continue
		}

		p, err := getPage(book, pageNumber)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			fmt.Println("\n" + separator)

			continue
		}

		return p
	}
}

func printBookDetails(book *Book) {
	fmt.Println()
	fmt.Println(book.Title)
	fmt.Println()
	fmt.Println("by")
	fmt.Println()
	fmt.Println(book.Author)
}

func printPage(p page) {
	fmt.Println()
	fmt.Println(separator)
	fmt.Println()
	fmt.Println("page", p.PageNumber)
	fmt.Println()
	fmt.Println(p.Text)

	if len(p.Choices) == 0 {
		return
	}

	fmt.Println(separator)
	fmt.Println()
	for _, choice := range p.Choices {
		fmt.Println(choice)
		fmt.Println()
	}
}

func getBook() (*Book, error) {
	book := &Book{}
	filename := os.Args[1]

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return book, err
	}

	err = yaml.Unmarshal([]byte(source), &book)
	if err != nil {
		return book, err
	}

	return book, nil
}

func isTerminalPage(p page) bool {
	return len(p.Choices) == 0
}
