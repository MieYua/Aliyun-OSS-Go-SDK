/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	SecurityToken's response's json struct.
//	SecurityToken的返回json属性。
type SecurityTokenResponseJSON struct {
	FederatedUser FederatedUser `json:"FederatedUser,omitempty"` //	联盟信息
	Credentials   Credentials   `json:"Credentials,omitempty"`   //	令牌信息
	RequestId     string        `json:"RequestId,omitempty"`     //	请求Id
}
