/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function UploadPartCopy.
//	拷贝Part用于Multipart Upload的返回XML解析结果。
type CopyPartResult struct {
	LastModified string
	ETag         string
}
