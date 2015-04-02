/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"Aliyun-OSS-Go-SDK/oss/consts"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

// 	Post up an object to replace putObject.
/*
 *	Example:
 *	err := c.PostObject(bucketName, filePath)
 */
func (c *Client) PostObject(bucketName, filePath string) (err error) {
	cc := ConvertClient(c)

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
	fileWriter, err := bodyWriter.CreateFormFile("file", filePath)
	if err != nil {
		return
	}

	fh, err := os.Open(filePath)
	if err != nil {
		return
	}

	defer fh.Close()
	io.Copy(fileWriter, fh)

	params := map[string]string{}
	params[consts.HH_CONTENT_TYPE] = "multipart/form-data; boundary=" + bodyWriter.Boundary()
	bodyWriter.Close()

	resp, err := cc.DoRequest("POST", bucketName, bucketName, params, buffer)
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
	fmt.Println("The object(" + bucketName + ") has been posted up.")
	return
}
