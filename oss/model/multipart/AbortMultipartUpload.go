/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

// 	Abort this multipart upload by uploadId.
//	放弃对应uploadId的MultipartUpload上传。
/*
 *	Example:
 *	You can use c.ListMultipartUpload to find all the uploadId
 *	err := c.AbortMultipartUpload(initObjectPath, imur1.UploadId)
 */
func (c *Client) AbortMultipartUpload(objectPath, uploadId string) (err error) {
	cc := c.CClient

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
		log.Println(string(body))
		return
	}

	//log.Println("The " + objectPath + " whose uploadId:" + uploadId + " has been aborted.")
	return
}
