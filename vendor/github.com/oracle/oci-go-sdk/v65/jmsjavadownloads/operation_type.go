// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the download engine of the Java Management Service.
//

package jmsjavadownloads

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateJavaDownloadToken  OperationTypeEnum = "CREATE_JAVA_DOWNLOAD_TOKEN"
	OperationTypeUpdateJavaDownloadToken  OperationTypeEnum = "UPDATE_JAVA_DOWNLOAD_TOKEN"
	OperationTypeDeleteJavaDownloadToken  OperationTypeEnum = "DELETE_JAVA_DOWNLOAD_TOKEN"
	OperationTypeCreateJavaDownloadReport OperationTypeEnum = "CREATE_JAVA_DOWNLOAD_REPORT"
	OperationTypeDeleteJavaDownloadReport OperationTypeEnum = "DELETE_JAVA_DOWNLOAD_REPORT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_JAVA_DOWNLOAD_TOKEN":  OperationTypeCreateJavaDownloadToken,
	"UPDATE_JAVA_DOWNLOAD_TOKEN":  OperationTypeUpdateJavaDownloadToken,
	"DELETE_JAVA_DOWNLOAD_TOKEN":  OperationTypeDeleteJavaDownloadToken,
	"CREATE_JAVA_DOWNLOAD_REPORT": OperationTypeCreateJavaDownloadReport,
	"DELETE_JAVA_DOWNLOAD_REPORT": OperationTypeDeleteJavaDownloadReport,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_java_download_token":  OperationTypeCreateJavaDownloadToken,
	"update_java_download_token":  OperationTypeUpdateJavaDownloadToken,
	"delete_java_download_token":  OperationTypeDeleteJavaDownloadToken,
	"create_java_download_report": OperationTypeCreateJavaDownloadReport,
	"delete_java_download_report": OperationTypeDeleteJavaDownloadReport,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_JAVA_DOWNLOAD_TOKEN",
		"UPDATE_JAVA_DOWNLOAD_TOKEN",
		"DELETE_JAVA_DOWNLOAD_TOKEN",
		"CREATE_JAVA_DOWNLOAD_REPORT",
		"DELETE_JAVA_DOWNLOAD_REPORT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
