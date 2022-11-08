/**
* @file: network.go ==> common/network
* @package: network
* @author: jingxiu
* @since: 2022/11/8
* @desc: 网络请求库
 */

package network

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 5 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 10 * time.Minute,
			}).DialContext,
		},
	}
}

// 发送GET请求
func HttpGet(url string) ([]byte, error) {
	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}

// 发送POST请求(JSON)
func HttpPostJson(url string, body interface{}) ([]byte, error) {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return httpPost(url, "application/json", string(bodyStr))
}

// 发送POST请求(XML)
func HttpPostXml(url string, xmlBody string) ([]byte, error) {
	return httpPost(url, "application/xml", xmlBody)
}

// 发送通用的POST请求
func httpPost(url string, contentType string, body string) ([]byte, error) {
	rsp, err := client.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}

func PostHeader(url string, msg []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(msg)))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
