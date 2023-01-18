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

// EnableExternalContainerDatabaseDatabaseManagementDetails Details to enable Database Management on an external container database.
type EnableExternalContainerDatabaseDatabaseManagementDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// CreateExternalDatabaseConnectorDetails.
	ExternalDatabaseConnectorId *string `mandatory:"true" json:"externalDatabaseConnectorId"`

	// The Oracle license model that applies to the external database.
	LicenseModel EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

func (m EnableExternalContainerDatabaseDatabaseManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableExternalContainerDatabaseDatabaseManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum Enum with underlying type: string
type EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum string

// Set of constants representing the allowable values for EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum
const (
	EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelLicenseIncluded     EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelBringYourOwnLicense EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum = map[string]EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelBringYourOwnLicense,
}

var mappingEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnumLowerCase = map[string]EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum{
	"license_included":       EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelBringYourOwnLicense,
}

// GetEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnumValues Enumerates the set of values for EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum
func GetEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnumValues() []EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum {
	values := make([]EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum, 0)
	for _, v := range mappingEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnumStringValues Enumerates the set of values in String for EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum
func GetEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum(val string) (EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum, bool) {
	enum, ok := mappingEnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
