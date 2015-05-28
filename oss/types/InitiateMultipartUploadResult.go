/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function InitiateMultipartUpload.
//	初始化Multipart Upload返回XML解析结果。
type InitiateMultipartUploadResult struct {
	Bucket   string // Bucket名称
	Key      string // 上传Object名称
	UploadId string // 生成的UploadId
}
