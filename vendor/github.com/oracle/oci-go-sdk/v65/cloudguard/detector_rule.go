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

// DetectorRule Information for a detector rule (DetectorRule object).
type DetectorRule struct {

	// The unique identifier of the detector rule.
	Id *string `mandatory:"true" json:"id"`

	// detector for the rule
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"true" json:"serviceType"`

	// resource type of the configuration to which the rule is applied
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Display name for DetectorRule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for DetectorRule.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectorRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectorEnumEnum(string(m.Detector)); !ok && m.Detector != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detector: %s. Supported values are: %s.", m.Detector, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}

	for _, val := range m.ManagedListTypes {
		if _, ok := GetMappingDetectorRuleManagedListTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedListTypes: %s. Supported values are: %s.", val, strings.Join(GetDetectorRuleManagedListTypesEnumStringValues(), ",")))
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

// DetectorRuleManagedListTypesEnum Enum with underlying type: string
type DetectorRuleManagedListTypesEnum string

// Set of constants representing the allowable values for DetectorRuleManagedListTypesEnum
const (
	DetectorRuleManagedListTypesCidrBlock            DetectorRuleManagedListTypesEnum = "CIDR_BLOCK"
	DetectorRuleManagedListTypesUsers                DetectorRuleManagedListTypesEnum = "USERS"
	DetectorRuleManagedListTypesGroups               DetectorRuleManagedListTypesEnum = "GROUPS"
	DetectorRuleManagedListTypesIpv4address          DetectorRuleManagedListTypesEnum = "IPV4ADDRESS"
	DetectorRuleManagedListTypesIpv6address          DetectorRuleManagedListTypesEnum = "IPV6ADDRESS"
	DetectorRuleManagedListTypesResourceOcid         DetectorRuleManagedListTypesEnum = "RESOURCE_OCID"
	DetectorRuleManagedListTypesRegion               DetectorRuleManagedListTypesEnum = "REGION"
	DetectorRuleManagedListTypesCountry              DetectorRuleManagedListTypesEnum = "COUNTRY"
	DetectorRuleManagedListTypesState                DetectorRuleManagedListTypesEnum = "STATE"
	DetectorRuleManagedListTypesCity                 DetectorRuleManagedListTypesEnum = "CITY"
	DetectorRuleManagedListTypesTags                 DetectorRuleManagedListTypesEnum = "TAGS"
	DetectorRuleManagedListTypesGeneric              DetectorRuleManagedListTypesEnum = "GENERIC"
	DetectorRuleManagedListTypesFusionAppsRole       DetectorRuleManagedListTypesEnum = "FUSION_APPS_ROLE"
	DetectorRuleManagedListTypesFusionAppsPermission DetectorRuleManagedListTypesEnum = "FUSION_APPS_PERMISSION"
)

var mappingDetectorRuleManagedListTypesEnum = map[string]DetectorRuleManagedListTypesEnum{
	"CIDR_BLOCK":             DetectorRuleManagedListTypesCidrBlock,
	"USERS":                  DetectorRuleManagedListTypesUsers,
	"GROUPS":                 DetectorRuleManagedListTypesGroups,
	"IPV4ADDRESS":            DetectorRuleManagedListTypesIpv4address,
	"IPV6ADDRESS":            DetectorRuleManagedListTypesIpv6address,
	"RESOURCE_OCID":          DetectorRuleManagedListTypesResourceOcid,
	"REGION":                 DetectorRuleManagedListTypesRegion,
	"COUNTRY":                DetectorRuleManagedListTypesCountry,
	"STATE":                  DetectorRuleManagedListTypesState,
	"CITY":                   DetectorRuleManagedListTypesCity,
	"TAGS":                   DetectorRuleManagedListTypesTags,
	"GENERIC":                DetectorRuleManagedListTypesGeneric,
	"FUSION_APPS_ROLE":       DetectorRuleManagedListTypesFusionAppsRole,
	"FUSION_APPS_PERMISSION": DetectorRuleManagedListTypesFusionAppsPermission,
}

var mappingDetectorRuleManagedListTypesEnumLowerCase = map[string]DetectorRuleManagedListTypesEnum{
	"cidr_block":             DetectorRuleManagedListTypesCidrBlock,
	"users":                  DetectorRuleManagedListTypesUsers,
	"groups":                 DetectorRuleManagedListTypesGroups,
	"ipv4address":            DetectorRuleManagedListTypesIpv4address,
	"ipv6address":            DetectorRuleManagedListTypesIpv6address,
	"resource_ocid":          DetectorRuleManagedListTypesResourceOcid,
	"region":                 DetectorRuleManagedListTypesRegion,
	"country":                DetectorRuleManagedListTypesCountry,
	"state":                  DetectorRuleManagedListTypesState,
	"city":                   DetectorRuleManagedListTypesCity,
	"tags":                   DetectorRuleManagedListTypesTags,
	"generic":                DetectorRuleManagedListTypesGeneric,
	"fusion_apps_role":       DetectorRuleManagedListTypesFusionAppsRole,
	"fusion_apps_permission": DetectorRuleManagedListTypesFusionAppsPermission,
}

// GetDetectorRuleManagedListTypesEnumValues Enumerates the set of values for DetectorRuleManagedListTypesEnum
func GetDetectorRuleManagedListTypesEnumValues() []DetectorRuleManagedListTypesEnum {
	values := make([]DetectorRuleManagedListTypesEnum, 0)
	for _, v := range mappingDetectorRuleManagedListTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectorRuleManagedListTypesEnumStringValues Enumerates the set of values in String for DetectorRuleManagedListTypesEnum
func GetDetectorRuleManagedListTypesEnumStringValues() []string {
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

// GetMappingDetectorRuleManagedListTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectorRuleManagedListTypesEnum(val string) (DetectorRuleManagedListTypesEnum, bool) {
	enum, ok := mappingDetectorRuleManagedListTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
