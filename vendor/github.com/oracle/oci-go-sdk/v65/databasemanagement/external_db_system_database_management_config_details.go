// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalDbSystemDatabaseManagementConfigDetails The details required to enable Database Management for an external DB system.
type ExternalDbSystemDatabaseManagementConfigDetails struct {

	// The Oracle license model that applies to the external database.
	LicenseModel ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

func (m ExternalDbSystemDatabaseManagementConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDbSystemDatabaseManagementConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum Enum with underlying type: string
type ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum string

// Set of constants representing the allowable values for ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum
const (
	ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelLicenseIncluded     ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelBringYourOwnLicense ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum = map[string]ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelBringYourOwnLicense,
}

var mappingExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnumLowerCase = map[string]ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum{
	"license_included":       ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelBringYourOwnLicense,
}

// GetExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnumValues Enumerates the set of values for ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum
func GetExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnumValues() []ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum {
	values := make([]ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum, 0)
	for _, v := range mappingExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnumStringValues Enumerates the set of values in String for ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum
func GetExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum(val string) (ExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnum, bool) {
	enum, ok := mappingExternalDbSystemDatabaseManagementConfigDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
