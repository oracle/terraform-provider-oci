// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content and Experience API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

// LicenseTypeEnum Enum with underlying type: string
type LicenseTypeEnum string

// Set of constants representing the allowable values for LicenseTypeEnum
const (
	LicenseTypeNew  LicenseTypeEnum = "NEW"
	LicenseTypeByol LicenseTypeEnum = "BYOL"
)

var mappingLicenseType = map[string]LicenseTypeEnum{
	"NEW":  LicenseTypeNew,
	"BYOL": LicenseTypeByol,
}

// GetLicenseTypeEnumValues Enumerates the set of values for LicenseTypeEnum
func GetLicenseTypeEnumValues() []LicenseTypeEnum {
	values := make([]LicenseTypeEnum, 0)
	for _, v := range mappingLicenseType {
		values = append(values, v)
	}
	return values
}
