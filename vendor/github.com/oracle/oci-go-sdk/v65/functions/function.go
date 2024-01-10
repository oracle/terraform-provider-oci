// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Function A function resource defines the code (Docker image) and configuration for a specific function. Functions are defined in applications. Avoid entering confidential information.
type Function struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the function. The display name is unique within the application containing the function.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the function.
	LifecycleState FunctionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The OCID of the application the function belongs to.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The OCID of the compartment that contains the function.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The qualified name of the Docker image to use in the function, including the image tag.
	// The image should be in the OCI Registry that is in the same region as the function itself.
	// Example: `phx.ocir.io/ten/functions/function:0.0.1`
	Image *string `mandatory:"false" json:"image"`

	// The image digest for the version of the image that will be pulled when invoking this function.
	// If no value is specified, the digest currently associated with the image in the OCI Registry will be used.
	// Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7`
	ImageDigest *string `mandatory:"false" json:"imageDigest"`

	SourceDetails FunctionSourceDetails `mandatory:"false" json:"sourceDetails"`

	// The processor shape (`GENERIC_X86`/`GENERIC_ARM`) on which to run functions in the application, extracted from the image manifest.
	Shape FunctionShapeEnum `mandatory:"false" json:"shape,omitempty"`

	// Maximum usable memory for the function (MiB).
	MemoryInMBs *int64 `mandatory:"false" json:"memoryInMBs"`

	// Function configuration. Overrides application configuration.
	// Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.
	// Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`
	// The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8.
	Config map[string]string `mandatory:"false" json:"config"`

	// Timeout for executions of the function. Value in seconds.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	ProvisionedConcurrencyConfig FunctionProvisionedConcurrencyConfig `mandatory:"false" json:"provisionedConcurrencyConfig"`

	TraceConfig *FunctionTraceConfig `mandatory:"false" json:"traceConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The base https invoke URL to set on a client in order to invoke a function. This URL will never change over the lifetime of the function and can be cached.
	InvokeEndpoint *string `mandatory:"false" json:"invokeEndpoint"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The time the function was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the function was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m Function) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Function) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFunctionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFunctionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFunctionShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetFunctionShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Function) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                  *string                              `json:"displayName"`
		LifecycleState               FunctionLifecycleStateEnum           `json:"lifecycleState"`
		ApplicationId                *string                              `json:"applicationId"`
		CompartmentId                *string                              `json:"compartmentId"`
		Image                        *string                              `json:"image"`
		ImageDigest                  *string                              `json:"imageDigest"`
		SourceDetails                functionsourcedetails                `json:"sourceDetails"`
		Shape                        FunctionShapeEnum                    `json:"shape"`
		MemoryInMBs                  *int64                               `json:"memoryInMBs"`
		Config                       map[string]string                    `json:"config"`
		TimeoutInSeconds             *int                                 `json:"timeoutInSeconds"`
		ProvisionedConcurrencyConfig functionprovisionedconcurrencyconfig `json:"provisionedConcurrencyConfig"`
		TraceConfig                  *FunctionTraceConfig                 `json:"traceConfig"`
		FreeformTags                 map[string]string                    `json:"freeformTags"`
		InvokeEndpoint               *string                              `json:"invokeEndpoint"`
		DefinedTags                  map[string]map[string]interface{}    `json:"definedTags"`
		TimeCreated                  *common.SDKTime                      `json:"timeCreated"`
		TimeUpdated                  *common.SDKTime                      `json:"timeUpdated"`
		Id                           *string                              `json:"id"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	m.ApplicationId = model.ApplicationId

	m.CompartmentId = model.CompartmentId

	m.Image = model.Image

	m.ImageDigest = model.ImageDigest

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(FunctionSourceDetails)
	} else {
		m.SourceDetails = nil
	}

	m.Shape = model.Shape

	m.MemoryInMBs = model.MemoryInMBs

	m.Config = model.Config

	m.TimeoutInSeconds = model.TimeoutInSeconds

	nn, e = model.ProvisionedConcurrencyConfig.UnmarshalPolymorphicJSON(model.ProvisionedConcurrencyConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ProvisionedConcurrencyConfig = nn.(FunctionProvisionedConcurrencyConfig)
	} else {
		m.ProvisionedConcurrencyConfig = nil
	}

	m.TraceConfig = model.TraceConfig

	m.FreeformTags = model.FreeformTags

	m.InvokeEndpoint = model.InvokeEndpoint

	m.DefinedTags = model.DefinedTags

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.Id = model.Id

	return
}

// FunctionLifecycleStateEnum Enum with underlying type: string
type FunctionLifecycleStateEnum string

// Set of constants representing the allowable values for FunctionLifecycleStateEnum
const (
	FunctionLifecycleStateCreating FunctionLifecycleStateEnum = "CREATING"
	FunctionLifecycleStateActive   FunctionLifecycleStateEnum = "ACTIVE"
	FunctionLifecycleStateInactive FunctionLifecycleStateEnum = "INACTIVE"
	FunctionLifecycleStateUpdating FunctionLifecycleStateEnum = "UPDATING"
	FunctionLifecycleStateDeleting FunctionLifecycleStateEnum = "DELETING"
	FunctionLifecycleStateDeleted  FunctionLifecycleStateEnum = "DELETED"
	FunctionLifecycleStateFailed   FunctionLifecycleStateEnum = "FAILED"
)

var mappingFunctionLifecycleStateEnum = map[string]FunctionLifecycleStateEnum{
	"CREATING": FunctionLifecycleStateCreating,
	"ACTIVE":   FunctionLifecycleStateActive,
	"INACTIVE": FunctionLifecycleStateInactive,
	"UPDATING": FunctionLifecycleStateUpdating,
	"DELETING": FunctionLifecycleStateDeleting,
	"DELETED":  FunctionLifecycleStateDeleted,
	"FAILED":   FunctionLifecycleStateFailed,
}

var mappingFunctionLifecycleStateEnumLowerCase = map[string]FunctionLifecycleStateEnum{
	"creating": FunctionLifecycleStateCreating,
	"active":   FunctionLifecycleStateActive,
	"inactive": FunctionLifecycleStateInactive,
	"updating": FunctionLifecycleStateUpdating,
	"deleting": FunctionLifecycleStateDeleting,
	"deleted":  FunctionLifecycleStateDeleted,
	"failed":   FunctionLifecycleStateFailed,
}

// GetFunctionLifecycleStateEnumValues Enumerates the set of values for FunctionLifecycleStateEnum
func GetFunctionLifecycleStateEnumValues() []FunctionLifecycleStateEnum {
	values := make([]FunctionLifecycleStateEnum, 0)
	for _, v := range mappingFunctionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionLifecycleStateEnumStringValues Enumerates the set of values in String for FunctionLifecycleStateEnum
func GetFunctionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFunctionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionLifecycleStateEnum(val string) (FunctionLifecycleStateEnum, bool) {
	enum, ok := mappingFunctionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FunctionShapeEnum Enum with underlying type: string
type FunctionShapeEnum string

// Set of constants representing the allowable values for FunctionShapeEnum
const (
	FunctionShapeX86    FunctionShapeEnum = "GENERIC_X86"
	FunctionShapeArm    FunctionShapeEnum = "GENERIC_ARM"
	FunctionShapeX86Arm FunctionShapeEnum = "GENERIC_X86_ARM"
)

var mappingFunctionShapeEnum = map[string]FunctionShapeEnum{
	"GENERIC_X86":     FunctionShapeX86,
	"GENERIC_ARM":     FunctionShapeArm,
	"GENERIC_X86_ARM": FunctionShapeX86Arm,
}

var mappingFunctionShapeEnumLowerCase = map[string]FunctionShapeEnum{
	"generic_x86":     FunctionShapeX86,
	"generic_arm":     FunctionShapeArm,
	"generic_x86_arm": FunctionShapeX86Arm,
}

// GetFunctionShapeEnumValues Enumerates the set of values for FunctionShapeEnum
func GetFunctionShapeEnumValues() []FunctionShapeEnum {
	values := make([]FunctionShapeEnum, 0)
	for _, v := range mappingFunctionShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionShapeEnumStringValues Enumerates the set of values in String for FunctionShapeEnum
func GetFunctionShapeEnumStringValues() []string {
	return []string{
		"GENERIC_X86",
		"GENERIC_ARM",
		"GENERIC_X86_ARM",
	}
}

// GetMappingFunctionShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionShapeEnum(val string) (FunctionShapeEnum, bool) {
	enum, ok := mappingFunctionShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
