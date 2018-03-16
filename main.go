package main

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"fmt"
	"strings"
)

type Herf struct {
	Type    string
	Content string
	Img     string
}

func main() {

	re := GetContent()
	println("--------------------------------------------------------------")
	print(re[1].Content)
	for a, value := range re {
		print(strconv.Itoa(a) + "---" + value.Content)
	}

}

func GetContent() []Herf {
	c := make([]Herf, 0)
	var url = "https://www.qiushibaike.com/"
	doc, _ := goquery.NewDocument(url)
	doc.Find("div[id^='qiushi']").Each(func(i int, selection *goquery.Selection) {
		println(i)

		s := selection.Find("div.content").Find("span").Text()
		//fmt.Printf(s)
		if s != "" {
			var d = Herf{}
			d.Content = s

			src, _ := selection.Find("img.illustration").Attr("src")
			d.Img = "http:" + src
			fmt.Println(strings.Replace(src, "//", "http://", -1))


			c = append(c, d)
		}
	})
	return c
}
