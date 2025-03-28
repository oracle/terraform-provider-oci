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

// TargetDetectorRecipeDetectorRule A TargetDetectorRecipeDetectorRule resource contains a specific instance of a
// single detector rule in one of the supported detector types (for example,
// activity, configuration, or threat).
// A TargetDetectorRecipeDetectorRule resource:
// * Is effectively a copy of a DetectorRecipeRule resource (made when
// a detector recipe is attached to a target) in which users can make
// certain changes if it’s Oracle-managed, and other changes if it’s user-managed.
// * Is visible on the Cloud Guard Targets, Target Details page.
// * Is effectively located in a specific OCI compartment, through the
// ThreatDetectorRecipe resource to which it belongs.
// * Can be modified by users, programmatically or through the UI.
// * Changes that can be made here apply locally, to resources in OCI compartments
// mapped to the target that attaches the associated detector recipe
// (in a TargetDetectorRecipe resource), and override any changes made in rules
// associated with the corresponding DetectorRecipe resource.
type TargetDetectorRecipeDetectorRule struct {

	// The unique identifier of the detector rule
	DetectorRuleId *string `mandatory:"true" json:"detectorRuleId"`

	// Detector type for the rule
	Detector DetectorEnumEnum `mandatory:"true" json:"detector"`

	// Service type of the configuration to which the rule is applied
	ServiceType *string `mandatory:"true" json:"serviceType"`

	// The type of resource which is monitored by the detector rule.
	// For example, Instance, Database, VCN, Policy. To find the resource type for a
	// particular rule, see Detector Recipe Reference
	//  (https://docs.oracle.com/iaas/cloud-guard/using/detect-recipes.htm#detect-recipes-reference).
	// Or try Detector Recipe Reference
	//  (https://docs.oracle.com/iaas/cloud-guard/using/detect-recipes.htm#detect-recipes-reference).
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Display name for TargetDetectorRecipeDetectorRule resource
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for TargetDetectorRecipeDetectorRule resource
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for TargetDetectorRecipeDetectorRule resource
	Recommendation *string `mandatory:"false" json:"recommendation"`

	Details *TargetDetectorDetails `mandatory:"false" json:"details"`

	// List of managed list types related to this rule
	ManagedListTypes []TargetDetectorRecipeDetectorRuleManagedListTypesEnum `mandatory:"false" json:"managedListTypes,omitempty"`

	// The date and time the target detector recipe rule was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target detector recipe rule was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current lifecycle state of the detector rule
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The ID of the attached data source
	DataSourceId *string `mandatory:"false" json:"dataSourceId"`

	// Data source entities mapping for a detector rule
	EntitiesMappings []EntitiesMapping `mandatory:"false" json:"entitiesMappings"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m TargetDetectorRecipeDetectorRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetDetectorRecipeDetectorRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectorEnumEnum(string(m.Detector)); !ok && m.Detector != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detector: %s. Supported values are: %s.", m.Detector, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}

	for _, val := range m.ManagedListTypes {
		if _, ok := GetMappingTargetDetectorRecipeDetectorRuleManagedListTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedListTypes: %s. Supported values are: %s.", val, strings.Join(GetTargetDetectorRecipeDetectorRuleManagedListTypesEnumStringValues(), ",")))
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

// TargetDetectorRecipeDetectorRuleManagedListTypesEnum Enum with underlying type: string
type TargetDetectorRecipeDetectorRuleManagedListTypesEnum string

// Set of constants representing the allowable values for TargetDetectorRecipeDetectorRuleManagedListTypesEnum
const (
	TargetDetectorRecipeDetectorRuleManagedListTypesCidrBlock            TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "CIDR_BLOCK"
	TargetDetectorRecipeDetectorRuleManagedListTypesUsers                TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "USERS"
	TargetDetectorRecipeDetectorRuleManagedListTypesGroups               TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "GROUPS"
	TargetDetectorRecipeDetectorRuleManagedListTypesIpv4address          TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "IPV4ADDRESS"
	TargetDetectorRecipeDetectorRuleManagedListTypesIpv6address          TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "IPV6ADDRESS"
	TargetDetectorRecipeDetectorRuleManagedListTypesResourceOcid         TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "RESOURCE_OCID"
	TargetDetectorRecipeDetectorRuleManagedListTypesRegion               TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "REGION"
	TargetDetectorRecipeDetectorRuleManagedListTypesCountry              TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "COUNTRY"
	TargetDetectorRecipeDetectorRuleManagedListTypesState                TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "STATE"
	TargetDetectorRecipeDetectorRuleManagedListTypesCity                 TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "CITY"
	TargetDetectorRecipeDetectorRuleManagedListTypesTags                 TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "TAGS"
	TargetDetectorRecipeDetectorRuleManagedListTypesGeneric              TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "GENERIC"
	TargetDetectorRecipeDetectorRuleManagedListTypesFusionAppsRole       TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "FUSION_APPS_ROLE"
	TargetDetectorRecipeDetectorRuleManagedListTypesFusionAppsPermission TargetDetectorRecipeDetectorRuleManagedListTypesEnum = "FUSION_APPS_PERMISSION"
)

var mappingTargetDetectorRecipeDetectorRuleManagedListTypesEnum = map[string]TargetDetectorRecipeDetectorRuleManagedListTypesEnum{
	"CIDR_BLOCK":             TargetDetectorRecipeDetectorRuleManagedListTypesCidrBlock,
	"USERS":                  TargetDetectorRecipeDetectorRuleManagedListTypesUsers,
	"GROUPS":                 TargetDetectorRecipeDetectorRuleManagedListTypesGroups,
	"IPV4ADDRESS":            TargetDetectorRecipeDetectorRuleManagedListTypesIpv4address,
	"IPV6ADDRESS":            TargetDetectorRecipeDetectorRuleManagedListTypesIpv6address,
	"RESOURCE_OCID":          TargetDetectorRecipeDetectorRuleManagedListTypesResourceOcid,
	"REGION":                 TargetDetectorRecipeDetectorRuleManagedListTypesRegion,
	"COUNTRY":                TargetDetectorRecipeDetectorRuleManagedListTypesCountry,
	"STATE":                  TargetDetectorRecipeDetectorRuleManagedListTypesState,
	"CITY":                   TargetDetectorRecipeDetectorRuleManagedListTypesCity,
	"TAGS":                   TargetDetectorRecipeDetectorRuleManagedListTypesTags,
	"GENERIC":                TargetDetectorRecipeDetectorRuleManagedListTypesGeneric,
	"FUSION_APPS_ROLE":       TargetDetectorRecipeDetectorRuleManagedListTypesFusionAppsRole,
	"FUSION_APPS_PERMISSION": TargetDetectorRecipeDetectorRuleManagedListTypesFusionAppsPermission,
}

var mappingTargetDetectorRecipeDetectorRuleManagedListTypesEnumLowerCase = map[string]TargetDetectorRecipeDetectorRuleManagedListTypesEnum{
	"cidr_block":             TargetDetectorRecipeDetectorRuleManagedListTypesCidrBlock,
	"users":                  TargetDetectorRecipeDetectorRuleManagedListTypesUsers,
	"groups":                 TargetDetectorRecipeDetectorRuleManagedListTypesGroups,
	"ipv4address":            TargetDetectorRecipeDetectorRuleManagedListTypesIpv4address,
	"ipv6address":            TargetDetectorRecipeDetectorRuleManagedListTypesIpv6address,
	"resource_ocid":          TargetDetectorRecipeDetectorRuleManagedListTypesResourceOcid,
	"region":                 TargetDetectorRecipeDetectorRuleManagedListTypesRegion,
	"country":                TargetDetectorRecipeDetectorRuleManagedListTypesCountry,
	"state":                  TargetDetectorRecipeDetectorRuleManagedListTypesState,
	"city":                   TargetDetectorRecipeDetectorRuleManagedListTypesCity,
	"tags":                   TargetDetectorRecipeDetectorRuleManagedListTypesTags,
	"generic":                TargetDetectorRecipeDetectorRuleManagedListTypesGeneric,
	"fusion_apps_role":       TargetDetectorRecipeDetectorRuleManagedListTypesFusionAppsRole,
	"fusion_apps_permission": TargetDetectorRecipeDetectorRuleManagedListTypesFusionAppsPermission,
}

// GetTargetDetectorRecipeDetectorRuleManagedListTypesEnumValues Enumerates the set of values for TargetDetectorRecipeDetectorRuleManagedListTypesEnum
func GetTargetDetectorRecipeDetectorRuleManagedListTypesEnumValues() []TargetDetectorRecipeDetectorRuleManagedListTypesEnum {
	values := make([]TargetDetectorRecipeDetectorRuleManagedListTypesEnum, 0)
	for _, v := range mappingTargetDetectorRecipeDetectorRuleManagedListTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetDetectorRecipeDetectorRuleManagedListTypesEnumStringValues Enumerates the set of values in String for TargetDetectorRecipeDetectorRuleManagedListTypesEnum
func GetTargetDetectorRecipeDetectorRuleManagedListTypesEnumStringValues() []string {
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

// GetMappingTargetDetectorRecipeDetectorRuleManagedListTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetDetectorRecipeDetectorRuleManagedListTypesEnum(val string) (TargetDetectorRecipeDetectorRuleManagedListTypesEnum, bool) {
	enum, ok := mappingTargetDetectorRecipeDetectorRuleManagedListTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
