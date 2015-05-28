/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package types

import (
	"encoding/xml"
)

//	The request's XML part of function DeleteMultipleObject.
//	删除多个Object的XML请求部分。
type DeleteXML struct {
	XMLName xml.Name       `xml:"Delete"`
	Object  []DeleteObject // 删除的所有Object
	Quiet   bool           // "安静"响应模式
}
