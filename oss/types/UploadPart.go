/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The informations of the parts which haven't been uploaded.
//	未上传的Part信息。
type UploadPart struct {
	Key       string // Object名称
	UploadId  string // 对应UploadId
	Initiated string // 初始化时间
}
