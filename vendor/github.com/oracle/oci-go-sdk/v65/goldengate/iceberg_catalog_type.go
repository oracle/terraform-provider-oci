// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// IcebergCatalogTypeEnum Enum with underlying type: string
type IcebergCatalogTypeEnum string

// Set of constants representing the allowable values for IcebergCatalogTypeEnum
const (
	IcebergCatalogTypeGlue    IcebergCatalogTypeEnum = "GLUE"
	IcebergCatalogTypeHadoop  IcebergCatalogTypeEnum = "HADOOP"
	IcebergCatalogTypeNessie  IcebergCatalogTypeEnum = "NESSIE"
	IcebergCatalogTypePolaris IcebergCatalogTypeEnum = "POLARIS"
	IcebergCatalogTypeRest    IcebergCatalogTypeEnum = "REST"
)

var mappingIcebergCatalogTypeEnum = map[string]IcebergCatalogTypeEnum{
	"GLUE":    IcebergCatalogTypeGlue,
	"HADOOP":  IcebergCatalogTypeHadoop,
	"NESSIE":  IcebergCatalogTypeNessie,
	"POLARIS": IcebergCatalogTypePolaris,
	"REST":    IcebergCatalogTypeRest,
}

var mappingIcebergCatalogTypeEnumLowerCase = map[string]IcebergCatalogTypeEnum{
	"glue":    IcebergCatalogTypeGlue,
	"hadoop":  IcebergCatalogTypeHadoop,
	"nessie":  IcebergCatalogTypeNessie,
	"polaris": IcebergCatalogTypePolaris,
	"rest":    IcebergCatalogTypeRest,
}

// GetIcebergCatalogTypeEnumValues Enumerates the set of values for IcebergCatalogTypeEnum
func GetIcebergCatalogTypeEnumValues() []IcebergCatalogTypeEnum {
	values := make([]IcebergCatalogTypeEnum, 0)
	for _, v := range mappingIcebergCatalogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIcebergCatalogTypeEnumStringValues Enumerates the set of values in String for IcebergCatalogTypeEnum
func GetIcebergCatalogTypeEnumStringValues() []string {
	return []string{
		"GLUE",
		"HADOOP",
		"NESSIE",
		"POLARIS",
		"REST",
	}
}

// GetMappingIcebergCatalogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIcebergCatalogTypeEnum(val string) (IcebergCatalogTypeEnum, bool) {
	enum, ok := mappingIcebergCatalogTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
