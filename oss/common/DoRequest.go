/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package common

import (
	"Aliyun-OSS-Go-SDK/oss/consts"
	"io"
	"net/http"
	"strconv"
	"time"
)

// 	Deal with some requests.
/*
 *	Example:
 *	resp,err := DoRequest(method, path, canonicalizedResource, params, data)
 *		method: Get，Put，Post...
 *		data: io file
 */
func (c *Client) DoRequest(method, path, canonicalizedResource string, params map[string]string, data io.Reader) (resp *http.Response, err error) {
	reqUrl := "http://" + c.Host + path
	req, _ := http.NewRequest(method, reqUrl, data)
	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	req.Header.Set("Date", date)
	req.Header.Set("Host", c.Host)

	if params != nil {
		for k, v := range params {
			req.Header.Set(k, v)
		}
	}

	if data != nil {
		req.Header.Set(consts.HH_CONTENT_LENGTH, strconv.Itoa(int(req.ContentLength)))
	}

	c.SignHeader(req, canonicalizedResource)

	//fmt.Println(req)
	resp, err = c.HttpClient.Do(req)
	//fmt.Println(resp)
	if method == "POST" {
		resp.Header.Set(consts.HH_AUTHORIZATION, req.Header.Get(consts.HH_AUTHORIZATION))
	}
	return
}