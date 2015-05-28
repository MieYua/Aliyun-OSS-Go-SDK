/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"encoding/xml"
)

//	The request's XML part of function PutBucketReferer.
//	修改Bucket的白名单设置XML请求部分。
type RefererXML struct {
	XMLName           xml.Name `xml:"RefererConfiguration"`
	AllowEmptyReferer bool     // 是否允许白名单为空
	RefererList       []RefererList
}
