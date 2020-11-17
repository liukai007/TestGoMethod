package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//定义api接口数据结构和序列化json字段
type Data struct {
	ID    int    `json:"id"`
	IP    string `json:"IP"`
	DESC  string `json:"desc"`
	OWNER string `json:"owner"`
}

type CloumnsData struct {
	NAME  string `json:"name"`
	ALIAS string `json:"alias"`
}

type Employee struct {
	CODE    int           `json:"code"`
	ISADMIN bool          `json:"isadmin"`
	DATA    []Data        `json:"data"`
	COLUMNS []CloumnsData `json:"columns"`
	MESSGAE string        `json:"messgae"`
}

type Employee1 struct {
	BaudRate  int `json:"BaudRate"`
	DataBit   int `json:"DataBit"`
	StopBit   int `json:"StopBit"`
	ParityBit int `json:"ParityBit"`
	StartBit  int `json:"StartBit"`
}

type Employee2 struct {
	Data Employee1 `json:"data"`
}

type Animal struct {
	Name  string
	Order string
	site  string
}

type Protocols struct {
	SerialPort string `json:"serialPort"`
	BaudRate   string `json:"baudRate"`
	DataBits   string `json:"dataBits"`
	StopBits   string `json:"stopBits"`
	Parity     string `json:"parity"`
}

//发送http请求和json 序列化并打印数据结构;

func main() {
	url := "http://192.168.2.10:8989/portType/equipmentEnName/7he1"
	resp, _ := http.Get(url)
	if resp.StatusCode == 200 {
		fmt.Println("200")
	} else if resp.StatusCode == 400 {
		fmt.Println("400")
	}
	s := Protocols{}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	//string(body) 转成字符串
	fmt.Println(string(body))
	// []byte(str) 转成字节组
	fmt.Println([]byte(body))
	//字符串替换
	content := string(body)
	//content=strings.Replace(content,"\\\"","\"",-1)
	//content=strings.Replace(content,"\"{","{",-1)
	//content=strings.Replace(content,"}\"","}",-1)
	fmt.Println(content)
	content=strings.Replace(content,"{\"data\":\"","",-1)
	content=strings.Replace(content,"\"}\"}","\"}",-1)
	content=strings.Replace(content,"\\\"","\"",-1)
	//body1:="{\"MESSGAE\":\"1\"}"
	fmt.Println(content)
	json.Unmarshal([]byte(content), &s)
	fmt.Println(fmt.Sprintf("%+v", s))



}
