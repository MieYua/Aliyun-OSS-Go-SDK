/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"encoding/xml"
)

//	The request's XML part of function PutBucketLogging.
//	修改Bucket的logging属性XML请求部分。
type LoggingXML struct {
	XMLName        xml.Name `xml:"BucketLoggingStatus"`
	LoggingEnabled LoggingEnabled
}
