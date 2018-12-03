package main

import (
	"encoding/json"
	"fmt"
	"go/token"
	"io"
	"log"
	"strings"
)

type Record struct {
	AuthorRaw interface{} `json:"author"`
	Title     string      `json:"title"`
	URL       string      `json:"url"`

	AuthorEmail string
	AuthorID    uint64
}

var ArrRecord struct {
	arr []Record
}

var (
	le = fmt.Sprintf("%v", token.LBRACE)
	re = fmt.Sprintf("%v", token.RBRACE)
	lk = fmt.Sprintf("%v", token.LBRACK)
	rk = fmt.Sprintf("%v", token.RBRACK)
)

func DecodeRecordArray(r io.Reader) (x []*Record, err error) {

	arr := make([]*Record, 0)

	dec := json.NewDecoder(r)

	for dec.More() {
		t, err := dec.Token()
		t = fmt.Sprintf("%v", t)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if t != le &&
			t != re &&
			t != lk &&
			t != rk {
			var rec = new(Record)
			dec.Decode(&rec)
			arr = append(arr, rec)
		}
	}

	return arr, nil
}

func main() {

	dt := `[{
		"author": "attila@attilaolah.eu",
		"title":  "My Blog",
		"url":    "http://attilaolah.eu"
	  }, {
		"author": 1234567890,
		"title":  "Westartup",
		"url":    "http://www.westartup.eu"
	  }]`

	r := strings.NewReader(dt)

	res, err := DecodeRecordArray(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(res[0])

}
