/*
 * 	Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * 	All rights reserved.
 *
 * 	version 1.0 released on 150330
 *	version 2.0 released on 150402
 *	version 3.0 released on 150424
 *	version 4.0 released on 150715
 */

//	Aliyun OSS Go(Golang) SDK.
//	阿里云OSS的Go语言SDK。
package oss

import (
	"errors"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/common"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/model/bucket"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/model/multipart"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/model/object"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/model/service"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/sts"
	"github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"net/http"
)

//	Type: Clent.
//	Convert bucket.Client to BClient, multipart.Client to MClient, object.Client to OClient and service.Client to SClient.
//	分别将bucket，multipart，object和service包的Client转换在Client类内。
type Client struct {
	BClient *bucket.Client
	MClient *multipart.Client
	OClient *object.Client
	SClient *service.Client
}

// 	STS: GetSecurityToken.
//	Get the securityToken(Please get again(refresh) before the expiration is coming).
//	获得临时安全令牌（同一个用户名请在失效前再重新获取，不然容易返回空）。
/*
 *	Example(oss_test.go-TestGetSecurityToken):
 *	strj, err := GetSecurityToken(accessKeyId, accessKeySecret, assumeRole)
 *		assumeRole.DurationSeconds: mainAccount:900-3600s/childAccount:900-129600s
 *	Please see its example in oss_test.go
 *	实例请见oss_test.go-TestGetSecurityToken。
 *	失效时间strj.Credentials.Expiration；
 *	请自行按失效时间刷新。
 */
func GetSecurityToken(accessKeyId, accessKeySecret string, assumeRole *types.AssumeRole) (strj types.SecurityTokenResponseJSON, err error) {
	return sts.GetSecurityToken(accessKeyId, accessKeySecret, assumeRole)
}

// 	STS: SetSTSCondition.
//	Set the condition of STS.
//	设置STS的Condition参数。
/*
 *	Example:
 *	condition := SetSTSCondition(userAgent, currentTime, secureTransport, prefix, delimiter, sourceIp)
 *	userAgent		指定http useragent头
 *	currentTime		指定合法的访问时间
 *	secureTransport	是否是https协议（http或https）
 *	prefix			用作ListObjects(对应GetBucketInfo)的prefix
 *	delimiter		用作ListObjects(对应GetBucketInfo)的delimiter
 *	sourceIp		指定ip网段
 */
func SetSTSCondition(userAgent, currentTime, secureTransport, prefix, delimiter, sourceIp string) (condition types.Condition) {
	if userAgent != "" {
		condition.StringEquals.UserAgent = userAgent
	}
	if currentTime != "" {
		condition.StringEquals.CurrentTime = currentTime
	}
	if secureTransport != "" {
		condition.StringEquals.SourceTransport = secureTransport
	}
	if prefix != "" {
		condition.StringEquals.Prefix = prefix
	}
	if delimiter != "" {
		condition.StringEquals.Delimiter = delimiter
	}
	if sourceIp != "" {
		condition.IpAddress.SourceIp = sourceIp
	}
	return
}

//	Client: InitiateClient.
//	Initiate a new client.
//	初始化客户端。
/*
 *	Example:
 *	c := InitiateClient((consts)endPoint, accessKeyId, accessKeySecret)
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

//	Client: InitiateTempClient.
//	Initiate a new tempClient.
//	初始化临时客户端。
/*
 *	Example:
 *	c := InitiateTempClient((consts)endPoint, tempAccessKeyId, tempAccessKeySecret, securityToken)
 */
func InitiateTempClient(endPoint, tempAccessKeyId, tempAccessKeySecret, securityToken, tempPrefix, tempDelimiter string) *Client {
	cc := common.NewClient(endPoint, tempAccessKeyId, tempAccessKeySecret)
	cc.TClient.UserProperty = "TempUser"
	cc.TClient.SecurityToken = securityToken
	if tempPrefix != "" {
		cc.TClient.TempPrefix = tempPrefix
	}
	if tempDelimiter != "" {
		cc.TClient.TempDelimiter = tempDelimiter
	}
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
 *	lambr, err := c.GetServiceInfo()
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
//	返回用于设置的loggingSetting参数。
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
//	返回用于设置的websiteSetting参数。
/*
 *	Example:
 *	websiteSetting := SetBucketWebsite("index_xxxx.html", "error_xxxx.html")
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

//	Bucket: SetBucketReferer.
//	Get a referers.
//	返回用于设置的referers参数。
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
//	返回用于设置的rules参数。
/*
 *	Example:
 *	If you want to create a new rule, please follow:
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
//	Get a CORSRules.
//	返回用于设置的CORSRules参数。
/*
 *	Example:
 *	If you want to create a new rule, please follow:
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

// 	Bucket: DeleteBucket.
// 	Delete the details of this bucket.
//	删除Bucket或其中的信息。
/*
 *	Example:
 *	err := c.DeleteBucket(bucketName, "logging")
 *	If you use other deleteSettingName, err will be "Bad deleteSettingName!"
 */
func (c *Client) DeleteBucket(bucketName, deleteSettingName string) (err error) {
	switch deleteSettingName {
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
		return errors.New("Bad deleteSettingName!")
	}
}

// 	Bucket: GetBucketInfo.
//	Get all informations about this bucket and its all objects.
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

// 	Bucket: CleanBucket.
//	Delete all files in this bucket.
//	清空Bucket的所有文件。
/*
 *	Example:
 *	err := c.CleanBucket(bucketName)
 */
func (c *Client) CleanBucket(bucketName string) (err error) {
	bc := c.BClient
	oc := c.OClient
	for {
		lbr, err := bc.GetBucket(bucketName, "", "", "", "50")
		if err != nil {
			return err
		} else {
			objects := lbr.Contents
			length := len(objects)
			if length == 0 {
				return err
			} else {
				for i := 0; i < length; i++ {
					err = oc.DeleteObject(bucketName + "/" + objects[i].Key)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return
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
	if aclSetting != "" {
		err = bc.PutBucketACL(bucketName, aclSetting)
		if err != nil {
			return
		}
	}
	if loggingSetting != nil {
		err = bc.PutBucketLogging(bucketName, loggingSetting["targetBucket"], loggingSetting["targetPrefix"])
		if err != nil {
			return
		}
	}
	if websiteSetting != nil {
		err = bc.PutBucketWebsite(bucketName, websiteSetting["indexDocument"], websiteSetting["errorDocument"])
		if err != nil {
			return
		}
	}
	if referers != nil {
		err = bc.PutBucketReferer(bucketName, referers)
		if err != nil {
			return
		}
	}
	if rules != nil {
		err = bc.PutBucketLifecycle(bucketName, rules)
		if err != nil {
			return
		}
	}
	if corsRules != nil {
		err = bc.PutBucketCORS(bucketName, corsRules)
		if err != nil {
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
 *		Can be just a name of file(bucketName/fileName),
 *		Can be names of filepacks(bucketName/filePack/../fileName).
 */
func (c *Client) CreateObject(objectPath, filePath string) (err error) {
	oc := c.OClient
	err = oc.PutObject(objectPath, filePath)
	return
}

// 	Object: AppendObject.
// 	Append a new object to a bucket.
// 	在Bucket中追加一个Object。
/*
 *	Example:
 *	nextAppendPosition, err := c.AppendObject(bucketName, fileName, filePath, appendPosition, downloadFileName)
 *
 *	URL查询参数还必须包含position，其值指定从何处进行追加。首次追加操作的position必须
 *	为0，后续追加操作的position是Object的当前长度。例如，第一次Append Object请求指定
 *	position值为0，content-length是65536；那么，第二次Append Object需要指定position为
 *	65536。每次操作成功后，响应头部x-oss-next-append-position也会标明下一次追加的position。
 *	firstAppend---->    								appendPosition = 0 (return nextAppendPosition...)
 *	followingAppend---->								appendPosition = nextAppendPosition
 *	downloadFileName为文件下载时显示的文件，为空时默认为上传的文件名。
 *	此参数需在第一次上传时使用（即appendPosition为0时），之后无效。
 *	downloadFileName == "" ---->  	 					downloadFile's name = fileName
 *	downloadFileName != ""&&appendPosition==0 ---->		downloadFile's name = downloadFileName
 */
func (c *Client) AppendObject(bucketName, fileName, filePath string, appendPosition int64, downloadFileName string) (nextAppendPosition int64, err error) {
	oc := c.OClient
	nextAppendPosition, err = oc.AppendObject(bucketName, fileName, filePath, appendPosition, downloadFileName)
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
 *	obytes, err := c.GetObject(objectPath, rangeStart(form 0), rangeEnd(from 0 and it is larger than rangeStart))
 *	If you want the whole file,
 *			rangeStart:	default 	-1,
 *			rangeEnd:	default 	-1.
 *	Example:
 *	obytes, err := c.GetObject("xxxx/test.txt", -1, -1)
 *	log.Println(string(obytes[:]), err)-->test <nil>
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
 *	err := c.DeleteMultipleObject("bucketName", []string{"copy_test1.txt", "copy_test2.txt"})
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
 *	log.Println(header, err)
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
 *	err := c.PostObject(bucketName, fileName, tempFileName)
 */
func (c *Client) PostObject(bucketName, filePath string, tempFileName string) (err error) {
	oc := c.OClient
	err = oc.PostObject(bucketName, filePath, tempFileName)
	return
}

//	Multipart Upload: MultipartUpload.
//	Upload a file by multipart upload.
//	用Multipart Upload方式上传文件。
/*
 *	Example:
 *	cmur, err, lastPoint, uploadId := c.MultipartUpload(bucketName+"/test_mu.pdf", "test.pdf", 1024000)
 */
func (c *Client) MultipartUpload(objectPath, filePath string, chunkSize int64) (cmu types.CompleteMultipartUpload, err error, lastPoint int64, uploadId string) {
	mc := c.MClient
	initObjectPath, imur, _ := mc.InitiateMultipartUpload(objectPath)
	isLastPart := false
	lastPoint = 0
	uploadId = imur.UploadId
	for i := 1; isLastPart == false; i++ {
		isLastPart, lastPoint, cmu, err = mc.UploadPart(imur, initObjectPath, filePath, cmu, lastPoint, chunkSize, i)
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
 *	cmu, err, lastPoint, uploadIdCon := c.MultipartUpload(bucketName+"/test_mu.pdf", "test.pdf", breakPoint, 1024000)
 */
func (c *Client) ContinueMultipartUpload(objectPath, filePath string, breakPoint, chunkSize int64, uploadId string) (cmu types.CompleteMultipartUpload, err error, lastPoint int64, uploadIdCon string) {
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
		isLastPart, lastPoint, cmu, err = mc.UploadPart(imur, objectPath, filePath, cmu, lastPoint, chunkSize, i)
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
 *	cmu, err, lastPoint, uploadId := c.MultipartUploadCopy(bucketName, "test_muc.pdf", "test.pdf", 1024000)
 */
func (c *Client) MultipartUploadCopy(bucketName, objectPath, copyPath string, chunkSize int64) (cmu types.CompleteMultipartUpload, err error, lastPoint int64, uploadId string) {
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
		isLastPart, lastPoint, _, cmu, err = mc.UploadPartCopy(imur, initObjectPath, bucketName+"/"+copyPath, cmu, lastPoint, chunkSize, length, i)
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
 *	cmur, err := c.CompleteMultipartUpload(cmu, initObjectPath, uploadId)
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
