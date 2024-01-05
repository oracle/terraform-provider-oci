// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// License Manager API
//
// Use the License Manager API to manage product licenses and license records. For more information, see License Manager Overview (https://docs.cloud.oracle.com/iaas/Content/LicenseManager/Concepts/licensemanageroverview.htm).
//

package licensemanager

import (
	"strings"
)

// ResourceUnitEnum Enum with underlying type: string
type ResourceUnitEnum string

// Set of constants representing the allowable values for ResourceUnitEnum
const (
	ResourceUnitOcpu ResourceUnitEnum = "OCPU"
	ResourceUnitEcpu ResourceUnitEnum = "ECPU"
)

var mappingResourceUnitEnum = map[string]ResourceUnitEnum{
	"OCPU": ResourceUnitOcpu,
	"ECPU": ResourceUnitEcpu,
}

var mappingResourceUnitEnumLowerCase = map[string]ResourceUnitEnum{
	"ocpu": ResourceUnitOcpu,
	"ecpu": ResourceUnitEcpu,
}

// GetResourceUnitEnumValues Enumerates the set of values for ResourceUnitEnum
func GetResourceUnitEnumValues() []ResourceUnitEnum {
	values := make([]ResourceUnitEnum, 0)
	for _, v := range mappingResourceUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceUnitEnumStringValues Enumerates the set of values in String for ResourceUnitEnum
func GetResourceUnitEnumStringValues() []string {
	return []string{
		"OCPU",
		"ECPU",
	}
}

// GetMappingResourceUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceUnitEnum(val string) (ResourceUnitEnum, bool) {
	enum, ok := mappingResourceUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
