package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"log"
)

//For testing purpose
func main() {
	http.HandleFunc("/api/v0/available_products", upload)
	http.HandleFunc("/uaa/oauth/token", uaa)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func uaa(w http.ResponseWriter, r *http.Request) {
	token := `{
		"access_token" : "hello_world"
	}
	`
	w.Write([]byte(token))
	fmt.Printf("Write token %s", token)
}
