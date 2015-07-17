/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package sts

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"net/url"
	"sort"
	"strings"
)

//	Rename types.HeaderSorter to HeaderSorter.
//	将types包的HeaderSorter命名成HeaderSorter类。
type HeaderSorter types.HeaderSorter

// 	Generate the right Authorization by this canonicalizedResource.
//	生成签名方法（直接设置请求的Header）。
/*
 *	Example:
 *	c.SignHeader(req, "/")
 */
func STSStringToSign(accessKeySecret, postBody string) (signature string) {
	temp := make(map[string]string)

	tempAnd := strings.Split(postBody, "&")
	l := len(tempAnd)
	for i := 0; i < l; i++ {
		tempEqual := strings.Split(tempAnd[i], "=")
		if len(tempEqual) == 2 {
			temp[tempEqual[0]] = tempEqual[1]
		}
	}
	hs := NewHeaderSorter(temp)
	hs.Sort()
	canonicalizedQueryString := ""
	for i := range hs.Keys {
		canonicalizedQueryString += "&" + hs.Keys[i] + "=" + percentEncode(hs.Vals[i])
	}

	cqsStr, _ := url.Parse(canonicalizedQueryString)
	canonicalizedQueryString = cqsStr.String()
	canonicalizedQueryString = strings.Replace(canonicalizedQueryString, "=", "%3D", -1)
	canonicalizedQueryString = strings.Replace(canonicalizedQueryString, "&", "%26", -1)
	canonicalizedQueryString = strings.Replace(canonicalizedQueryString, ":", "%3A", -1)
	canonicalizedQueryString = strings.Replace(canonicalizedQueryString, ",", "%2C", -1)

	dotFirst := strings.Index(canonicalizedQueryString, "Policy")
	dotLast := strings.Index(canonicalizedQueryString, "%7D%5D%7D%26")
	dotFirstTS := strings.Index(canonicalizedQueryString, "Timestamp")
	bc := []byte(canonicalizedQueryString)
	lc := len(bc)
	policyStr := string(bc[dotFirst+9 : dotLast+9])
	policyStr = strings.Replace(policyStr, "/", "%2F", -1)
	policyStr = strings.Replace(policyStr, "%", "%25", -1)
	timestampStr := string(bc[dotFirstTS+12 : dotFirstTS+36])
	timestampStr = strings.Replace(timestampStr, "%", "%25", -1)

	signStr := "POST&%2F&" + percentEncode(string(bc[3:dotFirst+9])+policyStr+string(bc[dotLast+9:dotFirstTS+12])+timestampStr+string(bc[dotFirstTS+36:lc]))

	h := hmac.New(sha1.New, []byte(accessKeySecret+"&"))
	h.Write([]byte(signStr))

	signature = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}

//	Additional function for function SignHeader.
func NewHeaderSorter(m map[string]string) *HeaderSorter {
	hs := &HeaderSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]string, 0, len(m)),
	}

	for k, v := range m {
		hs.Keys = append(hs.Keys, k)
		hs.Vals = append(hs.Vals, v)
	}
	return hs
}

//	Additional function for function SignHeader.
func (hs *HeaderSorter) Sort() {
	sort.Sort(hs)
}

//	Additional function for function SignHeader.
func (hs *HeaderSorter) Len() int {
	return len(hs.Vals)
}

//	Additional function for function SignHeader.
func (hs *HeaderSorter) Less(i, j int) bool {
	return bytes.Compare([]byte(hs.Keys[i]), []byte(hs.Keys[j])) < 0
}

//	Additional function for function SignHeader.
func (hs *HeaderSorter) Swap(i, j int) {
	hs.Vals[i], hs.Vals[j] = hs.Vals[j], hs.Vals[i]
	hs.Keys[i], hs.Keys[j] = hs.Keys[j], hs.Keys[i]
}
