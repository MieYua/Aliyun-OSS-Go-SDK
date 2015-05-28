/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function GetBucketCORS.
//	获得Bucket的跨域资源设置的返回XML解析结果。
type CORSConfiguration struct {
	CORSRule []CORSRule // 返回的跨域资源设置规则
}
