/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function CompleteMultipartUpload.
//	完成Multipart Upload的返回XML解析结果。
type CompleteMultipartUploadResult struct {
	Location string
	Bucket   string
	ETag     string
	Key      string
}
