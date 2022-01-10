package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
)

type Row struct {
	ID   string `csv:"id"`
	Name string `csv:"名前"`
	Inci string `csv:"INCI"`
}

// func NewRow(id string, name string, inci string) *Row {
// 	return &Row{
// 		ID:   id,
// 		Name: name,
// 		Inci: inci,
// 	}
// }

func (row *Row) isValid() bool {
	// NOTE(okubo): ルール決めておくべきかも
	if &row.Name == nil || &row.Inci == nil {
		return false
	}

	return true
}

func main() {
	row := &Row{}
	rows := []*Row{}

	file, _ := os.OpenFile("koreandata.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()

	doc, err := goquery.NewDocument("https://kcia.or.kr/cid/search/ingd_list.php?page=1")
	if err != nil {
		fmt.Print("url scarappting faild")
	}

	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indextd int, cell *goquery.Selection) {
				// fmt.Println("index is ", indextd)
				if indextd == 0 {
					row.ID = cell.Text()
				} else if indextd == 1 {
					row.Name = cell.Text()
				} else if indextd == 2 {
					row.Inci = cell.Text()
				}
				// row = append(row, cell.Text())
			})

			if row.isValid() {
				fmt.Println(row)
				rows = append(rows, row)
			}
			row = &Row{ID: "", Name: "", Inci: ""}

		})
	})

	gocsv.MarshalFile(&rows, file)
	fmt.Println(&rows)
}
