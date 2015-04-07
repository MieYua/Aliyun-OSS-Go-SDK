/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/common"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
)

//	Import types.Client.
/*
 *
 */
type Client types.Client

//	Covert this Client into common's
/*
 *	Example:
 *	cc := ConvertClient(c)
 */
func ConvertClient(c *Client) *common.Client {
	cc := common.Client{
		Host:            c.Host,
		AccessKeyId:     c.AccessKeyId,
		AccessKeySecret: c.AccessKeySecret,
		HttpClient:      c.HttpClient,
		FileIOLocker:    c.FileIOLocker,
	}
	return &cc
}
