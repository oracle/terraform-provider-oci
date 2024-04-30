// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnableExternalDbSystemDatabaseManagementDetails The details required to enable Database Management for an external DB system.
type EnableExternalDbSystemDatabaseManagementDetails struct {

	// The Oracle license model that applies to the external database.
	LicenseModel EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

func (m EnableExternalDbSystemDatabaseManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableExternalDbSystemDatabaseManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum Enum with underlying type: string
type EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum string

// Set of constants representing the allowable values for EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum
const (
	EnableExternalDbSystemDatabaseManagementDetailsLicenseModelLicenseIncluded     EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	EnableExternalDbSystemDatabaseManagementDetailsLicenseModelBringYourOwnLicense EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum = map[string]EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       EnableExternalDbSystemDatabaseManagementDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": EnableExternalDbSystemDatabaseManagementDetailsLicenseModelBringYourOwnLicense,
}

var mappingEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnumLowerCase = map[string]EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum{
	"license_included":       EnableExternalDbSystemDatabaseManagementDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": EnableExternalDbSystemDatabaseManagementDetailsLicenseModelBringYourOwnLicense,
}

// GetEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnumValues Enumerates the set of values for EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum
func GetEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnumValues() []EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum {
	values := make([]EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum, 0)
	for _, v := range mappingEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnumStringValues Enumerates the set of values in String for EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum
func GetEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum(val string) (EnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnum, bool) {
	enum, ok := mappingEnableExternalDbSystemDatabaseManagementDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
