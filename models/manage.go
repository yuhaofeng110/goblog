package models

import (
	"net/http"
	"strings"
	"time"

	"github.com/deepzz0/go-common/log"
	db "github.com/deepzz0/go-common/mongo"
	tm "github.com/deepzz0/go-common/time"
	"github.com/deepzz0/go-common/useragent"
	"gopkg.in/mgo.v2/bson"
)

///////////////////////////////////////////////////////////////////////////
type Leftbar struct {
	ID    string // 内部ID
	Title string // 说明
	Extra string // 链接
	Text  string // 显示名称
}

///////////////////////////////////////////////////////////////////////////
type Request struct {
	Referer    string               // 请求来源
	URL        string               // 访问页面
	Major      int                  // 主版本
	RemoteAddr string               // 请求IP
	SessionID  string               // 请求session
	UserAgent  *useragent.UserAgent //
	Time       time.Time            // 请求时间
}

func NewRequest(r *http.Request) *Request {
	request := &Request{Time: time.Now()}
	request.Referer = r.Referer()
	request.URL = r.URL.String()
	request.Major = r.ProtoMajor
	request.RemoteAddr = r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")]
	request.UserAgent = useragent.ParseByRequest(r)
	sessionid, err := r.Cookie("SESSIONID")
	if err != nil {
		log.Error(err)
	}
	request.SessionID = sessionid.Value
	return request
}

type RequestManage struct {
	// lock sync.RWMutex
	// SSidToTime map[string]time.Time
	Ch chan *Request
}

var RequestM = NewRequestM()

func NewRequestM() *RequestManage {
	return &RequestManage{Ch: make(chan *Request, 20)}
}

func (m *RequestManage) Saver() {
	t := time.NewTicker(time.Minute * 1)
	for {
		select {
		case request := <-m.Ch:
			err := db.Insert(DB, C_REQUEST, request)
			if err != nil {
				log.Error(err)
			}
		case <-t.C:
			ManageData.LoadData()
			ManageData.LoadLatestRequst(1)
		}
	}
}

///////////////////////////////////////////////////////////////////////////
const (
	YESTERDAY = "yesterday"
	TODAY     = "today"
)

var ManageData = NewBaseData()

type BaseData struct {
	PV        map[string]int
	UV        map[string]int
	IP        map[string]int
	TimePV    map[string][]int
	EngineTop map[string]int
	PageTop   map[string]int
	Area      [][]int
	Latest    []*Request
}

func NewBaseData() *BaseData {
	bd := &BaseData{}
	bd.LoadData()
	bd.LoadLatestRequst(1)
	return bd
}

func (b *BaseData) LoadData() {
	b.PV = make(map[string]int)
	b.UV = make(map[string]int)
	b.IP = make(map[string]int)
	b.TimePV = make(map[string][]int)

	now := tm.New(time.Now())
	todayBegin := now.BeginningOfDay()
	todayEnd := now.EndOfDay()

	yestdBegin := todayBegin.Add(-24 * time.Hour)
	yestdEnd := todayEnd.Add(-24 * time.Hour)

	ms, c := db.Connect(DB, C_REQUEST)
	c.EnsureIndexKey("time")
	defer ms.Close()
	count, err := c.Find(bson.M{"time": bson.M{"$gte": yestdBegin, "$lt": yestdEnd}}).Count()
	if err != nil {
		log.Error(err)
	}
	b.PV[YESTERDAY] = count
	count, err = c.Find(bson.M{"time": bson.M{"$gte": todayBegin, "$lt": todayEnd}}).Count()
	if err != nil {
		log.Error(err)
	}
	b.PV[TODAY] = count
	var sessions []string
	err = c.Find(bson.M{"time": bson.M{"$gte": todayBegin, "$lt": todayEnd}}).Distinct("sessionid", &sessions)
	if err != nil {
		log.Error(err)
	}
	b.UV[TODAY] = len(sessions)
	err = c.Find(bson.M{"time": bson.M{"$gte": yestdBegin, "$lt": yestdEnd}}).Distinct("sessionid", &sessions)
	if err != nil {
		log.Error(err)
	}
	b.UV[YESTERDAY] = len(sessions)
	var ips []string
	err = c.Find(bson.M{"time": bson.M{"$gte": yestdBegin, "$lt": yestdEnd}}).Distinct("remoteaddr", &ips)
	if err != nil {
		log.Error(err)
	}
	b.IP[YESTERDAY] = len(ips)
	err = c.Find(bson.M{"time": bson.M{"$gte": todayBegin, "$lt": todayEnd}}).Distinct("remoteaddr", &ips)
	if err != nil {
		log.Error(err)
	}
	b.IP[TODAY] = len(ips)
	var ts []*Request
	err = c.Find(bson.M{"time": bson.M{"$gte": todayBegin, "$lt": todayEnd}}).Select(bson.M{"time": 1}).All(&ts)
	if err != nil {
		log.Error(err)
	}
	b.TimePV[TODAY] = make([]int, 145)
	for _, v := range ts {
		b.TimePV[TODAY][ParseTime(v.Time)]++
	}
	err = c.Find(bson.M{"time": bson.M{"$gte": yestdBegin, "$lt": yestdEnd}}).Select(bson.M{"time": true}).All(&ts)
	if err != nil {
		log.Error(err)
	}
	b.TimePV[YESTERDAY] = make([]int, 145)
	for _, v := range ts {
		b.TimePV[YESTERDAY][ParseTime(v.Time)]++
	}

}

const (
	pageCount = 30
)

func (b *BaseData) LoadLatestRequst(page int) {
	ms, c := db.Connect(DB, C_REQUEST)
	defer ms.Close()
	err := c.Find(nil).Sort("-time").Skip(pageCount * (page - 1)).Limit(pageCount).All(&b.Latest)
	if err != nil {
		log.Error(err)
	}
}

func ParseTime(t time.Time) int { // 第几个十分钟
	return (t.Hour()*60+t.Minute())/10 + 1
}

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
	if err != nil && !strings.Contains(err.Error(), "not found") {
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
	err := db.Update(DB, C_CONFIG, bson.M{"name": "config"}, bson.M{"$set": bson.M{"siteverify": conf.SiteVerify}})
	if err != nil {
		log.Error(err)
	}
}
