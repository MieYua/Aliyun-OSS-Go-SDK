/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"bytes"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/common"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//	Convert common.Client to Client.
//	将common包的Client转换成Client类。
type Client struct {
	CClient *common.Client
}

// 	Create a new object to a bucket.
//	新建一个Object。
/*
 *	Example:
 *	err := c.PutObject(objectPath, filePath)
 *	objectPath:
 *			Can be just a name of file(bucketName/fileName),
 *			Can be names of filepacks(bucketName/filepack/../file).
 */
func (c *Client) PutObject(objectPath, filePath string) (err error) {
	cc := c.CClient

	if strings.HasPrefix(objectPath, "/") == false {
		objectPath = "/" + objectPath
	}
	buffer := new(bytes.Buffer)

	fh, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer fh.Close()
	io.Copy(buffer, fh)

	contentType := http.DetectContentType(buffer.Bytes())
	params := map[string]string{}
	params[consts.HH_CONTENT_TYPE] = contentType

	resp, err := cc.DoRequest("PUT", objectPath, objectPath, params, buffer)
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

	//log.Println("A new object(" + objectPath + ") has been put into this bucket.")
	return
}
