package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	phantomjspath = "./phantomjs"
)

var result map[string]map[string]string //保存最终数据

var wg sync.WaitGroup //定义一个同步等待的组

var searchDay int = 7

//md5加密
func md5Encode(code string) string {
	h := md5.New()
	h.Write([]byte(code))
	return hex.EncodeToString(h.Sum(nil))
}

//创建phantomjs运行的js文件
func setPhantomjs(url string) {
	content := `
	var page = require('webpage').create();
	var url = "` + url + `";
	page.open(url, function (status) {
		if (status !== 'success') {  
			console.log('Unable to post!');  
		} else {
			console.log(page.content)
		}
		phantom.exit();
	})
	`
	ioutil.WriteFile(md5Encode(url)+".js", []byte(content), 0777)
}

//抓包获取信息
func search(url string, domDeal func(*goquery.Document)) {
	//生成运行的js文件
	setPhantomjs(url)
	defer os.Remove(md5Encode(url) + ".js")
	//运行phantomjs
	cmd := exec.Command(phantomjspath, md5Encode(url)+".js")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	// //debug
	// ioutil.WriteFile(md5Encode(url), opBytes, 0777)

	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(opBytes))
	if err != nil {
		log.Fatalln(err)
	}
	//准备工作完毕，开始筛选dom元素
	domDeal(dom)
	//删除运行js的脚本
	os.Remove(url)
}

func main() {
	//获取待搜查帖子的时间戳
	futureStamp := time.Now().Unix() - int64(3600*24*searchDay)
	result = make(map[string]map[string]string)
	search("http://blockgeek.org/latest?no_definitions=true&no_subcategories=false&page=0", func(dom *goquery.Document) {

		dom.Find(".topic-list-item").Each(func(i int, subDom *goquery.Selection) {
			//获取每个帖子的时间戳
			invitationTimeStr, _ := subDom.Find("td .post-activity .relative-date").Attr("data-time")
			invitationTimestamp, _ := strconv.Atoi(invitationTimeStr)
			//是7天前的帖子
			if int64(invitationTimestamp/1000) > futureStamp {
				// //帖子id
				invitationId, _ := subDom.Attr("data-topic-id")
				// //帖子标题
				if result[invitationId] == nil {
					result[invitationId] = make(map[string]string)
				}
				result[invitationId]["title"] = subDom.Find("td .link-top-line .title").Text()
				result[invitationId]["category"] = subDom.Find(".category .category-name").Text()
				result[invitationId]["time"] = time.Unix(int64(invitationTimestamp/1000), 0).Format("2006-01-02 15:04:05")
				fmt.Println(result[invitationId]["time"])
			}
		})
	})
}
