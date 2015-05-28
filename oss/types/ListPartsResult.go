/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function ListParts.
//	列出已上传的Part的返回XML解析结果。
type ListPartsResult struct {
	Bucket               string         // Bucket名称
	Key                  string         // Object名称
	UploadId             string         // 上传Id
	NextPartNumberMarker string         // 下一个Part的位置
	MaxParts             int            // 最大Part个数
	IsTruncated          bool           // 是否完全上传完成
	Part                 []UploadedPart // 已完成的Part
}
