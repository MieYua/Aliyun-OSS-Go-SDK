/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package bucket

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/common"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//	Import common.Client.
/*
 *
 */
type Client struct {
	CClient *common.Client
}

// 	Bucket: Create a new bucket.
/*
 *	Example:
 *	err := PutBucket(bucketName)
 */
func (c *Client) PutBucket(bucketName string) (err error) {
	cc := c.CClient

	reqStr := "/" + bucketName

	bxml := types.BucketXML{}

	// Get the correct region
	if strings.HasPrefix(c.CClient.TClient.Host, "-internal.aliyuncs.com") {
		urlSuffixInternal := "-internal.aliyuncs.com"
		bxml.LocationConstraint = strings.TrimSuffix(c.CClient.TClient.Host, urlSuffixInternal)
	} else {
		urlSuffix := ".aliyuncs.com"
		bxml.LocationConstraint = strings.TrimSuffix(c.CClient.TClient.Host, urlSuffix)
	}

	b, err := xml.Marshal(bxml)
	if err != nil {
		return
	}
	buffer := new(bytes.Buffer)
	buffer.Write(b)

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
	}
	log.Println("A new bucket(" + bucketName + ") has been created.")
	return
}
