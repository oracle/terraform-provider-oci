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

// CreateDisApplicationDetails Properties used in application create operations.
type CreateDisApplicationDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// OCID of the compartment that this resource belongs to. Defaults to compartment of the Workspace.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Currently not used on application creation. Reserved for future.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The type of the application.
	ModelType CreateDisApplicationDetailsModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current state of the workspace.
	LifecycleState CreateDisApplicationDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	SourceApplicationInfo *CreateSourceApplicationInfo `mandatory:"false" json:"sourceApplicationInfo"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m CreateDisApplicationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDisApplicationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDisApplicationDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetCreateDisApplicationDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDisApplicationDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCreateDisApplicationDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDisApplicationDetailsModelTypeEnum Enum with underlying type: string
type CreateDisApplicationDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateDisApplicationDetailsModelTypeEnum
const (
	CreateDisApplicationDetailsModelTypeIntegrationApplication CreateDisApplicationDetailsModelTypeEnum = "INTEGRATION_APPLICATION"
)

var mappingCreateDisApplicationDetailsModelTypeEnum = map[string]CreateDisApplicationDetailsModelTypeEnum{
	"INTEGRATION_APPLICATION": CreateDisApplicationDetailsModelTypeIntegrationApplication,
}

var mappingCreateDisApplicationDetailsModelTypeEnumLowerCase = map[string]CreateDisApplicationDetailsModelTypeEnum{
	"integration_application": CreateDisApplicationDetailsModelTypeIntegrationApplication,
}

// GetCreateDisApplicationDetailsModelTypeEnumValues Enumerates the set of values for CreateDisApplicationDetailsModelTypeEnum
func GetCreateDisApplicationDetailsModelTypeEnumValues() []CreateDisApplicationDetailsModelTypeEnum {
	values := make([]CreateDisApplicationDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateDisApplicationDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDisApplicationDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateDisApplicationDetailsModelTypeEnum
func GetCreateDisApplicationDetailsModelTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_APPLICATION",
	}
}

// GetMappingCreateDisApplicationDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDisApplicationDetailsModelTypeEnum(val string) (CreateDisApplicationDetailsModelTypeEnum, bool) {
	enum, ok := mappingCreateDisApplicationDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDisApplicationDetailsLifecycleStateEnum Enum with underlying type: string
type CreateDisApplicationDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for CreateDisApplicationDetailsLifecycleStateEnum
const (
	CreateDisApplicationDetailsLifecycleStateCreating CreateDisApplicationDetailsLifecycleStateEnum = "CREATING"
	CreateDisApplicationDetailsLifecycleStateActive   CreateDisApplicationDetailsLifecycleStateEnum = "ACTIVE"
	CreateDisApplicationDetailsLifecycleStateUpdating CreateDisApplicationDetailsLifecycleStateEnum = "UPDATING"
	CreateDisApplicationDetailsLifecycleStateDeleting CreateDisApplicationDetailsLifecycleStateEnum = "DELETING"
	CreateDisApplicationDetailsLifecycleStateDeleted  CreateDisApplicationDetailsLifecycleStateEnum = "DELETED"
	CreateDisApplicationDetailsLifecycleStateFailed   CreateDisApplicationDetailsLifecycleStateEnum = "FAILED"
)

var mappingCreateDisApplicationDetailsLifecycleStateEnum = map[string]CreateDisApplicationDetailsLifecycleStateEnum{
	"CREATING": CreateDisApplicationDetailsLifecycleStateCreating,
	"ACTIVE":   CreateDisApplicationDetailsLifecycleStateActive,
	"UPDATING": CreateDisApplicationDetailsLifecycleStateUpdating,
	"DELETING": CreateDisApplicationDetailsLifecycleStateDeleting,
	"DELETED":  CreateDisApplicationDetailsLifecycleStateDeleted,
	"FAILED":   CreateDisApplicationDetailsLifecycleStateFailed,
}

var mappingCreateDisApplicationDetailsLifecycleStateEnumLowerCase = map[string]CreateDisApplicationDetailsLifecycleStateEnum{
	"creating": CreateDisApplicationDetailsLifecycleStateCreating,
	"active":   CreateDisApplicationDetailsLifecycleStateActive,
	"updating": CreateDisApplicationDetailsLifecycleStateUpdating,
	"deleting": CreateDisApplicationDetailsLifecycleStateDeleting,
	"deleted":  CreateDisApplicationDetailsLifecycleStateDeleted,
	"failed":   CreateDisApplicationDetailsLifecycleStateFailed,
}

// GetCreateDisApplicationDetailsLifecycleStateEnumValues Enumerates the set of values for CreateDisApplicationDetailsLifecycleStateEnum
func GetCreateDisApplicationDetailsLifecycleStateEnumValues() []CreateDisApplicationDetailsLifecycleStateEnum {
	values := make([]CreateDisApplicationDetailsLifecycleStateEnum, 0)
	for _, v := range mappingCreateDisApplicationDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDisApplicationDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for CreateDisApplicationDetailsLifecycleStateEnum
func GetCreateDisApplicationDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCreateDisApplicationDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDisApplicationDetailsLifecycleStateEnum(val string) (CreateDisApplicationDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingCreateDisApplicationDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
