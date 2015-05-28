/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package bucket

import (
	"encoding/xml"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"log"
)

// 	Get the details of this bucket with some parameters.
//	获得Bucket下筛选后所有的object的列表。
/*
 *	Example:
 *	lbr(ObjectList), err := c.GetBucket(bucketName, prefix, marker, delimiter, maxkeys)
 *	prefix: Choose that contain this string (default:"")
 *	marker: Return after this letter (default:"")
 *	delimiter: Common Prefixes (default:"")
 *	maxkeys: The maximum of objects (default:"100")
 */
func (c *Client) GetBucket(bucketName, prefix, marker, delimiter, maxkeys string) (lbr types.ListBucketResult, err error) {
	cc := c.CClient

	reqStr := "/" + bucketName
	resStr := reqStr
	query := map[string]string{}
	if prefix != "" {
		query["prefix"] = prefix
	}

	if marker != "" {
		query["marker"] = marker
	}

	if delimiter != "" {
		query["delimiter"] = delimiter
	}

	if maxkeys != "" {
		query["max-keys"] = maxkeys
	}

	if len(query) > 0 {
		reqStr += "?"
		for k, v := range query {
			reqStr += k + "=" + v + "&"
		}
	}

	resp, err := cc.DoRequest("GET", reqStr, resStr, nil, nil)
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

	err = xml.Unmarshal(body, &lbr)
	if err != nil {
		return
	}

	//log.Println("You have got all the objects' settings of " + bucketName + ".")
	return
}
