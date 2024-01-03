// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplicationDetails The information about the application.
type ApplicationDetails struct {

	// Generated key that can be used in API calls to identify application.
	Key *string `mandatory:"true" json:"key"`

	// The object type.
	ModelType *string `mandatory:"true" json:"modelType"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// version
	ApplicationVersion *int `mandatory:"false" json:"applicationVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current state of the workspace.
	LifecycleState ApplicationDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ApplicationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApplicationDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApplicationDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApplicationDetailsLifecycleStateEnum Enum with underlying type: string
type ApplicationDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for ApplicationDetailsLifecycleStateEnum
const (
	ApplicationDetailsLifecycleStateCreating ApplicationDetailsLifecycleStateEnum = "CREATING"
	ApplicationDetailsLifecycleStateActive   ApplicationDetailsLifecycleStateEnum = "ACTIVE"
	ApplicationDetailsLifecycleStateUpdating ApplicationDetailsLifecycleStateEnum = "UPDATING"
	ApplicationDetailsLifecycleStateDeleting ApplicationDetailsLifecycleStateEnum = "DELETING"
	ApplicationDetailsLifecycleStateDeleted  ApplicationDetailsLifecycleStateEnum = "DELETED"
	ApplicationDetailsLifecycleStateFailed   ApplicationDetailsLifecycleStateEnum = "FAILED"
)

var mappingApplicationDetailsLifecycleStateEnum = map[string]ApplicationDetailsLifecycleStateEnum{
	"CREATING": ApplicationDetailsLifecycleStateCreating,
	"ACTIVE":   ApplicationDetailsLifecycleStateActive,
	"UPDATING": ApplicationDetailsLifecycleStateUpdating,
	"DELETING": ApplicationDetailsLifecycleStateDeleting,
	"DELETED":  ApplicationDetailsLifecycleStateDeleted,
	"FAILED":   ApplicationDetailsLifecycleStateFailed,
}

var mappingApplicationDetailsLifecycleStateEnumLowerCase = map[string]ApplicationDetailsLifecycleStateEnum{
	"creating": ApplicationDetailsLifecycleStateCreating,
	"active":   ApplicationDetailsLifecycleStateActive,
	"updating": ApplicationDetailsLifecycleStateUpdating,
	"deleting": ApplicationDetailsLifecycleStateDeleting,
	"deleted":  ApplicationDetailsLifecycleStateDeleted,
	"failed":   ApplicationDetailsLifecycleStateFailed,
}

// GetApplicationDetailsLifecycleStateEnumValues Enumerates the set of values for ApplicationDetailsLifecycleStateEnum
func GetApplicationDetailsLifecycleStateEnumValues() []ApplicationDetailsLifecycleStateEnum {
	values := make([]ApplicationDetailsLifecycleStateEnum, 0)
	for _, v := range mappingApplicationDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for ApplicationDetailsLifecycleStateEnum
func GetApplicationDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingApplicationDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationDetailsLifecycleStateEnum(val string) (ApplicationDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingApplicationDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
