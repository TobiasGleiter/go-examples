package main

import (
	"log"
	"os"
	"fmt"
	"net/http"
	"html/template"
	"regexp"
	
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := "./data/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

var pgs []string

func pagesHandler(w http.ResponseWriter, r *http.Request) {
	folder := "./data/"

	dir, err := os.Open(folder)
	if err != nil {
        fmt.Println("Error opening directory:", err)
        return
    }
    defer dir.Close()


	files, err := dir.Readdir(-1)
    if err != nil {
        fmt.Println("Error reading directory contents:", err)
        return
    }

	pgs := make([]string, 0)
	for _, file := range files {
		l := len(file.Name()) - len(".txt")
		fiNa := file.Name()[:l]
		pgs = append(pgs, fiNa)
	}

	t, _ := template.ParseFiles("./tmpl/pages" + ".html")
    t.Execute(w, pgs)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }

    renderTemplate(w, "./tmpl/view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
	if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "./tmpl/edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
    }
}

func main() {
	http.HandleFunc("/", pagesHandler)
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

	fmt.Println("Server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}