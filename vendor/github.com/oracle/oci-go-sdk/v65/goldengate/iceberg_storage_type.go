// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// IcebergStorageTypeEnum Enum with underlying type: string
type IcebergStorageTypeEnum string

// Set of constants representing the allowable values for IcebergStorageTypeEnum
const (
	IcebergStorageTypeAmazonS3             IcebergStorageTypeEnum = "AMAZON_S3"
	IcebergStorageTypeGoogleCloudStorage   IcebergStorageTypeEnum = "GOOGLE_CLOUD_STORAGE"
	IcebergStorageTypeAzureDataLakeStorage IcebergStorageTypeEnum = "AZURE_DATA_LAKE_STORAGE"
)

var mappingIcebergStorageTypeEnum = map[string]IcebergStorageTypeEnum{
	"AMAZON_S3":               IcebergStorageTypeAmazonS3,
	"GOOGLE_CLOUD_STORAGE":    IcebergStorageTypeGoogleCloudStorage,
	"AZURE_DATA_LAKE_STORAGE": IcebergStorageTypeAzureDataLakeStorage,
}

var mappingIcebergStorageTypeEnumLowerCase = map[string]IcebergStorageTypeEnum{
	"amazon_s3":               IcebergStorageTypeAmazonS3,
	"google_cloud_storage":    IcebergStorageTypeGoogleCloudStorage,
	"azure_data_lake_storage": IcebergStorageTypeAzureDataLakeStorage,
}

// GetIcebergStorageTypeEnumValues Enumerates the set of values for IcebergStorageTypeEnum
func GetIcebergStorageTypeEnumValues() []IcebergStorageTypeEnum {
	values := make([]IcebergStorageTypeEnum, 0)
	for _, v := range mappingIcebergStorageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIcebergStorageTypeEnumStringValues Enumerates the set of values in String for IcebergStorageTypeEnum
func GetIcebergStorageTypeEnumStringValues() []string {
	return []string{
		"AMAZON_S3",
		"GOOGLE_CLOUD_STORAGE",
		"AZURE_DATA_LAKE_STORAGE",
	}
}

// GetMappingIcebergStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIcebergStorageTypeEnum(val string) (IcebergStorageTypeEnum, bool) {
	enum, ok := mappingIcebergStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
