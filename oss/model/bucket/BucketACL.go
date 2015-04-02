/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package bucket

import (
	"Aliyun-OSS-Go-SDK/oss/consts"
	"Aliyun-OSS-Go-SDK/oss/types"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
)

// 	Change the setting of this bucket' acl.
/*
 *	Example:
 *	err := PutBucketACL(bucketName, (consts)ACL)
 */
func (c *Client) PutBucketACL(bucketName, acl string) (err error) {
	cc := ConvertClient(c)

	params := map[string]string{consts.OH_OSS_CANNED_ACL: acl}
	reqStr := "/" + bucketName
	resp, err := cc.DoRequest("PUT", reqStr, reqStr, params, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	fmt.Println("The ACL's setting of " + bucketName + " has been changed.")
	return
}

// 	Get the setting of this bucket' acl.
/*
 *	Example:
 *	acl,err := GetBucketACL(bucketName)
 */
func (c *Client) GetBucketACL(bucketName string) (acl types.AccessControlPolicy, err error) {
	cc := ConvertClient(c)

	reqStr := "/" + bucketName + "?acl"
	resp, err := cc.DoRequest("GET", reqStr, reqStr, nil, nil)
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

	err = xml.Unmarshal(body, &acl)
	if err == nil {
		fmt.Println("You have got the ACL's setting of " + bucketName + ".")
	}
	return
}
