package cron

import (
	"github.com/astaxie/beego/httplib"
	"crypto/tls"
	"github.com/liudng/godump"
	"log"
	"time"
	"encoding/json"
	"github.com/astaxie/beego/toolbox"
)

const (
	CorpID = "ww5436abd0c7595f3e"
	AgentId = 1000002
	Secret = "-JlVjTKC4AvEOXNwAeVjxNgYSG5jZKzm74H33qN03zE"
	Expires = 7200
)

type TokenRes struct {
	Errcode int
	Errmsg string
	Access_token string
	Expires_in int
}
type Token struct {
	token string
	start_time int64
}
type NewMsg struct {
	Agentid int    `json:"agentid"`
	Msgtype string `json:"msgtype"`
	Safe    int    `json:"safe"`
	Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Touser  string `json:"touser"`
}


const (
	GetTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	SendMsgURL = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
)
var token Token
func mainSend(){
	weather := Weather{}
	err :=weather.getWeather()
	if err != nil {
		log.Print(err)
		return
	}
	tianqi_day := weather.Results[0].Daily[0].TextDay
	tianqi_night := weather.Results[0].Daily[0].TextNight

	SendMsg("报告队长,天气预报说明天白天 ["+tianqi_day+"] 晚上 ["+tianqi_night+"]")
}


func (token *Token)GetToken() (string,error) {
	if token == nil || time.Now().Unix() - token.start_time > Expires {
		var result TokenRes
		//var now = time.Now().Unix()
		req := httplib.Get(GetTokenURL+"?corpid="+CorpID+"&corpsecret="+Secret)
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		if err := req.ToJSON(&result); err != nil{
			log.Print(err)
			return "",err
		}
		if result.Errcode != 0 {
			log.Print(result.Errmsg)
		}
		token.token  = result.Access_token
		token.start_time = time.Now().Unix()
	}
	return token.token,nil
}
func SendMsg(msg string ){
	Msg := NewMsg{Agentid:AgentId,Msgtype:"text",Safe:0,
		Text:struct {
			Content string `json:"content"`
		}{ Content: msg },
		Toparty:"",
		Totag:"",
		Touser:"@all",
	}
	msgjson,err  := json.Marshal(Msg)
	if err != nil{
		log.Fatal(err)
	}
	mytoken ,err := token.GetToken()
	if err != nil {
		log.Print(err)
		return
	}
	req := httplib.Post(SendMsgURL+"?access_token="+ mytoken)
	req.Body(msgjson)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	str, err := req.String()
	if err != nil {
		log.Print(err)
		return
	}
	godump.Dump(str)
}

type Weather struct {
	Results []struct {
		Daily []struct {
			CodeDay             string `json:"code_day"`
			CodeNight           string `json:"code_night"`
			Date                string `json:"date"`
			High                string `json:"high"`
			Low                 string `json:"low"`
			Precip              string `json:"precip"`
			TextDay             string `json:"text_day"`
			TextNight           string `json:"text_night"`
			WindDirection       string `json:"wind_direction"`
			WindDirectionDegree string `json:"wind_direction_degree"`
			WindScale           string `json:"wind_scale"`
			WindSpeed           string `json:"wind_speed"`
		} `json:"daily"`
		LastUpdate string `json:"last_update"`
		Location   struct {
			      Country        string `json:"country"`
			      ID             string `json:"id"`
			      Name           string `json:"name"`
			      Path           string `json:"path"`
			      Timezone       string `json:"timezone"`
			      TimezoneOffset string `json:"timezone_offset"`
		      } `json:"location"`
	} `json:"results"`
}
const (
	WeatherURL="https://api.seniverse.com/v3/weather/daily.json?key=wzohy6883cphimee&location=ningbo&language=zh-Hans&unit=c&start=1&days=5"
)


func (weather *Weather)getWeather() error{
	req := httplib.Get(WeatherURL)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if err := req.ToJSON(weather); err != nil{
		//log.Print(err)
		return err
	}
	return nil
}
func init(){

	tk1 := toolbox.NewTask("tk1","0 0 21 * *",
		func() error{
			mainSend()
			return nil
		})

	toolbox.AddTask("tk1",tk1)
	toolbox.StartTask()
	//defer  toolbox.StopTask()
}
