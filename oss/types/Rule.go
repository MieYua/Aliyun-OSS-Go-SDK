/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The rule of the lifecycle.
//	生命周期的规则。
type Rule struct {
	Id         string // 生命周期标记名称
	Prefix     string // 对应的Bucket名称
	Status     string // 生命周期状态: Enabled
	Expiration Expiration
}
