package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// var headings []string
	var row []string
	var rows [][]string

	doc, err := goquery.NewDocument("https://kcia.or.kr/cid/search/ingd_list.php?page=1")
	if err != nil {
		fmt.Print("url scarappting faild")
	}

	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indextd int, cell *goquery.Selection) {
				row = append(row, cell.Text())
			})

			rows = append(rows, row)
			row = nil
		})
	})

	fmt.Println(rows)
}
