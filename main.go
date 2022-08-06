package main

import ( // 套件引入可以直接用括號包起來，看起來很舒服
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) { // 看起來func的需要傳入http讀寫器，http.Request用了*不知道是什麼意思
	// 解析參數，預設不會解析
	r.ParseForm() // Q：留個疑問，不解析會怎樣？

	//這些資訊是輸出到Server的
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello ISAAC!") // 這個寫入到 w 的是輸出到客戶端的
}

func main() {
	http.HandleFunc("/", sayhelloName)     // 設定路由與處理func
	err := http.ListenAndServe(":9453", nil) // 設定監聽的port
	if err != nil {
		log.Fatal("ListenAndServe: ", err) // Q：這裏報錯不知道會記錄到哪裏去？
	}
}
