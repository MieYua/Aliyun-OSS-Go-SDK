/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Find the object's head's meta information and show it out
/*
 *	Example:
 *	header, err := c.HeadObject("bucketName/test.txt")
 *	fmt.Println(header, err)
 *	--> map[Accept-Ranges:[bytes] Etag:["xxxxx"] Server:[AliyunOSS] Date:[xxxx GMT] Content-Type:[text/plain; charset=utf-8] Content-Length:[x] Last-Modified:[xxxx GMT] X-Oss-Request-Id:[xxxxxx]] <nil>
 */
func (c *Client) HeadObject(objectPath string) (header http.Header, err error) {
	cc := ConvertClient(c)

	if strings.HasPrefix(objectPath, "/") == false {
		objectPath = "/" + objectPath
	}
	resp, err := cc.DoRequest("HEAD", objectPath, objectPath, nil, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}
	header = resp.Header
	fmt.Println("You have got the header's meta of (" + objectPath + ").")
	return
}
