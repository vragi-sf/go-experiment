package main

import (
	"salesforce.com/ohana/wiki/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

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
	var url = "http://"+os.Getenv("SERVICE_ID")+"/receive?message="+r.URL.Query()["message"][0]
	response, _ := http.Get(url)
	if response.Status ==  "200 OK" {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "All Good", "Message Set")
	} else {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Not Good", "Message could not be Set")
	}
}

func receiveMessage(w http.ResponseWriter, r *http.Request) {
	var msg = model.Message{Msg: r.URL.Query()["message"][0]}
	model.AddMessage(msg)
}

func fetchMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Hello", model.GetMessage().Msg)
}

func main() {
	http.HandleFunc("/", homeUrlHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/fetch", fetchMessage)
	http.HandleFunc("/set", setMessage)
	http.HandleFunc("/receive", receiveMessage)
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}