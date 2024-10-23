// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Onboarding Fleet Application Management onboarding resource.
// The onboarding process lets Fleet Application Management create a few required policies that you need to start using it and its features.
type Onboarding struct {

	// The unique id of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Onboarding.
	LifecycleState OnboardingLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A value determining if the Fleet Application Management tagging is enabled or not.
	// Allow Fleet Application Management to tag resources with fleet name using "Oracle$FAMS-Tags.FleetName" tag.
	IsFamsTagEnabled *bool `mandatory:"false" json:"isFamsTagEnabled"`

	// The version of Fleet Application Management that the tenant is onboarded to.
	Version *string `mandatory:"false" json:"version"`

	// A value determining if the cost tracking tag is enabled or not.
	// Allow Fleet Application Management to tag resources with cost tracking tag using "Oracle$FAMS-Tags.FAMSManaged" tag.
	IsCostTrackingTagEnabled *bool `mandatory:"false" json:"isCostTrackingTagEnabled"`

	AppliedPolicies *OnboardingPolicySummary `mandatory:"false" json:"appliedPolicies"`

	// Provide discovery frequency.
	DiscoveryFrequency *string `mandatory:"false" json:"discoveryFrequency"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Onboarding) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Onboarding) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOnboardingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOnboardingLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OnboardingLifecycleStateEnum Enum with underlying type: string
type OnboardingLifecycleStateEnum string

// Set of constants representing the allowable values for OnboardingLifecycleStateEnum
const (
	OnboardingLifecycleStateActive         OnboardingLifecycleStateEnum = "ACTIVE"
	OnboardingLifecycleStateInactive       OnboardingLifecycleStateEnum = "INACTIVE"
	OnboardingLifecycleStateCreating       OnboardingLifecycleStateEnum = "CREATING"
	OnboardingLifecycleStateDeleted        OnboardingLifecycleStateEnum = "DELETED"
	OnboardingLifecycleStateDeleting       OnboardingLifecycleStateEnum = "DELETING"
	OnboardingLifecycleStateFailed         OnboardingLifecycleStateEnum = "FAILED"
	OnboardingLifecycleStateUpdating       OnboardingLifecycleStateEnum = "UPDATING"
	OnboardingLifecycleStateNeedsAttention OnboardingLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingOnboardingLifecycleStateEnum = map[string]OnboardingLifecycleStateEnum{
	"ACTIVE":          OnboardingLifecycleStateActive,
	"INACTIVE":        OnboardingLifecycleStateInactive,
	"CREATING":        OnboardingLifecycleStateCreating,
	"DELETED":         OnboardingLifecycleStateDeleted,
	"DELETING":        OnboardingLifecycleStateDeleting,
	"FAILED":          OnboardingLifecycleStateFailed,
	"UPDATING":        OnboardingLifecycleStateUpdating,
	"NEEDS_ATTENTION": OnboardingLifecycleStateNeedsAttention,
}

var mappingOnboardingLifecycleStateEnumLowerCase = map[string]OnboardingLifecycleStateEnum{
	"active":          OnboardingLifecycleStateActive,
	"inactive":        OnboardingLifecycleStateInactive,
	"creating":        OnboardingLifecycleStateCreating,
	"deleted":         OnboardingLifecycleStateDeleted,
	"deleting":        OnboardingLifecycleStateDeleting,
	"failed":          OnboardingLifecycleStateFailed,
	"updating":        OnboardingLifecycleStateUpdating,
	"needs_attention": OnboardingLifecycleStateNeedsAttention,
}

// GetOnboardingLifecycleStateEnumValues Enumerates the set of values for OnboardingLifecycleStateEnum
func GetOnboardingLifecycleStateEnumValues() []OnboardingLifecycleStateEnum {
	values := make([]OnboardingLifecycleStateEnum, 0)
	for _, v := range mappingOnboardingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOnboardingLifecycleStateEnumStringValues Enumerates the set of values in String for OnboardingLifecycleStateEnum
func GetOnboardingLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
		"NEEDS_ATTENTION",
	}
}

// GetMappingOnboardingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOnboardingLifecycleStateEnum(val string) (OnboardingLifecycleStateEnum, bool) {
	enum, ok := mappingOnboardingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
