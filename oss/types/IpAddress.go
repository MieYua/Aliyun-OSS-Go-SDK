/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	IpAddress's struct.
//	IpAddress属性。
type IpAddress struct {
	SourceIp string `json:"acs:SourceIp,omitempty"` //	指定ip网段 普通的ip，支持*通配
}
