/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The deault error page(*.html only).
//	默认错误页（目前只支持HTML网页）。
type ErrorDocument struct {
	Key string // 默认错误页地址
}
