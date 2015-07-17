/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"time"
)

//	Credentials's struct.
//	Credentials属性。
type Credentials struct {
	AccessKeySecret string    `json:"AccessKeySecret,omitempty"` //	临时secret
	AccessKeyId     string    `json:"AccessKeyId,omitempty"`     //	临时id
	Expiration      time.Time `json:"Expiration,omitempty"`      //	失效时间
	SecurityToken   string    `json:"SecurityToken,omitempty"`   //	安全令牌
}
