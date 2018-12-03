package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Record struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	URL    string `json:"url"`
}

func Decode(r io.Reader) (x *Record, err error) {
	x = new(Record)
	if err = json.NewDecoder(r).Decode(x); err != nil {
		return x, err
	}
	return x, nil
}

func main() {

	dt := `{
		"author": "attila@attilaolah.eu",
		"title":  "My Blog",
		"url":    "http://attilaolah.eu"
	  }`

	var res *Record
	r := strings.NewReader(dt)

	res, err := Decode(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
