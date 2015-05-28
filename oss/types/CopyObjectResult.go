/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function CopyObject.
//	CopyObject返回值的XML解析结果。
type CopyObjectResult struct {
	LastModified string
	ETag         string
}
