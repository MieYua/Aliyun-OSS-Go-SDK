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

// 	Change the setting of this bucket's lifecycle.
//	修改Bucket的生命周期设置。
/*
 *	Example:
 *	rule1 := (types)Rule{
 *		Id:         "delete after one year",
 *		Prefix:     "bucket1logs/",
 *		Status:     "Enabled",
 *		Expiration: Expiration{Days: 365},
 *	}
 *	rule2 := (types)Rule{
 *		Id:         "delete after one month",
 *		Prefix:     "bucket2logs/",
 *		Status:     "Enabled",
 *		Expiration: Expiration{Days: 30},
 *	}
 *	...
 *	rules := [](types)Rule{rule1,rule2,...}
 *	err := PutBucketReferer(bucketName, rules)
 *	If the referers is null, its AllowEmptyReferer'default is true.
 */
func (c *Client) PutBucketLifecycle(bucketName string, rules []types.Rule) (err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?lifecycle"

	lcxml := types.LifecycleXML{}

	for _, v := range rules {
		r := types.Rule{}
		r.Id = v.Id
		r.Prefix = v.Prefix
		r.Status = v.Status
		r.Expiration.Days = v.Expiration.Days
		lcxml.Rule = append(lcxml.Rule, r)
	}

	bs, err := xml.Marshal(lcxml)
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

	//log.Println("The lifecycle's setting of " + bucketName + " has been changed.")
	return
}

// 	Get the details od this bucket's lifecycle.
//	获得Bucket的生命周期设置。
/*
 *	Example:
 *	lfc, err := c.GetBucketLifecycle(bucketName)
 */
func (c *Client) GetBucketLifecycle(bucketName string) (lfc types.LifecycleConfiguration, err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?lifecycle"
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

	err = xml.Unmarshal(body, &lfc)
	if err != nil {
		return
	}

	// log.Println("You have got the lifecycle's setting of " + bucketName + ".")
	return
}
