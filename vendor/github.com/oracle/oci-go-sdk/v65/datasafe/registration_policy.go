// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RegistrationPolicy A registration policy.
// This object contains detailed information about a registration policy, including its ID, compartment ID, display name, and features.
type RegistrationPolicy struct {

	// The OCID for the compartment containing the registration policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the registration policy.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the registration policy.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the resource used in the registration policy.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The Data Safe features granted to the databases registering under the registration policy.
	Features []RegistrationPolicyFeaturesEnum `mandatory:"true" json:"features"`

	// The lifecycle state of the registration policy.
	// - CREATING - The registration policy is getting created.
	// - ACTIVE  - The registration policy has been successfully created.
	// - UPDATING - The registration policy is getting updated.
	// - NEEDS_ATTENTION - The registration policy needs attention.
	// - FAILED - The registration policy creation or update has failed.
	// - DELETING - The registration policy is getting deleted.
	LifecycleState RegistrationPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the registration policy was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time when the registration policy was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// A description of the registration policy.
	Description *string `mandatory:"false" json:"description"`

	// The resource type which has been opted in for the registration policy.
	// - DATABASE - The registration policy will be opted in at the Container Database level.
	EnablementLevel RegistrationPolicyEnablementLevelEnum `mandatory:"false" json:"enablementLevel,omitempty"`

	// Indicates whether features will be overridden.
	CanOverrideFeatures *bool `mandatory:"false" json:"canOverrideFeatures"`

	// Details about the lifecycle state of the registration policy
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	ConnectionOption *PolicyConnectionOption `mandatory:"false" json:"connectionOption"`
}

func (m RegistrationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegistrationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Features {
		if _, ok := GetMappingRegistrationPolicyFeaturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Features: %s. Supported values are: %s.", val, strings.Join(GetRegistrationPolicyFeaturesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingRegistrationPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRegistrationPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRegistrationPolicyEnablementLevelEnum(string(m.EnablementLevel)); !ok && m.EnablementLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnablementLevel: %s. Supported values are: %s.", m.EnablementLevel, strings.Join(GetRegistrationPolicyEnablementLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RegistrationPolicyEnablementLevelEnum Enum with underlying type: string
type RegistrationPolicyEnablementLevelEnum string

// Set of constants representing the allowable values for RegistrationPolicyEnablementLevelEnum
const (
	RegistrationPolicyEnablementLevelDatabase RegistrationPolicyEnablementLevelEnum = "DATABASE"
)

var mappingRegistrationPolicyEnablementLevelEnum = map[string]RegistrationPolicyEnablementLevelEnum{
	"DATABASE": RegistrationPolicyEnablementLevelDatabase,
}

var mappingRegistrationPolicyEnablementLevelEnumLowerCase = map[string]RegistrationPolicyEnablementLevelEnum{
	"database": RegistrationPolicyEnablementLevelDatabase,
}

// GetRegistrationPolicyEnablementLevelEnumValues Enumerates the set of values for RegistrationPolicyEnablementLevelEnum
func GetRegistrationPolicyEnablementLevelEnumValues() []RegistrationPolicyEnablementLevelEnum {
	values := make([]RegistrationPolicyEnablementLevelEnum, 0)
	for _, v := range mappingRegistrationPolicyEnablementLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetRegistrationPolicyEnablementLevelEnumStringValues Enumerates the set of values in String for RegistrationPolicyEnablementLevelEnum
func GetRegistrationPolicyEnablementLevelEnumStringValues() []string {
	return []string{
		"DATABASE",
	}
}

// GetMappingRegistrationPolicyEnablementLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegistrationPolicyEnablementLevelEnum(val string) (RegistrationPolicyEnablementLevelEnum, bool) {
	enum, ok := mappingRegistrationPolicyEnablementLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RegistrationPolicyLifecycleStateEnum Enum with underlying type: string
type RegistrationPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for RegistrationPolicyLifecycleStateEnum
const (
	RegistrationPolicyLifecycleStateCreating       RegistrationPolicyLifecycleStateEnum = "CREATING"
	RegistrationPolicyLifecycleStateActive         RegistrationPolicyLifecycleStateEnum = "ACTIVE"
	RegistrationPolicyLifecycleStateUpdating       RegistrationPolicyLifecycleStateEnum = "UPDATING"
	RegistrationPolicyLifecycleStateNeedsAttention RegistrationPolicyLifecycleStateEnum = "NEEDS_ATTENTION"
	RegistrationPolicyLifecycleStateFailed         RegistrationPolicyLifecycleStateEnum = "FAILED"
	RegistrationPolicyLifecycleStateDeleting       RegistrationPolicyLifecycleStateEnum = "DELETING"
)

var mappingRegistrationPolicyLifecycleStateEnum = map[string]RegistrationPolicyLifecycleStateEnum{
	"CREATING":        RegistrationPolicyLifecycleStateCreating,
	"ACTIVE":          RegistrationPolicyLifecycleStateActive,
	"UPDATING":        RegistrationPolicyLifecycleStateUpdating,
	"NEEDS_ATTENTION": RegistrationPolicyLifecycleStateNeedsAttention,
	"FAILED":          RegistrationPolicyLifecycleStateFailed,
	"DELETING":        RegistrationPolicyLifecycleStateDeleting,
}

var mappingRegistrationPolicyLifecycleStateEnumLowerCase = map[string]RegistrationPolicyLifecycleStateEnum{
	"creating":        RegistrationPolicyLifecycleStateCreating,
	"active":          RegistrationPolicyLifecycleStateActive,
	"updating":        RegistrationPolicyLifecycleStateUpdating,
	"needs_attention": RegistrationPolicyLifecycleStateNeedsAttention,
	"failed":          RegistrationPolicyLifecycleStateFailed,
	"deleting":        RegistrationPolicyLifecycleStateDeleting,
}

// GetRegistrationPolicyLifecycleStateEnumValues Enumerates the set of values for RegistrationPolicyLifecycleStateEnum
func GetRegistrationPolicyLifecycleStateEnumValues() []RegistrationPolicyLifecycleStateEnum {
	values := make([]RegistrationPolicyLifecycleStateEnum, 0)
	for _, v := range mappingRegistrationPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRegistrationPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for RegistrationPolicyLifecycleStateEnum
func GetRegistrationPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
	}
}

// GetMappingRegistrationPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegistrationPolicyLifecycleStateEnum(val string) (RegistrationPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingRegistrationPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
