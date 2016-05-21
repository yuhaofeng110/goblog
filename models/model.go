package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/deepzz0/go-common/log"
	"github.com/deepzz0/go-common/monitor"
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
	FeedFile     = "./static/feed.xml"
	SiteFile     = "./static/sitemap.xml"
	RobotsFile   = "./static/robots.txt"
)

var Blogger *User

func init() {
	path, _ := os.Getwd()
	if err := ip17mon.Init(path + "/conf/17monipdb.dat"); err != nil {
		log.Fatal(err)
	}
	// 以下三句保证调用顺序
	UMgr.loadUsers()
	Blogger = UMgr.Get("deepzz")
	if Blogger == nil { // 从配置初始化用户
		f, err := os.Open(path + "/conf/backup/user.json")
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
		UMgr.RegisterUser(&user)
		code := UMgr.UpdateUsers()
		if code != RS.RS_success {
			panic("更新用户数据失败。")
		}
		Blogger = UMgr.Get("deepzz")
	}
	// 开启警告邮件
	log.SetEmail(Blogger.Email)
	TMgr.loadTopics()
	ManageData.LoadData()
	monitor.HookOnExit("flushdata", flushdata)
	monitor.Startup()
	go RequestM.Saver()
	go scheduleTopic()
	go scheduleUser()
	go NewDay()
}

// 新的一天
func NewDay() {
	t := time.NewTicker(time.Minute)
	Today := time.Now()
	for {
		select {
		case <-t.C:
			if time.Now().Day() != Today.Day() {
				Today = time.Now()
				ManageData.LoadData()
				ManageData.CleanData(Today)
			}
		}
	}
}

func flushdata() {
	UMgr.UpdateUsers()
	TMgr.UpdateTopics()
	ManageConf.UpdateConf()
}
