/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// 	Get an object.
//	获得一个Object的信息。
/*
 *	obytes, err := c.GetObject(objectPath, rangeStart(form 0), rangeEnd(from 0 and it is larger than rangeStart))
 *	If you want the whole file,
 *			rangeStart:	default 	-1,
 *			rangeEnd:	default 	-1.
 *	Example:
 *	obytes, err := c.GetObject("xxxx/test.txt", -1, -1)
 *	fmt.Println(string(obytes[:]), err)-->test <nil>
 */
func (c *Client) GetObject(objectPath string, rangeStart, rangeEnd int) (obytes []byte, err error) {
	cc := c.CClient

	if strings.HasPrefix(objectPath, "/") == false {
		objectPath = "/" + objectPath
	}

	params := map[string]string{}
	if rangeStart > -1 && rangeEnd > -1 {
		params[consts.HH_RANGE] = "bytes=" + strconv.Itoa(rangeStart) + "-" + strconv.Itoa(rangeEnd)
	}

	resp, err := cc.DoRequest("GET", objectPath, objectPath, params, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 206 {
		err = errors.New(resp.Status)
		log.Println(string(body))
		return
	} else {
		obytes = body
	}

	//log.Println("You have got the details of this object(" + objectPath + ").")
	return
}
