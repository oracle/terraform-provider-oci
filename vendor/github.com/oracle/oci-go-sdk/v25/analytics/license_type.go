// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

// LicenseTypeEnum Enum with underlying type: string
type LicenseTypeEnum string

// Set of constants representing the allowable values for LicenseTypeEnum
const (
	LicenseTypeLicenseIncluded     LicenseTypeEnum = "LICENSE_INCLUDED"
	LicenseTypeBringYourOwnLicense LicenseTypeEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingLicenseType = map[string]LicenseTypeEnum{
	"LICENSE_INCLUDED":       LicenseTypeLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": LicenseTypeBringYourOwnLicense,
}

// GetLicenseTypeEnumValues Enumerates the set of values for LicenseTypeEnum
func GetLicenseTypeEnumValues() []LicenseTypeEnum {
	values := make([]LicenseTypeEnum, 0)
	for _, v := range mappingLicenseType {
		values = append(values, v)
	}
	return values
}
