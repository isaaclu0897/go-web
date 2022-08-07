package main

import (
	"Unit3/tool"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, _ *http.Request) {
	var homePageHTML string
	homePageHTML = `<!DOCTYPE html>
<html>
<meta charset="UTF-8">
<body>

<!-- Content will go here -->

<h1>Hello ISAAC!</h1>

</body>
</html>
`

	w.Write([]byte(homePageHTML))
}

func main() {
	http.HandleFunc("/", homePage)

	tool.Open("http://127.0.0.1:9453")

	err := http.ListenAndServe(":9453", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
