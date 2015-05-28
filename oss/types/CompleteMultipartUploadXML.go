/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"encoding/xml"
)

//	The request's XML part of function CompleteMultipartUpload.
//	完成Multipart Upload的XML请求部分。
type CompleteMultipartUploadXML struct {
	XMLName xml.Name `xml:"CompleteMultipartUpload"`
	Part    []Part
}
