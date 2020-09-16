// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

// InstallationRequirementsEnum Enum with underlying type: string
type InstallationRequirementsEnum string

// Set of constants representing the allowable values for InstallationRequirementsEnum
const (
	InstallationRequirementsEulaAcceptanceRequired  InstallationRequirementsEnum = "EULA_ACCEPTANCE_REQUIRED"
	InstallationRequirementsSoftwareMediaRequired   InstallationRequirementsEnum = "SOFTWARE_MEDIA_REQUIRED"
	InstallationRequirementsUserInteractionRequired InstallationRequirementsEnum = "USER_INTERACTION_REQUIRED"
)

var mappingInstallationRequirements = map[string]InstallationRequirementsEnum{
	"EULA_ACCEPTANCE_REQUIRED":  InstallationRequirementsEulaAcceptanceRequired,
	"SOFTWARE_MEDIA_REQUIRED":   InstallationRequirementsSoftwareMediaRequired,
	"USER_INTERACTION_REQUIRED": InstallationRequirementsUserInteractionRequired,
}

// GetInstallationRequirementsEnumValues Enumerates the set of values for InstallationRequirementsEnum
func GetInstallationRequirementsEnumValues() []InstallationRequirementsEnum {
	values := make([]InstallationRequirementsEnum, 0)
	for _, v := range mappingInstallationRequirements {
		values = append(values, v)
	}
	return values
}
