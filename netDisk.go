package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/levigross/grequests"
)

var wg sync.WaitGroup

//GetAddr 获取下载地址
func GetAddr() {

	now := time.Now()
	initalUrls := []string{"http://disk.bjsasc.com:8180/NetDisk/"}
	for _, url := range initalUrls {
		doc, err := goquery.NewDocument(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		doc.Find(".loginLogo").Each(func(i int, s *goquery.Selection) {
			src, exists := s.Find("img").Attr("src")
			if exists {
				wg.Add(1)
				go func() {
					defer wg.Done()
					fmt.Println(src)
					time.Sleep(3 * time.Second)
				}()
			}
		})

	}

	wg.Wait()
	//4M的带宽下载，需要16m36s，总大小202M,10个文件夹,560个文件
	fmt.Printf("下载任务完成，耗时:%#v\n", time.Now().Sub(now))
}

//GetToken get the token
func GetToken() string {
	username := "fengshaomin"
	pass := "1"
	url := "http://disk.bjsasc.com:8180/NetDisk/rest/mobile"
	//url := "http://127.0.0.1:8080/list"
	paras := &grequests.RequestOptions{Params: map[string]string{"userName": username, "passWord": EncryptPass(pass), "method": "login"}}
	res, err := grequests.Get(url, paras)
	if err != nil {
		fmt.Println(err)
	}
	var token Token

	fmt.Println(res)
	json.Unmarshal(res.Bytes(), &token)

	fmt.Println(token.Token)
	return token.Token
}

//EncryptPass enctypt the password
func EncryptPass(orig string) string {
	s, _ := TripleDesEncrypt([]byte(orig))
	encStr := base64.StdEncoding.EncodeToString(s)
	fmt.Println(encStr)
	return encStr
}

//Token token
type Token struct {
	Token string `json:"token"`
}
