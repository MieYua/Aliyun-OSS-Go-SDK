/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"bytes"
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 	Append an object to bucket.
//	用Append方法上传文件。
/*
 *	URL查询参数还必须包含position，其值指定从何处进行追加。首次追加操作的position必须为
0，后续追加操作的position是Object的当前长度。例如，第一次Append Object请求指定
position值为0，content-length是65536；那么，第二次Append Object需要指定position为
65536。每次操作成功后，响应头部x-oss-next-append-position也会标明下一次追加的
position。
 *	downloadFileName为文件下载时显示的文件，为空时默认为上传的文件名
 *	Example:
 *	nextAppendPosition, err := c.AppendObject(bucketName, fileName, filePath, appendPosition, downloadFileName)
*/
func (c *Client) AppendObject(bucketName, fileName, filePath string, appendPosition int64, downloadFileName string) (nextAppendPosition int64, err error) {
	cc := c.CClient

	objectPath := bucketName + "/" + fileName
	if strings.HasPrefix(objectPath, "/") == false {
		objectPath = "/" + objectPath
	}
	buffer := new(bytes.Buffer)

	fh, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer fh.Close()
	io.Copy(buffer, fh)

	contentType := http.DetectContentType(buffer.Bytes())
	params := map[string]string{}
	params[consts.HH_CONTENT_TYPE] = contentType
	params[consts.HH_CONTENT_LENGTH] = strconv.Itoa(buffer.Len())
	params[consts.HH_CONTENT_ENCODING] = "utf-8"
	if downloadFileName != "" {
		params[consts.HH_CONTENT_DISPOSITION] = "attachment;filename=" + downloadFileName
	} else {
		params[consts.HH_CONTENT_DISPOSITION] = "attachment;filename=" + fileName
	}

	if appendPosition <= 0 && appendPosition > int64(buffer.Len()) {
		appendPosition = 0
	}

	url := objectPath + "?append&position=" + strconv.Itoa(int(appendPosition))

	resp, err := cc.DoRequest("POST", url, url, params, buffer)
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

	napint, err := strconv.Atoi(resp.Header.Get(consts.OH_OSS_NEXT_APPEND_POSITION))
	if err != nil {
		return
	}
	nextAppendPosition = int64(napint)

	//log.Println("A new object(" + objectPath + ") has been appended into this bucket.")
	return
}
