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

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
)

const (
	phantomjspath = "./phantomjs"
)

var result map[string][]string //保存最终数据

var wg sync.WaitGroup //定义一个同步等待的组

var searchDay int = 7

var futureStamp, nowStamp int64 //待搜查以及现在的时间戳

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
	defer wg.Done()
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

//将获取数据存入xlsx表中
func saveToExecl(data map[string][]string) {
	var alphabet []string
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	hang := 1
	line := 0
	xlsx := excelize.NewFile()
	title := map[string]string{"A1": "标题", "B1": "分类", "C1": "发帖时间", "D1": "作者（英文）", "E1": "作者（中文）", "F1": "文章链接"}
	for k, v := range title {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for _, v1 := range data {
		line = 0
		hang++
		for _, v2 := range v1 {
			lineName := alphabet[line]
			xlsx.SetCellValue("Sheet1", lineName+strconv.Itoa(hang), v2)
			line++
		}
	}
	fileName := time.Unix(futureStamp, 0).Format("2006年1月2日15时4分5秒") + "至" + time.Unix(nowStamp, 0).Format("2006年1月2日15时4分5秒")
	err := xlsx.SaveAs(fileName + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	fmt.Println("爬虫正在努力爬取，行走速度可能较慢~")
	//获取待搜查帖子的时间戳
	nowStamp = time.Now().Unix()
	futureStamp = nowStamp - int64(3600*24*searchDay)
	result = make(map[string][]string)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		seachUrl := fmt.Sprintf("http://blockgeek.org/latest?no_definitions=true&no_subcategories=false&page=%d", i)
		go search(seachUrl, func(dom *goquery.Document) {
			dom.Find(".topic-list-item").Each(func(i int, subDom *goquery.Selection) {
				//获取每个帖子的时间戳
				invitationTimeStr, _ := subDom.Find("td .post-activity .relative-date").Attr("data-time")
				invitationTimestamp, _ := strconv.Atoi(invitationTimeStr)
				//是7天前的帖子
				if int64(invitationTimestamp/1000) > futureStamp {
					//帖子id
					invitationId, _ := subDom.Attr("data-topic-id")
					//帖子标题
					if result[invitationId] == nil {
						result[invitationId] = make([]string, 6)
					}
					result[invitationId][0] = subDom.Find("td .link-top-line .title").Text()
					result[invitationId][1] = subDom.Find(".category .category-name").Text()
					result[invitationId][2] = time.Unix(int64(invitationTimestamp/1000), 0).Format("2006-01-02 15:04:05")
					result[invitationId][3], _ = subDom.Find(".posters a").Attr("data-user-card")
					result[invitationId][4] = "懒，咕咕咕"
					result[invitationId][5] = "http://blockgeek.org/t/topic/" + invitationId
				}
			})
		})
	}
	wg.Wait()
	saveToExecl(result)
}
