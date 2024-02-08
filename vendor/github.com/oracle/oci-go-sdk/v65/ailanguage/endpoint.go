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

// Endpoint Description of the endpoint.
type Endpoint struct {

	// Unique identifier endpoint OCID of an endpoint that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the resource. It should be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the endpoint compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the Endpoint.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The time the the endpoint was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the endpoint.
	LifecycleState EndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model to associate with the endpoint.
	ModelId *string `mandatory:"true" json:"modelId"`

	// A short description of the endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The time the endpoint was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Number of replicas required for this endpoint.
	InferenceUnits *int `mandatory:"false" json:"inferenceUnits"`

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

func (m Endpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Endpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EndpointLifecycleStateEnum Enum with underlying type: string
type EndpointLifecycleStateEnum string

// Set of constants representing the allowable values for EndpointLifecycleStateEnum
const (
	EndpointLifecycleStateDeleting EndpointLifecycleStateEnum = "DELETING"
	EndpointLifecycleStateDeleted  EndpointLifecycleStateEnum = "DELETED"
	EndpointLifecycleStateFailed   EndpointLifecycleStateEnum = "FAILED"
	EndpointLifecycleStateCreating EndpointLifecycleStateEnum = "CREATING"
	EndpointLifecycleStateActive   EndpointLifecycleStateEnum = "ACTIVE"
	EndpointLifecycleStateUpdating EndpointLifecycleStateEnum = "UPDATING"
)

var mappingEndpointLifecycleStateEnum = map[string]EndpointLifecycleStateEnum{
	"DELETING": EndpointLifecycleStateDeleting,
	"DELETED":  EndpointLifecycleStateDeleted,
	"FAILED":   EndpointLifecycleStateFailed,
	"CREATING": EndpointLifecycleStateCreating,
	"ACTIVE":   EndpointLifecycleStateActive,
	"UPDATING": EndpointLifecycleStateUpdating,
}

var mappingEndpointLifecycleStateEnumLowerCase = map[string]EndpointLifecycleStateEnum{
	"deleting": EndpointLifecycleStateDeleting,
	"deleted":  EndpointLifecycleStateDeleted,
	"failed":   EndpointLifecycleStateFailed,
	"creating": EndpointLifecycleStateCreating,
	"active":   EndpointLifecycleStateActive,
	"updating": EndpointLifecycleStateUpdating,
}

// GetEndpointLifecycleStateEnumValues Enumerates the set of values for EndpointLifecycleStateEnum
func GetEndpointLifecycleStateEnumValues() []EndpointLifecycleStateEnum {
	values := make([]EndpointLifecycleStateEnum, 0)
	for _, v := range mappingEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for EndpointLifecycleStateEnum
func GetEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"DELETING",
		"DELETED",
		"FAILED",
		"CREATING",
		"ACTIVE",
		"UPDATING",
	}
}

// GetMappingEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointLifecycleStateEnum(val string) (EndpointLifecycleStateEnum, bool) {
	enum, ok := mappingEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
