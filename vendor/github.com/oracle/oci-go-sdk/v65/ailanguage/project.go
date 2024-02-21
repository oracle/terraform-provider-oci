// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Project Project enable users to organize their project resources.
type Project struct {

	// Unique identifier OCID of the project
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  for the project's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the resource was created in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the project.
	LifecycleState ProjectLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A short description of a project.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the resource was updated in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
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
	ProjectLifecycleStateDeleting ProjectLifecycleStateEnum = "DELETING"
	ProjectLifecycleStateDeleted  ProjectLifecycleStateEnum = "DELETED"
	ProjectLifecycleStateFailed   ProjectLifecycleStateEnum = "FAILED"
	ProjectLifecycleStateCreating ProjectLifecycleStateEnum = "CREATING"
	ProjectLifecycleStateActive   ProjectLifecycleStateEnum = "ACTIVE"
	ProjectLifecycleStateUpdating ProjectLifecycleStateEnum = "UPDATING"
)

var mappingProjectLifecycleStateEnum = map[string]ProjectLifecycleStateEnum{
	"DELETING": ProjectLifecycleStateDeleting,
	"DELETED":  ProjectLifecycleStateDeleted,
	"FAILED":   ProjectLifecycleStateFailed,
	"CREATING": ProjectLifecycleStateCreating,
	"ACTIVE":   ProjectLifecycleStateActive,
	"UPDATING": ProjectLifecycleStateUpdating,
}

var mappingProjectLifecycleStateEnumLowerCase = map[string]ProjectLifecycleStateEnum{
	"deleting": ProjectLifecycleStateDeleting,
	"deleted":  ProjectLifecycleStateDeleted,
	"failed":   ProjectLifecycleStateFailed,
	"creating": ProjectLifecycleStateCreating,
	"active":   ProjectLifecycleStateActive,
	"updating": ProjectLifecycleStateUpdating,
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
		"DELETING",
		"DELETED",
		"FAILED",
		"CREATING",
		"ACTIVE",
		"UPDATING",
	}
}

// GetMappingProjectLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProjectLifecycleStateEnum(val string) (ProjectLifecycleStateEnum, bool) {
	enum, ok := mappingProjectLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
