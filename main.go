package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	jsonFile, err := os.Open("../InputJson.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteValue)
	if err != nil {
		return
	}

	defer jsonFile.Close()

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	_, err := fmt.Fprintf(w, "Hello")
	if err != nil {
		return
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		_, err := fmt.Fprintf(w, "Error %v", err)
		if err != nil {
			return
		}
	}
	_, err := fmt.Fprintf(w, "Post Request succesfull")
	if err != nil {
		return
	}

	name := r.FormValue("name")

	_, err = fmt.Fprintf(w, "Hello  %s\n", name)
	if err != nil {
		return

	}

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/json/", jsonHandler)

	fmt.Printf("Server started at port 8080")

	//err := http.ListenAndServeTLS(":443", "/Users/mind/go/src/gofree/server.crt", "/Users/mind/go/src/gofree/server.key", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
