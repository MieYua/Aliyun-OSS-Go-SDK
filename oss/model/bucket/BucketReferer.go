/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package bucket

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"log"
	"net/http"
)

// 	Change the referers of this bucket(default: allow empty referer).
//	修改Bucket的访问白名单（不设置默认为空）。
/*
 *	Example:
 *	err := PutBucketReferer(bucketName, []string{(consts)REFERER,or other "http(s)://*.*.*"'s addresses})
 *	If the referers is null, its AllowEmptyReferer'default is true.
 */
func (c *Client) PutBucketReferer(bucketName string, referers []string) (err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?referer"

	rxml := types.RefererXML{}
	rxml.AllowEmptyReferer = true
	if referers == nil {
		rxml.RefererList = append(rxml.RefererList, types.RefererList{[]string{""}})
		rxml.AllowEmptyReferer = true
	} else {
		for _, v := range referers {
			rxml.RefererList = append(rxml.RefererList, types.RefererList{[]string{v}})
			rxml.AllowEmptyReferer = false
		}
	}

	bs, err := xml.Marshal(rxml)
	if err != nil {
		return
	}
	buffer := new(bytes.Buffer)
	buffer.Write(bs)

	contentType := http.DetectContentType(buffer.Bytes())
	params := map[string]string{}
	params[consts.HH_CONTENT_TYPE] = contentType

	resp, err := cc.DoRequest("PUT", reqStr, reqStr, params, buffer)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		log.Println(string(body))
		return
	}

	//log.Println("The referer's setting of " + bucketName + " has been changed.")
	return
}

//	Get the referers' addresses of this bucket.
//	获得Bucket的白名单地址。
/*
 *	Example:
 *	rc, err := c.GetBucektReferer(bucketName)
 */
func (c *Client) GetBucketReferer(bucketName string) (rc types.RefererConfiguration, err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?referer"
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

	err = xml.Unmarshal(body, &rc)
	if err != nil {
		return
	}

	// log.Println("You have got the referer's setting of " + bucketName + ".")
	return
}
