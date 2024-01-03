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

// EnableExternalExadataInfrastructureManagementDetails The details required to enable Database Management on the Exadata infrastructure.
type EnableExternalExadataInfrastructureManagementDetails struct {

	// The Oracle license model.
	LicenseModel EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

func (m EnableExternalExadataInfrastructureManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableExternalExadataInfrastructureManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum Enum with underlying type: string
type EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum string

// Set of constants representing the allowable values for EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum
const (
	EnableExternalExadataInfrastructureManagementDetailsLicenseModelLicenseIncluded     EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	EnableExternalExadataInfrastructureManagementDetailsLicenseModelBringYourOwnLicense EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum = map[string]EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       EnableExternalExadataInfrastructureManagementDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": EnableExternalExadataInfrastructureManagementDetailsLicenseModelBringYourOwnLicense,
}

var mappingEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnumLowerCase = map[string]EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum{
	"license_included":       EnableExternalExadataInfrastructureManagementDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": EnableExternalExadataInfrastructureManagementDetailsLicenseModelBringYourOwnLicense,
}

// GetEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnumValues Enumerates the set of values for EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum
func GetEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnumValues() []EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum {
	values := make([]EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum, 0)
	for _, v := range mappingEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnumStringValues Enumerates the set of values in String for EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum
func GetEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum(val string) (EnableExternalExadataInfrastructureManagementDetailsLicenseModelEnum, bool) {
	enum, ok := mappingEnableExternalExadataInfrastructureManagementDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
