// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

// LicenseModelEnum Enum with underlying type: string
type LicenseModelEnum string

// Set of constants representing the allowable values for LicenseModelEnum
const (
	LicenseModelLicenseIncluded     LicenseModelEnum = "LICENSE_INCLUDED"
	LicenseModelBringYourOwnLicense LicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingLicenseModel = map[string]LicenseModelEnum{
	"LICENSE_INCLUDED":       LicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": LicenseModelBringYourOwnLicense,
}

// GetLicenseModelEnumValues Enumerates the set of values for LicenseModelEnum
func GetLicenseModelEnumValues() []LicenseModelEnum {
	values := make([]LicenseModelEnum, 0)
	for _, v := range mappingLicenseModel {
		values = append(values, v)
	}
	return values
}
