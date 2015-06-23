/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"bytes"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// 	Upload a new file part.
//	上传一个新part。
/*
 *	Example:
 *	isLastPart, end, cmuNew, err := c.UploadPart(imur, initobjectPath, filePath, cmu, start, chunkSize, partNumber)
 *
 *	chunkSize must be larger than 102400
 *	If chunkSize is smaller than 102400, chunkSize will be 102400
 */
func (c *Client) UploadPart(imur types.InitiateMultipartUploadResult, initObjectPath, filePath string, cmu types.CompleteMultipartUpload, start, chunkSize int64, partNumber int) (isLastPart bool, end int64, cmuNew types.CompleteMultipartUpload, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := stat.Size()
	if start >= fileSize {
		err = errors.New("invalid argument")
		return
	}
	_, err = file.Seek(start, 0)
	if err != nil {
		return
	}
	var written int64
	isLastPart, cmuNew, written, err = c.UploadPartFromReader(imur, initObjectPath, file, cmu, chunkSize, partNumber)
	end = start + written
	return
}

// 	Upload a new file part.
//	上传一个新part。
/*
 *	Example:
 *	isLastPart, cmuNew,written, err := c.UploadPart(imur, initobjectPath, reader, cmu, chunkSize, partNumber)
 *
 *	chunkSize must be larger than 102400
 *	If chunkSize is smaller than 102400, chunkSize will be 102400
 */
func (c *Client) UploadPartFromReader(imur types.InitiateMultipartUploadResult, initObjectPath string, reader io.Reader, cmu types.CompleteMultipartUpload, chunkSize int64, partNumber int) (isLastPart bool, cmuNew types.CompleteMultipartUpload, written int64, err error) {

	if partNumber < 1 {
		partNumber = 1
	}

	if chunkSize < 102400 {
		chunkSize = 102400
	}

	buffer := new(bytes.Buffer)
	written, err = io.CopyN(buffer, reader, chunkSize)
	if err != nil {
		if err != io.EOF {
			return
		} else {
			isLastPart = true
			err = nil
		}
	} else {
		isLastPart = false
	}

	cc := c.CClient
	if strings.HasPrefix(initObjectPath, "/") == false {
		initObjectPath = "/" + initObjectPath
	}

	reqStr := initObjectPath + "?partNumber=" + strconv.Itoa(partNumber) + "&uploadId=" + imur.UploadId

	resp, err := cc.DoRequest("PUT", reqStr, reqStr, nil, buffer)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println(string(body))
		return
	}

	newPart := types.Part{}
	newPart.ETag = resp.Header.Get(consts.HH_ETAG)
	newPart.PartNumber = partNumber
	cmuNew.Part = append(cmu.Part, newPart)

	//log.Println("Part number " + strconv.Itoa(partNumber) + " of the " + initObjectPath + " has been uploaded.")
	return

}
