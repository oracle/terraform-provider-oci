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

	// The id of the attached DataSource.
	DataSourceId *string `mandatory:"false" json:"dataSourceId"`

	// Data Source entities mapping for a Detector Rule
	EntitiesMappings []EntitiesMapping `mandatory:"false" json:"entitiesMappings"`
}

func (m TargetDetectorRecipeDetectorRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetDetectorRecipeDetectorRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectorEnumEnum(string(m.Detector)); !ok && m.Detector != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detector: %s. Supported values are: %s.", m.Detector, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}

	for _, val := range m.ManagedListTypes {
		if _, ok := GetMappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedListTypes: %s. Supported values are: %s.", val, strings.Join(GetTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumStringValues(), ",")))
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

// TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum Enum with underlying type: string
type TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum string

// Set of constants representing the allowable values for TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum
const (
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock            TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "CIDR_BLOCK"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesUsers                TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "USERS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGroups               TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "GROUPS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address          TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "IPV4ADDRESS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address          TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "IPV6ADDRESS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid         TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "RESOURCE_OCID"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesRegion               TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "REGION"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCountry              TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "COUNTRY"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesState                TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "STATE"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCity                 TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "CITY"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesTags                 TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "TAGS"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGeneric              TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "GENERIC"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsRole       TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "FUSION_APPS_ROLE"
	TargetDetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsPermission TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "FUSION_APPS_PERMISSION"
)

var mappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = map[string]TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum{
	"CIDR_BLOCK":             TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock,
	"USERS":                  TargetDetectorRecipeDetectorRuleSummaryManagedListTypesUsers,
	"GROUPS":                 TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGroups,
	"IPV4ADDRESS":            TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address,
	"IPV6ADDRESS":            TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address,
	"RESOURCE_OCID":          TargetDetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid,
	"REGION":                 TargetDetectorRecipeDetectorRuleSummaryManagedListTypesRegion,
	"COUNTRY":                TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCountry,
	"STATE":                  TargetDetectorRecipeDetectorRuleSummaryManagedListTypesState,
	"CITY":                   TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCity,
	"TAGS":                   TargetDetectorRecipeDetectorRuleSummaryManagedListTypesTags,
	"GENERIC":                TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGeneric,
	"FUSION_APPS_ROLE":       TargetDetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsRole,
	"FUSION_APPS_PERMISSION": TargetDetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsPermission,
}

var mappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumLowerCase = map[string]TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum{
	"cidr_block":             TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock,
	"users":                  TargetDetectorRecipeDetectorRuleSummaryManagedListTypesUsers,
	"groups":                 TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGroups,
	"ipv4address":            TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address,
	"ipv6address":            TargetDetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address,
	"resource_ocid":          TargetDetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid,
	"region":                 TargetDetectorRecipeDetectorRuleSummaryManagedListTypesRegion,
	"country":                TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCountry,
	"state":                  TargetDetectorRecipeDetectorRuleSummaryManagedListTypesState,
	"city":                   TargetDetectorRecipeDetectorRuleSummaryManagedListTypesCity,
	"tags":                   TargetDetectorRecipeDetectorRuleSummaryManagedListTypesTags,
	"generic":                TargetDetectorRecipeDetectorRuleSummaryManagedListTypesGeneric,
	"fusion_apps_role":       TargetDetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsRole,
	"fusion_apps_permission": TargetDetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsPermission,
}

// GetTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumValues Enumerates the set of values for TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum
func GetTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumValues() []TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum {
	values := make([]TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum, 0)
	for _, v := range mappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumStringValues Enumerates the set of values in String for TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum
func GetTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumStringValues() []string {
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

// GetMappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum(val string) (TargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnum, bool) {
	enum, ok := mappingTargetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
