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

// EndpointSummary Summary of the language endpoint.
type EndpointSummary struct {

	// Unique identifier endpoint OCID of an endpoint that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the resource. It should be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  for the Endpoint compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the endpoint.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The time the the endpoint was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the endpoint.
	LifecycleState EndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model to associate with the endpoint.
	ModelId *string `mandatory:"true" json:"modelId"`

	// Unique name across user tenancy in a region to identify an endpoint to be used for inferencing.
	Alias *string `mandatory:"false" json:"alias"`

	// Compute infra type for endpoint.
	ComputeType EndpointSummaryComputeTypeEnum `mandatory:"false" json:"computeType,omitempty"`

	// A short description of the endpoint.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Number of replicas required for this endpoint. This will be optional parameter. Default will be 1.
	InferenceUnits *int `mandatory:"false" json:"inferenceUnits"`

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

func (m EndpointSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EndpointSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEndpointSummaryComputeTypeEnum(string(m.ComputeType)); !ok && m.ComputeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeType: %s. Supported values are: %s.", m.ComputeType, strings.Join(GetEndpointSummaryComputeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EndpointSummaryComputeTypeEnum Enum with underlying type: string
type EndpointSummaryComputeTypeEnum string

// Set of constants representing the allowable values for EndpointSummaryComputeTypeEnum
const (
	EndpointSummaryComputeTypeCpu EndpointSummaryComputeTypeEnum = "CPU"
	EndpointSummaryComputeTypeGpu EndpointSummaryComputeTypeEnum = "GPU"
)

var mappingEndpointSummaryComputeTypeEnum = map[string]EndpointSummaryComputeTypeEnum{
	"CPU": EndpointSummaryComputeTypeCpu,
	"GPU": EndpointSummaryComputeTypeGpu,
}

var mappingEndpointSummaryComputeTypeEnumLowerCase = map[string]EndpointSummaryComputeTypeEnum{
	"cpu": EndpointSummaryComputeTypeCpu,
	"gpu": EndpointSummaryComputeTypeGpu,
}

// GetEndpointSummaryComputeTypeEnumValues Enumerates the set of values for EndpointSummaryComputeTypeEnum
func GetEndpointSummaryComputeTypeEnumValues() []EndpointSummaryComputeTypeEnum {
	values := make([]EndpointSummaryComputeTypeEnum, 0)
	for _, v := range mappingEndpointSummaryComputeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointSummaryComputeTypeEnumStringValues Enumerates the set of values in String for EndpointSummaryComputeTypeEnum
func GetEndpointSummaryComputeTypeEnumStringValues() []string {
	return []string{
		"CPU",
		"GPU",
	}
}

// GetMappingEndpointSummaryComputeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointSummaryComputeTypeEnum(val string) (EndpointSummaryComputeTypeEnum, bool) {
	enum, ok := mappingEndpointSummaryComputeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
