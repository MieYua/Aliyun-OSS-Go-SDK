/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

// 	The analysis of response's XML part of function GetBucketLifecycle.
//	GetBucketLifecycle返回值的XML解析结果。
type LifecycleConfiguration struct {
	Rule []Rule
}
