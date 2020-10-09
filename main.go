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
	Pages  []Page `yaml:"pages"`
}

type Page struct {
	PageNumber int      `yaml:"pageNumber"`
	Text       string   `yaml:"text"`
	Choices    []string `yaml:"choices"`
}

const firstPageNumber = 1
const separator = "~~~~~~~~~~~~~~~~~~~~~~~~~~~"

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

	page, _ := getPage(book, firstPageNumber)

	for {
		printPage(page)

		if isTerminalPage(page) {
			break
		}

		page = chooseNextPage(book)
	}
}

func getPage(book *Book, pageNumber int) (Page, error) {
	var matchingPage Page
	for _, page := range book.Pages {
		if page.PageNumber == pageNumber {
			matchingPage = page
			return matchingPage, nil
		}
	}

	return matchingPage, fmt.Errorf("There is no page %v in this book.", pageNumber)
}

func chooseNextPage(book *Book) Page {
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

		page, err := getPage(book, pageNumber)

		if err != nil {
			fmt.Println()
			fmt.Println(err)
			fmt.Println("\n" + separator)
			continue
		}

		return page
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

func printPage(page Page) {
	fmt.Println()
	fmt.Println(separator)
	fmt.Println()
	fmt.Println("page", page.PageNumber)
	fmt.Println()
	fmt.Println(page.Text)

	if len(page.Choices) == 0 {
		return
	}

	fmt.Println(separator)
	fmt.Println()
	for _, choice := range page.Choices {
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

func isTerminalPage(page Page) bool {
	return len(page.Choices) == 0
}
