/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The analysis of response's XML part of function GetBucket(List Object).
//	GetBucket(List Object)返回值的XML解析结果。
type ListBucketResult struct {
	Name        string   // Bucket名称
	Prefix      string   // 查询前缀，默认为空查询所有
	Marker      string   // 查询起点，默认为空查询所有
	MaxKeys     int      // 返回结果的最大数量,默认100
	Delimiter   string   // 用于分组，默认为空
	IsTruncated bool     // 所有结果是否均返回
	Contents    []Object // 所有Object的属性
}
