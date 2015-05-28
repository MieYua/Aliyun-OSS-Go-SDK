/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"encoding/xml"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"log"
	"strings"
)

// 	Give the list of the uploaded multipart upload mission by uploadId.
//	获得对应uploadId的已经上传的multipart任务列表。
/*
 *	Example:
 *	lpr, err := c.ListParts(objectName, uploadId)
 */
func (c *Client) ListParts(objectName, uploadId string) (lpr types.ListPartsResult, err error) {
	cc := c.CClient

	if strings.HasPrefix(objectName, "/") == false {
		objectName = "/" + objectName
	}

	reqStr := objectName + "?uploadId=" + uploadId
	resp, err := cc.DoRequest("GET", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &lpr)
	if err != nil {
		return
	}

	//log.Println("You have got all the uploaded files' details of " + objectName + " by uploadId:" + uploadId + ".")
	return
}
