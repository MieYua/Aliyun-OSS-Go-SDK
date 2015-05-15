/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"bytes"
	//"errors"
	//"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	//"io"
	//"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func (c *Client) UploadPartWeb(imur types.InitiateMultipartUploadResult, initObjectPath string, data []byte, cmu types.CompleteMultipartUpload, startPoint, cutLength int64, partNumber int) (isLastPart bool, endPoint int64, cmuNew types.CompleteMultipartUpload, err error) {
	cc := c.CClient

	bufferLength := new(bytes.Buffer)
	bufferLength.Write(data)
	buffer := new(bytes.Buffer)

	length := int64(bufferLength.Len())

	if partNumber < 1 {
		partNumber = 1
	}

	if cutLength < 102400 {
		cutLength = 102400
	}

	if length < (startPoint + cutLength - 1) {
		cutLength = length - startPoint
		endPoint = length - 1
		isLastPart = true
	} else {
		endPoint = startPoint + cutLength
		isLastPart = false
	}

	bufCutFile := make([]byte, cutLength)

	newFileName := bufferLength.Bytes()
	file, err := os.Create(string(newFileName[0:15]))
	file.ReadAt(bufCutFile, startPoint)
	var i int64 = 0
	for i = 0; i < cutLength; i++ {
		buffer.WriteByte(bufCutFile[i])
	}

	if strings.HasPrefix(initObjectPath, "/") == false {
		initObjectPath = "/" + initObjectPath
	}

	reqStr := initObjectPath + "?partNumber=" + strconv.Itoa(partNumber) + "&uploadId=" + imur.UploadId

	resp, err := cc.DoRequest("PUT", reqStr, reqStr, nil, buffer)
	if err != nil {
		return
	}

	// body, _ := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	err = errors.New(resp.Status)
	// 	fmt.Println(string(body))
	// 	return
	// }

	newPart := types.Part{}
	newPart.ETag = resp.Header.Get(consts.HH_ETAG)
	newPart.PartNumber = partNumber
	cmuNew.Part = append(cmu.Part, newPart)
	//fmt.Println("Part number " + strconv.Itoa(partNumber) + " of the " + initObjectPath + " has been uploaded.")
	return

}
