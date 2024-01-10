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

// CreateFunctionDetails Properties to create a new function.
type CreateFunctionDetails struct {

	// The display name of the function. The display name must be unique within the application containing the function. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the application this function belongs to.
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Maximum usable memory for the function (MiB).
	MemoryInMBs *int64 `mandatory:"true" json:"memoryInMBs"`

	// The qualified name of the Docker image to use in the function, including the image tag.
	// The image should be in the OCI Registry that is in the same region as the function itself.
	// Example: `phx.ocir.io/ten/functions/function:0.0.1`
	Image *string `mandatory:"false" json:"image"`

	// The image digest for the version of the image that will be pulled when invoking this function.
	// If no value is specified, the digest currently associated with the image in the OCI Registry will be used.
	// Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7`
	ImageDigest *string `mandatory:"false" json:"imageDigest"`

	SourceDetails FunctionSourceDetails `mandatory:"false" json:"sourceDetails"`

	// Function configuration. These values are passed on to the function as environment variables, this overrides application configuration values.
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

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateFunctionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFunctionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateFunctionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Image                        *string                              `json:"image"`
		ImageDigest                  *string                              `json:"imageDigest"`
		SourceDetails                functionsourcedetails                `json:"sourceDetails"`
		Config                       map[string]string                    `json:"config"`
		TimeoutInSeconds             *int                                 `json:"timeoutInSeconds"`
		ProvisionedConcurrencyConfig functionprovisionedconcurrencyconfig `json:"provisionedConcurrencyConfig"`
		TraceConfig                  *FunctionTraceConfig                 `json:"traceConfig"`
		FreeformTags                 map[string]string                    `json:"freeformTags"`
		DefinedTags                  map[string]map[string]interface{}    `json:"definedTags"`
		DisplayName                  *string                              `json:"displayName"`
		ApplicationId                *string                              `json:"applicationId"`
		MemoryInMBs                  *int64                               `json:"memoryInMBs"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
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

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.ApplicationId = model.ApplicationId

	m.MemoryInMBs = model.MemoryInMBs

	return
}
