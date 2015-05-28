/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function GetBucketWebsite.
//	GetBucketWebsite返回值的XML解析结果。
type WebsiteConfiguration struct {
	IndexDocument IndexDocument
	ErrorDocument ErrorDocument
}
