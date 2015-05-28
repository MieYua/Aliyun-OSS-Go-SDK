/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The informations of the parts which have been uploaded.
//	已上传得到Part信息。
type UploadedPart struct {
	PartNumber   int    // Part编号
	LastModified string // 最后一次提交时间
	ETag         string // ETag缓存码
	Size         int    // Part大小
}
