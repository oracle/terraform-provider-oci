// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudDatabaseManagementConfig The configuration of the Database Management service.
type CloudDatabaseManagementConfig struct {

	// The status of the Database Management service.
	ManagementStatus CloudDatabaseManagementConfigManagementStatusEnum `mandatory:"true" json:"managementStatus"`

	// The Database Management type.
	ManagementType CloudDatabaseManagementConfigManagementTypeEnum `mandatory:"true" json:"managementType"`
}

func (m CloudDatabaseManagementConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudDatabaseManagementConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudDatabaseManagementConfigManagementStatusEnum(string(m.ManagementStatus)); !ok && m.ManagementStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementStatus: %s. Supported values are: %s.", m.ManagementStatus, strings.Join(GetCloudDatabaseManagementConfigManagementStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudDatabaseManagementConfigManagementTypeEnum(string(m.ManagementType)); !ok && m.ManagementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementType: %s. Supported values are: %s.", m.ManagementType, strings.Join(GetCloudDatabaseManagementConfigManagementTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudDatabaseManagementConfigManagementStatusEnum Enum with underlying type: string
type CloudDatabaseManagementConfigManagementStatusEnum string

// Set of constants representing the allowable values for CloudDatabaseManagementConfigManagementStatusEnum
const (
	CloudDatabaseManagementConfigManagementStatusEnabling        CloudDatabaseManagementConfigManagementStatusEnum = "ENABLING"
	CloudDatabaseManagementConfigManagementStatusEnabled         CloudDatabaseManagementConfigManagementStatusEnum = "ENABLED"
	CloudDatabaseManagementConfigManagementStatusDisabling       CloudDatabaseManagementConfigManagementStatusEnum = "DISABLING"
	CloudDatabaseManagementConfigManagementStatusDisabled        CloudDatabaseManagementConfigManagementStatusEnum = "DISABLED"
	CloudDatabaseManagementConfigManagementStatusUpdating        CloudDatabaseManagementConfigManagementStatusEnum = "UPDATING"
	CloudDatabaseManagementConfigManagementStatusFailedEnabling  CloudDatabaseManagementConfigManagementStatusEnum = "FAILED_ENABLING"
	CloudDatabaseManagementConfigManagementStatusFailedDisabling CloudDatabaseManagementConfigManagementStatusEnum = "FAILED_DISABLING"
	CloudDatabaseManagementConfigManagementStatusFailedUpdating  CloudDatabaseManagementConfigManagementStatusEnum = "FAILED_UPDATING"
)

var mappingCloudDatabaseManagementConfigManagementStatusEnum = map[string]CloudDatabaseManagementConfigManagementStatusEnum{
	"ENABLING":         CloudDatabaseManagementConfigManagementStatusEnabling,
	"ENABLED":          CloudDatabaseManagementConfigManagementStatusEnabled,
	"DISABLING":        CloudDatabaseManagementConfigManagementStatusDisabling,
	"DISABLED":         CloudDatabaseManagementConfigManagementStatusDisabled,
	"UPDATING":         CloudDatabaseManagementConfigManagementStatusUpdating,
	"FAILED_ENABLING":  CloudDatabaseManagementConfigManagementStatusFailedEnabling,
	"FAILED_DISABLING": CloudDatabaseManagementConfigManagementStatusFailedDisabling,
	"FAILED_UPDATING":  CloudDatabaseManagementConfigManagementStatusFailedUpdating,
}

var mappingCloudDatabaseManagementConfigManagementStatusEnumLowerCase = map[string]CloudDatabaseManagementConfigManagementStatusEnum{
	"enabling":         CloudDatabaseManagementConfigManagementStatusEnabling,
	"enabled":          CloudDatabaseManagementConfigManagementStatusEnabled,
	"disabling":        CloudDatabaseManagementConfigManagementStatusDisabling,
	"disabled":         CloudDatabaseManagementConfigManagementStatusDisabled,
	"updating":         CloudDatabaseManagementConfigManagementStatusUpdating,
	"failed_enabling":  CloudDatabaseManagementConfigManagementStatusFailedEnabling,
	"failed_disabling": CloudDatabaseManagementConfigManagementStatusFailedDisabling,
	"failed_updating":  CloudDatabaseManagementConfigManagementStatusFailedUpdating,
}

// GetCloudDatabaseManagementConfigManagementStatusEnumValues Enumerates the set of values for CloudDatabaseManagementConfigManagementStatusEnum
func GetCloudDatabaseManagementConfigManagementStatusEnumValues() []CloudDatabaseManagementConfigManagementStatusEnum {
	values := make([]CloudDatabaseManagementConfigManagementStatusEnum, 0)
	for _, v := range mappingCloudDatabaseManagementConfigManagementStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDatabaseManagementConfigManagementStatusEnumStringValues Enumerates the set of values in String for CloudDatabaseManagementConfigManagementStatusEnum
func GetCloudDatabaseManagementConfigManagementStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"DISABLED",
		"UPDATING",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
		"FAILED_UPDATING",
	}
}

// GetMappingCloudDatabaseManagementConfigManagementStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDatabaseManagementConfigManagementStatusEnum(val string) (CloudDatabaseManagementConfigManagementStatusEnum, bool) {
	enum, ok := mappingCloudDatabaseManagementConfigManagementStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudDatabaseManagementConfigManagementTypeEnum Enum with underlying type: string
type CloudDatabaseManagementConfigManagementTypeEnum string

// Set of constants representing the allowable values for CloudDatabaseManagementConfigManagementTypeEnum
const (
	CloudDatabaseManagementConfigManagementTypeBasic    CloudDatabaseManagementConfigManagementTypeEnum = "BASIC"
	CloudDatabaseManagementConfigManagementTypeAdvanced CloudDatabaseManagementConfigManagementTypeEnum = "ADVANCED"
)

var mappingCloudDatabaseManagementConfigManagementTypeEnum = map[string]CloudDatabaseManagementConfigManagementTypeEnum{
	"BASIC":    CloudDatabaseManagementConfigManagementTypeBasic,
	"ADVANCED": CloudDatabaseManagementConfigManagementTypeAdvanced,
}

var mappingCloudDatabaseManagementConfigManagementTypeEnumLowerCase = map[string]CloudDatabaseManagementConfigManagementTypeEnum{
	"basic":    CloudDatabaseManagementConfigManagementTypeBasic,
	"advanced": CloudDatabaseManagementConfigManagementTypeAdvanced,
}

// GetCloudDatabaseManagementConfigManagementTypeEnumValues Enumerates the set of values for CloudDatabaseManagementConfigManagementTypeEnum
func GetCloudDatabaseManagementConfigManagementTypeEnumValues() []CloudDatabaseManagementConfigManagementTypeEnum {
	values := make([]CloudDatabaseManagementConfigManagementTypeEnum, 0)
	for _, v := range mappingCloudDatabaseManagementConfigManagementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDatabaseManagementConfigManagementTypeEnumStringValues Enumerates the set of values in String for CloudDatabaseManagementConfigManagementTypeEnum
func GetCloudDatabaseManagementConfigManagementTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
	}
}

// GetMappingCloudDatabaseManagementConfigManagementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDatabaseManagementConfigManagementTypeEnum(val string) (CloudDatabaseManagementConfigManagementTypeEnum, bool) {
	enum, ok := mappingCloudDatabaseManagementConfigManagementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
