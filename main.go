package main

import (
	"fmt"
	"html/template"
	"net/http"

	// renaming our package
	myFunc "ascii/ascii"
)

func StartServer(s string) {
	// Creating Local Server
	fmt.Println("Server is Working succesfully....")
	fmt.Println(http.ListenAndServe(s, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	// Check the reciving method is it valid or not
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Handling the wrong path error.
	if r.URL.Path != "/" {
		// Changing status of the Header
		w.WriteHeader(http.StatusNotFound)
		temp, err := template.ParseFiles("404.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
		return
	}
	// make the template and parse it
	temp, err := template.ParseFiles("index.html")
	// handlling the server error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func asciiArt(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()

	// Getting Form Values Text And Banner :)
	text := r.Form.Get("name")
	banner := r.Form.Get("banner")
	if text == "" || banner == "" {
		http.Error(w, "<h1>Bad Request</h1>", http.StatusBadRequest)
		return
	}
	// Proccessing Data and Getting ascii art result
	result, er := myFunc.Batata(banner, text)
	// Handling the case if the banner is not exists
	if er {
		w.WriteHeader(http.StatusNotFound)
		temp, _ := template.ParseFiles("404-banner.html")
		temp.Execute(w, nil)
		return
	}

	// Creating new Object from the template
	temp, errr := template.ParseFiles("ascii-art.html")
	if errr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, result)
}

func main() {
	// Handling Home (root) Page
	http.HandleFunc("/", index)
	// Handling /Ascci-art path
	http.HandleFunc("/ascii-art", asciiArt)
	// host style file to serve it later
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	// we intialized the local server with 8080 port
	StartServer(":8080")
}
