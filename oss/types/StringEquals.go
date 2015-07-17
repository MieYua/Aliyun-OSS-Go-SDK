/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	StringEquals's struct.
//	StringEquals属性。
type StringEquals struct {
	UserAgent       string `json:"acs:UserAgent,omitempty"`       //	指定http useragent头 字符串
	CurrentTime     string `json:"acs:CurrentTime,omitempty"`     //	指定合法的访问时间 ISO8601格式
	SourceTransport string `json:"acs:SecureTransport,omitempty"` // 	是否是https协议 "http"或者"https"
	Prefix          string `json:"oss:Prefix,omitempty"`          //	用作ListObjects时的prefix 合法的object name
	Delimiter       string `json:"oss:Delimiter,omitempty"`       //	用作ListObject时的delimiter 合法的delimiter值
}
