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

// DetectorRule Detector
type DetectorRule struct {

	// The unique identifier of the detector rule
	Id *string `mandatory:"true" json:"id"`

	// detector for the rule
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"true" json:"serviceType"`

	// resource type of the configuration to which the rule is applied
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// displayName
	DisplayName *string `mandatory:"false" json:"displayName"`

	// description for DetectorRule
	Description *string `mandatory:"false" json:"description"`

	// recommendation for DetectorRule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	DetectorDetails *DetectorDetails `mandatory:"false" json:"detectorDetails"`

	// List of cloudguard managed list types related to this rule
	ManagedListTypes []DetectorRuleManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	// List of CandidateResponderRule related to this rule
	CandidateResponderRules []CandidateResponderRule `mandatory:"false" json:"candidateResponderRules"`

	// The date and time the detector rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the detector rule was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the DetectorRule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m DetectorRule) String() string {
	return common.PointerString(m)
}

// DetectorRuleManagedListTypesEnum Enum with underlying type: string
type DetectorRuleManagedListTypesEnum string

// Set of constants representing the allowable values for DetectorRuleManagedListTypesEnum
const (
	DetectorRuleManagedListTypesCidrBlock    DetectorRuleManagedListTypesEnum = "CIDR_BLOCK"
	DetectorRuleManagedListTypesUsers        DetectorRuleManagedListTypesEnum = "USERS"
	DetectorRuleManagedListTypesGroups       DetectorRuleManagedListTypesEnum = "GROUPS"
	DetectorRuleManagedListTypesIpv4address  DetectorRuleManagedListTypesEnum = "IPV4ADDRESS"
	DetectorRuleManagedListTypesIpv6address  DetectorRuleManagedListTypesEnum = "IPV6ADDRESS"
	DetectorRuleManagedListTypesResourceOcid DetectorRuleManagedListTypesEnum = "RESOURCE_OCID"
	DetectorRuleManagedListTypesRegion       DetectorRuleManagedListTypesEnum = "REGION"
	DetectorRuleManagedListTypesCountry      DetectorRuleManagedListTypesEnum = "COUNTRY"
	DetectorRuleManagedListTypesState        DetectorRuleManagedListTypesEnum = "STATE"
	DetectorRuleManagedListTypesCity         DetectorRuleManagedListTypesEnum = "CITY"
	DetectorRuleManagedListTypesTags         DetectorRuleManagedListTypesEnum = "TAGS"
	DetectorRuleManagedListTypesGeneric      DetectorRuleManagedListTypesEnum = "GENERIC"
)

var mappingDetectorRuleManagedListTypes = map[string]DetectorRuleManagedListTypesEnum{
	"CIDR_BLOCK":    DetectorRuleManagedListTypesCidrBlock,
	"USERS":         DetectorRuleManagedListTypesUsers,
	"GROUPS":        DetectorRuleManagedListTypesGroups,
	"IPV4ADDRESS":   DetectorRuleManagedListTypesIpv4address,
	"IPV6ADDRESS":   DetectorRuleManagedListTypesIpv6address,
	"RESOURCE_OCID": DetectorRuleManagedListTypesResourceOcid,
	"REGION":        DetectorRuleManagedListTypesRegion,
	"COUNTRY":       DetectorRuleManagedListTypesCountry,
	"STATE":         DetectorRuleManagedListTypesState,
	"CITY":          DetectorRuleManagedListTypesCity,
	"TAGS":          DetectorRuleManagedListTypesTags,
	"GENERIC":       DetectorRuleManagedListTypesGeneric,
}

// GetDetectorRuleManagedListTypesEnumValues Enumerates the set of values for DetectorRuleManagedListTypesEnum
func GetDetectorRuleManagedListTypesEnumValues() []DetectorRuleManagedListTypesEnum {
	values := make([]DetectorRuleManagedListTypesEnum, 0)
	for _, v := range mappingDetectorRuleManagedListTypes {
		values = append(values, v)
	}
	return values
}
