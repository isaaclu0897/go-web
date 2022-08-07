package main

import (
	"Unit3/tool"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, _ *http.Request) {
	var pageHTML []byte
	pageHTML, err := os.ReadFile("html/home.html")
	if err != nil {
		log.Fatal("os.ReadFile: ", err)
	}

	w.Write([]byte(pageHTML))
}

func contactPage(w http.ResponseWriter, _ *http.Request) {
	var pageHTML []byte
	pageHTML, err := os.ReadFile("html/contact.html")
	if err != nil {
		log.Fatal("os.ReadFile: ", err)
	}
	w.Write([]byte(pageHTML))

}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contact", contactPage)

	tool.Open("http://127.0.0.1:9453")

	err := http.ListenAndServe(":9453", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
