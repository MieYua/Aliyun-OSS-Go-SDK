/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	The informations of the bucket's logging setting.
//	日志设置信息。
type LoggingEnabled struct {
	TargetBucket string // 日志存放Bucket名称
	TargetPrefix string // 日志名称前缀
}
