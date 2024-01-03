// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Model Description of Model.
type Model struct {

	// The OCID of the model that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID for the model's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ModelTrainingDetails *ModelTrainingDetails `mandatory:"true" json:"modelTrainingDetails"`

	ModelTrainingResults *ModelTrainingResults `mandatory:"true" json:"modelTrainingResults"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The time the the Model was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A short description of the Model.
	Description *string `mandatory:"false" json:"description"`

	// The time the Model was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Model) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Model) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetModelLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelLifecycleStateEnum Enum with underlying type: string
type ModelLifecycleStateEnum string

// Set of constants representing the allowable values for ModelLifecycleStateEnum
const (
	ModelLifecycleStateDeleting ModelLifecycleStateEnum = "DELETING"
	ModelLifecycleStateDeleted  ModelLifecycleStateEnum = "DELETED"
	ModelLifecycleStateFailed   ModelLifecycleStateEnum = "FAILED"
	ModelLifecycleStateCreating ModelLifecycleStateEnum = "CREATING"
	ModelLifecycleStateActive   ModelLifecycleStateEnum = "ACTIVE"
	ModelLifecycleStateUpdating ModelLifecycleStateEnum = "UPDATING"
)

var mappingModelLifecycleStateEnum = map[string]ModelLifecycleStateEnum{
	"DELETING": ModelLifecycleStateDeleting,
	"DELETED":  ModelLifecycleStateDeleted,
	"FAILED":   ModelLifecycleStateFailed,
	"CREATING": ModelLifecycleStateCreating,
	"ACTIVE":   ModelLifecycleStateActive,
	"UPDATING": ModelLifecycleStateUpdating,
}

var mappingModelLifecycleStateEnumLowerCase = map[string]ModelLifecycleStateEnum{
	"deleting": ModelLifecycleStateDeleting,
	"deleted":  ModelLifecycleStateDeleted,
	"failed":   ModelLifecycleStateFailed,
	"creating": ModelLifecycleStateCreating,
	"active":   ModelLifecycleStateActive,
	"updating": ModelLifecycleStateUpdating,
}

// GetModelLifecycleStateEnumValues Enumerates the set of values for ModelLifecycleStateEnum
func GetModelLifecycleStateEnumValues() []ModelLifecycleStateEnum {
	values := make([]ModelLifecycleStateEnum, 0)
	for _, v := range mappingModelLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetModelLifecycleStateEnumStringValues Enumerates the set of values in String for ModelLifecycleStateEnum
func GetModelLifecycleStateEnumStringValues() []string {
	return []string{
		"DELETING",
		"DELETED",
		"FAILED",
		"CREATING",
		"ACTIVE",
		"UPDATING",
	}
}

// GetMappingModelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelLifecycleStateEnum(val string) (ModelLifecycleStateEnum, bool) {
	enum, ok := mappingModelLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
