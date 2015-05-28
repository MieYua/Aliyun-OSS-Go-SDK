/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package service

import (
	"encoding/xml"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/common"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"log"
)

//	Convert common.Client to Client.
//	将common包的Client转换成Client类。
type Client struct {
	CClient *common.Client
}

// 	Get service's details.
//	获得所有Bukcets的信息。
/*
 *	Example:
 *	lambr, err := c.GetService()
 *		lambr:	{
 *					Owner{ID,DisplayName}
 *					[]Bucket{Location,Name,CreationDate}
 *				}
 */
func (c *Client) GetService() (lambr types.ListAllMyBucketsResult, err error) {
	cc := c.CClient
	resp, err := cc.DoRequest("GET", "/", "/", nil, nil)
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

	err = xml.Unmarshal(body, &lambr)

	if err == nil {
		return
	}

	// log.Println("You have got this service's details.")
	return
}
