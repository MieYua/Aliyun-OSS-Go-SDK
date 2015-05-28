/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function ListMultipartUpload.
//	列出未上传的Part的返回XML解析结果。
type ListMultipartUploadResult struct {
	Bucket             string       // Bucket名称
	KeyMarker          string       // 起始Object位置
	UploadIdMarker     string       // 起始UploadId位置
	NextKeyMarker      string       // 如果没有全部返回，标明接下去的KeyMarker位置
	NextUploadIdMarker string       // 如果没有全部返回，标明接下去的UploadId位置
	MaxUploads         int          // 返回最大Upload数目
	IsTruncated        bool         // 是否完全返回
	Upload             []UploadPart // 未上传的部分
}
