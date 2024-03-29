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

	// The id of the attached DataSource.
	DataSourceId *string `mandatory:"false" json:"dataSourceId"`

	// Data Source entities mapping for a Detector Rule
	EntitiesMappings []EntitiesMapping `mandatory:"false" json:"entitiesMappings"`
}

func (m DetectorRecipeDetectorRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectorRecipeDetectorRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectorEnumEnum(string(m.Detector)); !ok && m.Detector != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detector: %s. Supported values are: %s.", m.Detector, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}

	for _, val := range m.ManagedListTypes {
		if _, ok := GetMappingDetectorRecipeDetectorRuleSummaryManagedListTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedListTypes: %s. Supported values are: %s.", val, strings.Join(GetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumStringValues(), ",")))
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

// DetectorRecipeDetectorRuleSummaryManagedListTypesEnum Enum with underlying type: string
type DetectorRecipeDetectorRuleSummaryManagedListTypesEnum string

// Set of constants representing the allowable values for DetectorRecipeDetectorRuleSummaryManagedListTypesEnum
const (
	DetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock            DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "CIDR_BLOCK"
	DetectorRecipeDetectorRuleSummaryManagedListTypesUsers                DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "USERS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesGroups               DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "GROUPS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address          DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "IPV4ADDRESS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address          DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "IPV6ADDRESS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid         DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "RESOURCE_OCID"
	DetectorRecipeDetectorRuleSummaryManagedListTypesRegion               DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "REGION"
	DetectorRecipeDetectorRuleSummaryManagedListTypesCountry              DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "COUNTRY"
	DetectorRecipeDetectorRuleSummaryManagedListTypesState                DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "STATE"
	DetectorRecipeDetectorRuleSummaryManagedListTypesCity                 DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "CITY"
	DetectorRecipeDetectorRuleSummaryManagedListTypesTags                 DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "TAGS"
	DetectorRecipeDetectorRuleSummaryManagedListTypesGeneric              DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "GENERIC"
	DetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsRole       DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "FUSION_APPS_ROLE"
	DetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsPermission DetectorRecipeDetectorRuleSummaryManagedListTypesEnum = "FUSION_APPS_PERMISSION"
)

var mappingDetectorRecipeDetectorRuleSummaryManagedListTypesEnum = map[string]DetectorRecipeDetectorRuleSummaryManagedListTypesEnum{
	"CIDR_BLOCK":             DetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock,
	"USERS":                  DetectorRecipeDetectorRuleSummaryManagedListTypesUsers,
	"GROUPS":                 DetectorRecipeDetectorRuleSummaryManagedListTypesGroups,
	"IPV4ADDRESS":            DetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address,
	"IPV6ADDRESS":            DetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address,
	"RESOURCE_OCID":          DetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid,
	"REGION":                 DetectorRecipeDetectorRuleSummaryManagedListTypesRegion,
	"COUNTRY":                DetectorRecipeDetectorRuleSummaryManagedListTypesCountry,
	"STATE":                  DetectorRecipeDetectorRuleSummaryManagedListTypesState,
	"CITY":                   DetectorRecipeDetectorRuleSummaryManagedListTypesCity,
	"TAGS":                   DetectorRecipeDetectorRuleSummaryManagedListTypesTags,
	"GENERIC":                DetectorRecipeDetectorRuleSummaryManagedListTypesGeneric,
	"FUSION_APPS_ROLE":       DetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsRole,
	"FUSION_APPS_PERMISSION": DetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsPermission,
}

var mappingDetectorRecipeDetectorRuleSummaryManagedListTypesEnumLowerCase = map[string]DetectorRecipeDetectorRuleSummaryManagedListTypesEnum{
	"cidr_block":             DetectorRecipeDetectorRuleSummaryManagedListTypesCidrBlock,
	"users":                  DetectorRecipeDetectorRuleSummaryManagedListTypesUsers,
	"groups":                 DetectorRecipeDetectorRuleSummaryManagedListTypesGroups,
	"ipv4address":            DetectorRecipeDetectorRuleSummaryManagedListTypesIpv4address,
	"ipv6address":            DetectorRecipeDetectorRuleSummaryManagedListTypesIpv6address,
	"resource_ocid":          DetectorRecipeDetectorRuleSummaryManagedListTypesResourceOcid,
	"region":                 DetectorRecipeDetectorRuleSummaryManagedListTypesRegion,
	"country":                DetectorRecipeDetectorRuleSummaryManagedListTypesCountry,
	"state":                  DetectorRecipeDetectorRuleSummaryManagedListTypesState,
	"city":                   DetectorRecipeDetectorRuleSummaryManagedListTypesCity,
	"tags":                   DetectorRecipeDetectorRuleSummaryManagedListTypesTags,
	"generic":                DetectorRecipeDetectorRuleSummaryManagedListTypesGeneric,
	"fusion_apps_role":       DetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsRole,
	"fusion_apps_permission": DetectorRecipeDetectorRuleSummaryManagedListTypesFusionAppsPermission,
}

// GetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumValues Enumerates the set of values for DetectorRecipeDetectorRuleSummaryManagedListTypesEnum
func GetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumValues() []DetectorRecipeDetectorRuleSummaryManagedListTypesEnum {
	values := make([]DetectorRecipeDetectorRuleSummaryManagedListTypesEnum, 0)
	for _, v := range mappingDetectorRecipeDetectorRuleSummaryManagedListTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumStringValues Enumerates the set of values in String for DetectorRecipeDetectorRuleSummaryManagedListTypesEnum
func GetDetectorRecipeDetectorRuleSummaryManagedListTypesEnumStringValues() []string {
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

// GetMappingDetectorRecipeDetectorRuleSummaryManagedListTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectorRecipeDetectorRuleSummaryManagedListTypesEnum(val string) (DetectorRecipeDetectorRuleSummaryManagedListTypesEnum, bool) {
	enum, ok := mappingDetectorRecipeDetectorRuleSummaryManagedListTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
