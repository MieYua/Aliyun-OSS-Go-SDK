/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The list of preparing for function CompleteMultipartUpload.
//	完成Multipart Upload的存放容器。
type CompleteMultipartUpload struct {
	Part []Part
}
