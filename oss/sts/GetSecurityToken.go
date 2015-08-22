/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package sts

import (
	"encoding/json"
	"errors"
	// "fmt"
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
func GetSecurityToken(accessKeyId, accessKeySecret, username string, durationSeconds int, policy *types.SecurityTokenJSON, regionId string) (securityTokenResponseJSON types.SecurityTokenResponseJSON, err error) {
	reqUrl := "https://sts.aliyuncs.com"

	bs, _ := json.Marshal(policy)
	policyUrl, _ := url.Parse(string(bs))
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
	queryMap.Add("StsVersion", "1")
	queryMap.Add("Name", username)
	queryMap.Add("DurationSeconds", strconv.Itoa(durationSeconds))
	queryMap.Add("Policy", string(bs))
	queryMap.Add("Action", "GetFederationToken")
	queryMap.Add("Format", "json")
	queryMap.Add("Version", "2015-04-01")
	queryMap.Add("SignatureMethod", "HMAC-SHA1")
	queryMap.Add("SignatureNonce", strconv.Itoa(randNumber))
	queryMap.Add("SignatureVersion", "1.0")
	queryMap.Add("AccessKeyId", accessKeyId)
	queryMap.Add("Timestamp", date)
	queryMap.Add("RegionId", regionId)

	signature := STSStringToSign(accessKeySecret, percentEncode(queryMap.Encode()))
	queryMap.Add("Signature", signature)
	reqUrl = reqUrl + "/?" + queryMap.Encode()

	req, _ := http.NewRequest("POST", reqUrl, nil)
	// fmt.Println(req)
	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		return
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			var errMap map[string]string
			err = json.Unmarshal(b, &errMap)
			if err != nil {
				return
			}
			err = errors.New("Get security token error! Code: " + errMap["Code"] + "; Message: " + errMap["Message"] + "; StatusCode: " + strconv.Itoa(resp.StatusCode))
			return
		}

		err = json.Unmarshal(b, &securityTokenResponseJSON)
		if err != nil {
			return
		}
	}
	return
}

func percentEncode(str string) (pestr string) {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	pestr = str
	return
}
