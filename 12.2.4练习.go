package main

import "io/ioutil"

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save()  {
	err := ioutil.WriteFile(p.Title, p.Body, 0666)
	if err != nil {
		panic(err)
	}
}

func (p *Page) load()  {
	bs, err := ioutil.ReadFile(p.Title)
	if err != nil {
		panic(err)
	}
	p.Body = bs
}

func main() {
	//
}
