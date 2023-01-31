package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient http.Client
	Endpoint   string
	Token      string
}

type ClusterInfo struct {
	ClusterName string `json: "cluster_name`
	ClusterAddr string `json: "cluster_addrrrr"`
}

func NewClient(endpoint, token string) Client {
	return Client{
		Endpoint:   endpoint,
		Token:      token,
		HttpClient: http.Client{},
	}
}

func (c Client) NewRequest(method, path string, body interface{}) *http.Request {
	bodyByte, _ := json.Marshal(body)
	url := fmt.Sprintf("%s%s", c.Endpoint, path)
	req, err := http.NewRequest(method, url, bytes.NewReader(bodyByte))
	if err != nil {
		fmt.Printf("http.NewRequest error: %v", err)
	}
	fmt.Printf("req.URL: %v", req.URL)
	req.Header.Set(Authorization, c.Token)
	return req
}

func (c Client) CreateClusterInfo(clusterInfo ClusterInfo) {
	req := c.NewRequest("POST", ClusterInfoPath, clusterInfo)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		fmt.Printf("do request error: %v", err)
	}
	defer res.Body.Close()
	resByte, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("res: %v\n", string(resByte))
}
