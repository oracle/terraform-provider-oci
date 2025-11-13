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

// InternalTenantThrottlingGroupAssociation Tenant group association
type InternalTenantThrottlingGroupAssociation struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of `tenant`
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The name of the tenancy group.
	ThrottlingTenantGroupName *string `mandatory:"true" json:"throttlingTenantGroupName"`

	// Data plane id of the tenant ocid
	TenantDpId *int64 `mandatory:"true" json:"tenantDpId"`

	// Specifies the category to which this group type belongs.
	GroupCategory InternalTenantThrottlingGroupAssociationGroupCategoryEnum `mandatory:"true" json:"groupCategory"`

	// The date and time when the tenant group association was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m InternalTenantThrottlingGroupAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalTenantThrottlingGroupAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalTenantThrottlingGroupAssociationGroupCategoryEnum(string(m.GroupCategory)); !ok && m.GroupCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupCategory: %s. Supported values are: %s.", m.GroupCategory, strings.Join(GetInternalTenantThrottlingGroupAssociationGroupCategoryEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalTenantThrottlingGroupAssociationGroupCategoryEnum Enum with underlying type: string
type InternalTenantThrottlingGroupAssociationGroupCategoryEnum string

// Set of constants representing the allowable values for InternalTenantThrottlingGroupAssociationGroupCategoryEnum
const (
	InternalTenantThrottlingGroupAssociationGroupCategoryTier   InternalTenantThrottlingGroupAssociationGroupCategoryEnum = "TIER"
	InternalTenantThrottlingGroupAssociationGroupCategoryCanary InternalTenantThrottlingGroupAssociationGroupCategoryEnum = "CANARY"
)

var mappingInternalTenantThrottlingGroupAssociationGroupCategoryEnum = map[string]InternalTenantThrottlingGroupAssociationGroupCategoryEnum{
	"TIER":   InternalTenantThrottlingGroupAssociationGroupCategoryTier,
	"CANARY": InternalTenantThrottlingGroupAssociationGroupCategoryCanary,
}

var mappingInternalTenantThrottlingGroupAssociationGroupCategoryEnumLowerCase = map[string]InternalTenantThrottlingGroupAssociationGroupCategoryEnum{
	"tier":   InternalTenantThrottlingGroupAssociationGroupCategoryTier,
	"canary": InternalTenantThrottlingGroupAssociationGroupCategoryCanary,
}

// GetInternalTenantThrottlingGroupAssociationGroupCategoryEnumValues Enumerates the set of values for InternalTenantThrottlingGroupAssociationGroupCategoryEnum
func GetInternalTenantThrottlingGroupAssociationGroupCategoryEnumValues() []InternalTenantThrottlingGroupAssociationGroupCategoryEnum {
	values := make([]InternalTenantThrottlingGroupAssociationGroupCategoryEnum, 0)
	for _, v := range mappingInternalTenantThrottlingGroupAssociationGroupCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalTenantThrottlingGroupAssociationGroupCategoryEnumStringValues Enumerates the set of values in String for InternalTenantThrottlingGroupAssociationGroupCategoryEnum
func GetInternalTenantThrottlingGroupAssociationGroupCategoryEnumStringValues() []string {
	return []string{
		"TIER",
		"CANARY",
	}
}

// GetMappingInternalTenantThrottlingGroupAssociationGroupCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalTenantThrottlingGroupAssociationGroupCategoryEnum(val string) (InternalTenantThrottlingGroupAssociationGroupCategoryEnum, bool) {
	enum, ok := mappingInternalTenantThrottlingGroupAssociationGroupCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
