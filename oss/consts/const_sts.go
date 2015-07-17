/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package consts

const (
	//	STS configuration.
	//	STS配置常量。
	//	STS Action.(Bucket level)
	//	STS Action配置。（Bucket级别）
	STS_ACTION_BUCKET_GETBUCKET             string = "oss:ListObjects"
	STS_ACTION_BUCKET_PUTBUCKETACL                 = "oss:PutBucketAcl"
	STS_ACTION_BUCKET_DELETEBUCKET                 = "oss:DeleteBucket"
	STS_ACTION_BUCKET_GETBUCKETLOCATION            = "oss:GetBucketLocation"
	STS_ACTION_BUCKET_GETBUCKETACL                 = "oss:GetBucketAcl"
	STS_ACTION_BUCKET_GETBUCKETLOGGING             = "oss:GetBucketLogging"
	STS_ACTION_BUCKET_PUTBUCKETLOGGING             = "oss:PutBucketLogging"
	STS_ACTION_BUCKET_DELETEBUCKETLOGGING          = "oss:DeleteBucketLogging"
	STS_ACTION_BUCKET_GETBUCKETWEBSITE             = "oss:GetBucketWebsite"
	STS_ACTION_BUCKET_PUTBUCKETWEBSITE             = "oss:PutBucketWebsite"
	STS_ACTION_BUCKET_DELETEBUCKETWEBSITE          = "oss:DeleteBucketWebsite"
	STS_ACTION_BUCKET_GETBUCKETREFERER             = "oss:GetBucketReferer"
	STS_ACTION_BUCKET_PUTBUCKETREFERER             = "oss:PutBucketReferer"
	STS_ACTION_BUCKET_GETBUCKETLIFECYCLE           = "oss:GetBucketLifecycle"
	STS_ACTION_BUCKET_PUTBUCKETLIFECYCLE           = "oss:PutBucketLifecycle"
	STS_ACTION_BUCKET_DELETEBUCKETLIFECYCLE        = "oss:DeleteBucketLifecycle"
	STS_ACTION_BUCKET_LISTMULTIPARTUPLOADS         = "oss:ListMultipartUploads"
	STS_ACTION_BUCKET_PUTBUCKETCORS                = "oss:PutBucketCors"
	STS_ACTION_BUCKET_GETBUCKETCORS                = "oss:GetBucketCors"
	STS_ACTION_BUCKET_DELETEBUCKETCORS             = "oss:DeleteBucketCors"
	//	STS Action.(Object level)
	//	STS Action配置。（Object级别）
	STS_ACTION_OBJECT_APPENDOBJECT            = "oss:GetObject,oss:PutObject"
	STS_ACTION_OBJECT_GETOBJECT               = "oss:GetObject"
	STS_ACTION_OBJECT_HEADOBJECT              = "oss:GetObject"
	STS_ACTION_OBJECT_PUTOBJECT               = "oss:PutObject"
	STS_ACTION_OBJECT_POSTOBJECT              = "oss:PutObject"
	STS_ACTION_OBJECT_INITIATEMULTIPARTUPLOAD = "oss:PutObject"
	STS_ACTION_OBJECT_UPLOADPART              = "oss:PutObject"
	STS_ACTION_OBJECT_COMPLETEMULTIPART       = "oss:PutObject"
	STS_ACTION_OBJECT_DELETEOBJECT            = "oss:DeleteObject"
	STS_ACTION_OBJECT_DELETEMULTIPARTOBJECTS  = "oss:DeleteObject"
	STS_ACTION_OBJECT_ABORTMULTIPARTUPLOAD    = "oss:AbortMultipartUpload"
	STS_ACTION_OBJECT_LISTPARTS               = "oss:ListParts"
	STS_ACTION_OBJECT_COPYOBJECT              = "oss:GetObject,oss:PutObject"
	STS_ACTION_OBJECT_UPLOADPARTCOPY          = "oss:GetObject,oss:PutObject"
	//	STS Condition.
	//	STS条件。
	STS_ACTION_CONDITION_SOURCEIP        = "acs:SourceIp"        //	指定ip网段:普通的ip，支持*通配
	STS_ACTION_CONDITION_USERAGENT       = "acs:UserAgent"       //	指定http useragent头:字符串
	STS_ACTION_CONDITION_CURRENTTIME     = "acs:CurrentTime"     //	指定合法的访问时间:ISO8601格式
	STS_ACTION_CONDITION_SECURETRANSPORT = "acs:SecureTransport" //	是否是https协议:"http"或者"https"
	STS_ACTION_CONDITION_PREFIX          = "oss:Prefix"          //	用作ListObjects时的prefix:合法的object name
	STS_ACTION_CONDITION_DELIMITER       = "oss:Delimiter"       //	用作ListObject时的delimiter:合法的delimiter值

)
