/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"encoding/xml"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/common"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"log"
	"strings"
)

//	Convert common.Client to Client.
//	将common包的Client转换成Client类。
type Client struct {
	CClient *common.Client
}

// 	Start the multipartUpload and get the uploadId.
//	开始MultipartUpload上传，并获得uploadId。
/*
 *	Example:
 *	initObjectPath, imur, err := c.InitiateMultipartUpload("bucketName/test.txt")
 */
func (c *Client) InitiateMultipartUpload(objectPath string) (initObjectPath string, imur types.InitiateMultipartUploadResult, err error) {
	cc := c.CClient

	if strings.HasPrefix(objectPath, "/") == false {
		objectPath = "/" + objectPath
	}
	initObjectPath = objectPath
	resp, err := cc.DoRequest("POST", objectPath+"?uploads", objectPath+"?uploads", nil, nil)
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

	err = xml.Unmarshal(body, &imur)
	if err != nil {
		return
	}

	//log.Println("The multipart upload has been intiated and you have got the UploadId.")
	return
}
