package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"text/template"
)

//缺少template(view/edit/save)

const lenPath = len("/view/")
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

var templates = make(map[string]*template.Template)

var err error

type Page struct {
	Title string
	Body []byte
}

func init() {
	for _, tmpl := range []string{"edit", "view"} {
		templates[tmpl] = template.Must(template.ParseFiles(tmpl + ".html"))
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	//创建一个只有当前用户拥有读写权限的文件
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)
	if err != nil {
		//找不到页面
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
	}
	renderTemplate(w,title, p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err:= load(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w,r,"/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}