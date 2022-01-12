// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

// HcxLicenseStatusEnum Enum with underlying type: string
type HcxLicenseStatusEnum string

// Set of constants representing the allowable values for HcxLicenseStatusEnum
const (
	HcxLicenseStatusAvailable   HcxLicenseStatusEnum = "AVAILABLE"
	HcxLicenseStatusConsumed    HcxLicenseStatusEnum = "CONSUMED"
	HcxLicenseStatusDeactivated HcxLicenseStatusEnum = "DEACTIVATED"
	HcxLicenseStatusDeleted     HcxLicenseStatusEnum = "DELETED"
)

var mappingHcxLicenseStatus = map[string]HcxLicenseStatusEnum{
	"AVAILABLE":   HcxLicenseStatusAvailable,
	"CONSUMED":    HcxLicenseStatusConsumed,
	"DEACTIVATED": HcxLicenseStatusDeactivated,
	"DELETED":     HcxLicenseStatusDeleted,
}

// GetHcxLicenseStatusEnumValues Enumerates the set of values for HcxLicenseStatusEnum
func GetHcxLicenseStatusEnumValues() []HcxLicenseStatusEnum {
	values := make([]HcxLicenseStatusEnum, 0)
	for _, v := range mappingHcxLicenseStatus {
		values = append(values, v)
	}
	return values
}
