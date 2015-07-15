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
	STS_ACTION_BUCKET_GetBucket             string = "oss:ListObjects"
	STS_ACTION_BUCKET_PutBucketAcl                 = "oss:PutBucketAcl"
	STS_ACTION_BUCKET_DeleteBucket                 = "oss:DeleteBucket"
	STS_ACTION_BUCKET_GetBucketLocation            = "oss:GetBucketLocation"
	STS_ACTION_BUCKET_GetBucketAcl                 = "oss:GetBucketAcl"
	STS_ACTION_BUCKET_GetBucketLogging             = "oss:GetBucketLogging"
	STS_ACTION_BUCKET_PutBucketLogging             = "oss:PutBucketLogging"
	STS_ACTION_BUCKET_DeleteBucketLogging          = "oss:DeleteBucketLogging"
	STS_ACTION_BUCKET_GetBucketWebsite             = "oss:GetBucketWebsite"
	STS_ACTION_BUCKET_PutBucketWebsite             = "oss:PutBucketWebsite"
	STS_ACTION_BUCKET_DeleteBucketWebsite          = "oss:DeleteBucketWebsite"
	STS_ACTION_BUCKET_GetBucketReferer             = "oss:GetBucketReferer"
	STS_ACTION_BUCKET_PutBucketReferer             = "oss:PutBucketReferer"
	STS_ACTION_BUCKET_GetBucketLifecycle           = "oss:GetBucketLifecycle"
	STS_ACTION_BUCKET_PutBucketLifecycle           = "oss:PutBucketLifecycle"
	STS_ACTION_BUCKET_DeleteBucketLifecycle        = "oss:DeleteBucketLifecycle"
	STS_ACTION_BUCKET_ListMultipartUploads         = "oss:ListMultipartUploads"
	STS_ACTION_BUCKET_PutBucketCors                = "oss:PutBucketCors"
	STS_ACTION_BUCKET_GetBucketCors                = "oss:GetBucketCors"
	STS_ACTION_BUCKET_DeleteBucketCors             = "oss:DeleteBucketCors"
	//	STS Action.(Object level)
	//	STS Action配置。（Object级别）
	STS_ACTION_OBJECT_GetObject               = "oss:GetObject"
	STS_ACTION_OBJECT_HeadObject              = "oss:GetObject"
	STS_ACTION_OBJECT_PutObject               = "oss:PutObject"
	STS_ACTION_OBJECT_PostObject              = "oss:PutObject"
	STS_ACTION_OBJECT_InitiateMultipartUpload = "oss:PutObject"
	STS_ACTION_OBJECT_UploadPart              = "oss:PutObject"
	STS_ACTION_OBJECT_CompleteMultipart       = "oss:PutObject"
	STS_ACTION_OBJECT_DeleteObject            = "oss:DeleteObject"
	STS_ACTION_OBJECT_DeleteMultipartObjects  = "oss:DeleteObject"
	STS_ACTION_OBJECT_AbortMultipartUpload    = "oss:AbortMultipartUpload"
	STS_ACTION_OBJECT_ListParts               = "oss:ListParts"
	STS_ACTION_OBJECT_CopyObject              = "oss:GetObject,oss:PutObject"
	STS_ACTION_OBJECT_UploadPartCopy          = "oss:GetObject,oss:PutObject"
	//	STS Condition.
	//	STS条件。
	STS_ACTION_CONDITION_SourceIp        = "acs:SourceIp"        //指定ip网段:普通的ip，支持*通配
	STS_ACTION_CONDITION_UserAgent       = "acs:UserAgent"       //指定http useragent头:字符串
	STS_ACTION_CONDITION_CurrentTime     = "acs:CurrentTime"     //指定合法的访问时间:ISO8601格式
	STS_ACTION_CONDITION_SecureTransport = "acs:SecureTransport" //是否是https协议:"http"或者"https"
	STS_ACTION_CONDITION_Prefix          = "oss:Prefix"          //用作ListObjects时的prefix:合法的object name
	STS_ACTION_CONDITION_Delimiter       = "oss:Delimiter"       //用作ListObject时的delimiter:合法的delimiter值

)
