/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package common

import (
	"Aliyun-OSS-Go-SDK/oss/types"
	"fmt"
	"net/http"
)

type Client types.Client

// 	Create a new client.
/*
 *	Example:
 *	c := NewClient((Const)ENDPOINT, "your oss's accessKeyId", "your oss's accessKeySecret")
 */
func NewClient(endPoint, accessKeyId, accessKeySecret string) *Client {
	client := Client{
		Host:            endPoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		HttpClient:      http.DefaultClient,
	}
	fmt.Println("This client is ready.")
	return &client
}
