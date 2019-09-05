package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// ApplicationJSON application Type 有关Http协议GET和POST请求的封装
const ApplicationJSON = "application/json"

//Get 发送GET请求
//url:请求地址
//response:请求返回的内容
func Get(url string) (response string) {
	log.Println("请求入参为：", url)
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	response = result.String()
	log.Println("返回参数为：", response)

	return
}

//Post 发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求放回的内容
func Post(url string, data interface{}, contentType string) (content string) {
	jsonStr, _ := json.Marshal(data)
	log.Println("请求入参为：", url, string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	log.Println("返回参数为：", content)
	return
}
