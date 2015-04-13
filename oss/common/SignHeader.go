/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"hash"
	"io"
	"net/http"
	"sort"
	"strings"
)

type HeaderSorter types.HeaderSorter

// 	Generate the right Authorization by this canonicalizedResource.
/*
 *	Example:
 *	c.SignHeader(req, "/")
 */
func (c *Client) SignHeader(req *http.Request, canonicalizedResource string) {
	// Find out the "x-oss-"'s address in this request'header
	temp := make(map[string]string)

	for k, v := range req.Header {
		if strings.HasPrefix(strings.ToLower(k), "x-oss-") {

			temp[strings.ToLower(k)] = v[0]
		}
	}
	hs := NewHeaderSorter(temp)

	// Sort the temp by the Ascending Order
	hs.Sort()

	// Get the CanonicalizedOSSHeaders
	canonicalizedOSSHeaders := ""
	for i := range hs.Keys {
		canonicalizedOSSHeaders += hs.Keys[i] + ":" + hs.Vals[i] + "\n"
	}

	// Give other parameters values
	date := req.Header.Get("Date")
	contentType := req.Header.Get("Content-Type")
	contentMd5 := req.Header.Get("Content-Md5")

	signStr := req.Method + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n" + canonicalizedOSSHeaders + canonicalizedResource
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(c.TClient.AccessKeySecret))
	io.WriteString(h, signStr)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Get the final Authorization' string
	authorizationStr := "OSS " + c.TClient.AccessKeyId + ":" + signedStr

	// Give the parameter "Authorization" value
	req.Header.Set("Authorization", authorizationStr)
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
