// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Project DevOps project groups resources needed to implement the CI/CD workload. DevOps resources include artifacts, pipelines, and environments.
type Project struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Project name (case-sensitive).
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment where the project is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	NotificationConfig *NotificationConfig `mandatory:"true" json:"notificationConfig"`

	// Project description.
	Description *string `mandatory:"false" json:"description"`

	// Namespace associated with the project.
	Namespace *string `mandatory:"false" json:"namespace"`

	// Time the project was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the project was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the project.
	LifecycleState ProjectLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Project) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Project) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProjectLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProjectLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProjectLifecycleStateEnum Enum with underlying type: string
type ProjectLifecycleStateEnum string

// Set of constants representing the allowable values for ProjectLifecycleStateEnum
const (
	ProjectLifecycleStateCreating ProjectLifecycleStateEnum = "CREATING"
	ProjectLifecycleStateUpdating ProjectLifecycleStateEnum = "UPDATING"
	ProjectLifecycleStateActive   ProjectLifecycleStateEnum = "ACTIVE"
	ProjectLifecycleStateDeleting ProjectLifecycleStateEnum = "DELETING"
	ProjectLifecycleStateDeleted  ProjectLifecycleStateEnum = "DELETED"
	ProjectLifecycleStateFailed   ProjectLifecycleStateEnum = "FAILED"
)

var mappingProjectLifecycleStateEnum = map[string]ProjectLifecycleStateEnum{
	"CREATING": ProjectLifecycleStateCreating,
	"UPDATING": ProjectLifecycleStateUpdating,
	"ACTIVE":   ProjectLifecycleStateActive,
	"DELETING": ProjectLifecycleStateDeleting,
	"DELETED":  ProjectLifecycleStateDeleted,
	"FAILED":   ProjectLifecycleStateFailed,
}

// GetProjectLifecycleStateEnumValues Enumerates the set of values for ProjectLifecycleStateEnum
func GetProjectLifecycleStateEnumValues() []ProjectLifecycleStateEnum {
	values := make([]ProjectLifecycleStateEnum, 0)
	for _, v := range mappingProjectLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProjectLifecycleStateEnumStringValues Enumerates the set of values in String for ProjectLifecycleStateEnum
func GetProjectLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingProjectLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProjectLifecycleStateEnum(val string) (ProjectLifecycleStateEnum, bool) {
	mappingProjectLifecycleStateEnumIgnoreCase := make(map[string]ProjectLifecycleStateEnum)
	for k, v := range mappingProjectLifecycleStateEnum {
		mappingProjectLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingProjectLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
