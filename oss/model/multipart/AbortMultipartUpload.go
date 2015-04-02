/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// 	Abort this multipart upload by uploadId.
/*
 *	Example:
 *	You can use c.ListMultipartUpload to find all the uploadId
 *	err := c.AbortMultipartUpload(initObjectPath, imur1.UploadId)
 */
func (c *Client) AbortMultipartUpload(objectPath, uploadId string) (err error) {
	cc := ConvertClient(c)

	if strings.HasPrefix(objectPath, "/") == false {
		objectPath = "/" + objectPath
	}

	reqStr := objectPath + "?uploadId=" + uploadId
	resp, err := cc.DoRequest("DELETE", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 204 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	fmt.Println("The " + objectPath + " whose uploadId:" + uploadId + " has been aborted.")
	return
}
