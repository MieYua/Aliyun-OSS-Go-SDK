/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package bucket

import (
	"encoding/xml"
	//"errors"
	"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
)

// 	Get the location of this bucket's endpoint.
/*
 *	Example:
 *	lc,err := GetBucketLocation(bucketName)
 */
func (c *Client) GetBucketLocation(bucketName string) (lc types.LocationConstraint, err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?location"
	resp, err := cc.DoRequest("GET", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	err = errors.New(resp.Status)
	// 	fmt.Println(string(body))
	// 	return
	// }

	err = xml.Unmarshal(body, &lc)
	if err == nil {
		fmt.Println("You have got the region's location of " + bucketName + ".")
	}
	return
}
