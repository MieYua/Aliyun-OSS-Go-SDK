/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	Bucket's struct.
//	Bucket属性。
type Bucket struct {
	Location     string `xml:"Location"`     // Bucket的节点地址
	Name         string `xml:"Name"`         // Bucket的名称
	CreationDate string `xml:"CreationDate"` // Bucket的创建时间
}
