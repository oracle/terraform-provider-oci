// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DetectorRecipeDetectorRule A DetectorRecipeDetectorRule resource defines a single recipe rule in the collection for a DetectorRecipe resource.
type DetectorRecipeDetectorRule struct {

	// The unique identifier of the detector rule.
	DetectorRuleId *string `mandatory:"true" json:"detectorRuleId"`

	// Detector recipe for the rule
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// Service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"true" json:"serviceType"`

	// Resource type of the configuration to which the rule is applied
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Display name for DetectorRecipeDetectorRule resource
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for DetectorRecipeDetectorRule resource
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for DetectorRecipeDetectorRule resource
	Recommendation *string `mandatory:"false" json:"recommendation"`

	Details *DetectorDetails `mandatory:"false" json:"details"`

	// List of managed list types related to this rule
	ManagedListTypes []DetectorRecipeDetectorRuleManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	// List of responder rules that can be used to remediate this detector rule
	CandidateResponderRules []CandidateResponderRule `mandatory:"false" json:"candidateResponderRules"`

	// The date and time the detector recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the detector recipe rule was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current lifecycle state of the detector rule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The unique identifier of the attached data source
	DataSourceId *string `mandatory:"false" json:"dataSourceId"`

	// Data source entities mapping for the detector rule
	EntitiesMappings []EntitiesMapping `mandatory:"false" json:"entitiesMappings"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m DetectorRecipeDetectorRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectorRecipeDetectorRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectorEnumEnum(string(m.Detector)); !ok && m.Detector != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detector: %s. Supported values are: %s.", m.Detector, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}

	for _, val := range m.ManagedListTypes {
		if _, ok := GetMappingDetectorRecipeDetectorRuleManagedListTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedListTypes: %s. Supported values are: %s.", val, strings.Join(GetDetectorRecipeDetectorRuleManagedListTypesEnumStringValues(), ",")))
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

// DetectorRecipeDetectorRuleManagedListTypesEnum Enum with underlying type: string
type DetectorRecipeDetectorRuleManagedListTypesEnum string

// Set of constants representing the allowable values for DetectorRecipeDetectorRuleManagedListTypesEnum
const (
	DetectorRecipeDetectorRuleManagedListTypesCidrBlock            DetectorRecipeDetectorRuleManagedListTypesEnum = "CIDR_BLOCK"
	DetectorRecipeDetectorRuleManagedListTypesUsers                DetectorRecipeDetectorRuleManagedListTypesEnum = "USERS"
	DetectorRecipeDetectorRuleManagedListTypesGroups               DetectorRecipeDetectorRuleManagedListTypesEnum = "GROUPS"
	DetectorRecipeDetectorRuleManagedListTypesIpv4address          DetectorRecipeDetectorRuleManagedListTypesEnum = "IPV4ADDRESS"
	DetectorRecipeDetectorRuleManagedListTypesIpv6address          DetectorRecipeDetectorRuleManagedListTypesEnum = "IPV6ADDRESS"
	DetectorRecipeDetectorRuleManagedListTypesResourceOcid         DetectorRecipeDetectorRuleManagedListTypesEnum = "RESOURCE_OCID"
	DetectorRecipeDetectorRuleManagedListTypesRegion               DetectorRecipeDetectorRuleManagedListTypesEnum = "REGION"
	DetectorRecipeDetectorRuleManagedListTypesCountry              DetectorRecipeDetectorRuleManagedListTypesEnum = "COUNTRY"
	DetectorRecipeDetectorRuleManagedListTypesState                DetectorRecipeDetectorRuleManagedListTypesEnum = "STATE"
	DetectorRecipeDetectorRuleManagedListTypesCity                 DetectorRecipeDetectorRuleManagedListTypesEnum = "CITY"
	DetectorRecipeDetectorRuleManagedListTypesTags                 DetectorRecipeDetectorRuleManagedListTypesEnum = "TAGS"
	DetectorRecipeDetectorRuleManagedListTypesGeneric              DetectorRecipeDetectorRuleManagedListTypesEnum = "GENERIC"
	DetectorRecipeDetectorRuleManagedListTypesFusionAppsRole       DetectorRecipeDetectorRuleManagedListTypesEnum = "FUSION_APPS_ROLE"
	DetectorRecipeDetectorRuleManagedListTypesFusionAppsPermission DetectorRecipeDetectorRuleManagedListTypesEnum = "FUSION_APPS_PERMISSION"
)

var mappingDetectorRecipeDetectorRuleManagedListTypesEnum = map[string]DetectorRecipeDetectorRuleManagedListTypesEnum{
	"CIDR_BLOCK":             DetectorRecipeDetectorRuleManagedListTypesCidrBlock,
	"USERS":                  DetectorRecipeDetectorRuleManagedListTypesUsers,
	"GROUPS":                 DetectorRecipeDetectorRuleManagedListTypesGroups,
	"IPV4ADDRESS":            DetectorRecipeDetectorRuleManagedListTypesIpv4address,
	"IPV6ADDRESS":            DetectorRecipeDetectorRuleManagedListTypesIpv6address,
	"RESOURCE_OCID":          DetectorRecipeDetectorRuleManagedListTypesResourceOcid,
	"REGION":                 DetectorRecipeDetectorRuleManagedListTypesRegion,
	"COUNTRY":                DetectorRecipeDetectorRuleManagedListTypesCountry,
	"STATE":                  DetectorRecipeDetectorRuleManagedListTypesState,
	"CITY":                   DetectorRecipeDetectorRuleManagedListTypesCity,
	"TAGS":                   DetectorRecipeDetectorRuleManagedListTypesTags,
	"GENERIC":                DetectorRecipeDetectorRuleManagedListTypesGeneric,
	"FUSION_APPS_ROLE":       DetectorRecipeDetectorRuleManagedListTypesFusionAppsRole,
	"FUSION_APPS_PERMISSION": DetectorRecipeDetectorRuleManagedListTypesFusionAppsPermission,
}

var mappingDetectorRecipeDetectorRuleManagedListTypesEnumLowerCase = map[string]DetectorRecipeDetectorRuleManagedListTypesEnum{
	"cidr_block":             DetectorRecipeDetectorRuleManagedListTypesCidrBlock,
	"users":                  DetectorRecipeDetectorRuleManagedListTypesUsers,
	"groups":                 DetectorRecipeDetectorRuleManagedListTypesGroups,
	"ipv4address":            DetectorRecipeDetectorRuleManagedListTypesIpv4address,
	"ipv6address":            DetectorRecipeDetectorRuleManagedListTypesIpv6address,
	"resource_ocid":          DetectorRecipeDetectorRuleManagedListTypesResourceOcid,
	"region":                 DetectorRecipeDetectorRuleManagedListTypesRegion,
	"country":                DetectorRecipeDetectorRuleManagedListTypesCountry,
	"state":                  DetectorRecipeDetectorRuleManagedListTypesState,
	"city":                   DetectorRecipeDetectorRuleManagedListTypesCity,
	"tags":                   DetectorRecipeDetectorRuleManagedListTypesTags,
	"generic":                DetectorRecipeDetectorRuleManagedListTypesGeneric,
	"fusion_apps_role":       DetectorRecipeDetectorRuleManagedListTypesFusionAppsRole,
	"fusion_apps_permission": DetectorRecipeDetectorRuleManagedListTypesFusionAppsPermission,
}

// GetDetectorRecipeDetectorRuleManagedListTypesEnumValues Enumerates the set of values for DetectorRecipeDetectorRuleManagedListTypesEnum
func GetDetectorRecipeDetectorRuleManagedListTypesEnumValues() []DetectorRecipeDetectorRuleManagedListTypesEnum {
	values := make([]DetectorRecipeDetectorRuleManagedListTypesEnum, 0)
	for _, v := range mappingDetectorRecipeDetectorRuleManagedListTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectorRecipeDetectorRuleManagedListTypesEnumStringValues Enumerates the set of values in String for DetectorRecipeDetectorRuleManagedListTypesEnum
func GetDetectorRecipeDetectorRuleManagedListTypesEnumStringValues() []string {
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

// GetMappingDetectorRecipeDetectorRuleManagedListTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectorRecipeDetectorRuleManagedListTypesEnum(val string) (DetectorRecipeDetectorRuleManagedListTypesEnum, bool) {
	enum, ok := mappingDetectorRecipeDetectorRuleManagedListTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
