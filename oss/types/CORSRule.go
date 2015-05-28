/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The rule of the CORS.
//	跨域资源请求规则。
type CORSRule struct {
	AllowedOrigin []string // 允许的来源，默认通配符"*"
	AllowedMethod []string // 允许的方法
	AllowedHeader []string // 允许的请求头
	ExposeHeader  []string // 允许的响应头
	MaxAgeSeconds int      // 最大的缓存时间
}
