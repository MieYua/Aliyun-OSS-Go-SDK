/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"strings"
)

// 	Complete to upload all the multiparts.
/*
 *	Example:
 *	cmur, err := CompleteMultipartUpload(cmu, initObjectPath, uploadId)
 */
func (c *Client) CompleteMultipartUpload(cmu types.CompleteMultipartUpload, initObjectPath, uploadId string) (cmur types.CompleteMultipartUploadResult, err error) {
	cc := c.CClient

	if strings.HasPrefix(initObjectPath, "/") == false {
		initObjectPath = "/" + initObjectPath
	}

	reqStr := initObjectPath + "?uploadId=" + uploadId

	cxml := types.CompleteMultipartUploadXML{}
	cxml.Part = cmu.Part

	bs, err := xml.Marshal(cxml)
	if err != nil {
		return
	}

	buffer := new(bytes.Buffer)
	buffer.Write(bs)

	resp, err := cc.DoRequest("POST", reqStr, reqStr, nil, buffer)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}
	err = xml.Unmarshal(body, &cmur)
	if err == nil {
		fmt.Println("The object(" + initObjectPath + ") has been uploaded successfully.")
	}
	return
}
