// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateEndpointDetails The information needed to create a new endpoint and expose to end users.
type CreateEndpointDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) compartment identifier for the endpoint
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model to associate with the endpoint.
	ModelId *string `mandatory:"true" json:"modelId"`

	// A user-friendly display name for the resource. It should be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Compute infra type for endpoint.
	ComputeType CreateEndpointDetailsComputeTypeEnum `mandatory:"false" json:"computeType,omitempty"`

	// Unique name across user tenancy in a region to identify an endpoint to be used for inferencing.
	Alias *string `mandatory:"false" json:"alias"`

	// A short description of the an endpoint.
	Description *string `mandatory:"false" json:"description"`

	// Number of replicas required for this endpoint. This will be optional parameter. Default will be 1.
	InferenceUnits *int `mandatory:"false" json:"inferenceUnits"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateEndpointDetailsComputeTypeEnum(string(m.ComputeType)); !ok && m.ComputeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeType: %s. Supported values are: %s.", m.ComputeType, strings.Join(GetCreateEndpointDetailsComputeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateEndpointDetailsComputeTypeEnum Enum with underlying type: string
type CreateEndpointDetailsComputeTypeEnum string

// Set of constants representing the allowable values for CreateEndpointDetailsComputeTypeEnum
const (
	CreateEndpointDetailsComputeTypeCpu CreateEndpointDetailsComputeTypeEnum = "CPU"
	CreateEndpointDetailsComputeTypeGpu CreateEndpointDetailsComputeTypeEnum = "GPU"
)

var mappingCreateEndpointDetailsComputeTypeEnum = map[string]CreateEndpointDetailsComputeTypeEnum{
	"CPU": CreateEndpointDetailsComputeTypeCpu,
	"GPU": CreateEndpointDetailsComputeTypeGpu,
}

var mappingCreateEndpointDetailsComputeTypeEnumLowerCase = map[string]CreateEndpointDetailsComputeTypeEnum{
	"cpu": CreateEndpointDetailsComputeTypeCpu,
	"gpu": CreateEndpointDetailsComputeTypeGpu,
}

// GetCreateEndpointDetailsComputeTypeEnumValues Enumerates the set of values for CreateEndpointDetailsComputeTypeEnum
func GetCreateEndpointDetailsComputeTypeEnumValues() []CreateEndpointDetailsComputeTypeEnum {
	values := make([]CreateEndpointDetailsComputeTypeEnum, 0)
	for _, v := range mappingCreateEndpointDetailsComputeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateEndpointDetailsComputeTypeEnumStringValues Enumerates the set of values in String for CreateEndpointDetailsComputeTypeEnum
func GetCreateEndpointDetailsComputeTypeEnumStringValues() []string {
	return []string{
		"CPU",
		"GPU",
	}
}

// GetMappingCreateEndpointDetailsComputeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateEndpointDetailsComputeTypeEnum(val string) (CreateEndpointDetailsComputeTypeEnum, bool) {
	enum, ok := mappingCreateEndpointDetailsComputeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
