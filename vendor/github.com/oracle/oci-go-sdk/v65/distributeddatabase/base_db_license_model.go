// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"strings"
)

// BaseDbLicenseModelEnum Enum with underlying type: string
type BaseDbLicenseModelEnum string

// Set of constants representing the allowable values for BaseDbLicenseModelEnum
const (
	BaseDbLicenseModelLicenseIncluded     BaseDbLicenseModelEnum = "LICENSE_INCLUDED"
	BaseDbLicenseModelBringYourOwnLicense BaseDbLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingBaseDbLicenseModelEnum = map[string]BaseDbLicenseModelEnum{
	"LICENSE_INCLUDED":       BaseDbLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": BaseDbLicenseModelBringYourOwnLicense,
}

var mappingBaseDbLicenseModelEnumLowerCase = map[string]BaseDbLicenseModelEnum{
	"license_included":       BaseDbLicenseModelLicenseIncluded,
	"bring_your_own_license": BaseDbLicenseModelBringYourOwnLicense,
}

// GetBaseDbLicenseModelEnumValues Enumerates the set of values for BaseDbLicenseModelEnum
func GetBaseDbLicenseModelEnumValues() []BaseDbLicenseModelEnum {
	values := make([]BaseDbLicenseModelEnum, 0)
	for _, v := range mappingBaseDbLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseDbLicenseModelEnumStringValues Enumerates the set of values in String for BaseDbLicenseModelEnum
func GetBaseDbLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingBaseDbLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseDbLicenseModelEnum(val string) (BaseDbLicenseModelEnum, bool) {
	enum, ok := mappingBaseDbLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
