/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"net/http"
	"sync"
)

//	Client's struct.
//	客户端属性。
type Client struct {
	AccessKeyId     string       // 访问Id
	AccessKeySecret string       // 访问密钥
	Host            string       // 主机节点
	HttpClient      *http.Client // http客户端
	FileIOLocker    sync.Mutex   // 排它锁
	UserProperty    string       //	用户性质（默认Owner，STS临时用户为TempUser）
	SecurityToken   string       //	当用户为TempUser时使用
	TempPrefix      string       //	临时prefix
	TempDelimiter   string       //	临时delimiter
}

type ConvertClient interface {
	PutBucket(string) error
}
