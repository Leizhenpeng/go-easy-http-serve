package main

import (
	"fmt"
	"net/http"
)

func loginHandler(writer http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	userName := r.Form.Get("username")
	//password := r.Form.Get("password")

	fmt.Println("username: ", userName)
	fmt.Fprintf(writer, "welcome %s", userName)

}

func main() {
	fileServe := http.FileServer(http.Dir("./www"))
	http.Handle("/", fileServe)
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Listening on port 8080", "open: http://localhost:8080/index.html")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
