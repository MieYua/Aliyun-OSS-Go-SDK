/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package consts

const (
	//	OSS Header.
	//	OSS头标签。
	OH_OSS_PREFIX                             string = "x-oss-"
	OH_OSS_USER_METADATA_PREFIX                      = "x-oss-meta-"
	OH_OSS_CANNED_ACL                                = "x-oss-acl"
	OH_STORAGE_CLASS                                 = "x-oss-storage-class"
	OH_OSS_SECURITY_TOKEN                            = "x-oss-security-token"
	OH_OSS_VERSION_ID                                = "x-oss-version-id"
	OH_OSS_SERVER_SIDE_ENCRYPTION                    = "x-oss-server-side-encryption"
	OH_GET_OBJECT_IF_MODIFIED_SINCE                  = "If-Modified-Since"
	OH_GET_OBJECT_IF_UNMODIFIED_SINCE                = "If-Unmodified-Since"
	OH_GET_OBJECT_IF_MATCH                           = "If-Match"
	OH_GET_OBJECT_IF_NONE_MATCH                      = "If-None-Match"
	OH_COPY_OBJECT_SOURCE                            = "x-oss-copy-source"
	OH_COPY_SOURCE_RANGE                             = "x-oss-copy-source-range"
	OH_COPY_OBJECT_SOURCE_IF_MATCH                   = "x-oss-copy-source-if-match"
	OH_COPY_OBJECT_SOURCE_IF_NONE_MATCH              = "x-oss-copy-source-if-none-match"
	OH_COPY_OBJECT_SOURCE_IF_UNMODIFIED_SINCE        = "x-oss-copy-source-if-unmodified-since"
	OH_COPY_OBJECT_SOURCE_IF_MODIFIED_SINCE          = "x-oss-copy-source-if-modified-since"
	OH_COPY_OBJECT_METADATA_DIRECTIVE                = "x-oss-metadata-directive"
	OH_OSS_HEADER_REQUEST_ID                         = "x-oss-request-id"
	OH_ORIGIN                                        = "Origin"
	OH_ACCESS_CONTROL_REQUEST_METHOD                 = "Access-Control-Request-Method"
	OH_ACCESS_CONTROL_REQUEST_HEADER                 = "Access-Control-Request-Headers"
	OH_ACCESS_CONTROL_ALLOW_ORIGIN                   = "Access-Control-Allow-Origin"
	OH_ACCESS_CONTROL_ALLOW_METHODS                  = "Access-Control-Allow-Methods"
	OH_ACCESS_CONTROL_ALLOW_HEADERS                  = "Access-Control-Allow-Headers"
	OH_ACCESS_CONTROL_EXPOSE_HEADERS                 = "Access-Control-Expose-Headers"
	OH_ACCESS_CONTROL_MAX_AGE                        = "Access-Control-Max-Age"
	OH_OSS_HASH_CRC64ECMA                            = "x-oss-hash-ecr64ecma"
	OH_OSS_NEXT_APPEND_POSITION                      = "x-oss-next-append-position"
)
