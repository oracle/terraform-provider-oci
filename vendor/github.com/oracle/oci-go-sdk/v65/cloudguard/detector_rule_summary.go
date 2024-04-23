// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DetectorRuleSummary Summary information for a detector rule.
type DetectorRuleSummary struct {

	// The unique identifier of the detector rule
	Id *string `mandatory:"true" json:"id"`

	// Possible types of detectors
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// Display name for the detector rule
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for the detector rule
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for the detector rule
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// Service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"false" json:"serviceType"`

	// Resource type of the configuration to which the rule is applied
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// List of managed list types related to this rule
	ManagedListTypes []DetectorRuleSummaryManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	// List of responder rules that can be used to remediate a problem triggered by this detector rule
	CandidateResponderRules []CandidateResponderRule `mandatory:"false" json:"candidateResponderRules"`

	DetectorDetails *DetectorDetails `mandatory:"false" json:"detectorDetails"`

	// The date and time the detector rule was first created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the detector rule was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current lifecycle state of the detector rule
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m DetectorRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectorRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectorEnumEnum(string(m.Detector)); !ok && m.Detector != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detector: %s. Supported values are: %s.", m.Detector, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}

	for _, val := range m.ManagedListTypes {
		if _, ok := GetMappingDetectorRuleSummaryManagedListTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedListTypes: %s. Supported values are: %s.", val, strings.Join(GetDetectorRuleSummaryManagedListTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DetectorRuleSummaryManagedListTypesEnum Enum with underlying type: string
type DetectorRuleSummaryManagedListTypesEnum string

// Set of constants representing the allowable values for DetectorRuleSummaryManagedListTypesEnum
const (
	DetectorRuleSummaryManagedListTypesCidrBlock            DetectorRuleSummaryManagedListTypesEnum = "CIDR_BLOCK"
	DetectorRuleSummaryManagedListTypesUsers                DetectorRuleSummaryManagedListTypesEnum = "USERS"
	DetectorRuleSummaryManagedListTypesGroups               DetectorRuleSummaryManagedListTypesEnum = "GROUPS"
	DetectorRuleSummaryManagedListTypesIpv4address          DetectorRuleSummaryManagedListTypesEnum = "IPV4ADDRESS"
	DetectorRuleSummaryManagedListTypesIpv6address          DetectorRuleSummaryManagedListTypesEnum = "IPV6ADDRESS"
	DetectorRuleSummaryManagedListTypesResourceOcid         DetectorRuleSummaryManagedListTypesEnum = "RESOURCE_OCID"
	DetectorRuleSummaryManagedListTypesRegion               DetectorRuleSummaryManagedListTypesEnum = "REGION"
	DetectorRuleSummaryManagedListTypesCountry              DetectorRuleSummaryManagedListTypesEnum = "COUNTRY"
	DetectorRuleSummaryManagedListTypesState                DetectorRuleSummaryManagedListTypesEnum = "STATE"
	DetectorRuleSummaryManagedListTypesCity                 DetectorRuleSummaryManagedListTypesEnum = "CITY"
	DetectorRuleSummaryManagedListTypesTags                 DetectorRuleSummaryManagedListTypesEnum = "TAGS"
	DetectorRuleSummaryManagedListTypesGeneric              DetectorRuleSummaryManagedListTypesEnum = "GENERIC"
	DetectorRuleSummaryManagedListTypesFusionAppsRole       DetectorRuleSummaryManagedListTypesEnum = "FUSION_APPS_ROLE"
	DetectorRuleSummaryManagedListTypesFusionAppsPermission DetectorRuleSummaryManagedListTypesEnum = "FUSION_APPS_PERMISSION"
)

var mappingDetectorRuleSummaryManagedListTypesEnum = map[string]DetectorRuleSummaryManagedListTypesEnum{
	"CIDR_BLOCK":             DetectorRuleSummaryManagedListTypesCidrBlock,
	"USERS":                  DetectorRuleSummaryManagedListTypesUsers,
	"GROUPS":                 DetectorRuleSummaryManagedListTypesGroups,
	"IPV4ADDRESS":            DetectorRuleSummaryManagedListTypesIpv4address,
	"IPV6ADDRESS":            DetectorRuleSummaryManagedListTypesIpv6address,
	"RESOURCE_OCID":          DetectorRuleSummaryManagedListTypesResourceOcid,
	"REGION":                 DetectorRuleSummaryManagedListTypesRegion,
	"COUNTRY":                DetectorRuleSummaryManagedListTypesCountry,
	"STATE":                  DetectorRuleSummaryManagedListTypesState,
	"CITY":                   DetectorRuleSummaryManagedListTypesCity,
	"TAGS":                   DetectorRuleSummaryManagedListTypesTags,
	"GENERIC":                DetectorRuleSummaryManagedListTypesGeneric,
	"FUSION_APPS_ROLE":       DetectorRuleSummaryManagedListTypesFusionAppsRole,
	"FUSION_APPS_PERMISSION": DetectorRuleSummaryManagedListTypesFusionAppsPermission,
}

var mappingDetectorRuleSummaryManagedListTypesEnumLowerCase = map[string]DetectorRuleSummaryManagedListTypesEnum{
	"cidr_block":             DetectorRuleSummaryManagedListTypesCidrBlock,
	"users":                  DetectorRuleSummaryManagedListTypesUsers,
	"groups":                 DetectorRuleSummaryManagedListTypesGroups,
	"ipv4address":            DetectorRuleSummaryManagedListTypesIpv4address,
	"ipv6address":            DetectorRuleSummaryManagedListTypesIpv6address,
	"resource_ocid":          DetectorRuleSummaryManagedListTypesResourceOcid,
	"region":                 DetectorRuleSummaryManagedListTypesRegion,
	"country":                DetectorRuleSummaryManagedListTypesCountry,
	"state":                  DetectorRuleSummaryManagedListTypesState,
	"city":                   DetectorRuleSummaryManagedListTypesCity,
	"tags":                   DetectorRuleSummaryManagedListTypesTags,
	"generic":                DetectorRuleSummaryManagedListTypesGeneric,
	"fusion_apps_role":       DetectorRuleSummaryManagedListTypesFusionAppsRole,
	"fusion_apps_permission": DetectorRuleSummaryManagedListTypesFusionAppsPermission,
}

// GetDetectorRuleSummaryManagedListTypesEnumValues Enumerates the set of values for DetectorRuleSummaryManagedListTypesEnum
func GetDetectorRuleSummaryManagedListTypesEnumValues() []DetectorRuleSummaryManagedListTypesEnum {
	values := make([]DetectorRuleSummaryManagedListTypesEnum, 0)
	for _, v := range mappingDetectorRuleSummaryManagedListTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectorRuleSummaryManagedListTypesEnumStringValues Enumerates the set of values in String for DetectorRuleSummaryManagedListTypesEnum
func GetDetectorRuleSummaryManagedListTypesEnumStringValues() []string {
	return []string{
		"CIDR_BLOCK",
		"USERS",
		"GROUPS",
		"IPV4ADDRESS",
		"IPV6ADDRESS",
		"RESOURCE_OCID",
		"REGION",
		"COUNTRY",
		"STATE",
		"CITY",
		"TAGS",
		"GENERIC",
		"FUSION_APPS_ROLE",
		"FUSION_APPS_PERMISSION",
	}
}

// GetMappingDetectorRuleSummaryManagedListTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectorRuleSummaryManagedListTypesEnum(val string) (DetectorRuleSummaryManagedListTypesEnum, bool) {
	enum, ok := mappingDetectorRuleSummaryManagedListTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
