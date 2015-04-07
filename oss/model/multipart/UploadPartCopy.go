/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io/ioutil"
	"strconv"
	"strings"
)

// 	Copy a file of the bucket to upload a part.
/*
 *	Example:
 *	var length int64 = 0
 *		// Get the length of copyFile
 *		lbr, _ := c.GetBucket("xxxx", "", "", "", "100")
 *		l := len(lbr.Contents)
 *		for i := 0; i < l; i++ {
 *			if (lbr.Contents[i].Key) == "test.pdf" {
 *				size := lbr.Contents[i].Size
 *				length = int64(size)
 *			}
 *		}
 *	isLastPart, endPoint, cpr, cmu, err = c.UploadPartCopy(imur, initObjectPath, "xxxx/test.pdf", cmu, endPoint, 1048576(1MB), length, i)
 *
 *	If file size is smaller than 1GB, please use function CopyObject.
 */
func (c *Client) UploadPartCopy(imur types.InitiateMultipartUploadResult, initObjectPath, copySrc string, cmu types.CompleteMultipartUpload, startPoint, cutLength, length int64, partNumber int) (isLastPart bool, endPoint int64, cpr types.CopyPartResult, cmuNew types.CompleteMultipartUpload, err error) {
	cc := ConvertClient(c)

	if strings.HasPrefix(copySrc, "/") == false {
		copySrc = "/" + copySrc
	}
	if strings.HasPrefix(initObjectPath, "/") == false {
		initObjectPath = "/" + initObjectPath
	}

	if partNumber < 1 {
		partNumber = 1
	}

	if cutLength < 102400 {
		cutLength = 102400 // min 100KB
	}

	if length <= (startPoint + cutLength) {
		cutLength = length - startPoint
		endPoint = length - 1
		isLastPart = true
	} else {
		endPoint = startPoint + cutLength
		isLastPart = false
	}

	reqStr := initObjectPath + "?partNumber=" + strconv.Itoa(partNumber) + "&uploadId=" + imur.UploadId

	params := map[string]string{consts.OH_COPY_OBJECT_SOURCE: copySrc}
	params[consts.OH_COPY_SOURCE_RANGE] = "bytes=" + strconv.Itoa(int(startPoint)) + "-" + strconv.Itoa(int(startPoint+cutLength-1))

	resp, err := cc.DoRequest("PUT", reqStr, reqStr, params, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &cpr)
	newPart := types.Part{}
	newPart.ETag = cpr.ETag
	newPart.PartNumber = partNumber
	cmuNew.Part = append(cmu.Part, newPart)
	fmt.Println("Partnumber " + strconv.Itoa(partNumber) + " of the " + initObjectPath + " has been copied.")
	return

}
