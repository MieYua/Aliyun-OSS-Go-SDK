/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package service

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
)

// 	Get service's details.
/*
 *	Example:
 *	lambr,err := c.GetService()
 *		lambr:	{
 *					Owner{ID,DisplayName}
 *					[]Bucket{Location,Name,CreationDate}
 *				}
 */
func (c *Client) GetService() (lambr types.ListAllMyBucketsResult, err error) {
	cc := ConvertClient(c)
	resp, err := cc.DoRequest("GET", "/", "/", nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		e := types.Error{}
		xml.Unmarshal(body, &e)
		err = errors.New(resp.Status + " - " + e.Code)
		//fmt.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &lambr)
	if err == nil {
		fmt.Println("You have got this service's details.")
	}
	return
}
