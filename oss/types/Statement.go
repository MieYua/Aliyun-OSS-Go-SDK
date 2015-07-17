/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

//	Statement's struct.
//	Statement属性。
type Statement struct {
	Action    []string  `json:"Action,omitempty"`
	Resource  []string  `json:"Resource,omitempty"`  //	Resource指代的是OSS上面的某个具体的资源或者某些资源（支持*通配），resource的规则是“acs:oss:{region}:{bucket_owner}:{bucket_name}/{object_name}”。对于所有bucket级别的操作来说不需要最后的斜杠和{object_name}，就像这样“acs:oss:{region}:{bucket_owner}:{bucket_name}”。Resource也是一个列表，可以有多个Resource。
	Effect    string    `json:"Effect,omitempty"`    //	本条Policy的Statement的授权的结果，分为Allow和Deny，分别指代通过和禁止。
	Condition Condition `json:"Condition,omitempty"` //	Condition代表Policy授权的一些条件。
}
