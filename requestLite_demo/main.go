package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

// Request请求工具类
type LiteRequest struct {
	clients    http.Client
	urls       string
	methods    string
	headers    map[string]string
	jsondata   []byte
	formdata   string
	maxretries int
	returndata map[string]interface{}
	usehttp2   bool
	jar        cookiejar.Jar
}

func (receiver *LiteRequest) Session() *LiteRequest {
	// 创建一个 CookieJar 来存储 cookies
	jar, _ := cookiejar.New(nil)
	receiver.clients.Jar = jar
	return receiver
}

func (receiver *LiteRequest) AddHeader(headers map[string]string) *LiteRequest {
	receiver.headers = headers
	// 添加请求头
	return receiver
}

func (receiver *LiteRequest) AddProxy(proxy string) *LiteRequest {
	// 设置请求代理
	proxyurl, _ := url.Parse(proxy)
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyurl),
	}
	receiver.clients.Transport = transport
	return receiver
}

func (receiver *LiteRequest) AddTimeout(timeout time.Duration) *LiteRequest {
	receiver.clients.Timeout = timeout * time.Second
	return receiver
}

func (receiver *LiteRequest) SetRequestParams(urls string, methods string, jsondata []byte, formdata map[string]string) *LiteRequest {
	receiver.urls = urls
	switch methods {
	case "POST":
		fmt.Println("this is a POST request")
	case "GET":
		fmt.Println("this is a GET request")
	case "DELETE":
		fmt.Println("this is a DELETE request")
	case "PUT":
		fmt.Println("this is a PUT request")
	default:
		panic("error methods")
	}
	receiver.methods = methods
	receiver.jsondata = jsondata
	var r http.Request
	err := r.ParseForm()
	if err != nil {
		return nil
	}
	for key, value := range formdata {
		r.Form.Add(key, value)
	}
	bodystr := strings.TrimSpace(r.Form.Encode())
	receiver.formdata = bodystr
	return receiver
}

func (receiver *LiteRequest) AddRetryNum(maxretries int) *LiteRequest {
	// 设置最大重试次数
	receiver.maxretries = maxretries
	return receiver
}

func (receiver *LiteRequest) SendGetRequest() (map[string]interface{}, error) {
	var maxretries int
	if receiver.maxretries == 0 {
		maxretries = 1
	} else {
		maxretries = receiver.maxretries
	}
	for i := 0; i < maxretries; i++ {
		req, err := http.NewRequest(receiver.methods, receiver.urls, nil)
		if err != nil {
			panic(err)
		}
		for key, value := range receiver.headers {
			req.Header.Set(key, value)
		}
		response, err := receiver.clients.Do(req)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)
		if err == nil && response.StatusCode == http.StatusOK {
			returndata := make(map[string]interface{})
			fmt.Println(response.StatusCode)
			returndata["StatusCode"] = http.StatusOK
			body, _ := io.ReadAll(response.Body)
			returndata["Body"] = string(body)
			receiver.returndata = returndata
			return receiver.returndata, nil
		}
		fmt.Println("重试次数:", i+1)
		time.Sleep(3 * time.Second)
	}

	return receiver.returndata, errors.New("request failed after multiple attempts")
}

func (receiver *LiteRequest) SendJsonRequest() (map[string]interface{}, error) {
	var maxretries int
	if receiver.maxretries == 0 {
		maxretries = 1
	} else {
		maxretries = receiver.maxretries
	}
	for i := 0; i < maxretries; i++ {
		req, err := http.NewRequest(receiver.methods, receiver.urls, bytes.NewBuffer(receiver.jsondata))
		if err != nil {
			panic(err)
		}
		for key, value := range receiver.headers {
			req.Header.Set(key, value)
		}
		response, err := receiver.clients.Do(req)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)
		if err == nil && response.StatusCode == http.StatusOK {
			receiver.returndata["StatusCode"] = response.StatusCode
			body, _ := io.ReadAll(response.Body)
			receiver.returndata["Body"] = string(body)
			return receiver.returndata, nil
		}
		time.Sleep(3 * time.Second)
	}

	return receiver.returndata, errors.New("request failed after multiple attempts")
}

func (receiver *LiteRequest) SendFormRequest() (map[string]interface{}, error) {
	var maxretries int
	if receiver.maxretries == 0 {
		maxretries = 1
	} else {
		maxretries = receiver.maxretries
	}
	for i := 0; i < maxretries; i++ {
		req, err := http.NewRequest(receiver.methods, receiver.urls, strings.NewReader(receiver.formdata))
		if err != nil {
			panic(err)
		}
		for key, value := range receiver.headers {
			req.Header.Set(key, value)
		}
		response, err := receiver.clients.Do(req)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)
		if err == nil && response.StatusCode == http.StatusOK {
			receiver.returndata["StatusCode"] = response.StatusCode
			body, _ := io.ReadAll(response.Body)
			receiver.returndata["Body"] = string(body)
			return receiver.returndata, nil
		}
		time.Sleep(3 * time.Second)
	}

	return receiver.returndata, errors.New("request failed after multiple attempts")
}

func (receiver *LiteRequest) MapToJson(result map[string]interface{}) (bool, string) {
	jsonStringString, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error marshalling mapStringString:", err)
	}
	isJson := json.Valid([]byte(fmt.Sprintf("%v", string(jsonStringString))))
	return isJson, string(jsonStringString)
}

func main() {
	var req LiteRequest
	data, err := req.Session().AddHeader(map[string]string{"Content-Type": "application/json"}).AddRetryNum(3).SetRequestParams("https://api.gumengya.com/Api/Mobile?format=json&tel=13888888888", "GET", nil, nil).SendGetRequest()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
		// 将返回值转换成json
		fmt.Println(req.MapToJson(data))
	}
}
