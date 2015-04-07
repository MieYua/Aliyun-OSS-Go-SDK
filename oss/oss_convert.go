/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package oss

import (
	"Aliyun-OSS-Go-SDK/oss/model/bucket"
	"Aliyun-OSS-Go-SDK/oss/model/multipart"
	"Aliyun-OSS-Go-SDK/oss/model/object"
	"Aliyun-OSS-Go-SDK/oss/model/service"
)

//	Convert this Client type into bucket's.
func ConvertClientBucket(c *Client) *bucket.Client {
	bc := bucket.Client{
		Host:            c.Host,
		AccessKeyId:     c.AccessKeyId,
		AccessKeySecret: c.AccessKeySecret,
		HttpClient:      c.HttpClient,
		FileIOLocker:    c.FileIOLocker,
	}
	return &bc
}

//	Convert this Client type into multipart's.
func ConvertClientMultipart(c *Client) *multipart.Client {
	mc := multipart.Client{
		Host:            c.Host,
		AccessKeyId:     c.AccessKeyId,
		AccessKeySecret: c.AccessKeySecret,
		HttpClient:      c.HttpClient,
		FileIOLocker:    c.FileIOLocker,
	}
	return &mc
}

//	Convert this Client type into object's.
func ConvertClientObject(c *Client) *object.Client {
	oc := object.Client{
		Host:            c.Host,
		AccessKeyId:     c.AccessKeyId,
		AccessKeySecret: c.AccessKeySecret,
		HttpClient:      c.HttpClient,
		FileIOLocker:    c.FileIOLocker,
	}
	return &oc
}

//	Convert this Client type into service's.
func ConvertClientService(c *Client) *service.Client {
	sc := service.Client{
		Host:            c.Host,
		AccessKeyId:     c.AccessKeyId,
		AccessKeySecret: c.AccessKeySecret,
		HttpClient:      c.HttpClient,
		FileIOLocker:    c.FileIOLocker,
	}
	return &sc
}
