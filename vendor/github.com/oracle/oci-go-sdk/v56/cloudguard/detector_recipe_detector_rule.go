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

// DetectorRecipeDetectorRule Detector Recipe Rule
type DetectorRecipeDetectorRule struct {

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

	// Description for DetectorRecipeDetectorRule
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for DetectorRecipeDetectorRule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	Details *DetectorDetails `mandatory:"false" json:"details"`

	// List of cloudguard managed list types related to this rule
	ManagedListTypes []DetectorRecipeDetectorRuleManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	// List of CandidateResponderRule related to this rule
	CandidateResponderRules []CandidateResponderRule `mandatory:"false" json:"candidateResponderRules"`

	// The date and time the detector recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the detector recipe rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the DetectorRule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m DetectorRecipeDetectorRule) String() string {
	return common.PointerString(m)
}

// DetectorRecipeDetectorRuleManagedListTypesEnum Enum with underlying type: string
type DetectorRecipeDetectorRuleManagedListTypesEnum string

// Set of constants representing the allowable values for DetectorRecipeDetectorRuleManagedListTypesEnum
const (
	DetectorRecipeDetectorRuleManagedListTypesCidrBlock    DetectorRecipeDetectorRuleManagedListTypesEnum = "CIDR_BLOCK"
	DetectorRecipeDetectorRuleManagedListTypesUsers        DetectorRecipeDetectorRuleManagedListTypesEnum = "USERS"
	DetectorRecipeDetectorRuleManagedListTypesGroups       DetectorRecipeDetectorRuleManagedListTypesEnum = "GROUPS"
	DetectorRecipeDetectorRuleManagedListTypesIpv4address  DetectorRecipeDetectorRuleManagedListTypesEnum = "IPV4ADDRESS"
	DetectorRecipeDetectorRuleManagedListTypesIpv6address  DetectorRecipeDetectorRuleManagedListTypesEnum = "IPV6ADDRESS"
	DetectorRecipeDetectorRuleManagedListTypesResourceOcid DetectorRecipeDetectorRuleManagedListTypesEnum = "RESOURCE_OCID"
	DetectorRecipeDetectorRuleManagedListTypesRegion       DetectorRecipeDetectorRuleManagedListTypesEnum = "REGION"
	DetectorRecipeDetectorRuleManagedListTypesCountry      DetectorRecipeDetectorRuleManagedListTypesEnum = "COUNTRY"
	DetectorRecipeDetectorRuleManagedListTypesState        DetectorRecipeDetectorRuleManagedListTypesEnum = "STATE"
	DetectorRecipeDetectorRuleManagedListTypesCity         DetectorRecipeDetectorRuleManagedListTypesEnum = "CITY"
	DetectorRecipeDetectorRuleManagedListTypesTags         DetectorRecipeDetectorRuleManagedListTypesEnum = "TAGS"
	DetectorRecipeDetectorRuleManagedListTypesGeneric      DetectorRecipeDetectorRuleManagedListTypesEnum = "GENERIC"
)

var mappingDetectorRecipeDetectorRuleManagedListTypes = map[string]DetectorRecipeDetectorRuleManagedListTypesEnum{
	"CIDR_BLOCK":    DetectorRecipeDetectorRuleManagedListTypesCidrBlock,
	"USERS":         DetectorRecipeDetectorRuleManagedListTypesUsers,
	"GROUPS":        DetectorRecipeDetectorRuleManagedListTypesGroups,
	"IPV4ADDRESS":   DetectorRecipeDetectorRuleManagedListTypesIpv4address,
	"IPV6ADDRESS":   DetectorRecipeDetectorRuleManagedListTypesIpv6address,
	"RESOURCE_OCID": DetectorRecipeDetectorRuleManagedListTypesResourceOcid,
	"REGION":        DetectorRecipeDetectorRuleManagedListTypesRegion,
	"COUNTRY":       DetectorRecipeDetectorRuleManagedListTypesCountry,
	"STATE":         DetectorRecipeDetectorRuleManagedListTypesState,
	"CITY":          DetectorRecipeDetectorRuleManagedListTypesCity,
	"TAGS":          DetectorRecipeDetectorRuleManagedListTypesTags,
	"GENERIC":       DetectorRecipeDetectorRuleManagedListTypesGeneric,
}

// GetDetectorRecipeDetectorRuleManagedListTypesEnumValues Enumerates the set of values for DetectorRecipeDetectorRuleManagedListTypesEnum
func GetDetectorRecipeDetectorRuleManagedListTypesEnumValues() []DetectorRecipeDetectorRuleManagedListTypesEnum {
	values := make([]DetectorRecipeDetectorRuleManagedListTypesEnum, 0)
	for _, v := range mappingDetectorRecipeDetectorRuleManagedListTypes {
		values = append(values, v)
	}
	return values
}
