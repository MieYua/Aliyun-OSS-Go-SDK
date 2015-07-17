/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package oss

import (
	"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	ENDPOINT             string = consts.ENDPOINT_HANGZHOU
	ACCESSKEYID                 = "xxxxxxxxxxxxx"
	ACCESSKEYSECRET             = "xxxxxxxxxxxxx"
	TESTBUCKETNAME              = "xxxxxxxxxxxxx"
	TESTFILENAME                = "test.txt"
	TESTFILEPATH                = "X://.../test.txt"
	TESTDOWNLOADFILENAME        = "test_download.txt"
	TESTUSERNAME                = "testUser"
)

func TestAppendObject(t *testing.T) {
	Convey("追加文件测试", t, func() {
		fmt.Println("")
		c := InitiateClient(ENDPOINT, ACCESSKEYID, ACCESSKEYSECRET)
		nextAppendPosition, err := c.AppendObject(TESTBUCKETNAME, TESTFILENAME, TESTFILEPATH, 0, TESTDOWNLOADFILENAME)
		fmt.Println(err)
		if err == nil {
			fmt.Println(nextAppendPosition)
		}
	})
}

func TestGetSecurityToken(t *testing.T) {
	Convey("获得安全令牌测试", t, func() {
		fmt.Println("")
		c := InitiateClient(ENDPOINT, ACCESSKEYID, ACCESSKEYSECRET)
		lambr, err := c.GetServiceInfo()
		if err != nil {
			fmt.Println(lambr.Owner.ID, err)
		} else {
			bucketOwner := lambr.Owner.ID
			username := "xxxxxxx"
			durationSeconds := 1800
			allowedActions := []string{consts.STS_ACTION_OBJECT_PUTOBJECT}
			allowedResources := []string{"acs:oss:*:" + bucketOwner + ":" + TESTBUCKETNAME}
			effect := "Allow"
			regionId := "cn-hangzhou"
			condition := SetSTSCondition("", "", "", "", "", "")
			strj, err := GetSecurityToken(ACCESSKEYID, ACCESSKEYSECRET, username, durationSeconds, allowedActions, allowedResources, effect, condition, regionId)
			fmt.Println(err)
			fmt.Println(strj.RequestId)                                             // 本次请求的Id
			fmt.Println(strj.FederatedUser.FederatedUserId, strj.FederatedUser.Arn) // Bucket拥有者的信息
			fmt.Println(strj.Credentials.AccessKeyId)                               // 	临时AccessKeyId
			fmt.Println(strj.Credentials.AccessKeySecret)                           //	临时AccessKeySecret
			fmt.Println(strj.Credentials.Expiration)                                //	令牌失效时间
			fmt.Println(strj.Credentials.SecurityToken)                             //	临时令牌
		}
	})
}

func TestInitiateDifferentClient(t *testing.T) {
	Convey("获得不同的Client服务测试", t, func() {
		fmt.Println("")
		fmt.Println("拥有者(开发者)获得Client服务")
		cOwner := InitiateClient(ENDPOINT, ACCESSKEYID, ACCESSKEYSECRET)
		lambr, err := cOwner.GetServiceInfo()
		if err == nil {
			fmt.Println("获得开发者Client（OwnerName：" + lambr.Owner.DisplayName + "）。")
		}
		fmt.Println("临时获得Client服务")
		if err == nil {
			bucketOwner := lambr.Owner.ID
			username := "xxxxxxxxxx"
			durationSeconds := 3600                                                                            //	有效时长
			allowedActions := []string{consts.STS_ACTION_BUCKET_GETBUCKET, consts.STS_ACTION_OBJECT_PUTOBJECT} //	允许的操作动作
			allowedResources := []string{"acs:oss:*:" + bucketOwner + ":" + TESTBUCKETNAME + "/*"}             //	允许操作的bucket
			//	"acs:oss:*:" + bucketOwner + ":" + TESTBUCKETNAME			---->bucket操作资源
			//	"acs:oss:*:" + bucketOwner + ":" + TESTBUCKETNAME + "/*"	---->bucket下object操作资源（支持通配符*）
			effect := "Allow"
			regionId := "cn-hangzhou"                            //	节点，形式如例子
			condition := SetSTSCondition("", "", "", "", "", "") //	Policy设置，详情见oss文档
			strj, err := GetSecurityToken(ACCESSKEYID, ACCESSKEYSECRET, username, durationSeconds, allowedActions, allowedResources, effect, condition, regionId)
			if err == nil {
				tempExpiration := strj.Credentials.Expiration           //	令牌失效时间（请自行重新获取）
				tempAccessKeyId := strj.Credentials.AccessKeyId         // 	临时AccessKeyId
				tempAccessKeySecret := strj.Credentials.AccessKeySecret //	临时AccessKeySecret
				securityToken := strj.Credentials.SecurityToken         //	临时令牌
				cTempUser := InitiateTempClient(consts.ENDPOINT_HANGZHOU, tempAccessKeyId, tempAccessKeySecret, securityToken, condition.StringEquals.Prefix, condition.StringEquals.Delimiter)
				//_, err := cTempUser.GetBucketInfo(TESTBUCKETNAME, "", "", "", "100")
				err := cTempUser.CreateObject(TESTBUCKETNAME+"/"+TESTFILENAME, TESTFILEPATH)
				if err == nil {
					fmt.Println("获得临时Client（失效时间：" + tempExpiration.Format("2006年01月02日15时04分05秒") + "）。")
				}
			}
		}
	})
}
