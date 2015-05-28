/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	Object's struct.
//	Object属性。
type Object struct {
	Key          string // 访问路径地址
	LastModified string // 最后一次提交
	ETag         string // ETag标记码
	Type         string // Object类型
	Size         int    // Object大小
	StorageClass string // Obejct存储类型,目前只能是"Standard"
	Owner        Owner
}
