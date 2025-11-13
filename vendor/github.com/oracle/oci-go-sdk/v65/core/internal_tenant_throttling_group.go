// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalTenantThrottlingGroup Tenant group
type InternalTenantThrottlingGroup struct {

	// The unique name of the tenancy throttling group.
	Name *string `mandatory:"true" json:"name"`

	// Specifies the type of the tenancy group.
	// - **FREE_TIER**: Tenants in the `ALWAYS_FREE_ONLY` state.
	// - **PAID**: Tenants in the `ACTIVE` state.
	// - **VCN_CP_CANARY**: Canary group for the VCN Control Plane.
	// - **VCN_DP_CANARY**: Canary group for the VCN Data Plane.
	GroupType InternalTenantThrottlingGroupGroupTypeEnum `mandatory:"true" json:"groupType"`

	// Specifies the category to which this group type belongs.
	// Categories help organize group types under broader logical groups.
	// - The **TIER** category includes `FREE_TIER` and `PAID` group types.
	// - The **CANARY** category includes `VCN_CP_CANARY` and `VCN_DP_CANARY` group types.
	// A tenant can only be associated with one group type within a given category.
	// For example, a tenant associated with a `FREE_TIER` group cannot also be associated with a `PAID` group, since both belong to the `TIER` category.
	GroupCategory InternalTenantThrottlingGroupGroupCategoryEnum `mandatory:"true" json:"groupCategory"`

	// Data plane id of the tenant group
	GroupDpId *int64 `mandatory:"false" json:"groupDpId"`

	Overrides *TenantGroupOverrides `mandatory:"false" json:"overrides"`

	// The date and time when the tenant group was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m InternalTenantThrottlingGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalTenantThrottlingGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalTenantThrottlingGroupGroupTypeEnum(string(m.GroupType)); !ok && m.GroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", m.GroupType, strings.Join(GetInternalTenantThrottlingGroupGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInternalTenantThrottlingGroupGroupCategoryEnum(string(m.GroupCategory)); !ok && m.GroupCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupCategory: %s. Supported values are: %s.", m.GroupCategory, strings.Join(GetInternalTenantThrottlingGroupGroupCategoryEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalTenantThrottlingGroupGroupTypeEnum Enum with underlying type: string
type InternalTenantThrottlingGroupGroupTypeEnum string

// Set of constants representing the allowable values for InternalTenantThrottlingGroupGroupTypeEnum
const (
	InternalTenantThrottlingGroupGroupTypeFreeTier    InternalTenantThrottlingGroupGroupTypeEnum = "FREE_TIER"
	InternalTenantThrottlingGroupGroupTypePaid        InternalTenantThrottlingGroupGroupTypeEnum = "PAID"
	InternalTenantThrottlingGroupGroupTypeVcnCpCanary InternalTenantThrottlingGroupGroupTypeEnum = "VCN_CP_CANARY"
	InternalTenantThrottlingGroupGroupTypeVcnDpCanary InternalTenantThrottlingGroupGroupTypeEnum = "VCN_DP_CANARY"
)

var mappingInternalTenantThrottlingGroupGroupTypeEnum = map[string]InternalTenantThrottlingGroupGroupTypeEnum{
	"FREE_TIER":     InternalTenantThrottlingGroupGroupTypeFreeTier,
	"PAID":          InternalTenantThrottlingGroupGroupTypePaid,
	"VCN_CP_CANARY": InternalTenantThrottlingGroupGroupTypeVcnCpCanary,
	"VCN_DP_CANARY": InternalTenantThrottlingGroupGroupTypeVcnDpCanary,
}

var mappingInternalTenantThrottlingGroupGroupTypeEnumLowerCase = map[string]InternalTenantThrottlingGroupGroupTypeEnum{
	"free_tier":     InternalTenantThrottlingGroupGroupTypeFreeTier,
	"paid":          InternalTenantThrottlingGroupGroupTypePaid,
	"vcn_cp_canary": InternalTenantThrottlingGroupGroupTypeVcnCpCanary,
	"vcn_dp_canary": InternalTenantThrottlingGroupGroupTypeVcnDpCanary,
}

// GetInternalTenantThrottlingGroupGroupTypeEnumValues Enumerates the set of values for InternalTenantThrottlingGroupGroupTypeEnum
func GetInternalTenantThrottlingGroupGroupTypeEnumValues() []InternalTenantThrottlingGroupGroupTypeEnum {
	values := make([]InternalTenantThrottlingGroupGroupTypeEnum, 0)
	for _, v := range mappingInternalTenantThrottlingGroupGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalTenantThrottlingGroupGroupTypeEnumStringValues Enumerates the set of values in String for InternalTenantThrottlingGroupGroupTypeEnum
func GetInternalTenantThrottlingGroupGroupTypeEnumStringValues() []string {
	return []string{
		"FREE_TIER",
		"PAID",
		"VCN_CP_CANARY",
		"VCN_DP_CANARY",
	}
}

// GetMappingInternalTenantThrottlingGroupGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalTenantThrottlingGroupGroupTypeEnum(val string) (InternalTenantThrottlingGroupGroupTypeEnum, bool) {
	enum, ok := mappingInternalTenantThrottlingGroupGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InternalTenantThrottlingGroupGroupCategoryEnum Enum with underlying type: string
type InternalTenantThrottlingGroupGroupCategoryEnum string

// Set of constants representing the allowable values for InternalTenantThrottlingGroupGroupCategoryEnum
const (
	InternalTenantThrottlingGroupGroupCategoryTier   InternalTenantThrottlingGroupGroupCategoryEnum = "TIER"
	InternalTenantThrottlingGroupGroupCategoryCanary InternalTenantThrottlingGroupGroupCategoryEnum = "CANARY"
)

var mappingInternalTenantThrottlingGroupGroupCategoryEnum = map[string]InternalTenantThrottlingGroupGroupCategoryEnum{
	"TIER":   InternalTenantThrottlingGroupGroupCategoryTier,
	"CANARY": InternalTenantThrottlingGroupGroupCategoryCanary,
}

var mappingInternalTenantThrottlingGroupGroupCategoryEnumLowerCase = map[string]InternalTenantThrottlingGroupGroupCategoryEnum{
	"tier":   InternalTenantThrottlingGroupGroupCategoryTier,
	"canary": InternalTenantThrottlingGroupGroupCategoryCanary,
}

// GetInternalTenantThrottlingGroupGroupCategoryEnumValues Enumerates the set of values for InternalTenantThrottlingGroupGroupCategoryEnum
func GetInternalTenantThrottlingGroupGroupCategoryEnumValues() []InternalTenantThrottlingGroupGroupCategoryEnum {
	values := make([]InternalTenantThrottlingGroupGroupCategoryEnum, 0)
	for _, v := range mappingInternalTenantThrottlingGroupGroupCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalTenantThrottlingGroupGroupCategoryEnumStringValues Enumerates the set of values in String for InternalTenantThrottlingGroupGroupCategoryEnum
func GetInternalTenantThrottlingGroupGroupCategoryEnumStringValues() []string {
	return []string{
		"TIER",
		"CANARY",
	}
}

// GetMappingInternalTenantThrottlingGroupGroupCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalTenantThrottlingGroupGroupCategoryEnum(val string) (InternalTenantThrottlingGroupGroupCategoryEnum, bool) {
	enum, ok := mappingInternalTenantThrottlingGroupGroupCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
