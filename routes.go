package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	Text string
}

var (
	staticDir = getAbsDirPath() + "/static/"
	templatesDir = staticDir + "/templates"
	templates = template.Must(template.ParseFiles(templatesDir + "/index.html",
	))
)

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

func getAbsDirPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("Ошибка при получении пути к статическим ресурсам: %v\n", err)
	}
	return path
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

	insertBookmark(alias, link)

	w.Write([]byte(aliasLink));
}

func bookmarksHandler(w http.ResponseWriter, r *http.Request) {
	links := getAllBookmarks()

	tmpl := template.Must(template.ParseFiles(templatesDir + "/bookmarks.html"))

	err := tmpl.ExecuteTemplate(w, "bookmarks.html", links)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} 
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templatesDir + "/map.html"))

	err := tmpl.ExecuteTemplate(w, "map.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} 
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templatesDir + "/bookmarks.html"))

	err := tmpl.ExecuteTemplate(w, "bookmarks.html", nil)
	if err != nil {
		fmt.Printf("Ошибка при загрузке страницы закладок: %v\n", err)
	}
}

func getBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	links := getAllBookmarks()
	w.Write(links)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/url/"):];
	url := getBookmark(path)

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Println(body)
}
