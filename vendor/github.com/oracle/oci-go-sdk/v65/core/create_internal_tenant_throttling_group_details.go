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

// CreateInternalTenantThrottlingGroupDetails Represents the details required to create a new tenant throttling group.
type CreateInternalTenantThrottlingGroupDetails struct {

	// The unique name of the tenancy throttling group.
	Name *string `mandatory:"true" json:"name"`

	// Specifies the type of the tenancy group.
	GroupType CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum `mandatory:"true" json:"groupType"`

	// Specifies the category to which this group type belongs.
	GroupCategory CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum `mandatory:"true" json:"groupCategory"`

	Overrides *TenantGroupOverrides `mandatory:"false" json:"overrides"`
}

func (m CreateInternalTenantThrottlingGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalTenantThrottlingGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateInternalTenantThrottlingGroupDetailsGroupTypeEnum(string(m.GroupType)); !ok && m.GroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", m.GroupType, strings.Join(GetCreateInternalTenantThrottlingGroupDetailsGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum(string(m.GroupCategory)); !ok && m.GroupCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupCategory: %s. Supported values are: %s.", m.GroupCategory, strings.Join(GetCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum Enum with underlying type: string
type CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum string

// Set of constants representing the allowable values for CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum
const (
	CreateInternalTenantThrottlingGroupDetailsGroupTypeFreeTier    CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum = "FREE_TIER"
	CreateInternalTenantThrottlingGroupDetailsGroupTypePaid        CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum = "PAID"
	CreateInternalTenantThrottlingGroupDetailsGroupTypeVcnCpCanary CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum = "VCN_CP_CANARY"
	CreateInternalTenantThrottlingGroupDetailsGroupTypeVcnDpCanary CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum = "VCN_DP_CANARY"
)

var mappingCreateInternalTenantThrottlingGroupDetailsGroupTypeEnum = map[string]CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum{
	"FREE_TIER":     CreateInternalTenantThrottlingGroupDetailsGroupTypeFreeTier,
	"PAID":          CreateInternalTenantThrottlingGroupDetailsGroupTypePaid,
	"VCN_CP_CANARY": CreateInternalTenantThrottlingGroupDetailsGroupTypeVcnCpCanary,
	"VCN_DP_CANARY": CreateInternalTenantThrottlingGroupDetailsGroupTypeVcnDpCanary,
}

var mappingCreateInternalTenantThrottlingGroupDetailsGroupTypeEnumLowerCase = map[string]CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum{
	"free_tier":     CreateInternalTenantThrottlingGroupDetailsGroupTypeFreeTier,
	"paid":          CreateInternalTenantThrottlingGroupDetailsGroupTypePaid,
	"vcn_cp_canary": CreateInternalTenantThrottlingGroupDetailsGroupTypeVcnCpCanary,
	"vcn_dp_canary": CreateInternalTenantThrottlingGroupDetailsGroupTypeVcnDpCanary,
}

// GetCreateInternalTenantThrottlingGroupDetailsGroupTypeEnumValues Enumerates the set of values for CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum
func GetCreateInternalTenantThrottlingGroupDetailsGroupTypeEnumValues() []CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum {
	values := make([]CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum, 0)
	for _, v := range mappingCreateInternalTenantThrottlingGroupDetailsGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalTenantThrottlingGroupDetailsGroupTypeEnumStringValues Enumerates the set of values in String for CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum
func GetCreateInternalTenantThrottlingGroupDetailsGroupTypeEnumStringValues() []string {
	return []string{
		"FREE_TIER",
		"PAID",
		"VCN_CP_CANARY",
		"VCN_DP_CANARY",
	}
}

// GetMappingCreateInternalTenantThrottlingGroupDetailsGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalTenantThrottlingGroupDetailsGroupTypeEnum(val string) (CreateInternalTenantThrottlingGroupDetailsGroupTypeEnum, bool) {
	enum, ok := mappingCreateInternalTenantThrottlingGroupDetailsGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum Enum with underlying type: string
type CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum string

// Set of constants representing the allowable values for CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum
const (
	CreateInternalTenantThrottlingGroupDetailsGroupCategoryTier   CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum = "TIER"
	CreateInternalTenantThrottlingGroupDetailsGroupCategoryCanary CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum = "CANARY"
)

var mappingCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum = map[string]CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum{
	"TIER":   CreateInternalTenantThrottlingGroupDetailsGroupCategoryTier,
	"CANARY": CreateInternalTenantThrottlingGroupDetailsGroupCategoryCanary,
}

var mappingCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnumLowerCase = map[string]CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum{
	"tier":   CreateInternalTenantThrottlingGroupDetailsGroupCategoryTier,
	"canary": CreateInternalTenantThrottlingGroupDetailsGroupCategoryCanary,
}

// GetCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnumValues Enumerates the set of values for CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum
func GetCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnumValues() []CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum {
	values := make([]CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum, 0)
	for _, v := range mappingCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnumStringValues Enumerates the set of values in String for CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum
func GetCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnumStringValues() []string {
	return []string{
		"TIER",
		"CANARY",
	}
}

// GetMappingCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum(val string) (CreateInternalTenantThrottlingGroupDetailsGroupCategoryEnum, bool) {
	enum, ok := mappingCreateInternalTenantThrottlingGroupDetailsGroupCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
