package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

type Message struct {
	Msg string
}

var msg Message

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func homeUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Hello", "Aloha")
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg = Message{Msg: r.URL.Query()["message"][0]}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Allo Good", "Message Set")
}

func fetchMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Hello", msg.Msg)
}

func main() {
	http.HandleFunc("/", homeUrlHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/fetch", fetchMessage)
	http.HandleFunc("/set", setMessage)
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}