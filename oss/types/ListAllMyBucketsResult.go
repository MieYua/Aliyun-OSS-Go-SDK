/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function GetService.
//	GetService返回值的XML解析结果。
type ListAllMyBucketsResult struct {
	Owner   Owner   `xml:"Owner"`
	Buckets Buckets `xml:"Buckets"`
}
