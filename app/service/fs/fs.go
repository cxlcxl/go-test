package fs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"goskeleton/app/global/variable"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Fs struct {
	appId              string
	appSecret          string
	departmentId       string
	accessTokenUrl     string
	departmentUsersUrl string
	batchSendMsgUrl    string
	batchSendCardUrl   string
	token              string
}

type TokenRet struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

type simpleRet struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// CreateFsFactory 取配置信息映射到结构体
func CreateFsFactory() *Fs {
	return &Fs{
		appId:              variable.ConfigYml.GetString("Feishu.appId"),
		appSecret:          variable.ConfigYml.GetString("Feishu.appSecret"),
		departmentId:       variable.ConfigYml.GetString("Feishu.departmentId"),
		accessTokenUrl:     variable.ConfigYml.GetString("Feishu.accessTokenUrl"),
		departmentUsersUrl: variable.ConfigYml.GetString("Feishu.departmentUsersUrl"),
		batchSendMsgUrl:    variable.ConfigYml.GetString("Feishu.batchSendMsgUrl"),
		batchSendCardUrl:   variable.ConfigYml.GetString("Feishu.batchSendCardUrl"),
	}
}

// AccessToken 获取 token
func (f *Fs) AccessToken() *Fs {
	var body = map[string]interface{}{
		"app_id":     f.appId,
		"app_secret": f.appSecret,
	}
	var params = make(map[string]string, 0)
	res, _ := f.request(f.accessTokenUrl, body, params)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var token = TokenRet{}
	json.Unmarshal(data, &token)
	if token.Code == 200 {
		f.token = token.TenantAccessToken
	}
	return f
}

func (f *Fs) SendMsg(userIds, msg string) error {
	fmt.Println(f)
	var body = map[string]interface{}{
		"user_ids": strings.Split(userIds, ","),
		"msg_type": "text",
		"content": map[string]string{
			"text": msg,
		},
	}
	var params = make(map[string]string, 0)
	ret, _ := f.request(f.batchSendMsgUrl, body, params)
	retBody, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		fmt.Println(err)
	}
	var s simpleRet
	json.Unmarshal(retBody, &s)
	if s.Code != 0 {
		return errors.New(s.Msg)
	}
	return nil
}

func (f *Fs) request(url string, body map[string]interface{}, params map[string]string) (*http.Response, error) {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	log.Println(f.token)
	if f.token != "" {
		req.Header.Set("Authorization", "Bearer "+f.token)
	}
	//http client
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())
	return client.Do(req)
}
