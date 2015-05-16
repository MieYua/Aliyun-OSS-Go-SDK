/*
 * 	Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * 	All rights reserved.
 *
 * 	version 1.0 released on 150330
 *	version 2.0 created on 150402
 */

//	Aliyun OSS Go(Golang) SDK.
package oss

import (
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/common"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/model/bucket"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/model/multipart"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/model/object"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/model/service"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"net/http"
)

//	Type: Clent.
//	Import common.Client.
//	导入common包下的Client。
/*
 *
 */
type Client struct {
	BClient *bucket.Client
	MClient *multipart.Client
	OClient *object.Client
	SClient *service.Client
}

//	Client: InitiateClient.
//	Initiate a new client.
//	初始化客户端。
/*
 *	Example:
 *	c := InitiateClient((consts)ENDPOINT, accessKeyId, accessKeySecret)
 */
func InitiateClient(endPoint, accessKeyId, accessKeySecret string) *Client {
	cc := common.NewClient(endPoint, accessKeyId, accessKeySecret)
	bc := bucket.Client{}
	mc := multipart.Client{}
	oc := object.Client{}
	sc := service.Client{}

	bc.CClient = cc
	mc.CClient = cc
	oc.CClient = cc
	sc.CClient = cc

	c := Client{
		BClient: &bc,
		MClient: &mc,
		OClient: &oc,
		SClient: &sc,
	}

	return &c
}

//	Service: GetServiceInfo.
//	Get the informations of this service.
//	获取服务信息。
/*
 *	Example:
 *	lambr,err := c.GetServiceInfo()
 */
func (c *Client) GetServiceInfo() (lambr types.ListAllMyBucketsResult, err error) {
	sc := c.SClient
	return sc.GetService()
}

//	Bucket: CreateBucket.
//	Create a new bucket.
//	新建Bucket。
/*
 *	Example:
 *	err := c.CreateBucket(bucketName)
 */
func (c *Client) CreateBucket(bucketName string) (err error) {
	bc := c.BClient
	return bc.PutBucket(bucketName)
}

// 	Bucket: SetBucketLogging.
//	Get a loggingSetting.
//	获得loggingSetting参数。
/*
 *	Example:
 *	loggingSetting := SetBucketLogging("xxxlogs", "Mylog-")
 */
func SetBucketLogging(targetBucket, targetPrefix string) (loggingSetting map[string]string) {
	temp := map[string]string{}
	temp["targetBucket"] = targetBucket
	temp["targetPrefix"] = targetPrefix
	return temp
}

// 	Bucket: SetBucketWebsite.
//	Get a websiteSetting.
//	获得websiteSetting参数。
/*
 *	Example:
 *	websiteSetting := SetBucketWebsite("inde_xxxx.html", "error_xxxx.html")
 */
func SetBucketWebsite(indexDocument, errorDocument string) (websiteSetting map[string]string) {
	temp := map[string]string{}
	if indexDocument == "" {
		temp["indexDocument"] = "index.html"
	} else {
		temp["indexDocument"] = indexDocument
	}
	if errorDocument == "" {
		temp["errorDocument"] = "error.html"
	} else {
		temp["errorDocument"] = errorDocument
	}
	return temp
}

// 	Bucket: SetBucketReferer.
//	Get a referers.
//	获得referers参数。
/*
 *	Example:
 *	referers := SetBucketReferer([]string{consts.REFERER_XX,...})
 */
func SetBucketReferer(refererlist []string) (referers []string) {
	temp := []string{}
	length := len(refererlist)
	for i := 0; i < length; i++ {
		temp = append(temp, refererlist[i])
	}
	return temp
}

// 	Bucket: SetBucketLifecycle.
//	Get a rules.
//	获得rules参数。
/*
 *	Example:
 *	rules := SetBucketLifecycle([]types.Rule{}, "Delete after one month", "xxxxlogs", "", 30)
 *	If you want to add another rule, please follow:
 *	rules = SetBucketLifecycle(rules, "Delete after one year", "xxxxlogs2", "", 365)
 */
func SetBucketLifecycle(rulesOld []types.Rule, id, prefix, status string, days int) (rules []types.Rule) {
	if status == "" {
		status = "Enabled"
	}
	ruleAdd := types.Rule{
		Id:         id,
		Prefix:     prefix,
		Status:     status,
		Expiration: types.Expiration{Days: days},
	}
	rules = append(rulesOld, ruleAdd)
	return
}

// 	Bucket: SetBucketCORS.
//	Get a corsRules.
//	获得corsRules参数。
/*
 *	Example:
 *	corsRules := SetBucketCORS([]types.CORSRule{}, []string{"*"}, []string{"GET", "POST", "PUT"}, []string{}, []string{}, 100)
 *	If you want to add another rule, please follow:
 *	corsRules = SetBucketCORS(corsRules, []string{"http://www.a.com"}, []string{"GET", "POST"}, []string{}, []string{}, 100)
 */
func SetBucketCORS(corsRulesOld []types.CORSRule, allowedOrigin, allowedMethod, allowedHeader, exposeHeader []string, maxAgeSeconds int) (corsRules []types.CORSRule) {
	corsRuleAdd := types.CORSRule{
		AllowedOrigin: allowedOrigin,
		AllowedMethod: allowedMethod,
		AllowedHeader: allowedHeader, // can be null
		ExposeHeader:  exposeHeader,  // can bu null
		MaxAgeSeconds: maxAgeSeconds, // this bucket's cache time(s)
	}
	corsRules = append(corsRulesOld, corsRuleAdd)
	return
}

// 	Bucket: DeleteBucket...
// 	Delete the details of this bucket.
//	删除Bucket或其中的信息。
/*
 *	Example:
 *	err := c.DeleteBucket(bucketName, "logging")
 *	If you use other deleteName, err will be "Bad deleteName!"
 */
func (c *Client) DeleteBucket(bucketName, deleteName string) (err error) {
	switch deleteName {
	case "bucket":
		return c.DeleteObject(bucketName)
	case "logging":
		return c.DeleteObject(bucketName + "?logging")
	case "website":
		return c.DeleteObject(bucketName + "?website")
	case "lifecycle":
		return c.DeleteObject(bucketName + "?lifecycle")
	case "cors":
		return c.DeleteObject(bucketName + "?cors")
	default:
		return errors.New("Bad deleteName!")
	}
}

// 	Bucket: GetBucketInfo.
//	Get all information about this bucket and its all objects.
//	获得该Bucket的所有信息以及所有Object信息。
/*
 *	Example:
 *	lbr, err := c.GetBucketInfo(bucketName, prefix, marker, delimiter, maxkeys)
 *	prefix: Choose that contain this string (default:"")
 *	marker: Return after this letter (default:"")
 *	delimiter: Common Prefixes (default:"")
 *	maxkeys: The maximum of objects (default:"100")
 */
func (c *Client) GetBucketInfo(bucketName, prefix, marker, delimiter, maxkeys string) (lbr types.ListBucketResult, err error) {
	bc := c.BClient
	return bc.GetBucket(bucketName, prefix, marker, delimiter, maxkeys)
}

// 	Bucket: SetBucket.
//	Change some settings of this bucket.
//	修改Bucket的一些设置。
/*
 *	Example:
 *	loggingSetting := SetBucketLogging("xxxxlogs", "")
 *	websiteSetting := SetBucketWebsite("index_test.html", "error_test.html")
 *	referers := SetBucketReferer([]string{consts.REFERER_XXXX,...})
 *	rules := SetBucketLifecycle([]types.Rule{}, "Delete after one month", "xxxxlogs", "", 30)
 *	corsRules := SetBucketCORS([]types.CORSRule{}, []string{"*"}, []string{"GET", "POST", "PUT"}, []string{}, []string{}, 100)
 *
 *	err := c.SetBucket(testBucketName, consts.ACL_PUBLIC_RW, loggingSetting, websiteSetting, referers, rules, corsRules)
 *
 *	If you don't want to change some settings, please set:
 *	aclSetting = "",
 *	loggingSetting = nil,
 *	websiteSetting = nil,
 *	referers = nil,
 *	rules = nil,
 *	corsRules = nil.
 */
func (c *Client) SetBucket(bucketName, aclSetting string, loggingSetting, websiteSetting map[string]string, referers []string, rules []types.Rule, corsRules []types.CORSRule) (err error) {
	bc := c.BClient
	err = nil
	if aclSetting != "" {
		errACL := bc.PutBucketACL(bucketName, aclSetting)
		if errACL != nil {
			err = errACL
			return
		}
	}
	if loggingSetting != nil {
		errLogging := bc.PutBucketLogging(bucketName, loggingSetting["targetBucket"], loggingSetting["targetPrefix"])
		if errLogging != nil {
			err = errLogging
			return
		}
	}
	if websiteSetting != nil {
		errWebsite := bc.PutBucketWebsite(bucketName, websiteSetting["indexDocument"], websiteSetting["errorDocument"])
		if errWebsite != nil {
			err = errWebsite
			return
		}
	}
	if referers != nil {
		errReferer := bc.PutBucketReferer(bucketName, referers)
		if errReferer != nil {
			err = errReferer
			return
		}
	}
	if rules != nil {
		errLifecycle := bc.PutBucketLifecycle(bucketName, rules)
		if errLifecycle != nil {
			err = errLifecycle
			return
		}
	}

	if corsRules != nil {
		errCORS := bc.PutBucketCORS(bucketName, corsRules)
		if errCORS != nil {
			err = errCORS
			return
		}
	}
	return
}

// 	Bucket: GetBucketSetting.
//	Get all settings of this bucket.
//	获取该Bucket的所有设置。
/*
 *	Example:
 *	acl, lc, bls, wc, rc, lfc, corsc, err := c.GetBucketSetting(testBucketName)
 *	acl:	AccessControlPolicy
 *	lc: 	LocationConstraint
 *	bls: 	BucketLoggingStatus
 *	wc:		WebsiteConfiguration
 *	lfc: 	RefererConfiguration
 *	lc:  	LifecycleConfiguration
 *	corsc: 	CORSConfiguration
 */
func (c *Client) GetBucketSetting(bucketName string) (acl types.AccessControlPolicy, lc types.LocationConstraint, bls types.BucketLoggingStatus, wc types.WebsiteConfiguration, rc types.RefererConfiguration, lfc types.LifecycleConfiguration, corsc types.CORSConfiguration, err error) {
	bc := c.BClient
	err = nil
	acl, err = bc.GetBucketACL(bucketName)
	if err != nil {
		return
	}
	lc, err = bc.GetBucketLocation(bucketName)
	if err != nil {
		return
	}
	bls, err = bc.GetBucketLogging(bucketName)
	if err != nil {
		return
	}
	wc, err = bc.GetBucketWebsite(bucketName)
	if err != nil {
		return
	}
	rc, err = bc.GetBucketReferer(bucketName)
	if err != nil {
		return
	}
	lfc, err = bc.GetBucketLifecycle(bucketName)
	if err != nil {
		return
	}
	corsc, err = bc.GetBucketCORS(bucketName)
	if err != nil {
		return
	}
	return
}

// 	Object: OptionObject.
// 	Choose the object by options.
// 	发送跨域请求，返回是否符合条件。
/*
 *	Example:
 *	err := c.OptionObject("bucketName/test.txt", "", "")
 *	fmt.Println(err)
 *
 *	Warning:
 *	If the bucket's cors is not available or its cors hasn't been set up,
 *	response will show 403 ERROR.
 */
func (c *Client) OptionObject(objectPath, accessControlRequestMethod, accessControlRequestHeader, origin string) (err error) {
	oc := c.OClient
	err = oc.OptionObject(objectPath, accessControlRequestMethod, accessControlRequestHeader, origin)
	return
}

// 	Object: CreateObject.
// 	Create a new object to a bucket.
// 	在Bucket中新建一个Object。
/*
 *	Example:
 *	err := c.PutObject(objectPath, filePath)
 *	objectPath:
 *			Can be just a name of file(bucketName/fileName),
 *			Can be names of filepacks(bucketName/filePack/../fileName).
 */
func (c *Client) CreateObject(objectPath, filePath string) (err error) {
	oc := c.OClient
	err = oc.PutObject(objectPath, filePath)
	return
}

// 	Object: CreateObjectWeb.
// 	Create a new object to a bucket in web service.
// 	网页中在Bucket中新建一个Object。
/*
 *	Example:
 *	err := c.PutObjectWeb(objectPath, file)
 *	objectPath:
 *			Can be just a name of file(bucketName/fileName),
 *			Can be names of filepacks(bucketName/filePack/../fileName).
 */
func (c *Client) CreateObjectWeb(objectPath string, file []byte) (err error) {
	oc := c.OClient
	err = oc.PutObjectWeb(objectPath, file)
	return
}

// 	Object: CopyObject.
// 	Copy an object of one bucket to another bucket or this one.
// 	在Bucket中拷贝一个Object。
/*
 *	Example:
 *	err := c.CopyObject(pasteSrc, copySrc)
 *	copySrc: 	the file need be copied(bucket1Name/../fileName)
 *	pasteSrc:	the file src need be pasted(bucket2Name/../copy_fileName)
 *
 *	If file size is larger than 1GB, please use function UploadPartCopy
 */
func (c *Client) CopyObject(pasteSrc, copySrc string) (cor types.CopyObjectResult, err error) {
	oc := c.OClient
	cor, err = oc.CopyObject(pasteSrc, copySrc)
	return
}

//	Object: GetObject.
//	Get bytes of object.
// 	获取Object的字节。
/*
 *	obytes,err := c.GetObject(objectPath, rangeStart(form 0), rangeEnd(from 0 and it is larger than rangeStart))
 *	If you want the whole file,
 *			rangeStart:	default 	-1,
 *			rangeEnd:	default 	-1.
 *	Example:
 *	obytes, err := c.GetObject("xxxx/test.txt", -1, -1)
 *	fmt.Println(string(obytes[:]), err)-->test <nil>
 */
func (c *Client) GetObject(objectPath string, rangeStart, rangeEnd int) (obytes []byte, err error) {
	oc := c.OClient
	obytes, err = oc.GetObject(objectPath, rangeStart, rangeEnd)
	return
}

//	Object: DeleteObject.
//	Delete an object.
// 	删除Object。
/*
 * 	Example:
 *	err := c.DeleteObject(objectPath)
 *	objectPath: bucketName/objectName
 *	Warning:
 *	If you want to delete a filepack, you need clear all files in this filepack and than delete this filepack.
 *		c.PutObject("bucketName/test/test.txt")
 *		The wrong way:
 *		c.DeleteObject("bucketName/test/") can't delete this filepack
 *		The right way:
 *		c.DeleteObject("bucketName/test/test.txt")
 *		c.DeleteObject("bucketName/test/")
 */
func (c *Client) DeleteObject(objectPath string) (err error) {
	oc := c.OClient
	err = oc.DeleteObject(objectPath)
	return
}

//	Object: DeleteMultipleObject.
//	Delete some objects at one time.
// 	删除一些Object。
/*
 *	Example:
 *	c.CopyObject("bucketName/copy_test1.txt", "bucketName/test.txt")
 *	c.CopyObject("bucketName/copy_test2.txt", "bucketName/test.txt")
 *	c.DeleteMultipleObject("bucketName", []string{"copy_test1.txt", "copy_test2.txt"})
 */
func (c *Client) DeleteMultipleObject(bucketName string, keys []string) (err error) {
	oc := c.OClient
	err = oc.DeleteMultipleObject(bucketName, keys)
	return
}

//	Object: HeadObject.
//	Find the object's head's meta information and show it out.
// 	查询Object的头信息。
/*
 *	Example:
 *	header, err := c.HeadObject("bucketName/test.txt")
 *	fmt.Println(header, err)
 *	--> map[Accept-Ranges:[bytes] Etag:["xxxxx"] Server:[AliyunOSS] Date:[xxxx GMT] Content-Type:[text/plain; charset=utf-8] Content-Length:[x] Last-Modified:[xxxx GMT] X-Oss-Request-Id:[xxxxxx]] <nil>
 */
func (c *Client) HeadObject(objectPath string) (header http.Header, err error) {
	oc := c.OClient
	header, err = oc.HeadObject(objectPath)
	return
}

//	Object: PostObject.
//	Post up an object to replace putObject.
// 	用Post方法上传Object。
/*
 *	Example:
 *	err := c.PostObject(bucketName, fileName, data)
 */
func (c *Client) PostObject(bucketName, filePath string, tempFileName string) (err error) {
	oc := c.OClient
	err = oc.PostObject(bucketName, filePath, tempFileName)
	return
}

//	Multipart Upload: MultipartUpload.
//	Upload a file by multipart upload.
// 	用Multipart Upload方式上传文件。
/*
 *	Example:
 *	cmur, err := c.MultipartUpload(bucketName+"/test_mu.pdf", "test.pdf", 1024000)
 */
func (c *Client) MultipartUpload(objectPath, filePath string, cutLength int64) (cmu types.CompleteMultipartUpload, err error, lastPoint int64, uploadId string) {
	mc := c.MClient
	initObjectPath, imur, _ := mc.InitiateMultipartUpload(objectPath)
	isLastPart := false
	lastPoint = 0
	uploadId = imur.UploadId
	for i := 1; isLastPart == false; i++ {
		isLastPart, lastPoint, cmu, err = mc.UploadPart(imur, initObjectPath, filePath, cmu, lastPoint, cutLength, i)
		if err != nil {
			return
		}
	}
	return
}

//	Multipart Upload: MultipartUploadWeb.
//	Upload a file by multipart upload in web service.
// 	在浏览器用Multipart Upload方式上传文件。
/*
 *	Example:
 *	cmur, err := c.MultipartUploadWeb(bucketName+"/test_mu.pdf", file, 1024000)
 */
func (c *Client) MultipartUploadWeb(objectPath string, file []byte, cutLength int64) (cmu types.CompleteMultipartUpload, err error, lastPoint int64, uploadId string) {
	mc := c.MClient
	initObjectPath, imur, _ := mc.InitiateMultipartUpload(objectPath)
	isLastPart := false
	lastPoint = 0
	uploadId = imur.UploadId
	for i := 1; isLastPart == false; i++ {
		isLastPart, lastPoint, cmu, err = mc.UploadPartWeb(imur, initObjectPath, file, cmu, lastPoint, cutLength, i)
		if err != nil {
			return
		}
	}
	return
}

//	Multipart Upload: ContinueMultipartUpload.
//	Upload a file by multipart upload.
// 	用Multipart Upload方式继续上传文件。
/*
 *	Example:
 *	cmur, err := c.MultipartUpload(bucketName+"/test_mu.pdf", "test.pdf", breakPoint, 1024000)
 */
func (c *Client) ContinueMultipartUpload(objectPath, filePath string, breakPoint, cutLength int64, uploadId string) (cmu types.CompleteMultipartUpload, err error, lastPoint int64, uploadIdCon string) {
	mc := c.MClient
	lpr, _ := mc.ListParts(objectPath, uploadId)
	maxPartNumber := 1
	length := len(lpr.Part)
	for i := 0; i < length; i++ {
		if lpr.Part[i].PartNumber > maxPartNumber {
			maxPartNumber = lpr.Part[0].PartNumber
		}
	}
	isLastPart := false
	imur := types.InitiateMultipartUploadResult{}
	imur.UploadId = uploadId
	uploadIdCon = uploadId
	lastPoint = breakPoint
	for i := maxPartNumber; isLastPart == false; i++ {
		isLastPart, lastPoint, cmu, err = mc.UploadPart(imur, objectPath, filePath, cmu, lastPoint, cutLength, i)
		if err != nil {
			return
		}
	}
	return
}

//	Multipart Upload: ContinueMultipartUploadWeb.
//	Upload a file by multipart upload in web service.
// 	在浏览器用Multipart Upload方式继续上传文件。
/*
 *	Example:
 *	cmur, err := c.MultipartUploadWeb(bucketName+"/test_mu.pdf", file, breakPoint, 1024000)
 */
func (c *Client) ContinueMultipartUploadWeb(objectPath string, file []byte, breakPoint, cutLength int64, uploadId string) (cmu types.CompleteMultipartUpload, err error, lastPoint int64, uploadIdCon string) {
	mc := c.MClient
	lpr, _ := mc.ListParts(objectPath, uploadId)
	maxPartNumber := 1
	length := len(lpr.Part)
	for i := 0; i < length; i++ {
		if lpr.Part[i].PartNumber > maxPartNumber {
			maxPartNumber = lpr.Part[0].PartNumber
		}
	}
	isLastPart := false
	imur := types.InitiateMultipartUploadResult{}
	imur.UploadId = uploadId
	uploadIdCon = uploadId
	lastPoint = breakPoint
	for i := maxPartNumber; isLastPart == false; i++ {
		isLastPart, lastPoint, cmu, err = mc.UploadPartWeb(imur, objectPath, file, cmu, lastPoint, cutLength, i)
		if err != nil {
			return
		}
	}
	return
}

//	Multipart Upload: MultipartUploadCopy.
//	Upload a file by multipart upload copy.
// 	用Multipart Upload Copy方式上传文件。
/*
 *	Example:
 *	cmur, err := c.MultipartUploadCopy(bucketName, "test_muc.pdf", "test.pdf", 1024000)
 */
func (c *Client) MultipartUploadCopy(bucketName, objectPath, copyPath string, cutLength int64) (cmu types.CompleteMultipartUpload, err error, lastPoint int64, uploadId string) {
	mc := c.MClient
	initObjectPath, imur, _ := mc.InitiateMultipartUpload(bucketName + "/" + objectPath)
	isLastPart := false
	var length int64 = 0
	// 	Get the length of copyFile
	lbr, _ := c.GetBucketInfo(bucketName, "", "", "", "100")
	l := len(lbr.Contents)
	for i := 0; i < l; i++ {
		if (lbr.Contents[i].Key) == copyPath {
			size := lbr.Contents[i].Size
			length = int64(size)
		}
	}
	lastPoint = 0
	for i := 1; isLastPart == false; i++ {
		isLastPart, lastPoint, _, cmu, err = mc.UploadPartCopy(imur, initObjectPath, bucketName+"/"+copyPath, cmu, lastPoint, cutLength, length, i)
		if err != nil {
			return
		}
	}
	uploadId = imur.UploadId
	return
}

//	Multipart Upload: CompleteMultipartUpload.
//	Complete the unuploaded multipart upload parts whose uploadId is thisId.
// 	完成指定UploadId的Multipart Upload上传。
/*
 *	Example:
 *	err := c.CompleteMultipartUpload(cmu, initObjectPath, uploadId)
 */
func (c *Client) CompleteMultipartUpload(cmu types.CompleteMultipartUpload, initObjectPath, uploadId string) (cmur types.CompleteMultipartUploadResult, err error) {
	mc := c.MClient
	cmur, err = mc.CompleteMultipartUpload(cmu, initObjectPath, uploadId)
	return
}

//	Multipart Upload: AbortMultipartUpload.
//	Delete the unuploaded multipart upload parts whose uploadId is thisId.
// 	放弃指定UploadId的Multipart Upload上传。
/*
 *	Example:
 *	err := c.AbortMultipartUpload(objectPath, uploadId)
 */
func (c *Client) AbortMultipartUpload(objectPath, uploadId string) (err error) {
	mc := c.MClient
	err = mc.AbortMultipartUpload(objectPath, uploadId)
	return
}

//	Multipart Upload: CleanMultipartUpload.
//	Delete all unuploaded multipart upload parts.
// 	清空Bucket内未上传的Part碎片。
/*
 *	Example:
 *	err := c.CleanMultipartUpload(bucketName)
 */
func (c *Client) CleanMultipartUpload(bucketName string) (err error) {
	mc := c.MClient
	lmur, _ := mc.ListMultipartUpload(bucketName, nil)
	length := len(lmur.Upload)
	for i := 0; i < length; i++ {
		err = mc.AbortMultipartUpload(bucketName+"/"+lmur.Upload[i].Key, lmur.Upload[i].UploadId)
	}
	return
}

//	Multipart Upload: ListMultipartUpload.
//	Give the list of the unuploaded multipart upload missions.
// 	列出Bucket内所有未上传的Part信息。
/*
 *	Example:
 *	lmur, err := c.ListMultipartUpload(bucketName, map[string]string or nil)
 *	If you want to search all results, second parameter is nil.
 */
func (c *Client) ListMultipartUpload(bucketName string, params map[string]string) (lmur types.ListMultipartUploadResult, err error) {
	mc := c.MClient
	lmur, err = mc.ListMultipartUpload(bucketName, params)
	return
}

//	Multipart Upload: ListParts.
//	Give the list of the uploaded multipart upload mission by uploadId.
// 	通过UploadId列出所有已经上传的Part信息。
//
/*
 *	Example:
 *	lpr, err := c.ListParts(bucketName, uploadId)
 */
func (c *Client) ListParts(objectName, uploadId string) (lpr types.ListPartsResult, err error) {
	mc := c.MClient
	lpr, err = mc.ListParts(objectName, uploadId)
	return
}

func (c *Client) InitiateMultipartUpload(objectPath string) {
	mc := c.MClient
	mc.InitiateMultipartUpload(objectPath)
}
