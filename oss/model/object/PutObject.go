/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"bytes"
	//"errors"
	//"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/common"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//	Import common.Client.
/*
 *
 */
type Client struct {
	CClient *common.Client
}

// 	Create a new object to a bucket.
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

	_, err = cc.DoRequest("PUT", objectPath, objectPath, params, buffer)
	if err != nil {
		return
	}

	// body, _ := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	err = errors.New(resp.Status)
	// 	fmt.Println(string(body))
	// 	return
	// }
	//fmt.Println("A new object(" + objectPath + ") has been put into this bucket.")
	return
}
