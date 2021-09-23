package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Book struct {
	title string
	price float64
	quantity int
}

func main() {
	res := make([]Book, 0)
	buf, inputError := ioutil.ReadFile("products.txt")
	if inputError != nil {
		panic(inputError)
	}
	r := csv.NewReader(strings.NewReader(string(buf)))
	r.Comma = ';'
	//r.Comment = '\n'
	for {
		records, err := r.ReadAll()
		if err != nil {
			panic(err)
		}
		if len(records) > 0 {
			fmt.Println(records)
			var v1 string
			var v2 float64
			var v3 int
			for _, record := range records {
				v1 = record[0]
				v2, _ = strconv.ParseFloat(record[1], 32)
				v3, _ = strconv.Atoi(record[2])
				res = append(res, Book{v1, v2, v3})
			}
			fmt.Println(res)
		}else {
			break
		}
	}
	/**
	[[The ABC of Go 25.5 1500] [Functional Programming with Go 56 280] [Go for It 45.9 356] [The Go Way 55 500]]
	[{The ABC of Go 25.5 1500} {Functional Programming with Go 56 280} {Go for It 45.900001525878906 356} {The Go Way 55 500}]
	*/
}
