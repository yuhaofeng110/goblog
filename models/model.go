package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/go-com/monitor"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/wangtuanjie/ip17mon"
)

const (
	DB         = "newblog" // database数据库
	C_USER     = "user"    // collections表
	C_TOPIC    = "topic"
	C_TOPIC_ID = "topic_id" // 文章ID计数
	C_REQUEST  = "request"  // 浏览记录
	C_CONFIG   = "config"   // 配置
)

const (
	TemplateFile = "./static/feedTemplate.xml"
	RobotsFile   = "./static/robots.txt"
	FeedFile     = "/data/goblog/feed.xml"
	SiteFile     = "/data/goblog/sitemap.xml"
)

var UMgr = NewUM()
var TMgr = NewTM()
var path, _ = os.Getwd()
var Blogger *User

func init() {
	if err := ip17mon.Init(path + "/conf/17monipdb.dat"); err != nil {
		log.Fatal(err)
	}
	// 以下三句保证调用顺序
	UMgr.loadUsers()
	Blogger = UMgr.Get("deepzz")
	if Blogger == nil { // 从配置初始化用户
		initAccount()
	}
	TMgr.loadTopics()
	// open error mail，email addr : Blogger.Email
	log.SetEmail(Blogger.Email)
	ManageData.LoadData()
	monitor.HookOnExit("flushdata", flushdata)
	monitor.Startup()

	go RequestM.Saver()
	go timer()
}

func initAccount() {
	f, err := os.Open(path + "/conf/init/user.json")
	if err != nil {
		panic(err)
	}
	user := User{}
	b, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(b, &user)
	if err != nil {
		panic(err)
	}
	user.PassWord = helper.EncryptPasswd(user.UserName, user.PassWord, user.Salt)
	UMgr.Register(&user)
	code := UMgr.Update()
	if code != RS.RS_success {
		panic("init failed。")
	}
	Blogger = UMgr.Get("deepzz")
}

// 新的一天
func timer() {
	t := time.NewTicker(time.Minute)
	Today := time.Now()

	tUser := time.NewTicker(time.Hour)
	tTopic := time.NewTicker(time.Minute * 10)
	for {
		select {
		case <-t.C:
			if time.Now().Day() != Today.Day() {
				Today = time.Now()
				ManageData.LoadData()
				ManageData.CleanData(Today)
			}
		case <-tUser.C:
			UMgr.Update()
		case <-tTopic.C:
			TMgr.DoDelete(time.Now())
		}
	}
}

func flushdata() {
	UMgr.Update()
	TMgr.Update()
	ManageConf.UpdateConf()
}
