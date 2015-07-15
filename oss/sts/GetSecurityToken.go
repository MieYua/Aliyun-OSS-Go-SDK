/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package sts

import (
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 	Deal with some requests.
//	处理请求，返回响应结果。
/*
 *	Example:
 *	securityToken, tempAccessKeyId, tempAccessKeySecret, err := GetSecurityToken(method, path, canonicalizedResource, params, data)
 *		durationSeconds: mainAccount:900-3600s/childAccount:900-129600s
 *
 */
func GetSecurityToken(durationSeconds int) (securityToken, tempAccessKeyId, tempAccessKeySecret string, err error) {
	reqUrl := "http://sts.aliyuncs.com"

	req, _ := http.NewRequest("POST", reqUrl, data)
	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	req.Header.Set("Date", date)
	req.Header.Set("Host", c.TClient.Host)

	if params != nil {
		for k, v := range params {
			req.Header.Set(k, v)
		}
	}

	if data != nil {
		req.Header.Set(consts.HH_CONTENT_LENGTH, strconv.Itoa(int(req.ContentLength)))
	}

	c.SignHeader(req, canonicalizedResource)
	resp, err = c.TClient.HttpClient.Do(req)

	if method == "POST" {
		resp.Header.Set(consts.HH_AUTHORIZATION, req.Header.Get(consts.HH_AUTHORIZATION))
	}
	return
}
