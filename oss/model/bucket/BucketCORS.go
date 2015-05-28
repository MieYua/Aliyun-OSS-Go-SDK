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

// 	Change the setting of this bukcet's cors.
//	修改Bucket的CORS设置。
/*
 *	Example:
 *	corsRule1 := (types)CORSRule{
 *		AllowedOrigin: []string{"*",...},
 *		AllowedMethod: []string{"PUT", "GET", "POST",...},
 *		AllowedHeader: []string{},						// can be null
 *		ExposeHeader:  []string{},						// can bu null
 *		// this bucket's cache time(s)
 *		MaxAgeSeconds: 100,
 *	}
 *	err := c.PutBucketCors(bucketName, [](types)CORSRule{corsRule1})
 */
func (c *Client) PutBucketCORS(bucketName string, corsRules []types.CORSRule) (err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?cors"

	corsxml := types.CORSXML{}

	for _, v := range corsRules {
		cr := types.CORSRule{}
		cr.AllowedMethod = v.AllowedMethod
		cr.AllowedOrigin = v.AllowedOrigin
		cr.AllowedHeader = v.AllowedHeader
		cr.ExposeHeader = v.ExposeHeader
		cr.MaxAgeSeconds = v.MaxAgeSeconds
		corsxml.CORSRule = append(corsxml.CORSRule, cr)
	}

	bs, err := xml.Marshal(corsxml)
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

	//log.Println("The CORS's setting of " + bucketName + " has been changed.")
	return
}

// 	Get the cors' setting of this bucket.
//	获得Bucket的CORS设置。
/*
 *	Example:
 *	corsc, err := c.GetBucketCors(bucketName)
 */
func (c *Client) GetBucketCORS(bucketName string) (corsc types.CORSConfiguration, err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?cors"
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

	err = xml.Unmarshal(body, &corsc)
	if err != nil {
		return
	}

	// log.Println("You have got the CORS's setting of " + bucketName + ".")
	return
}
