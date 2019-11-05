// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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
