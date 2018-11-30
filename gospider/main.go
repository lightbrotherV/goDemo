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
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/lightGoExcel"
)

const (
	phantomjspath = "./phantomjs"
)

type mapSp1 map[string][]string
type mapSp2 map[string]string

var mapSp1Lock sync.Mutex
var mapSp2Lock sync.Mutex
var wg sync.WaitGroup

//读写加锁
func (temp mapSp1) Get(index1 string, index2 int) string {
	mapSp1Lock.Lock()
	defer mapSp1Lock.Unlock()
	return temp[index1][index2]
}
func (temp mapSp1) Set(index1 string, index2 int, val string) {
	mapSp1Lock.Lock()
	defer mapSp1Lock.Unlock()
	temp[index1][index2] = val
}

func (temp mapSp1) GetAS(index1 string) []string {
	mapSp1Lock.Lock()
	defer mapSp1Lock.Unlock()
	return temp[index1]
}

func (temp mapSp1) SetAS(index1 string, value []string) {
	mapSp1Lock.Lock()
	defer mapSp1Lock.Unlock()
	temp[index1] = value
}

func (temp mapSp2) Get(index1 string) string {
	mapSp2Lock.Lock()
	defer mapSp2Lock.Unlock()
	return temp[index1]
}
func (temp mapSp2) Set(index1 string, val string) {
	mapSp2Lock.Lock()
	defer mapSp2Lock.Unlock()
	temp[index1] = val
}

var result mapSp1 //保存帖子数据

var userResult mapSp2 //保持用户数据

var searchDay float32 = 7

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
func getInvitation(seachUrl string, i *int) {
	search(seachUrl, func(dom *goquery.Document) {
		dom.Find(".topic-list-item").Each(func(indexI int, subDom *goquery.Selection) {
			resultIndex := 0
			//获取每个帖子的时间戳
			invitationTimeStr, _ := subDom.Find("td .post-activity .relative-date").Attr("data-time")
			invitationTimestamp, _ := strconv.Atoi(invitationTimeStr)
			//是否是指定时间
			if int64(invitationTimestamp/1000) > pastStamp {
				//帖子id
				invitationId, _ := subDom.Attr("data-topic-id")
				//帖子标题
				if result.GetAS(invitationId) == nil {
					result.SetAS(invitationId, make([]string, 100))
				}
				result.Set(invitationId, resultIndex, subDom.Find("td .link-top-line .title").Text()) //帖子名字
				resultIndex++
				result.Set(invitationId, resultIndex, subDom.Find(".category .category-name").Text()) //帖子分类
				resultIndex++
				result.Set(invitationId, resultIndex, time.Unix(int64(invitationTimestamp/1000), 0).Format("2006-01-02 15:04:05")) //发布时间
				resultIndex++
				if utf8.RuneCountInString(subDom.Find(".cooked").Text()) > 500 {
					result.Set(invitationId, resultIndex, "文章")
					resultIndex++
				} else {
					result.Set(invitationId, resultIndex, "提问")
					resultIndex++
				}
				wg.Add(1)
				go getInvitationInfo(invitationId, resultIndex)
			} else if !subDom.HasClass("pinned") {
				*i = -2
			}

		})
	})
}

//获取帖子信息
func getInvitationInfo(invitationId string, resultIndex int) {
	defer wg.Done()
	//帖子信息
	search("http://blockgeek.org/t/topic/"+invitationId, func(dom *goquery.Document) {
		var username string
		dom.Find(".post-stream .topic-post").Each(func(i int, subDom *goquery.Selection) {
			username = subDom.Find(".username a").Text()
			result.Set(invitationId, resultIndex, username)
			resultIndex++
			getUserInfo(username, invitationId)
		})
	})
}

//根据用户英文名查询用户信息
func getUserInfo(username string, invitationId string) {
	//根据已有数据去除不必要的请求
	if userResult.Get(username) != "" {
		fmt.Println("已获取：", username)
		return
	}
	search(fmt.Sprintf("http://blockgeek.org/u/%s.json", username), func(dom *goquery.Document) {
		buf := dom.Find("pre").Text()
		jsonObj, _ := simplejson.NewJson([]byte(buf))
		var name, wallet string
		if jsonObj != nil {
			name, _ = jsonObj.Get("user").Get("name").String()
			wallet, _ = jsonObj.Get("user").Get("user_fields").Get("1").String()
			author := fmt.Sprintf("中文名：%s\n英文名：%s\n钱包地址：%s\n\n", name, username, wallet)
			fmt.Println(author)
			userResult.Set(username, author)
		} else {
			fmt.Println("未获取：", username)
		}
	})
}

func exportData() {
	var res [][]string = make([][]string, 100)
	i := 0
	for _, v1 := range result {
		i++
		j := 0
		for k2, v2 := range v1 {
			if res[i] == nil {
				res[i] = make([]string, 200)
			}
			if k2 < 4 {
				res[i][j] = v2
			} else {
				if userResult[v2] == "" {
					res[i][j] = v2
				} else {
					res[i][j] = userResult[v2]
				}
			}
			j++
		}
	}
	e := lightGoExcel.LightExecl{}
	e.Init()
	// e.AddTitle([]string{"t1", "t2", "t3"})
	e.SaveFile("test.xlsx", res)
}

func main() {
	fmt.Println("爬虫正在努力爬取，行走速度可能较慢~")
	//获取待搜查帖子的时间戳
	nowStamp = time.Now().Unix()
	pastStamp = nowStamp - int64(3600*24*searchDay)
	// pastStamp = nowStamp - int64(3600*2)
	result = make(mapSp1)
	userResult = make(mapSp2)
	for i := 0; i != -1; i++ {
		seachUrl := fmt.Sprintf("http://blockgeek.org/latest?no_definitions=true&no_subcategories=false&page=%d", i)
		fmt.Println(i)
		getInvitation(seachUrl, &i)
	}
	wg.Wait()
	exportData()
	fmt.Println(len(result))
	// fmt.Println(result)
}
