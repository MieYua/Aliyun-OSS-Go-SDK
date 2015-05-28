/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of error.
//	Error返回XML解析结果。
type Error struct {
	Code    string
	Message string
}
