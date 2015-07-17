/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	Codition's struct.
//	Condition属性。
type Condition struct {
	StringEquals StringEquals `json:"StringEquals,omitempty"`
	IpAddress    IpAddress    `json:"IpAddress,omitempty"`
}
