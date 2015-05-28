/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function GetBucketReferer.
//	GetBucketReferer返回值的XML解析结果。
type RefererConfiguration struct {
	AllowEmptyReferer string      `xml:"AllowEmptyReferer"`
	RefererList       RefererList `xml:"RefererList"`
}
