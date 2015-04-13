/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
)

// 	Delete some objects at one time.
/*
 *	Example:
 *	c.CopyObject("bucketName/copy_test1.txt", "bucketName/test.txt")
 *	c.CopyObject("bucketName/copy_test2.txt", "bucketName/test.txt")
 *	c.DeleteMultipleObject("bucketName", []string{"copy_test1.txt", "copy_test2.txt"})
 */
func (c *Client) DeleteMultipleObject(bucketName string, keys []string) (err error) {
	cc := c.CClient

	dxml := types.DeleteXML{}

	for _, v := range keys {
		dxml.Object = append(dxml.Object, types.DeleteObject{v})
	}
	dxml.Quiet = true

	bs, err := xml.Marshal(dxml)
	if err != nil {
		return
	}

	reqStr := "/" + bucketName + "?delete"
	buffer := new(bytes.Buffer)
	buffer.Write(bs)

	h := md5.New()
	h.Write(bs)
	md5sum := base64.StdEncoding.EncodeToString(h.Sum(nil))
	params := map[string]string{}
	params[consts.HH_CONTENT_MD5] = md5sum

	resp, err := cc.DoRequest("POST", reqStr, reqStr, params, buffer)
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
	length := len(keys)
	list := keys[0]
	for i := 1; i < length; i++ {
		list += ", " + keys[i]
	}
	fmt.Println("The (" + list + ") of " + bucketName + " have been deleted.")
	return
}
