/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"bytes"
	//"errors"
	//"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"io"
	//"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

// 	Post up an object to replace putObject.
/*
 *	Example:
 *	err := c.PostObject(bucketName, fileName, tempFileName)
 */
func (c *Client) PostObject(bucketName, filePath string, tempFileName string) (err error) {
	cc := c.CClient

	if strings.HasPrefix(bucketName, "/") == false {
		bucketName = "/" + bucketName
	}
	buffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buffer)

	bodyWriter.CreateFormField("key")
	buffer.WriteString(filePath)
	bodyWriter.CreateFormField("success_action_status")
	buffer.WriteString("200")
	bodyWriter.CreateFormField("Content-Disposition")
	buffer.WriteString("content-disposition")
	/*
	 *	Can expand next step, for those buckets whose acl is private and public-read
	 *	//bodyWriter.CreateFormField("x-oss-neta-uuid")
	 *	//bodyWriter.CreateFormField("x-oss-meta-tag")
	 *	bodyWriter.CreateFormField("OSSAccessKeyId")
	 *	bodyWriter.CreateFormField("policy")
	 *	bodyWriter.CreateFormField("Signature")
	 *	//bodyWriter.CreateFormField("submit")
	 */
	fileWriter, err := bodyWriter.CreateFormFile("file", tempFileName)
	if err != nil {
		return
	}

	fh, err := os.Open(tempFileName)
	if err != nil {
		return
	}
	io.Copy(fileWriter, fh)
	defer fh.Close()

	params := map[string]string{}
	params[consts.HH_CONTENT_TYPE] = "multipart/form-data; boundary=" + bodyWriter.Boundary()
	bodyWriter.Close()

	_, err = cc.DoRequest("POST", bucketName, bucketName, params, buffer)
	if err != nil {
		return
	}

	return
}
