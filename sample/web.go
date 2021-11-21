package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"html/template"
)

type Page struct {
	Title	string
	Body	[]byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
	p, err := loadPage(title)
	if err != nil {
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
	p, _ := loadPage(title)
	if p != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, htmlFileName string, p *Page){
	t, err := template.ExecuteTemplate(w, htmlFileName + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, p)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err = p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is body")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}