/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package common

import (
	"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"net/http"
)

type Client struct {
	TClient *types.Client
}

// 	Create a new client.
/*
 *	Example:
 *	c := NewClient((Const)ENDPOINT, "your oss's accessKeyId", "your oss's accessKeySecret")
 */
func NewClient(endPoint, accessKeyId, accessKeySecret string) *Client {
	client := &types.Client{
		Host:            endPoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		HttpClient:      http.DefaultClient,
	}
	fmt.Println("This client is ready.")
	c := Client{}
	c.TClient = client
	return &c
}
