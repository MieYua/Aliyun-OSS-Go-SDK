/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// 	Delete an object.
/*
 * 	Example:
 * 	err := c.DeleteObject(objectPath)
 *	objectPath: bucketName/objectName
 *	Warning:
 *	If you want to delete a filepack, you need clear all files in this filepack and than delete this filepack.
 *		c.PutObject("bucketName/test/test.txt")
 *		The wrong way:
 *		c.DeleteObject("bucketName/test/") can't delete this filepack
 *		The right way:
 *		c.DeleteObject("bucketName/test/test.txt")
 *		c.DeleteObject("bucketName/test/")
 */
func (c *Client) DeleteObject(objectPath string) (err error) {
	cc := ConvertClient(c)

	if strings.HasPrefix(objectPath, "/") == false {
		objectPath = "/" + objectPath
	}
	resp, err := cc.DoRequest("DELETE", objectPath, objectPath, nil, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 204 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	fmt.Println("The (" + objectPath + ") has been deleted.")
	return
}
