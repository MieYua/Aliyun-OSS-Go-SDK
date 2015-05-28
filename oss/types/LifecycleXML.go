/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"encoding/xml"
)

//	The request's XML part of function PutBucketLifecycle.
//	修改Bucket的生命周期设置XML请求部分。
type LifecycleXML struct {
	XMLName xml.Name `xml:"LifecycleConfiguration"`
	Rule    []Rule
}
