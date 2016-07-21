package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/models"
)

type cache struct {
	BackgroundLeftBar  map[string]string
	BackgroundLeftBars []*models.Leftbar
	BuildVersion       string
	Icons              map[string]*Icon
}

var Cache = NewCache()

func NewCache() *cache {
	return &cache{BackgroundLeftBar: make(map[string]string), Icons: make(map[string]*Icon, 500)}
}

func init() {
	doReadBackLeftBarConfig()
	doReadBuildVersionConfig()
	go timer()
}

func timer() {
	cleanIcons()
	time.AfterFunc(time.Hour*12, timer)
}

var path, _ = os.Getwd()

func doReadBackLeftBarConfig() {
	b, err := ioutil.ReadFile(path + "/conf/backleft.conf")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &Cache.BackgroundLeftBars)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range Cache.BackgroundLeftBars {
		if v.ID != "" {
			Cache.BackgroundLeftBar[v.ID] = v.ID
		}
	}
}

func doReadBuildVersionConfig() {
	b, err := ioutil.ReadFile(path + "/version")
	if err != nil {
		log.Error(err)
	}
	Cache.BuildVersion = string(b)
}

type Icon struct {
	Data []byte
	Time time.Time
}

func cleanIcons() {
	for k, v := range Cache.Icons {
		if v.Time.Before(time.Now().AddDate(0, 0, -2)) {
			Cache.Icons[k] = nil
			delete(Cache.Icons, k)
		}
	}
}
