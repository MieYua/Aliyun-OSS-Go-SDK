/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"encoding/xml"
)

//	The request's XML part of function PutBucketCORS.
//	跨域资源共享的XML请求部分。
type CORSXML struct {
	XMLName  xml.Name   `xml:"CORSConfiguration"`
	CORSRule []CORSRule // 跨域资源请求规则
}
