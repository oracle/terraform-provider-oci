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

// DetectorRuleSummary Summary of the Detector Rules.
type DetectorRuleSummary struct {

	// The unique identifier of the detector rule
	Id *string `mandatory:"true" json:"id"`

	// possible type of detectors
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// DetectorTemplate Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for detector rule
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for detector rule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"false" json:"serviceType"`

	// resource type of the configuration to which the rule is applied
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// List of cloudguard managed list types related to this rule
	ManagedListTypes []DetectorRuleSummaryManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	// List of CandidateResponderRule related to this rule
	CandidateResponderRules []CandidateResponderRule `mandatory:"false" json:"candidateResponderRules"`

	DetectorDetails *DetectorDetails `mandatory:"false" json:"detectorDetails"`

	// The date and time the detector rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the detector rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the detector rule
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m DetectorRuleSummary) String() string {
	return common.PointerString(m)
}

// DetectorRuleSummaryManagedListTypesEnum Enum with underlying type: string
type DetectorRuleSummaryManagedListTypesEnum string

// Set of constants representing the allowable values for DetectorRuleSummaryManagedListTypesEnum
const (
	DetectorRuleSummaryManagedListTypesCidrBlock    DetectorRuleSummaryManagedListTypesEnum = "CIDR_BLOCK"
	DetectorRuleSummaryManagedListTypesUsers        DetectorRuleSummaryManagedListTypesEnum = "USERS"
	DetectorRuleSummaryManagedListTypesGroups       DetectorRuleSummaryManagedListTypesEnum = "GROUPS"
	DetectorRuleSummaryManagedListTypesIpv4address  DetectorRuleSummaryManagedListTypesEnum = "IPV4ADDRESS"
	DetectorRuleSummaryManagedListTypesIpv6address  DetectorRuleSummaryManagedListTypesEnum = "IPV6ADDRESS"
	DetectorRuleSummaryManagedListTypesResourceOcid DetectorRuleSummaryManagedListTypesEnum = "RESOURCE_OCID"
	DetectorRuleSummaryManagedListTypesRegion       DetectorRuleSummaryManagedListTypesEnum = "REGION"
	DetectorRuleSummaryManagedListTypesCountry      DetectorRuleSummaryManagedListTypesEnum = "COUNTRY"
	DetectorRuleSummaryManagedListTypesState        DetectorRuleSummaryManagedListTypesEnum = "STATE"
	DetectorRuleSummaryManagedListTypesCity         DetectorRuleSummaryManagedListTypesEnum = "CITY"
	DetectorRuleSummaryManagedListTypesTags         DetectorRuleSummaryManagedListTypesEnum = "TAGS"
	DetectorRuleSummaryManagedListTypesGeneric      DetectorRuleSummaryManagedListTypesEnum = "GENERIC"
)

var mappingDetectorRuleSummaryManagedListTypes = map[string]DetectorRuleSummaryManagedListTypesEnum{
	"CIDR_BLOCK":    DetectorRuleSummaryManagedListTypesCidrBlock,
	"USERS":         DetectorRuleSummaryManagedListTypesUsers,
	"GROUPS":        DetectorRuleSummaryManagedListTypesGroups,
	"IPV4ADDRESS":   DetectorRuleSummaryManagedListTypesIpv4address,
	"IPV6ADDRESS":   DetectorRuleSummaryManagedListTypesIpv6address,
	"RESOURCE_OCID": DetectorRuleSummaryManagedListTypesResourceOcid,
	"REGION":        DetectorRuleSummaryManagedListTypesRegion,
	"COUNTRY":       DetectorRuleSummaryManagedListTypesCountry,
	"STATE":         DetectorRuleSummaryManagedListTypesState,
	"CITY":          DetectorRuleSummaryManagedListTypesCity,
	"TAGS":          DetectorRuleSummaryManagedListTypesTags,
	"GENERIC":       DetectorRuleSummaryManagedListTypesGeneric,
}

// GetDetectorRuleSummaryManagedListTypesEnumValues Enumerates the set of values for DetectorRuleSummaryManagedListTypesEnum
func GetDetectorRuleSummaryManagedListTypesEnumValues() []DetectorRuleSummaryManagedListTypesEnum {
	values := make([]DetectorRuleSummaryManagedListTypesEnum, 0)
	for _, v := range mappingDetectorRuleSummaryManagedListTypes {
		values = append(values, v)
	}
	return values
}
