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
	ACCESSKEYID                 = "XXXXXXXXXXXXXXXX"
	ACCESSKEYSECRET             = "XXXXXXXXXXXXXXXX"
	TESTBUCKETNAME              = "testbucket"
	TESTFILENAME                = "test.txt"
	TESTFILEPATH                = "xxx/test.txt"
	TESTDOWNLOADFILENAME        = "test_download.txt"
)

func TestAppendObject(t *testing.T) {
	Convey("获得所有文件Id测试", t, func() {
		fmt.Println("")
		c := InitiateClient(ENDPOINT, ACCESSKEYID, ACCESSKEYSECRET)
		nextAppendPosition, err := c.AppendObject(TESTBUCKETNAME, TESTFILENAME, TESTFILEPATH, 0, TESTDOWNLOADFILENAME)
		fmt.Println(err)
		if err == nil {
			fmt.Println(nextAppendPosition)
		}
	})
}
