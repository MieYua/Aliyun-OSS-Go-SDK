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
}

type ConvertClient interface {
	PutBucket(string) error
}
