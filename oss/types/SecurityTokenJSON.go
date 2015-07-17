/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	SecurityToken's json struct.
//	SecurityToken的json属性。
type SecurityTokenJSON struct {
	Version   string      `json:"Version"` // 	版本，一般为"1"
	Statement []Statement `json:"Statement,omitempty"`
}
