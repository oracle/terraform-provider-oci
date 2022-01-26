// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TargetDetectorRecipeDetectorRule Detector Recipe Rule
type TargetDetectorRecipeDetectorRule struct {

	// The unique identifier of the detector rule
	DetectorRuleId *string `mandatory:"true" json:"detectorRuleId"`

	// detector for the rule
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"true" json:"serviceType"`

	// resource type of the configuration to which the rule is applied
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// displayName
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for TargetDetectorRecipeDetectorRule
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for TargetDetectorRecipeDetectorRule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	Details *TargetDetectorDetails `mandatory:"false" json:"details"`

	// List of cloudguard managed list types related to this rule
	ManagedListTypes []TargetDetectorRecipeDetectorRuleManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	// The date and time the target detector recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target detector recipe rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the DetectorRule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TargetDetectorRecipeDetectorRule) String() string {
	return common.PointerString(m)
}

// TargetDetectorRecipeDetectorRuleManagedListTypesEnum Enum with underlying type: string
type TargetDetectorRecipeDetectorRuleManagedListTypesEnum string

// Set of constants representing the allowable values for TargetDetectorRecipeDetectorRuleManagedListTypesEnum
const (
	TargetDetectorRecipeDetectorRuleManagedListTypesCidrBlock    TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "CIDR_BLOCK"
	TargetDetectorRecipeDetectorRuleManagedListTypesUsers        TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "USERS"
	TargetDetectorRecipeDetectorRuleManagedListTypesGroups       TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "GROUPS"
	TargetDetectorRecipeDetectorRuleManagedListTypesIpv4address  TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "IPV4ADDRESS"
	TargetDetectorRecipeDetectorRuleManagedListTypesIpv6address  TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "IPV6ADDRESS"
	TargetDetectorRecipeDetectorRuleManagedListTypesResourceOcid TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "RESOURCE_OCID"
	TargetDetectorRecipeDetectorRuleManagedListTypesRegion       TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "REGION"
	TargetDetectorRecipeDetectorRuleManagedListTypesCountry      TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "COUNTRY"
	TargetDetectorRecipeDetectorRuleManagedListTypesState        TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "STATE"
	TargetDetectorRecipeDetectorRuleManagedListTypesCity         TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "CITY"
	TargetDetectorRecipeDetectorRuleManagedListTypesTags         TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "TAGS"
	TargetDetectorRecipeDetectorRuleManagedListTypesGeneric      TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "GENERIC"
)

var mappingTargetDetectorRecipeDetectorRuleManagedListTypes = map[string]TargetDetectorRecipeDetectorRuleManagedListTypesEnum{
	"CIDR_BLOCK":    TargetDetectorRecipeDetectorRuleManagedListTypesCidrBlock,
	"USERS":         TargetDetectorRecipeDetectorRuleManagedListTypesUsers,
	"GROUPS":        TargetDetectorRecipeDetectorRuleManagedListTypesGroups,
	"IPV4ADDRESS":   TargetDetectorRecipeDetectorRuleManagedListTypesIpv4address,
	"IPV6ADDRESS":   TargetDetectorRecipeDetectorRuleManagedListTypesIpv6address,
	"RESOURCE_OCID": TargetDetectorRecipeDetectorRuleManagedListTypesResourceOcid,
	"REGION":        TargetDetectorRecipeDetectorRuleManagedListTypesRegion,
	"COUNTRY":       TargetDetectorRecipeDetectorRuleManagedListTypesCountry,
	"STATE":         TargetDetectorRecipeDetectorRuleManagedListTypesState,
	"CITY":          TargetDetectorRecipeDetectorRuleManagedListTypesCity,
	"TAGS":          TargetDetectorRecipeDetectorRuleManagedListTypesTags,
	"GENERIC":       TargetDetectorRecipeDetectorRuleManagedListTypesGeneric,
}

// GetTargetDetectorRecipeDetectorRuleManagedListTypesEnumValues Enumerates the set of values for TargetDetectorRecipeDetectorRuleManagedListTypesEnum
func GetTargetDetectorRecipeDetectorRuleManagedListTypesEnumValues() []TargetDetectorRecipeDetectorRuleManagedListTypesEnum {
	values := make([]TargetDetectorRecipeDetectorRuleManagedListTypesEnum, 0)
	for _, v := range mappingTargetDetectorRecipeDetectorRuleManagedListTypes {
		values = append(values, v)
	}
	return values
}
