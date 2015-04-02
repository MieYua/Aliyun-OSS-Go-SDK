/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package oss

import (
	// "Aliyun-OSS-Go-SDK/oss/model/bucket"
	"github.com/MieYua/Aliyun-OSS-Go-SDK-v2/oss/bucket"
	// "Aliyun-OSS-Go-SDK/oss/model/multipart"
	"github.com/MieYua/Aliyun-OSS-Go-SDK-v2/oss/multipart"
	// "Aliyun-OSS-Go-SDK/oss/model/object"
	"github.com/MieYua/Aliyun-OSS-Go-SDK-v2/oss/object"
	// "Aliyun-OSS-Go-SDK/oss/model/service"
	"github.com/MieYua/Aliyun-OSS-Go-SDK-v2/oss/service"
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
