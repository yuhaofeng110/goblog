package models

import (
	// "net/http"
	// "regexp"
	// "strings"
	"time"

	"github.com/deepzz0/go-common/log"
	db "github.com/deepzz0/go-common/mongo"
	// "github.com/deepzz0/go-common/useragent"
	"gopkg.in/mgo.v2/bson"
)

/////////////////////////////////////////////////////////////////////////
const (
	BingBot      = "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
	BaiduSpider  = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
	GoogleBot    = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	YahooChina   = "Mozilla/5.0 (compatible; Yahoo! Slurp China; http://misc.yahoo.com.cn/help.html)"
	YodaoBot     = "Mozilla/5.0 (compatible; YodaoBot/1.0; http://www.yodao.com/help/webmaster/spider/; )"
	YoudaoBot    = "Mozilla/5.0 (compatible; YoudaoBot/1.0; http://www.youdao.com/help/webmaster/spider/; )"
	MSNBot       = "msnbot/2.1"
	SougouSpider = "sogou spider"
	SousouSpider = "Sosospider+(+http://help.soso.com/webspider.htm)"
	YahooSeeker  = "YahooSeeker/1.2 (compatible; Mozilla 4.0; MSIE 5.5; yahooseeker at yahoo-inc dot com ; http://help.yahoo.com/help/us/shop/merchant/)"

	android  = "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.23 Mobile Safari/537.36"
	ios      = "Mozilla/5.0 (iPhone; CPU iPhone OS 9_3_1 like Mac OS X) AppleWebKit/601.1 (KHTML, like Gecko) CriOS/50.0.2661.77 Mobile/13E238 Safari/601.1.46"
	ipad     = "Mozilla/5.0 (iPad; CPU OS 9_0_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13A452 Safari/601.1"
	winphone = "Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 630) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537"
	uc       = "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_2 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5 UCBrowser/10.9.8.738"
	macox    = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.86 Safari/537.36"
)

// type Viewer struct {

// }

// func NewViewer(request *http.Request, sessionid string) *Viewer {
// 	// header := request.Header
// 	// userAgent := header.Get("User-Agent")

// 	// reg, err := regexp.Compile(``)

// 	// v := &Viewer{RequestURI: header.Get("RequestURI"), RemoteAddr: header.Get("RemoteAddr"), Referer: header.Get("Referer"), SessionId: sessionid, Time: time.Now()}
// 	return nil
// }

// type ViewerManage struct {
// 	SidToTime map[string]time.Time
// 	Ch        chan *Viewer
// }

// var ViewM = NewViewM()

// func NewViewM() *ViewerManage {
// 	return &ViewerManage{SidToTime: make(map[string]time.Time)}
// }

// func (m *ViewerManage) Update(sid string) {
// 	m.SidToTime[sid] = time.Now()
// }

// func (m *ViewerManage) Get(sid string) time.Time {
// 	return m.SidToTime[sid]
// }

// func (m *ViewerManage) Flush() {
// 	for k, v := range m.SidToTime {
// 		if v.Add(30 * time.Minute).Before(time.Now()) {
// 			delete(m.SidToTime, k)
// 		}
// 	}
// }

// func (m *ViewerManage) Saver() {
// 	t := time.NewTicker(time.Minute)
// 	for {
// 		select {
// 		case viewer := <-m.Ch:
// 			err := db.Insert(DB, C_VIEWER, viewer)
// 			if err != nil {
// 				log.Error(err)
// 			}
// 		case <-t.C:
// 			m.Flush()
// 		}
// 	}
// }

///////////////////////////////////////////////////////////////////////////
type Leftbar struct {
	ID    string // 内部ID
	Title string // 说明
	Extra string // 链接
	Text  string // 显示名称
}

///////////////////////////////////////////////////////////////////////////
// type Request struct {
// 	Referer    string              // 请求来源
// 	URL        string              // 访问页面
//  Major	   string			   // 主版本
// 	RemoteAddr string              // 请求IP
// 	SessionID  string              // 请求session
// 	UserAgent  useragent.UserAgent //
// 	Time       time.Time           // 请求时间
// }

///////////////////////////////////////////////////////////////////////////
type Verification struct {
	Name       string // pk
	Content    string
	CreateTime time.Time
}

func NewVerify() *Verification {
	return &Verification{CreateTime: time.Now()}
}

var ManageConf = LoadConf()

type Config struct {
	Name       string
	SiteVerify map[string]*Verification
}

func LoadConf() *Config {
	conf := &Config{Name: "config", SiteVerify: make(map[string]*Verification)}
	err := db.FindOne(DB, C_CONFIG, bson.M{"name": "config"}, conf)
	if err != nil {
		log.Error(err)
	}
	return conf
}

func (conf *Config) GetVerification(name string) *Verification {
	return conf.SiteVerify[name]
}

func (conf *Config) AddVerification(verify *Verification) {
	conf.SiteVerify[verify.Name] = verify
}

func (conf *Config) DelVerification(name string) {
	conf.SiteVerify[name] = nil
	delete(conf.SiteVerify, name)
}

func (conf *Config) UpdateConf() {
	err := db.Update(DB, C_CONFIG, bson.M{"name": "config"}, conf)
	if err != nil {
		log.Error(err)
	}
}
