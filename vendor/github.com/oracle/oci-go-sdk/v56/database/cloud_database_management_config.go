// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingCloudDatabaseManagementConfigManagementStatus = map[string]CloudDatabaseManagementConfigManagementStatusEnum{
	"ENABLING":         CloudDatabaseManagementConfigManagementStatusEnabling,
	"ENABLED":          CloudDatabaseManagementConfigManagementStatusEnabled,
	"DISABLING":        CloudDatabaseManagementConfigManagementStatusDisabling,
	"DISABLED":         CloudDatabaseManagementConfigManagementStatusDisabled,
	"UPDATING":         CloudDatabaseManagementConfigManagementStatusUpdating,
	"FAILED_ENABLING":  CloudDatabaseManagementConfigManagementStatusFailedEnabling,
	"FAILED_DISABLING": CloudDatabaseManagementConfigManagementStatusFailedDisabling,
	"FAILED_UPDATING":  CloudDatabaseManagementConfigManagementStatusFailedUpdating,
}

// GetCloudDatabaseManagementConfigManagementStatusEnumValues Enumerates the set of values for CloudDatabaseManagementConfigManagementStatusEnum
func GetCloudDatabaseManagementConfigManagementStatusEnumValues() []CloudDatabaseManagementConfigManagementStatusEnum {
	values := make([]CloudDatabaseManagementConfigManagementStatusEnum, 0)
	for _, v := range mappingCloudDatabaseManagementConfigManagementStatus {
		values = append(values, v)
	}
	return values
}

// CloudDatabaseManagementConfigManagementTypeEnum Enum with underlying type: string
type CloudDatabaseManagementConfigManagementTypeEnum string

// Set of constants representing the allowable values for CloudDatabaseManagementConfigManagementTypeEnum
const (
	CloudDatabaseManagementConfigManagementTypeBasic    CloudDatabaseManagementConfigManagementTypeEnum = "BASIC"
	CloudDatabaseManagementConfigManagementTypeAdvanced CloudDatabaseManagementConfigManagementTypeEnum = "ADVANCED"
)

var mappingCloudDatabaseManagementConfigManagementType = map[string]CloudDatabaseManagementConfigManagementTypeEnum{
	"BASIC":    CloudDatabaseManagementConfigManagementTypeBasic,
	"ADVANCED": CloudDatabaseManagementConfigManagementTypeAdvanced,
}

// GetCloudDatabaseManagementConfigManagementTypeEnumValues Enumerates the set of values for CloudDatabaseManagementConfigManagementTypeEnum
func GetCloudDatabaseManagementConfigManagementTypeEnumValues() []CloudDatabaseManagementConfigManagementTypeEnum {
	values := make([]CloudDatabaseManagementConfigManagementTypeEnum, 0)
	for _, v := range mappingCloudDatabaseManagementConfigManagementType {
		values = append(values, v)
	}
	return values
}
