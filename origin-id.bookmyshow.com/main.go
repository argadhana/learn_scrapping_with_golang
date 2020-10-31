package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	Title string
	Book  string
}

func main() {
	res, err := http.Get("https://origin-id.bookmyshow.com/movies")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]Product, 0)

	doc.Find(".mv-row").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(Product)
		row.Title = sel.Find(".__name").Text()
		row.Book, _ = sel.Find(".__movie-name").Attr("href")
		rows = append(rows, *row)
	})

	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bts))
}
