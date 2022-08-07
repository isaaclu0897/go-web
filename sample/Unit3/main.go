package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func open(url string) error { // open url from browser
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

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

	open("http://127.0.0.1:9453")

	err := http.ListenAndServe(":9453", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
