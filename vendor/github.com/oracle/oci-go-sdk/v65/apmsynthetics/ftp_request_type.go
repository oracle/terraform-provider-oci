// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// FtpRequestTypeEnum Enum with underlying type: string
type FtpRequestTypeEnum string

// Set of constants representing the allowable values for FtpRequestTypeEnum
const (
	FtpRequestTypeList     FtpRequestTypeEnum = "LIST"
	FtpRequestTypeUpload   FtpRequestTypeEnum = "UPLOAD"
	FtpRequestTypeDownload FtpRequestTypeEnum = "DOWNLOAD"
)

var mappingFtpRequestTypeEnum = map[string]FtpRequestTypeEnum{
	"LIST":     FtpRequestTypeList,
	"UPLOAD":   FtpRequestTypeUpload,
	"DOWNLOAD": FtpRequestTypeDownload,
}

var mappingFtpRequestTypeEnumLowerCase = map[string]FtpRequestTypeEnum{
	"list":     FtpRequestTypeList,
	"upload":   FtpRequestTypeUpload,
	"download": FtpRequestTypeDownload,
}

// GetFtpRequestTypeEnumValues Enumerates the set of values for FtpRequestTypeEnum
func GetFtpRequestTypeEnumValues() []FtpRequestTypeEnum {
	values := make([]FtpRequestTypeEnum, 0)
	for _, v := range mappingFtpRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFtpRequestTypeEnumStringValues Enumerates the set of values in String for FtpRequestTypeEnum
func GetFtpRequestTypeEnumStringValues() []string {
	return []string{
		"LIST",
		"UPLOAD",
		"DOWNLOAD",
	}
}

// GetMappingFtpRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFtpRequestTypeEnum(val string) (FtpRequestTypeEnum, bool) {
	enum, ok := mappingFtpRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
