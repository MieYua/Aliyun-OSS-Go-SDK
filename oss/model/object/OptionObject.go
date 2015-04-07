/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"errors"
	"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"io/ioutil"
)

// 	Choose the object by options.
/*
 *	Example:
 *	err := c.OptionObject("bucketName/test.txt", "", "")
 *	fmt.Println(err)
 *
 *	Warning:
 *	If the bucket's cors is not available or its cors hasn't been set up,
 *	response will show 403 ERROR.
 */
func (c *Client) OptionObject(opath, accessControlRequestMethod, accessControlRequestHeader, origin string) (err error) {
	cc := ConvertClient(c)

	reqStr := "/" + opath

	params := map[string]string{consts.HH_CONTENT_TYPE: "application/xml"}
	if accessControlRequestMethod != "" {
		params[consts.OH_ACCESS_CONTROL_REQUEST_METHOD] = accessControlRequestMethod
	} else {
		params[consts.OH_ACCESS_CONTROL_REQUEST_METHOD] = "PUT"
	}
	if accessControlRequestHeader != "" {
		params[consts.OH_ACCESS_CONTROL_REQUEST_HEADER] = accessControlRequestHeader
	} else {
		params[consts.OH_ACCESS_CONTROL_REQUEST_HEADER] = "x-oss-test"
	}
	if origin != "" {
		params[consts.OH_ORIGIN] = origin
	} else {
		params[consts.OH_ORIGIN] = "http://www.example.com"
	}
	fmt.Println(params)
	resp, err := cc.DoRequest("OPTIONS", reqStr, reqStr, params, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
		return
	}
	fmt.Println("CORS's request has passed by options.")
	return
}
