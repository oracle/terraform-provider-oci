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

// EnableExternalNonContainerDatabaseDatabaseManagementDetails Details to enable Database Management on an external non-container database.
type EnableExternalNonContainerDatabaseDatabaseManagementDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// CreateExternalDatabaseConnectorDetails.
	ExternalDatabaseConnectorId *string `mandatory:"true" json:"externalDatabaseConnectorId"`

	// The Oracle license model that applies to the external database.
	LicenseModel EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

func (m EnableExternalNonContainerDatabaseDatabaseManagementDetails) String() string {
	return common.PointerString(m)
}

// EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum Enum with underlying type: string
type EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum string

// Set of constants representing the allowable values for EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum
const (
	EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelLicenseIncluded     EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelBringYourOwnLicense EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingEnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModel = map[string]EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelBringYourOwnLicense,
}

// GetEnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnumValues Enumerates the set of values for EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum
func GetEnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnumValues() []EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum {
	values := make([]EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum, 0)
	for _, v := range mappingEnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
