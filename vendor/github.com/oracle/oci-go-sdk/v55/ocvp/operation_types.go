// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateSddc              OperationTypesEnum = "CREATE_SDDC"
	OperationTypesDeleteSddc              OperationTypesEnum = "DELETE_SDDC"
	OperationTypesCreateEsxiHost          OperationTypesEnum = "CREATE_ESXI_HOST"
	OperationTypesDeleteEsxiHost          OperationTypesEnum = "DELETE_ESXI_HOST"
	OperationTypesUpgradeHcx              OperationTypesEnum = "UPGRADE_HCX"
	OperationTypesDowngradeHcx            OperationTypesEnum = "DOWNGRADE_HCX"
	OperationTypesCancelDowngradeHcx      OperationTypesEnum = "CANCEL_DOWNGRADE_HCX"
	OperationTypesRefreshHcxLicenseStatus OperationTypesEnum = "REFRESH_HCX_LICENSE_STATUS"
)

var mappingOperationTypes = map[string]OperationTypesEnum{
	"CREATE_SDDC":                OperationTypesCreateSddc,
	"DELETE_SDDC":                OperationTypesDeleteSddc,
	"CREATE_ESXI_HOST":           OperationTypesCreateEsxiHost,
	"DELETE_ESXI_HOST":           OperationTypesDeleteEsxiHost,
	"UPGRADE_HCX":                OperationTypesUpgradeHcx,
	"DOWNGRADE_HCX":              OperationTypesDowngradeHcx,
	"CANCEL_DOWNGRADE_HCX":       OperationTypesCancelDowngradeHcx,
	"REFRESH_HCX_LICENSE_STATUS": OperationTypesRefreshHcxLicenseStatus,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypes {
		values = append(values, v)
	}
	return values
}
