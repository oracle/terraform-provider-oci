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

// CreateApplicationDetails Properties used in application create operations.
type CreateApplicationDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Currently not used on application creation. Reserved for future.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The type of the application.
	ModelType CreateApplicationDetailsModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

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
	LifecycleState CreateApplicationDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	SourceApplicationInfo *CreateSourceApplicationInfo `mandatory:"false" json:"sourceApplicationInfo"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m CreateApplicationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateApplicationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateApplicationDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetCreateApplicationDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateApplicationDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCreateApplicationDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateApplicationDetailsModelTypeEnum Enum with underlying type: string
type CreateApplicationDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateApplicationDetailsModelTypeEnum
const (
	CreateApplicationDetailsModelTypeIntegrationApplication CreateApplicationDetailsModelTypeEnum = "INTEGRATION_APPLICATION"
)

var mappingCreateApplicationDetailsModelTypeEnum = map[string]CreateApplicationDetailsModelTypeEnum{
	"INTEGRATION_APPLICATION": CreateApplicationDetailsModelTypeIntegrationApplication,
}

var mappingCreateApplicationDetailsModelTypeEnumLowerCase = map[string]CreateApplicationDetailsModelTypeEnum{
	"integration_application": CreateApplicationDetailsModelTypeIntegrationApplication,
}

// GetCreateApplicationDetailsModelTypeEnumValues Enumerates the set of values for CreateApplicationDetailsModelTypeEnum
func GetCreateApplicationDetailsModelTypeEnumValues() []CreateApplicationDetailsModelTypeEnum {
	values := make([]CreateApplicationDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateApplicationDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateApplicationDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateApplicationDetailsModelTypeEnum
func GetCreateApplicationDetailsModelTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_APPLICATION",
	}
}

// GetMappingCreateApplicationDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateApplicationDetailsModelTypeEnum(val string) (CreateApplicationDetailsModelTypeEnum, bool) {
	enum, ok := mappingCreateApplicationDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateApplicationDetailsLifecycleStateEnum Enum with underlying type: string
type CreateApplicationDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for CreateApplicationDetailsLifecycleStateEnum
const (
	CreateApplicationDetailsLifecycleStateCreating CreateApplicationDetailsLifecycleStateEnum = "CREATING"
	CreateApplicationDetailsLifecycleStateActive   CreateApplicationDetailsLifecycleStateEnum = "ACTIVE"
	CreateApplicationDetailsLifecycleStateUpdating CreateApplicationDetailsLifecycleStateEnum = "UPDATING"
	CreateApplicationDetailsLifecycleStateDeleting CreateApplicationDetailsLifecycleStateEnum = "DELETING"
	CreateApplicationDetailsLifecycleStateDeleted  CreateApplicationDetailsLifecycleStateEnum = "DELETED"
	CreateApplicationDetailsLifecycleStateFailed   CreateApplicationDetailsLifecycleStateEnum = "FAILED"
)

var mappingCreateApplicationDetailsLifecycleStateEnum = map[string]CreateApplicationDetailsLifecycleStateEnum{
	"CREATING": CreateApplicationDetailsLifecycleStateCreating,
	"ACTIVE":   CreateApplicationDetailsLifecycleStateActive,
	"UPDATING": CreateApplicationDetailsLifecycleStateUpdating,
	"DELETING": CreateApplicationDetailsLifecycleStateDeleting,
	"DELETED":  CreateApplicationDetailsLifecycleStateDeleted,
	"FAILED":   CreateApplicationDetailsLifecycleStateFailed,
}

var mappingCreateApplicationDetailsLifecycleStateEnumLowerCase = map[string]CreateApplicationDetailsLifecycleStateEnum{
	"creating": CreateApplicationDetailsLifecycleStateCreating,
	"active":   CreateApplicationDetailsLifecycleStateActive,
	"updating": CreateApplicationDetailsLifecycleStateUpdating,
	"deleting": CreateApplicationDetailsLifecycleStateDeleting,
	"deleted":  CreateApplicationDetailsLifecycleStateDeleted,
	"failed":   CreateApplicationDetailsLifecycleStateFailed,
}

// GetCreateApplicationDetailsLifecycleStateEnumValues Enumerates the set of values for CreateApplicationDetailsLifecycleStateEnum
func GetCreateApplicationDetailsLifecycleStateEnumValues() []CreateApplicationDetailsLifecycleStateEnum {
	values := make([]CreateApplicationDetailsLifecycleStateEnum, 0)
	for _, v := range mappingCreateApplicationDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateApplicationDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for CreateApplicationDetailsLifecycleStateEnum
func GetCreateApplicationDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCreateApplicationDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateApplicationDetailsLifecycleStateEnum(val string) (CreateApplicationDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingCreateApplicationDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
