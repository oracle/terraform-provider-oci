// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateCccProvisionedPackage    OperationTypeEnum = "CREATE_CCC_PROVISIONED_PACKAGE"
	OperationTypeDeleteCccProvisionedPackage    OperationTypeEnum = "DELETE_CCC_PROVISIONED_PACKAGE"
	OperationTypeUpdateCccProvisionedPackage    OperationTypeEnum = "UPDATE_CCC_PROVISIONED_PACKAGE"
	OperationTypeCreateCccFlexNetwork           OperationTypeEnum = "CREATE_CCC_FLEX_NETWORK"
	OperationTypeDeleteCccFlexNetwork           OperationTypeEnum = "DELETE_CCC_FLEX_NETWORK"
	OperationTypeUpdateCccFlexNetwork           OperationTypeEnum = "UPDATE_CCC_FLEX_NETWORK"
	OperationTypeCreateCccFlexNetworkAttachment OperationTypeEnum = "CREATE_CCC_FLEX_NETWORK_ATTACHMENT"
	OperationTypeDeleteCccFlexNetworkAttachment OperationTypeEnum = "DELETE_CCC_FLEX_NETWORK_ATTACHMENT"
	OperationTypeUpdateCccFlexNetworkAttachment OperationTypeEnum = "UPDATE_CCC_FLEX_NETWORK_ATTACHMENT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_CCC_PROVISIONED_PACKAGE":     OperationTypeCreateCccProvisionedPackage,
	"DELETE_CCC_PROVISIONED_PACKAGE":     OperationTypeDeleteCccProvisionedPackage,
	"UPDATE_CCC_PROVISIONED_PACKAGE":     OperationTypeUpdateCccProvisionedPackage,
	"CREATE_CCC_FLEX_NETWORK":            OperationTypeCreateCccFlexNetwork,
	"DELETE_CCC_FLEX_NETWORK":            OperationTypeDeleteCccFlexNetwork,
	"UPDATE_CCC_FLEX_NETWORK":            OperationTypeUpdateCccFlexNetwork,
	"CREATE_CCC_FLEX_NETWORK_ATTACHMENT": OperationTypeCreateCccFlexNetworkAttachment,
	"DELETE_CCC_FLEX_NETWORK_ATTACHMENT": OperationTypeDeleteCccFlexNetworkAttachment,
	"UPDATE_CCC_FLEX_NETWORK_ATTACHMENT": OperationTypeUpdateCccFlexNetworkAttachment,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_ccc_provisioned_package":     OperationTypeCreateCccProvisionedPackage,
	"delete_ccc_provisioned_package":     OperationTypeDeleteCccProvisionedPackage,
	"update_ccc_provisioned_package":     OperationTypeUpdateCccProvisionedPackage,
	"create_ccc_flex_network":            OperationTypeCreateCccFlexNetwork,
	"delete_ccc_flex_network":            OperationTypeDeleteCccFlexNetwork,
	"update_ccc_flex_network":            OperationTypeUpdateCccFlexNetwork,
	"create_ccc_flex_network_attachment": OperationTypeCreateCccFlexNetworkAttachment,
	"delete_ccc_flex_network_attachment": OperationTypeDeleteCccFlexNetworkAttachment,
	"update_ccc_flex_network_attachment": OperationTypeUpdateCccFlexNetworkAttachment,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_CCC_PROVISIONED_PACKAGE",
		"DELETE_CCC_PROVISIONED_PACKAGE",
		"UPDATE_CCC_PROVISIONED_PACKAGE",
		"CREATE_CCC_FLEX_NETWORK",
		"DELETE_CCC_FLEX_NETWORK",
		"UPDATE_CCC_FLEX_NETWORK",
		"CREATE_CCC_FLEX_NETWORK_ATTACHMENT",
		"DELETE_CCC_FLEX_NETWORK_ATTACHMENT",
		"UPDATE_CCC_FLEX_NETWORK_ATTACHMENT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
