// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmDataCenter An OLVM Data Center.
type OlvmDataCenter struct {

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	// A human-readable description in plain text.
	Description *string `mandatory:"false" json:"description"`

	// A unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// A human-readable name in plain text.
	Name *string `mandatory:"false" json:"name"`

	// Whether the data center is local.
	IsLocal *bool `mandatory:"false" json:"isLocal"`

	// The type of quota mode
	QuotaModeType OlvmDataCenterQuotaModeTypeEnum `mandatory:"false" json:"quotaModeType,omitempty"`

	// The status of data center
	DataCenterStatus OlvmDataCenterDataCenterStatusEnum `mandatory:"false" json:"dataCenterStatus,omitempty"`

	// Type which represents a format of storage domain
	StorageFormat OlvmDataCenterStorageFormatEnum `mandatory:"false" json:"storageFormat,omitempty"`

	// List of supported versions.
	SupportedVersions []OlvmVersion `mandatory:"false" json:"supportedVersions"`

	Version *OlvmVersion `mandatory:"false" json:"version"`
}

func (m OlvmDataCenter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmDataCenter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmDataCenterQuotaModeTypeEnum(string(m.QuotaModeType)); !ok && m.QuotaModeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuotaModeType: %s. Supported values are: %s.", m.QuotaModeType, strings.Join(GetOlvmDataCenterQuotaModeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmDataCenterDataCenterStatusEnum(string(m.DataCenterStatus)); !ok && m.DataCenterStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataCenterStatus: %s. Supported values are: %s.", m.DataCenterStatus, strings.Join(GetOlvmDataCenterDataCenterStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmDataCenterStorageFormatEnum(string(m.StorageFormat)); !ok && m.StorageFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageFormat: %s. Supported values are: %s.", m.StorageFormat, strings.Join(GetOlvmDataCenterStorageFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmDataCenterQuotaModeTypeEnum Enum with underlying type: string
type OlvmDataCenterQuotaModeTypeEnum string

// Set of constants representing the allowable values for OlvmDataCenterQuotaModeTypeEnum
const (
	OlvmDataCenterQuotaModeTypeAudit    OlvmDataCenterQuotaModeTypeEnum = "AUDIT"
	OlvmDataCenterQuotaModeTypeDisabled OlvmDataCenterQuotaModeTypeEnum = "DISABLED"
	OlvmDataCenterQuotaModeTypeEnabled  OlvmDataCenterQuotaModeTypeEnum = "ENABLED"
)

var mappingOlvmDataCenterQuotaModeTypeEnum = map[string]OlvmDataCenterQuotaModeTypeEnum{
	"AUDIT":    OlvmDataCenterQuotaModeTypeAudit,
	"DISABLED": OlvmDataCenterQuotaModeTypeDisabled,
	"ENABLED":  OlvmDataCenterQuotaModeTypeEnabled,
}

var mappingOlvmDataCenterQuotaModeTypeEnumLowerCase = map[string]OlvmDataCenterQuotaModeTypeEnum{
	"audit":    OlvmDataCenterQuotaModeTypeAudit,
	"disabled": OlvmDataCenterQuotaModeTypeDisabled,
	"enabled":  OlvmDataCenterQuotaModeTypeEnabled,
}

// GetOlvmDataCenterQuotaModeTypeEnumValues Enumerates the set of values for OlvmDataCenterQuotaModeTypeEnum
func GetOlvmDataCenterQuotaModeTypeEnumValues() []OlvmDataCenterQuotaModeTypeEnum {
	values := make([]OlvmDataCenterQuotaModeTypeEnum, 0)
	for _, v := range mappingOlvmDataCenterQuotaModeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmDataCenterQuotaModeTypeEnumStringValues Enumerates the set of values in String for OlvmDataCenterQuotaModeTypeEnum
func GetOlvmDataCenterQuotaModeTypeEnumStringValues() []string {
	return []string{
		"AUDIT",
		"DISABLED",
		"ENABLED",
	}
}

// GetMappingOlvmDataCenterQuotaModeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmDataCenterQuotaModeTypeEnum(val string) (OlvmDataCenterQuotaModeTypeEnum, bool) {
	enum, ok := mappingOlvmDataCenterQuotaModeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmDataCenterDataCenterStatusEnum Enum with underlying type: string
type OlvmDataCenterDataCenterStatusEnum string

// Set of constants representing the allowable values for OlvmDataCenterDataCenterStatusEnum
const (
	OlvmDataCenterDataCenterStatusContend        OlvmDataCenterDataCenterStatusEnum = "CONTEND"
	OlvmDataCenterDataCenterStatusMaintenance    OlvmDataCenterDataCenterStatusEnum = "MAINTENANCE"
	OlvmDataCenterDataCenterStatusNotOperational OlvmDataCenterDataCenterStatusEnum = "NOT_OPERATIONAL"
	OlvmDataCenterDataCenterStatusProblematic    OlvmDataCenterDataCenterStatusEnum = "PROBLEMATIC"
	OlvmDataCenterDataCenterStatusUninitialized  OlvmDataCenterDataCenterStatusEnum = "UNINITIALIZED"
	OlvmDataCenterDataCenterStatusUp             OlvmDataCenterDataCenterStatusEnum = "UP"
)

var mappingOlvmDataCenterDataCenterStatusEnum = map[string]OlvmDataCenterDataCenterStatusEnum{
	"CONTEND":         OlvmDataCenterDataCenterStatusContend,
	"MAINTENANCE":     OlvmDataCenterDataCenterStatusMaintenance,
	"NOT_OPERATIONAL": OlvmDataCenterDataCenterStatusNotOperational,
	"PROBLEMATIC":     OlvmDataCenterDataCenterStatusProblematic,
	"UNINITIALIZED":   OlvmDataCenterDataCenterStatusUninitialized,
	"UP":              OlvmDataCenterDataCenterStatusUp,
}

var mappingOlvmDataCenterDataCenterStatusEnumLowerCase = map[string]OlvmDataCenterDataCenterStatusEnum{
	"contend":         OlvmDataCenterDataCenterStatusContend,
	"maintenance":     OlvmDataCenterDataCenterStatusMaintenance,
	"not_operational": OlvmDataCenterDataCenterStatusNotOperational,
	"problematic":     OlvmDataCenterDataCenterStatusProblematic,
	"uninitialized":   OlvmDataCenterDataCenterStatusUninitialized,
	"up":              OlvmDataCenterDataCenterStatusUp,
}

// GetOlvmDataCenterDataCenterStatusEnumValues Enumerates the set of values for OlvmDataCenterDataCenterStatusEnum
func GetOlvmDataCenterDataCenterStatusEnumValues() []OlvmDataCenterDataCenterStatusEnum {
	values := make([]OlvmDataCenterDataCenterStatusEnum, 0)
	for _, v := range mappingOlvmDataCenterDataCenterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmDataCenterDataCenterStatusEnumStringValues Enumerates the set of values in String for OlvmDataCenterDataCenterStatusEnum
func GetOlvmDataCenterDataCenterStatusEnumStringValues() []string {
	return []string{
		"CONTEND",
		"MAINTENANCE",
		"NOT_OPERATIONAL",
		"PROBLEMATIC",
		"UNINITIALIZED",
		"UP",
	}
}

// GetMappingOlvmDataCenterDataCenterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmDataCenterDataCenterStatusEnum(val string) (OlvmDataCenterDataCenterStatusEnum, bool) {
	enum, ok := mappingOlvmDataCenterDataCenterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmDataCenterStorageFormatEnum Enum with underlying type: string
type OlvmDataCenterStorageFormatEnum string

// Set of constants representing the allowable values for OlvmDataCenterStorageFormatEnum
const (
	OlvmDataCenterStorageFormatV1 OlvmDataCenterStorageFormatEnum = "V1"
	OlvmDataCenterStorageFormatV2 OlvmDataCenterStorageFormatEnum = "V2"
	OlvmDataCenterStorageFormatV3 OlvmDataCenterStorageFormatEnum = "V3"
	OlvmDataCenterStorageFormatV4 OlvmDataCenterStorageFormatEnum = "V4"
	OlvmDataCenterStorageFormatV5 OlvmDataCenterStorageFormatEnum = "V5"
)

var mappingOlvmDataCenterStorageFormatEnum = map[string]OlvmDataCenterStorageFormatEnum{
	"V1": OlvmDataCenterStorageFormatV1,
	"V2": OlvmDataCenterStorageFormatV2,
	"V3": OlvmDataCenterStorageFormatV3,
	"V4": OlvmDataCenterStorageFormatV4,
	"V5": OlvmDataCenterStorageFormatV5,
}

var mappingOlvmDataCenterStorageFormatEnumLowerCase = map[string]OlvmDataCenterStorageFormatEnum{
	"v1": OlvmDataCenterStorageFormatV1,
	"v2": OlvmDataCenterStorageFormatV2,
	"v3": OlvmDataCenterStorageFormatV3,
	"v4": OlvmDataCenterStorageFormatV4,
	"v5": OlvmDataCenterStorageFormatV5,
}

// GetOlvmDataCenterStorageFormatEnumValues Enumerates the set of values for OlvmDataCenterStorageFormatEnum
func GetOlvmDataCenterStorageFormatEnumValues() []OlvmDataCenterStorageFormatEnum {
	values := make([]OlvmDataCenterStorageFormatEnum, 0)
	for _, v := range mappingOlvmDataCenterStorageFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmDataCenterStorageFormatEnumStringValues Enumerates the set of values in String for OlvmDataCenterStorageFormatEnum
func GetOlvmDataCenterStorageFormatEnumStringValues() []string {
	return []string{
		"V1",
		"V2",
		"V3",
		"V4",
		"V5",
	}
}

// GetMappingOlvmDataCenterStorageFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmDataCenterStorageFormatEnum(val string) (OlvmDataCenterStorageFormatEnum, bool) {
	enum, ok := mappingOlvmDataCenterStorageFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
