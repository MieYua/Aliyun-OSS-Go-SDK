/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"encoding/xml"
)

//	The request's XML part of function PutBucketWebsite.
//	修改Bucket的Website设置XML请求部分。
type WebsiteXML struct {
	XMLName       xml.Name `xml:"WebsiteConfiguration"`
	IndexDocument IndexDocument
	ErrorDocument ErrorDocument
}
