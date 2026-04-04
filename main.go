package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)	

type Bookmark struct {
	// Link string `json:"link"`
 	// Alias string `json:"alias"`
	Link string
	Alias string
}

type Page struct {
	Title string
	Text string
}

var listOfBookmarks = make(map[string]string, 100_000)
var input string

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func consoleHandler() {
	fmt.Printf("Введите ссылку: ")
	fmt.Scan(&input)
	input = checkLink(input)
	
	alias := generateRandomString(6)
	listOfBookmarks[alias] = input 
	fmt.Printf("Ваша короткая ссылка: %s\n", alias)
	fmt.Printf("Cохранённая ссылка: %s", listOfBookmarks[alias])

	/*
	bookmark := Bookmark{input, generateRandomString(7)}

	bookmarkJson, err := marshalJson(bookmark)
	fileSaver(bookmarkJson, err)
	jsonData := fileReader()
	bookmarkForRecovery := unmarshalJson(jsonData)

	fmt.Println(bookmarkForRecovery.Link)
	*/
}

func checkLink(input string) string {
	if (strings.Contains(input, "http://") || strings.Contains(input, "https://")) {
		return input
	}
		return "http://" + input
} 

func marshalJson(bookmark Bookmark) ([]byte, error) {
	bookmarkJson, err := json.Marshal(bookmark)
	if err != nil {
		panic(err)
	}
	return bookmarkJson, err
}

func unmarshalJson(jsonData []byte) Bookmark {
	var bookmarkForRecovery Bookmark
	err := json.Unmarshal(jsonData, &bookmarkForRecovery)
	if err != nil {
		panic(err)
	}
	return bookmarkForRecovery
}

func fileSaver(bookmarkJson []byte, err error) {
		err = os.WriteFile("user.json", bookmarkJson, 0644)
	if err != nil {
		panic(err)
	}
}

func fileReader() []byte {
	jsonData, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}
	return jsonData
}

func generateRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte

	for i := 0; i < length; i++ {
		index := seededRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title: "Cокращатель ссылок",
		Text: "Давайте поместим ссылку ниже и посмотрим на то, как она станет короче!",
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, data)
} 

func redirectToUrl(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/url/"):]
	url := listOfBookmarks[path]
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func main() {
	// consoleHandler()
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/url/", redirectToUrl)
	log.Fatal(http.ListenAndServe(":80", nil))
}