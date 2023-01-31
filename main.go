package main

import (
	"fmt"
)

func main() {
	const port = "58080"
	go StartHTTPServer(port)

	endpoint := fmt.Sprintf("http://localhost:%s", port)

	c := NewClient(endpoint, "iamtoken")
	clusterInfo := ClusterInfo{
		ClusterName: "andyscluster",
		ClusterAddr: "1.2.3.4",
	}
	c.CreateClusterInfo(clusterInfo)
}
