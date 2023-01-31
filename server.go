package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	rnd := rand.Int()
	s := fmt.Sprintf("hello %v", rnd)
	w.Write([]byte(s))
}

func clusterInfo(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body error in server, err: %v", err)
	}
	fmt.Printf("get body: %v\n", string(body))

	authorization := req.Header.Get(Authorization)
	fmt.Printf("get authorization: %v\n", authorization)
	w.Write([]byte("success"))
}

func StartHTTPServer(port string) {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: nil,
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc(ClusterInfoPath, clusterInfo)
	server.ListenAndServe()
	//	http.ListenAndServe(":58080", nil)
}
