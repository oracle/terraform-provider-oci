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

// TargetDetectorRecipeDetectorRuleSummary Summary of the Detector Recipe Rule.
type TargetDetectorRecipeDetectorRuleSummary struct {

	// The unique identifier of the detector rule
	Id *string `mandatory:"true" json:"id"`

	// possible type of detectors
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// DetectorTemplate Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// DetectorTemplate Identifier, can be renamed
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for TargetDetectorRecipeDetectorRule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"false" json:"serviceType"`

	// resource type of the configuration to which the rule is applied
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// List of cloudguard managed list types related to this rule
	ManagedListTypes []TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	DetectorDetails *TargetDetectorDetails `mandatory:"false" json:"detectorDetails"`

	// The date and time the target detector recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target detector recipe rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the target detector recipe rule
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TargetDetectorRecipeDetectorRuleSummary) String() string {
	return common.PointerString(m)
}

// TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum Enum with underlying type: string
type TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum string

// Set of constants representing the allowable values for TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum
const (
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock    TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "CIDR_BLOCK"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesUsers        TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "USERS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGroups       TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "GROUPS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address  TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "IPV4ADDRESS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address  TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "IPV6ADDRESS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "RESOURCE_OCID"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesRegion       TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "REGION"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCountry      TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "COUNTRY"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesState        TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "STATE"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCity         TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "CITY"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesTags         TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "TAGS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGeneric      TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "GENERIC"
)

var mappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypes = map[string]TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum{
	"CIDR_BLOCK":    TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock,
	"USERS":         TargetDetectorRecipeDetectorRuleSummaryManagedListTypesUsers,
	"GROUPS":        TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGroups,
	"IPV4ADDRESS":   TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address,
	"IPV6ADDRESS":   TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address,
	"RESOURCE_OCID": TargetDetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid,
	"REGION":        TargetDetectorRecipeDetectorRuleSummaryManagedListTypesRegion,
	"COUNTRY":       TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCountry,
	"STATE":         TargetDetectorRecipeDetectorRuleSummaryManagedListTypesState,
	"CITY":          TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCity,
	"TAGS":          TargetDetectorRecipeDetectorRuleSummaryManagedListTypesTags,
	"GENERIC":       TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGeneric,
}

// GetTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumValues Enumerates the set of values for TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum
func GetTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumValues() []TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum {
	values := make([]TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum, 0)
	for _, v := range mappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypes {
		values = append(values, v)
	}
	return values
}
