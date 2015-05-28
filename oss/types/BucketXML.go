/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"encoding/xml"
)

//	The request's XML part of function PutBucket.
//	新建Bucket的XML请求部分。
type BucketXML struct {
	XMLName            xml.Name `xml:"CreateBucketConfiguration"`
	LocationConstraint string   // 当前访问节点地址
}
