package cron

import (
	"github.com/astaxie/beego/httplib"
	"crypto/tls"
	"github.com/liudng/godump"
	"log"
	"time"
	"encoding/json"
	"github.com/astaxie/beego/toolbox"
	"strconv"
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
var Scale map[int]string =  map[int]string{
	0:"[无风],地面现象 [静，烟直上]",
	1:"[软风],地面现象 [烟示风向]",
	2:"[轻风],地面现象 [感觉有风]",
	3:"[微风],地面现象 [旌旗展开]",
	4:"[和风],地面现象 [吹起尘土]",
	5:"[清风],地面现象 [小树摇摆]",
	6:"[强风],地面现象 [电线有声]",
	7:"[疾风],地面现象 [步行困难]",
	8:"[大风],地面现象 [折毁树枝]",
	9:"[烈风],地面现象 [小损房屋]",
	10:"[狂风],地面现象 [拔起树木]",
	11:"[暴风],地面现象 [损毁重大]",
	12:"[台风],地面现象 [摧毁极大]",
	13:"[一级飓风]",
	14:"[二级飓风]",
	15:"[三级飓风]",
	16:"[超强台风]",
	17:"[四级飓风]",
	18:"[吊炸天风]",
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
	low ,err := strconv.Atoi(weather.Results[0].Daily[0].Low)
	if err != nil {
		return
	}
	high ,err := strconv.Atoi(weather.Results[0].Daily[0].High)
	if err != nil {
		return
	}
	direction := weather.Results[0].Daily[0].WindDirection
	scale,err := strconv.Atoi(weather.Results[0].Daily[0].WindScale)
	if err != nil {
		return
	}
	SendMsg("报告队长,天气预报说明天白天 ["+tianqi_day+"] 晚上 ["+tianqi_night+"],气温 "+weather.Results[0].Daily[0].Low+"~"+weather.Results[0].Daily[0].High+"℃,平均气温"+strconv.Itoa((low+high)/2)+"℃,"+direction+"风 ,"+weather.Results[0].Daily[0].WindScale+"级,属于 "+Scale[scale]+"。" )
}


func (token *Token)GetToken() (string,error) {
	if token == nil || time.Now().Unix() - token.start_time > Expires {
		var result TokenRes
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
		Touser:"YuHaoFeng",
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
