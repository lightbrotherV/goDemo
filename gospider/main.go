package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/lightbrotherV/lightGoJson"

	"github.com/PuerkitoBio/goquery"
)

const (
	phantomjspath = "./phantomjs"
)

var result map[string][]string //保存帖子数据

var userResult map[string]map[string]string //保持用户数据

var wg sync.WaitGroup //定义一个同步等待的组

var searchDay int = 7

var pastStamp, nowStamp int64 //待搜查以及现在的时间戳

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

//获取帖子列表
func getInvitation(seachUrl string, i int) {
	search(seachUrl, func(dom *goquery.Document) {
		resultIndex := 0
		dom.Find(".topic-list-item").Each(func(i int, subDom *goquery.Selection) {
			//获取每个帖子的时间戳
			invitationTimeStr, _ := subDom.Find("td .post-activity .relative-date").Attr("data-time")
			invitationTimestamp, _ := strconv.Atoi(invitationTimeStr)
			//是7天前的帖子
			if int64(invitationTimestamp/1000) > pastStamp {
				//帖子id
				invitationId, _ := subDom.Attr("data-topic-id")
				//帖子标题
				if result[invitationId] == nil {
					result[invitationId] = make([]string, 1000)
				}
				result[invitationId][resultIndex] = subDom.Find("td .link-top-line .title").Text() //帖子名字
				resultIndex++
				result[invitationId][resultIndex] = subDom.Find(".category .category-name").Text() //帖子分类
				resultIndex++
				result[invitationId][resultIndex] = time.Unix(int64(invitationTimestamp/1000), 0).Format("2006-01-02 15:04:05") //发布时间
				resultIndex++
				result[invitationId][resultIndex] = "http://blockgeek.org/t/topic/" + invitationId
				resultIndex++
				wg.Add(1)
				go getInvitationInfo(invitationId)
			}
		})
	})
}

//获取帖子用户名字
func getInvitationInfo(invitationId string) {
	//帖子信息
	search("http://blockgeek.org/t/topic/"+invitationId, func(dom *goquery.Document) {
		var username string
		dom.Find(".post-stream .topic-post").Each(func(i int, subDom *goquery.Selection) {
			username = subDom.Find(".username a").Text()
			wg.Add(1)
			go getUserInfo(username, invitationId)
		})
	})
}

//根据用户英文名查询用户信息
func getUserInfo(username string, invitationId string) {
	defer wg.Done()
	resp, err := http.Get(fmt.Sprintf("http://blockgeek.org/u/%s.json", username))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}
	defer resp.Body.Close()
	buf := make([]byte, 10000)
	for {
		n, _ := resp.Body.Read(buf)
		if 0 == n {
			break
		}
	}
	var jsonStr lightGoJson.LightJsonByte
	jsonStr = buf
	jsonMap := jsonStr.LightDecode()
	fmt.Println(jsonMap)
}

func main() {
	fmt.Println("爬虫正在努力爬取，行走速度可能较慢~")
	//获取待搜查帖子的时间戳
	nowStamp = time.Now().Unix()
	pastStamp = nowStamp - int64(3600*24*searchDay)
	result = make(map[string][]string)
	for i := 0; i < 1; i++ {
		wg.Add(1)
		seachUrl := fmt.Sprintf("http://blockgeek.org/latest?no_definitions=true&no_subcategories=false&page=%d", i)
		go getInvitation(seachUrl, i)
	}
	wg.Wait()
	fmt.Println(result)
}
