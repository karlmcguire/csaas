package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Page struct {
	Title string
	Desc  string
	Month string
	Day   string
	Year  string
	Date  time.Time
}

func (p *Page) Render() []byte {
	var buf bytes.Buffer

	p.Title = bluemonday.UGCPolicy().Sanitize(p.Title)

	buf.Write(T0)
	buf.WriteString(p.Title)
	buf.Write(T1)
	buf.WriteString(p.Title)
	buf.Write(T2)
	buf.WriteString(
		humanize.Time(p.Date),
	)
	buf.Write(T3)

	buf.Write(
		bluemonday.UGCPolicy().SanitizeBytes(
			blackfriday.MarkdownCommon([]byte(p.Desc)),
		),
	)

	return buf.Bytes()
}

// Handler handles all requests to api.csaas.xyz. It validates URL parameters
// and responds with a coming soon page.
func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		page = &Page{}
		err  error
	)

	if page.Title = r.FormValue("title"); page.Title == "" {
		w.Write([]byte("Title length must be > 0."))
		return
	}
	if page.Desc = r.FormValue("desc"); page.Desc == "" {
		w.Write([]byte("Desc length must be > 0."))
		return
	}
	if page.Month = r.FormValue("month"); page.Month == "" {
		w.Write([]byte("Month length must be > 0."))
		return
	}
	if page.Day = r.FormValue("day"); page.Day == "" {
		w.Write([]byte("Day length must be > 0."))
		return
	}
	if page.Year = r.FormValue("year"); page.Year == "" {
		w.Write([]byte("Year length must be > 0."))
		return
	}

	if page.Date, err = time.Parse(
		"01/02/2006",
		fmt.Sprintf(
			"%s/%s/%s",
			page.Month,
			page.Day,
			page.Year,
		),
	); err != nil {
		w.Write([]byte("Invalid date provided."))
		return
	}

	w.Write(page.Render())
}

var FAVICON []byte

func init() {
	var err error

	if FAVICON, err = ioutil.ReadFile("static/favicon.ico"); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Write(FAVICON)
	})

	http.HandleFunc("/", Handler)

	println("listening")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
