package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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
	Text string
}

var (
	staticDir = getAbsDirPath() + "/static/"
	templatesDir = staticDir + "/templates"
	templates = template.Must(template.ParseFiles(templatesDir + "/index.html",
	))
)

var listOfBookmarks = make(map[string]string, 100_000)
var input string

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func consoleHandler() {
	fmt.Println(getAbsDirPath())
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

func redirectLinkFormer(alias string) string {
	aliasLink := "localhost/url/" + alias;
	
	return aliasLink;
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Text: "Укоротим ссылку без потери качества!",
	}
	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else {
		w.Header().Set("Content-Type", "text/css")
	}
	data, err := os.ReadFile(staticDir + path[1:])
	if err != nil {
		fmt.Print(err)
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Print(err)
	}
}

func shortHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	link := checkLink(string(body[:]));
	alias := generateRandomString(8);
	aliasLink := redirectLinkFormer(alias);
	listOfBookmarks[alias] = link;

	w.Write([]byte(aliasLink));
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/url/"):];
	url := listOfBookmarks[path];

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func getAbsDirPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func main() {
	http.HandleFunc("/scripts/", staticHandler)
	http.HandleFunc("/css/", staticHandler)
	http.HandleFunc("/", indexHandler)	
	http.HandleFunc("/short/", shortHandler)
	http.HandleFunc("/url/", redirectHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}