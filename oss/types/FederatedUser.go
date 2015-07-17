/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	FederatedUser's struct.
//	FederatedUser属性。
type FederatedUser struct {
	FederatedUserId string `json:"FederatedUserId,omitempty"` //	Bucket拥有者Id(数字:名称)
	Arn             string `json:"Arn,omitempty"`             //	Bucket的Arn
}
