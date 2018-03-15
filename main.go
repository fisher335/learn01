package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Person struct {
	Name string
	Age  int
	Part string
}

func main() {

	//var c = grequests.RequestOptions{Headers: map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36"}}

	var url = "https://www.qiushibaike.com/"
	doc, _ := goquery.NewDocument(url)
	doc.Find("div[id^='qiushi']").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("div.content").Find("span").Text()

		fmt.Printf(s)

		src, _ := selection.Find("img.illustration").Attr("src")
		fmt.Println(strings.Replace(src,"//","http://",-1))

	})

}
