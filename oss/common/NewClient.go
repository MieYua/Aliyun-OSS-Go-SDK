/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package common

import (
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"net/http"
)

//	Convert types.Client to Client.
//	将types包的Client转换成Client类。
type Client struct {
	TClient *types.Client
}

// 	Create a new client.
//	新建客户端。
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
	c := Client{}
	c.TClient = client
	c.TClient.UserProperty = "Owner"
	return &c
}
