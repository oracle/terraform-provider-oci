// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateInternalTenantThrottlingGroupDetails Represents the details required to update tenant throttling group.
type UpdateInternalTenantThrottlingGroupDetails struct {

	// Specifies the type of the tenancy group.
	GroupType UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum `mandatory:"false" json:"groupType,omitempty"`

	// Specifies the category to which this group type belongs.
	GroupCategory UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum `mandatory:"false" json:"groupCategory,omitempty"`

	Overrides *TenantGroupOverrides `mandatory:"false" json:"overrides"`
}

func (m UpdateInternalTenantThrottlingGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInternalTenantThrottlingGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum(string(m.GroupType)); !ok && m.GroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", m.GroupType, strings.Join(GetUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum(string(m.GroupCategory)); !ok && m.GroupCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupCategory: %s. Supported values are: %s.", m.GroupCategory, strings.Join(GetUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum Enum with underlying type: string
type UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum string

// Set of constants representing the allowable values for UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum
const (
	UpdateInternalTenantThrottlingGroupDetailsGroupTypeFreeTier    UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum = "FREE_TIER"
	UpdateInternalTenantThrottlingGroupDetailsGroupTypePaid        UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum = "PAID"
	UpdateInternalTenantThrottlingGroupDetailsGroupTypeVcnCpCanary UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum = "VCN_CP_CANARY"
	UpdateInternalTenantThrottlingGroupDetailsGroupTypeVcnDpCanary UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum = "VCN_DP_CANARY"
)

var mappingUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum = map[string]UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum{
	"FREE_TIER":     UpdateInternalTenantThrottlingGroupDetailsGroupTypeFreeTier,
	"PAID":          UpdateInternalTenantThrottlingGroupDetailsGroupTypePaid,
	"VCN_CP_CANARY": UpdateInternalTenantThrottlingGroupDetailsGroupTypeVcnCpCanary,
	"VCN_DP_CANARY": UpdateInternalTenantThrottlingGroupDetailsGroupTypeVcnDpCanary,
}

var mappingUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnumLowerCase = map[string]UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum{
	"free_tier":     UpdateInternalTenantThrottlingGroupDetailsGroupTypeFreeTier,
	"paid":          UpdateInternalTenantThrottlingGroupDetailsGroupTypePaid,
	"vcn_cp_canary": UpdateInternalTenantThrottlingGroupDetailsGroupTypeVcnCpCanary,
	"vcn_dp_canary": UpdateInternalTenantThrottlingGroupDetailsGroupTypeVcnDpCanary,
}

// GetUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnumValues Enumerates the set of values for UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum
func GetUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnumValues() []UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum {
	values := make([]UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum, 0)
	for _, v := range mappingUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnumStringValues Enumerates the set of values in String for UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum
func GetUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnumStringValues() []string {
	return []string{
		"FREE_TIER",
		"PAID",
		"VCN_CP_CANARY",
		"VCN_DP_CANARY",
	}
}

// GetMappingUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum(val string) (UpdateInternalTenantThrottlingGroupDetailsGroupTypeEnum, bool) {
	enum, ok := mappingUpdateInternalTenantThrottlingGroupDetailsGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum Enum with underlying type: string
type UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum string

// Set of constants representing the allowable values for UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum
const (
	UpdateInternalTenantThrottlingGroupDetailsGroupCategoryTier   UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum = "TIER"
	UpdateInternalTenantThrottlingGroupDetailsGroupCategoryCanary UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum = "CANARY"
)

var mappingUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum = map[string]UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum{
	"TIER":   UpdateInternalTenantThrottlingGroupDetailsGroupCategoryTier,
	"CANARY": UpdateInternalTenantThrottlingGroupDetailsGroupCategoryCanary,
}

var mappingUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnumLowerCase = map[string]UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum{
	"tier":   UpdateInternalTenantThrottlingGroupDetailsGroupCategoryTier,
	"canary": UpdateInternalTenantThrottlingGroupDetailsGroupCategoryCanary,
}

// GetUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnumValues Enumerates the set of values for UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum
func GetUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnumValues() []UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum {
	values := make([]UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum, 0)
	for _, v := range mappingUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnumStringValues Enumerates the set of values in String for UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum
func GetUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnumStringValues() []string {
	return []string{
		"TIER",
		"CANARY",
	}
}

// GetMappingUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum(val string) (UpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnum, bool) {
	enum, ok := mappingUpdateInternalTenantThrottlingGroupDetailsGroupCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
