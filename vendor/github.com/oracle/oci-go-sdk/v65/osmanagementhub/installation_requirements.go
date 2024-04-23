// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// InstallationRequirementsEnum Enum with underlying type: string
type InstallationRequirementsEnum string

// Set of constants representing the allowable values for InstallationRequirementsEnum
const (
	InstallationRequirementsEulaAcceptanceRequired  InstallationRequirementsEnum = "EULA_ACCEPTANCE_REQUIRED"
	InstallationRequirementsSoftwareMediaRequired   InstallationRequirementsEnum = "SOFTWARE_MEDIA_REQUIRED"
	InstallationRequirementsUserInteractionRequired InstallationRequirementsEnum = "USER_INTERACTION_REQUIRED"
)

var mappingInstallationRequirementsEnum = map[string]InstallationRequirementsEnum{
	"EULA_ACCEPTANCE_REQUIRED":  InstallationRequirementsEulaAcceptanceRequired,
	"SOFTWARE_MEDIA_REQUIRED":   InstallationRequirementsSoftwareMediaRequired,
	"USER_INTERACTION_REQUIRED": InstallationRequirementsUserInteractionRequired,
}

var mappingInstallationRequirementsEnumLowerCase = map[string]InstallationRequirementsEnum{
	"eula_acceptance_required":  InstallationRequirementsEulaAcceptanceRequired,
	"software_media_required":   InstallationRequirementsSoftwareMediaRequired,
	"user_interaction_required": InstallationRequirementsUserInteractionRequired,
}

// GetInstallationRequirementsEnumValues Enumerates the set of values for InstallationRequirementsEnum
func GetInstallationRequirementsEnumValues() []InstallationRequirementsEnum {
	values := make([]InstallationRequirementsEnum, 0)
	for _, v := range mappingInstallationRequirementsEnum {
		values = append(values, v)
	}
	return values
}

// GetInstallationRequirementsEnumStringValues Enumerates the set of values in String for InstallationRequirementsEnum
func GetInstallationRequirementsEnumStringValues() []string {
	return []string{
		"EULA_ACCEPTANCE_REQUIRED",
		"SOFTWARE_MEDIA_REQUIRED",
		"USER_INTERACTION_REQUIRED",
	}
}

// GetMappingInstallationRequirementsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstallationRequirementsEnum(val string) (InstallationRequirementsEnum, bool) {
	enum, ok := mappingInstallationRequirementsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
