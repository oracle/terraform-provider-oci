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

// DetectorRecipeDetectorRuleSummary Summary of the Detector Recipe Rule.
type DetectorRecipeDetectorRuleSummary struct {

	// The unique identifier of the detector rule
	Id *string `mandatory:"true" json:"id"`

	// possible type of detectors
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// DetectorTemplate Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// DetectorTemplate Identifier, can be renamed
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for DetectorRecipeDetectorRule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"false" json:"serviceType"`

	// resource type of the configuration to which the rule is applied
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// List of cloudguard managed list types related to this rule
	ManagedListTypes []DetectorRecipeDetectorRuleSummaryManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	// List of CandidateResponderRule related to this rule
	CandidateResponderRules []CandidateResponderRule `mandatory:"false" json:"candidateResponderRules"`

	DetectorDetails *DetectorDetails `mandatory:"false" json:"detectorDetails"`

	// The date and time the detector recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the detector recipe rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the detector recipe rule
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m DetectorRecipeDetectorRuleSummary) String() string {
	return common.PointerString(m)
}

// DetectorRecipeDetectorRuleSummaryManagedListTypesEnum Enum with underlying type: string
type DetectorRecipeDetectorRuleSummaryManagedListTypesEnum string

// Set of constants representing the allowable values for DetectorRecipeDetectorRuleSummaryManagedListTypesEnum
const (
	DetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock    DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "CIDR_BLOCK"
	DetectorRecipeDetectorRuleSummaryManagedListTypesUsers        DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "USERS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesGroups       DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "GROUPS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address  DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "IPV4ADDRESS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address  DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "IPV6ADDRESS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "RESOURCE_OCID"
	DetectorRecipeDetectorRuleSummaryManagedListTypesRegion       DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "REGION"
	DetectorRecipeDetectorRuleSummaryManagedListTypesCountry      DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "COUNTRY"
	DetectorRecipeDetectorRuleSummaryManagedListTypesState        DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "STATE"
	DetectorRecipeDetectorRuleSummaryManagedListTypesCity         DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "CITY"
	DetectorRecipeDetectorRuleSummaryManagedListTypesTags         DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "TAGS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesGeneric      DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "GENERIC"
)

var mappingDetectorRecipeDetectorRuleSummaryManagedListTypes = map[string]DetectorRecipeDetectorRuleSummaryManagedListTypesEnum{
	"CIDR_BLOCK":    DetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock,
	"USERS":         DetectorRecipeDetectorRuleSummaryManagedListTypesUsers,
	"GROUPS":        DetectorRecipeDetectorRuleSummaryManagedListTypesGroups,
	"IPV4ADDRESS":   DetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address,
	"IPV6ADDRESS":   DetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address,
	"RESOURCE_OCID": DetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid,
	"REGION":        DetectorRecipeDetectorRuleSummaryManagedListTypesRegion,
	"COUNTRY":       DetectorRecipeDetectorRuleSummaryManagedListTypesCountry,
	"STATE":         DetectorRecipeDetectorRuleSummaryManagedListTypesState,
	"CITY":          DetectorRecipeDetectorRuleSummaryManagedListTypesCity,
	"TAGS":          DetectorRecipeDetectorRuleSummaryManagedListTypesTags,
	"GENERIC":       DetectorRecipeDetectorRuleSummaryManagedListTypesGeneric,
}

// GetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumValues Enumerates the set of values for DetectorRecipeDetectorRuleSummaryManagedListTypesEnum
func GetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumValues() []DetectorRecipeDetectorRuleSummaryManagedListTypesEnum {
	values := make([]DetectorRecipeDetectorRuleSummaryManagedListTypesEnum, 0)
	for _, v := range mappingDetectorRecipeDetectorRuleSummaryManagedListTypes {
		values = append(values, v)
	}
	return values
}
