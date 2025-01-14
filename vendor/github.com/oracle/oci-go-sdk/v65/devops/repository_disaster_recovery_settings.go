// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RepositoryDisasterRecoverySettings Users can enable the disaster recovery by creating Disaster Recovery Setting.
type RepositoryDisasterRecoverySettings struct {

	// The compartment OCID where the Disaster Recovery Setting is created. This is always the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the DR Setting.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Primary region identifier referred by the deployment environment. Region identifiers are listed at https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	PrimaryRegion *string `mandatory:"true" json:"primaryRegion"`

	// Standby Region identifier referred by the deployment environment. Region identifiers are listed at https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	StandbyRegion *string `mandatory:"true" json:"standbyRegion"`

	// Disaster Recovery Setting Id.
	Id *string `mandatory:"false" json:"id"`

	// Setting status.
	LifecycleState RepositoryDisasterRecoverySettingsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time the repository was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the repository was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The List of compartment not included in the Disaster Recovery Plan.
	ExcludeCompartments []string `mandatory:"false" json:"excludeCompartments"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RepositoryDisasterRecoverySettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryDisasterRecoverySettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRepositoryDisasterRecoverySettingsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRepositoryDisasterRecoverySettingsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RepositoryDisasterRecoverySettingsLifecycleStateEnum Enum with underlying type: string
type RepositoryDisasterRecoverySettingsLifecycleStateEnum string

// Set of constants representing the allowable values for RepositoryDisasterRecoverySettingsLifecycleStateEnum
const (
	RepositoryDisasterRecoverySettingsLifecycleStateActive         RepositoryDisasterRecoverySettingsLifecycleStateEnum = "ACTIVE"
	RepositoryDisasterRecoverySettingsLifecycleStateCreating       RepositoryDisasterRecoverySettingsLifecycleStateEnum = "CREATING"
	RepositoryDisasterRecoverySettingsLifecycleStateFailed         RepositoryDisasterRecoverySettingsLifecycleStateEnum = "FAILED"
	RepositoryDisasterRecoverySettingsLifecycleStateNeedsAttention RepositoryDisasterRecoverySettingsLifecycleStateEnum = "NEEDS_ATTENTION"
	RepositoryDisasterRecoverySettingsLifecycleStateDeleting       RepositoryDisasterRecoverySettingsLifecycleStateEnum = "DELETING"
)

var mappingRepositoryDisasterRecoverySettingsLifecycleStateEnum = map[string]RepositoryDisasterRecoverySettingsLifecycleStateEnum{
	"ACTIVE":          RepositoryDisasterRecoverySettingsLifecycleStateActive,
	"CREATING":        RepositoryDisasterRecoverySettingsLifecycleStateCreating,
	"FAILED":          RepositoryDisasterRecoverySettingsLifecycleStateFailed,
	"NEEDS_ATTENTION": RepositoryDisasterRecoverySettingsLifecycleStateNeedsAttention,
	"DELETING":        RepositoryDisasterRecoverySettingsLifecycleStateDeleting,
}

var mappingRepositoryDisasterRecoverySettingsLifecycleStateEnumLowerCase = map[string]RepositoryDisasterRecoverySettingsLifecycleStateEnum{
	"active":          RepositoryDisasterRecoverySettingsLifecycleStateActive,
	"creating":        RepositoryDisasterRecoverySettingsLifecycleStateCreating,
	"failed":          RepositoryDisasterRecoverySettingsLifecycleStateFailed,
	"needs_attention": RepositoryDisasterRecoverySettingsLifecycleStateNeedsAttention,
	"deleting":        RepositoryDisasterRecoverySettingsLifecycleStateDeleting,
}

// GetRepositoryDisasterRecoverySettingsLifecycleStateEnumValues Enumerates the set of values for RepositoryDisasterRecoverySettingsLifecycleStateEnum
func GetRepositoryDisasterRecoverySettingsLifecycleStateEnumValues() []RepositoryDisasterRecoverySettingsLifecycleStateEnum {
	values := make([]RepositoryDisasterRecoverySettingsLifecycleStateEnum, 0)
	for _, v := range mappingRepositoryDisasterRecoverySettingsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryDisasterRecoverySettingsLifecycleStateEnumStringValues Enumerates the set of values in String for RepositoryDisasterRecoverySettingsLifecycleStateEnum
func GetRepositoryDisasterRecoverySettingsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"FAILED",
		"NEEDS_ATTENTION",
		"DELETING",
	}
}

// GetMappingRepositoryDisasterRecoverySettingsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryDisasterRecoverySettingsLifecycleStateEnum(val string) (RepositoryDisasterRecoverySettingsLifecycleStateEnum, bool) {
	enum, ok := mappingRepositoryDisasterRecoverySettingsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
