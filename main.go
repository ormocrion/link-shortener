package main

import (
	"log"
	"net/http"
)	

func main() {
	dbConnect()
	defer dbClose()
	http.HandleFunc("/scripts/", staticHandler)
	http.HandleFunc("/css/", staticHandler)
	http.HandleFunc("/", indexHandler)	
	http.HandleFunc("/short/", shortHandler)
	http.HandleFunc("/url/", redirectHandler)
	http.HandleFunc("/bookmarks/", bookmarksHandler)
	http.HandleFunc("/service/", serviceHandler)
	http.HandleFunc("/test/", testHandler)
	http.HandleFunc("/map/", mapHandler)
	http.HandleFunc("/getBookmarks/", getBookmarksHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}