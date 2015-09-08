/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package sts

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 	Get the securityToken.
//	获得临时安全令牌。
/*
 *	Example:
 *	strj, err := GetSecurityToken(accessKeyId, accessKeySecret, username, durationSeconds, allowedActions, allowedResources, effect, condition, regionId)
 *		durationSeconds: mainAccount:900-3600s/childAccount:900-129600s
 */
func GetSecurityToken(accessKeyId, accessKeySecret string, assumeRole *types.AssumeRole) (securityTokenResponseJSON types.SecurityTokenResponseJSON, err error) {
	reqUrl := "https://sts.aliyuncs.com"

	bs, err := json.Marshal(assumeRole.Policy)
	if err != nil {
		return
	}
	policyUrl, err := url.Parse(string(bs))
	if err != nil {
		return
	}
	policyEncode := policyUrl.String()
	policyEncode = strings.Replace(policyEncode, "=", "%3D", -1)
	policyEncode = strings.Replace(policyEncode, "&", "%26", -1)

	date := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	//	用时间生成随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNumber := r.Intn(99999999)
	if randNumber <= 0 {
		randNumber = -randNumber
	}

	queryMap := url.Values{}
	queryMap.Add("Action", "AssumeRole")
	queryMap.Add("RoleArn", assumeRole.RoleArn)
	queryMap.Add("RoleSessionName", assumeRole.RoleSessionName)
	queryMap.Add("DurationSeconds", strconv.Itoa(assumeRole.DurationSeconds))
	queryMap.Add("Policy", string(bs))

	queryMap.Add("Format", "json")
	queryMap.Add("Version", "2015-04-01")
	queryMap.Add("SignatureMethod", "HMAC-SHA1")
	queryMap.Add("SignatureNonce", strconv.Itoa(randNumber))
	queryMap.Add("SignatureVersion", "1.0")
	queryMap.Add("AccessKeyId", accessKeyId)
	queryMap.Add("Timestamp", date)

	signature := stsStringToSign(accessKeySecret, queryMap)
	queryMap.Add("Signature", signature)
	reqUrl = reqUrl + "/?" + queryMap.Encode()
	req, err := http.NewRequest("POST", reqUrl, nil)
	if err != nil {
		return
	}
	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var errMap map[string]string
		err = json.Unmarshal(b, &errMap)
		if err != nil {
			return
		}
		err = errors.New("Get security token error! Code: " + errMap["Code"] + "; Message: " + errMap["Message"] + "; StatusCode: " + strconv.Itoa(resp.StatusCode) + "; RequestUrl: " + reqUrl)
		return
	}

	err = json.Unmarshal(b, &securityTokenResponseJSON)
	if err != nil {
		return
	}

	return
}

func percentEncode(str string) (pestr string) {
	str = url.QueryEscape(str)
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	pestr = str
	return
}

func stsStringToSign(accessKeySecret string, postBody url.Values) (signature string) {
	signStr := "POST&%2F&" + percentEncode(postBody.Encode())

	h := hmac.New(sha1.New, []byte(accessKeySecret+"&"))
	h.Write([]byte(signStr))

	signature = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}
