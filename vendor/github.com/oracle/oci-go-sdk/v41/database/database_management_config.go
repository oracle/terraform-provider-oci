// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// DatabaseManagementConfig The configuration of the Database Management service.
type DatabaseManagementConfig struct {

	// The status of the Database Management service.
	DatabaseManagementStatus DatabaseManagementConfigDatabaseManagementStatusEnum `mandatory:"true" json:"databaseManagementStatus"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// CreateExternalDatabaseConnectorDetails.
	DatabaseManagementConnectionId *string `mandatory:"false" json:"databaseManagementConnectionId"`

	// The Oracle license model that applies to the external database.
	LicenseModel DatabaseManagementConfigLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
}

func (m DatabaseManagementConfig) String() string {
	return common.PointerString(m)
}

// DatabaseManagementConfigDatabaseManagementStatusEnum Enum with underlying type: string
type DatabaseManagementConfigDatabaseManagementStatusEnum string

// Set of constants representing the allowable values for DatabaseManagementConfigDatabaseManagementStatusEnum
const (
	DatabaseManagementConfigDatabaseManagementStatusEnabling        DatabaseManagementConfigDatabaseManagementStatusEnum = "ENABLING"
	DatabaseManagementConfigDatabaseManagementStatusEnabled         DatabaseManagementConfigDatabaseManagementStatusEnum = "ENABLED"
	DatabaseManagementConfigDatabaseManagementStatusDisabling       DatabaseManagementConfigDatabaseManagementStatusEnum = "DISABLING"
	DatabaseManagementConfigDatabaseManagementStatusNotEnabled      DatabaseManagementConfigDatabaseManagementStatusEnum = "NOT_ENABLED"
	DatabaseManagementConfigDatabaseManagementStatusFailedEnabling  DatabaseManagementConfigDatabaseManagementStatusEnum = "FAILED_ENABLING"
	DatabaseManagementConfigDatabaseManagementStatusFailedDisabling DatabaseManagementConfigDatabaseManagementStatusEnum = "FAILED_DISABLING"
)

var mappingDatabaseManagementConfigDatabaseManagementStatus = map[string]DatabaseManagementConfigDatabaseManagementStatusEnum{
	"ENABLING":         DatabaseManagementConfigDatabaseManagementStatusEnabling,
	"ENABLED":          DatabaseManagementConfigDatabaseManagementStatusEnabled,
	"DISABLING":        DatabaseManagementConfigDatabaseManagementStatusDisabling,
	"NOT_ENABLED":      DatabaseManagementConfigDatabaseManagementStatusNotEnabled,
	"FAILED_ENABLING":  DatabaseManagementConfigDatabaseManagementStatusFailedEnabling,
	"FAILED_DISABLING": DatabaseManagementConfigDatabaseManagementStatusFailedDisabling,
}

// GetDatabaseManagementConfigDatabaseManagementStatusEnumValues Enumerates the set of values for DatabaseManagementConfigDatabaseManagementStatusEnum
func GetDatabaseManagementConfigDatabaseManagementStatusEnumValues() []DatabaseManagementConfigDatabaseManagementStatusEnum {
	values := make([]DatabaseManagementConfigDatabaseManagementStatusEnum, 0)
	for _, v := range mappingDatabaseManagementConfigDatabaseManagementStatus {
		values = append(values, v)
	}
	return values
}

// DatabaseManagementConfigLicenseModelEnum Enum with underlying type: string
type DatabaseManagementConfigLicenseModelEnum string

// Set of constants representing the allowable values for DatabaseManagementConfigLicenseModelEnum
const (
	DatabaseManagementConfigLicenseModelLicenseIncluded     DatabaseManagementConfigLicenseModelEnum = "LICENSE_INCLUDED"
	DatabaseManagementConfigLicenseModelBringYourOwnLicense DatabaseManagementConfigLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingDatabaseManagementConfigLicenseModel = map[string]DatabaseManagementConfigLicenseModelEnum{
	"LICENSE_INCLUDED":       DatabaseManagementConfigLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": DatabaseManagementConfigLicenseModelBringYourOwnLicense,
}

// GetDatabaseManagementConfigLicenseModelEnumValues Enumerates the set of values for DatabaseManagementConfigLicenseModelEnum
func GetDatabaseManagementConfigLicenseModelEnumValues() []DatabaseManagementConfigLicenseModelEnum {
	values := make([]DatabaseManagementConfigLicenseModelEnum, 0)
	for _, v := range mappingDatabaseManagementConfigLicenseModel {
		values = append(values, v)
	}
	return values
}
