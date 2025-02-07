// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"strings"
)

// BillingModelEnum Enum with underlying type: string
type BillingModelEnum string

// Set of constants representing the allowable values for BillingModelEnum
const (
	BillingModelCommitment BillingModelEnum = "COMMITMENT"
	BillingModelPaygo      BillingModelEnum = "PAYGO"
	BillingModelPromotion  BillingModelEnum = "PROMOTION"
)

var mappingBillingModelEnum = map[string]BillingModelEnum{
	"COMMITMENT": BillingModelCommitment,
	"PAYGO":      BillingModelPaygo,
	"PROMOTION":  BillingModelPromotion,
}

var mappingBillingModelEnumLowerCase = map[string]BillingModelEnum{
	"commitment": BillingModelCommitment,
	"paygo":      BillingModelPaygo,
	"promotion":  BillingModelPromotion,
}

// GetBillingModelEnumValues Enumerates the set of values for BillingModelEnum
func GetBillingModelEnumValues() []BillingModelEnum {
	values := make([]BillingModelEnum, 0)
	for _, v := range mappingBillingModelEnum {
		values = append(values, v)
	}
	return values
}

// GetBillingModelEnumStringValues Enumerates the set of values in String for BillingModelEnum
func GetBillingModelEnumStringValues() []string {
	return []string{
		"COMMITMENT",
		"PAYGO",
		"PROMOTION",
	}
}

// GetMappingBillingModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBillingModelEnum(val string) (BillingModelEnum, bool) {
	enum, ok := mappingBillingModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
