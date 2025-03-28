// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AgentDependency Description of the AgentDependency, which is a sub-resource of the external environment.
type AgentDependency struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Display name of the Agent dependency.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the dependency type. This should match the whitelisted enum of dependency names.
	DependencyName *string `mandatory:"true" json:"dependencyName"`

	// Object storage namespace associated with the customer's tenancy.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Object storage bucket where the Agent dependency is uploaded.
	Bucket *string `mandatory:"true" json:"bucket"`

	// Name of the dependency object uploaded by the customer.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Version of the Agent dependency.
	DependencyVersion *string `mandatory:"false" json:"dependencyVersion"`

	// Description about the Agent dependency.
	Description *string `mandatory:"false" json:"description"`

	// The time when the AgentDependency was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The eTag associated with the dependency object returned by Object Storage.
	ETag *string `mandatory:"false" json:"eTag"`

	// The checksum associated with the dependency object returned by Object Storage.
	Checksum *string `mandatory:"false" json:"checksum"`

	// The current state of AgentDependency.
	LifecycleState AgentDependencyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AgentDependency) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AgentDependency) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAgentDependencyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAgentDependencyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AgentDependencyLifecycleStateEnum Enum with underlying type: string
type AgentDependencyLifecycleStateEnum string

// Set of constants representing the allowable values for AgentDependencyLifecycleStateEnum
const (
	AgentDependencyLifecycleStateCreating AgentDependencyLifecycleStateEnum = "CREATING"
	AgentDependencyLifecycleStateUpdating AgentDependencyLifecycleStateEnum = "UPDATING"
	AgentDependencyLifecycleStateActive   AgentDependencyLifecycleStateEnum = "ACTIVE"
	AgentDependencyLifecycleStateDeleting AgentDependencyLifecycleStateEnum = "DELETING"
	AgentDependencyLifecycleStateDeleted  AgentDependencyLifecycleStateEnum = "DELETED"
	AgentDependencyLifecycleStateFailed   AgentDependencyLifecycleStateEnum = "FAILED"
)

var mappingAgentDependencyLifecycleStateEnum = map[string]AgentDependencyLifecycleStateEnum{
	"CREATING": AgentDependencyLifecycleStateCreating,
	"UPDATING": AgentDependencyLifecycleStateUpdating,
	"ACTIVE":   AgentDependencyLifecycleStateActive,
	"DELETING": AgentDependencyLifecycleStateDeleting,
	"DELETED":  AgentDependencyLifecycleStateDeleted,
	"FAILED":   AgentDependencyLifecycleStateFailed,
}

var mappingAgentDependencyLifecycleStateEnumLowerCase = map[string]AgentDependencyLifecycleStateEnum{
	"creating": AgentDependencyLifecycleStateCreating,
	"updating": AgentDependencyLifecycleStateUpdating,
	"active":   AgentDependencyLifecycleStateActive,
	"deleting": AgentDependencyLifecycleStateDeleting,
	"deleted":  AgentDependencyLifecycleStateDeleted,
	"failed":   AgentDependencyLifecycleStateFailed,
}

// GetAgentDependencyLifecycleStateEnumValues Enumerates the set of values for AgentDependencyLifecycleStateEnum
func GetAgentDependencyLifecycleStateEnumValues() []AgentDependencyLifecycleStateEnum {
	values := make([]AgentDependencyLifecycleStateEnum, 0)
	for _, v := range mappingAgentDependencyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAgentDependencyLifecycleStateEnumStringValues Enumerates the set of values in String for AgentDependencyLifecycleStateEnum
func GetAgentDependencyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAgentDependencyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgentDependencyLifecycleStateEnum(val string) (AgentDependencyLifecycleStateEnum, bool) {
	enum, ok := mappingAgentDependencyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
