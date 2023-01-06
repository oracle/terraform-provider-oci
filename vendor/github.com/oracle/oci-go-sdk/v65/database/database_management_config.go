// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseManagementConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseManagementConfigDatabaseManagementStatusEnum(string(m.DatabaseManagementStatus)); !ok && m.DatabaseManagementStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseManagementStatus: %s. Supported values are: %s.", m.DatabaseManagementStatus, strings.Join(GetDatabaseManagementConfigDatabaseManagementStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseManagementConfigLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetDatabaseManagementConfigLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingDatabaseManagementConfigDatabaseManagementStatusEnum = map[string]DatabaseManagementConfigDatabaseManagementStatusEnum{
	"ENABLING":         DatabaseManagementConfigDatabaseManagementStatusEnabling,
	"ENABLED":          DatabaseManagementConfigDatabaseManagementStatusEnabled,
	"DISABLING":        DatabaseManagementConfigDatabaseManagementStatusDisabling,
	"NOT_ENABLED":      DatabaseManagementConfigDatabaseManagementStatusNotEnabled,
	"FAILED_ENABLING":  DatabaseManagementConfigDatabaseManagementStatusFailedEnabling,
	"FAILED_DISABLING": DatabaseManagementConfigDatabaseManagementStatusFailedDisabling,
}

var mappingDatabaseManagementConfigDatabaseManagementStatusEnumLowerCase = map[string]DatabaseManagementConfigDatabaseManagementStatusEnum{
	"enabling":         DatabaseManagementConfigDatabaseManagementStatusEnabling,
	"enabled":          DatabaseManagementConfigDatabaseManagementStatusEnabled,
	"disabling":        DatabaseManagementConfigDatabaseManagementStatusDisabling,
	"not_enabled":      DatabaseManagementConfigDatabaseManagementStatusNotEnabled,
	"failed_enabling":  DatabaseManagementConfigDatabaseManagementStatusFailedEnabling,
	"failed_disabling": DatabaseManagementConfigDatabaseManagementStatusFailedDisabling,
}

// GetDatabaseManagementConfigDatabaseManagementStatusEnumValues Enumerates the set of values for DatabaseManagementConfigDatabaseManagementStatusEnum
func GetDatabaseManagementConfigDatabaseManagementStatusEnumValues() []DatabaseManagementConfigDatabaseManagementStatusEnum {
	values := make([]DatabaseManagementConfigDatabaseManagementStatusEnum, 0)
	for _, v := range mappingDatabaseManagementConfigDatabaseManagementStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseManagementConfigDatabaseManagementStatusEnumStringValues Enumerates the set of values in String for DatabaseManagementConfigDatabaseManagementStatusEnum
func GetDatabaseManagementConfigDatabaseManagementStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"NOT_ENABLED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
	}
}

// GetMappingDatabaseManagementConfigDatabaseManagementStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseManagementConfigDatabaseManagementStatusEnum(val string) (DatabaseManagementConfigDatabaseManagementStatusEnum, bool) {
	enum, ok := mappingDatabaseManagementConfigDatabaseManagementStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseManagementConfigLicenseModelEnum Enum with underlying type: string
type DatabaseManagementConfigLicenseModelEnum string

// Set of constants representing the allowable values for DatabaseManagementConfigLicenseModelEnum
const (
	DatabaseManagementConfigLicenseModelLicenseIncluded     DatabaseManagementConfigLicenseModelEnum = "LICENSE_INCLUDED"
	DatabaseManagementConfigLicenseModelBringYourOwnLicense DatabaseManagementConfigLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingDatabaseManagementConfigLicenseModelEnum = map[string]DatabaseManagementConfigLicenseModelEnum{
	"LICENSE_INCLUDED":       DatabaseManagementConfigLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": DatabaseManagementConfigLicenseModelBringYourOwnLicense,
}

var mappingDatabaseManagementConfigLicenseModelEnumLowerCase = map[string]DatabaseManagementConfigLicenseModelEnum{
	"license_included":       DatabaseManagementConfigLicenseModelLicenseIncluded,
	"bring_your_own_license": DatabaseManagementConfigLicenseModelBringYourOwnLicense,
}

// GetDatabaseManagementConfigLicenseModelEnumValues Enumerates the set of values for DatabaseManagementConfigLicenseModelEnum
func GetDatabaseManagementConfigLicenseModelEnumValues() []DatabaseManagementConfigLicenseModelEnum {
	values := make([]DatabaseManagementConfigLicenseModelEnum, 0)
	for _, v := range mappingDatabaseManagementConfigLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseManagementConfigLicenseModelEnumStringValues Enumerates the set of values in String for DatabaseManagementConfigLicenseModelEnum
func GetDatabaseManagementConfigLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingDatabaseManagementConfigLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseManagementConfigLicenseModelEnum(val string) (DatabaseManagementConfigLicenseModelEnum, bool) {
	enum, ok := mappingDatabaseManagementConfigLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
