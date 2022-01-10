package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://kcia.or.kr/cid/search/ingd_list.php?page=1")
	if err != nil {
		fmt.Print("url scarappting faild")
	}

	doc.Find("tbody > tr").Each(func(_ int, s *goquery.Selection) {
		fmt.Println("-------")
		fmt.Println(s.Text())
		fmt.Println("-------")
	})
}
