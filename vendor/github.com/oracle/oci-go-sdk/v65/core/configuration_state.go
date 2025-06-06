// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"strings"
)

// ConfigurationStateEnum Enum with underlying type: string
type ConfigurationStateEnum string

// Set of constants representing the allowable values for ConfigurationStateEnum
const (
	ConfigurationStateConformant    ConfigurationStateEnum = "CONFORMANT"
	ConfigurationStateNonConformant ConfigurationStateEnum = "NON_CONFORMANT"
	ConfigurationStateChecking      ConfigurationStateEnum = "CHECKING"
	ConfigurationStatePreApplying   ConfigurationStateEnum = "PRE_APPLYING"
	ConfigurationStateApplying      ConfigurationStateEnum = "APPLYING"
	ConfigurationStateUnknown       ConfigurationStateEnum = "UNKNOWN"
)

var mappingConfigurationStateEnum = map[string]ConfigurationStateEnum{
	"CONFORMANT":     ConfigurationStateConformant,
	"NON_CONFORMANT": ConfigurationStateNonConformant,
	"CHECKING":       ConfigurationStateChecking,
	"PRE_APPLYING":   ConfigurationStatePreApplying,
	"APPLYING":       ConfigurationStateApplying,
	"UNKNOWN":        ConfigurationStateUnknown,
}

var mappingConfigurationStateEnumLowerCase = map[string]ConfigurationStateEnum{
	"conformant":     ConfigurationStateConformant,
	"non_conformant": ConfigurationStateNonConformant,
	"checking":       ConfigurationStateChecking,
	"pre_applying":   ConfigurationStatePreApplying,
	"applying":       ConfigurationStateApplying,
	"unknown":        ConfigurationStateUnknown,
}

// GetConfigurationStateEnumValues Enumerates the set of values for ConfigurationStateEnum
func GetConfigurationStateEnumValues() []ConfigurationStateEnum {
	values := make([]ConfigurationStateEnum, 0)
	for _, v := range mappingConfigurationStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationStateEnumStringValues Enumerates the set of values in String for ConfigurationStateEnum
func GetConfigurationStateEnumStringValues() []string {
	return []string{
		"CONFORMANT",
		"NON_CONFORMANT",
		"CHECKING",
		"PRE_APPLYING",
		"APPLYING",
		"UNKNOWN",
	}
}

// GetMappingConfigurationStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationStateEnum(val string) (ConfigurationStateEnum, bool) {
	enum, ok := mappingConfigurationStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
