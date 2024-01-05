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

// FunctionSummary Summary of a function.
type FunctionSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the function. The display name is unique within the application containing the function.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the application the function belongs to.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The OCID of the compartment that contains the function.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The current state of the function.
	LifecycleState FunctionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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
	Shape FunctionSummaryShapeEnum `mandatory:"false" json:"shape,omitempty"`

	// Maximum usable memory for the function (MiB).
	MemoryInMBs *int64 `mandatory:"false" json:"memoryInMBs"`

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

func (m FunctionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFunctionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFunctionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFunctionSummaryShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetFunctionSummaryShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FunctionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                  *string                              `json:"displayName"`
		ApplicationId                *string                              `json:"applicationId"`
		CompartmentId                *string                              `json:"compartmentId"`
		LifecycleState               FunctionLifecycleStateEnum           `json:"lifecycleState"`
		Image                        *string                              `json:"image"`
		ImageDigest                  *string                              `json:"imageDigest"`
		SourceDetails                functionsourcedetails                `json:"sourceDetails"`
		Shape                        FunctionSummaryShapeEnum             `json:"shape"`
		MemoryInMBs                  *int64                               `json:"memoryInMBs"`
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

	m.ApplicationId = model.ApplicationId

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

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

// FunctionSummaryShapeEnum Enum with underlying type: string
type FunctionSummaryShapeEnum string

// Set of constants representing the allowable values for FunctionSummaryShapeEnum
const (
	FunctionSummaryShapeX86    FunctionSummaryShapeEnum = "GENERIC_X86"
	FunctionSummaryShapeArm    FunctionSummaryShapeEnum = "GENERIC_ARM"
	FunctionSummaryShapeX86Arm FunctionSummaryShapeEnum = "GENERIC_X86_ARM"
)

var mappingFunctionSummaryShapeEnum = map[string]FunctionSummaryShapeEnum{
	"GENERIC_X86":     FunctionSummaryShapeX86,
	"GENERIC_ARM":     FunctionSummaryShapeArm,
	"GENERIC_X86_ARM": FunctionSummaryShapeX86Arm,
}

var mappingFunctionSummaryShapeEnumLowerCase = map[string]FunctionSummaryShapeEnum{
	"generic_x86":     FunctionSummaryShapeX86,
	"generic_arm":     FunctionSummaryShapeArm,
	"generic_x86_arm": FunctionSummaryShapeX86Arm,
}

// GetFunctionSummaryShapeEnumValues Enumerates the set of values for FunctionSummaryShapeEnum
func GetFunctionSummaryShapeEnumValues() []FunctionSummaryShapeEnum {
	values := make([]FunctionSummaryShapeEnum, 0)
	for _, v := range mappingFunctionSummaryShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionSummaryShapeEnumStringValues Enumerates the set of values in String for FunctionSummaryShapeEnum
func GetFunctionSummaryShapeEnumStringValues() []string {
	return []string{
		"GENERIC_X86",
		"GENERIC_ARM",
		"GENERIC_X86_ARM",
	}
}

// GetMappingFunctionSummaryShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionSummaryShapeEnum(val string) (FunctionSummaryShapeEnum, bool) {
	enum, ok := mappingFunctionSummaryShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
