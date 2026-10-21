// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSandboxWarmPoolDetails Create Sandbox Warm Pool Details.
type CreateSandboxWarmPoolDetails struct {

	// Unique immutable name within the Project.
	Name *string `mandatory:"true" json:"name"`

	// MANAGED uses the service-provided python-3.11 runtime. CUSTOM uses a customer-provided containerImage. Both types support named pool creation.
	Type CreateSandboxWarmPoolDetailsTypeEnum `mandatory:"true" json:"type"`

	// Compute shape for each idle Sandbox member.
	Shape CreateSandboxWarmPoolDetailsShapeEnum `mandatory:"true" json:"shape"`

	// Desired number of idle members. Zero pauses idle prewarming. Once the pool is ACTIVE, named-pool Sandbox creation can still cold-start from its available applied template when no eligible warm member is available. Existing customer Sandboxes continue. A positive target does not guarantee ready capacity. Existing service capacity validation still applies.
	DesiredCapacity *int `mandatory:"true" json:"desiredCapacity"`

	// Required for CUSTOM and must be omitted for MANAGED. A CUSTOM image must be a digest-pinned same-region OCIR reference, using sha256 with exactly 64 lowercase hexadecimal characters. The CP workflow prepares the template through the data plane, reusing current BYOC image and Project identity validation. Successful configuration does not guarantee image pull or runtime readiness.
	ContainerImage *string `mandatory:"false" json:"containerImage"`
}

func (m CreateSandboxWarmPoolDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSandboxWarmPoolDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateSandboxWarmPoolDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreateSandboxWarmPoolDetailsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateSandboxWarmPoolDetailsShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetCreateSandboxWarmPoolDetailsShapeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSandboxWarmPoolDetailsTypeEnum Enum with underlying type: string
type CreateSandboxWarmPoolDetailsTypeEnum string

// Set of constants representing the allowable values for CreateSandboxWarmPoolDetailsTypeEnum
const (
	CreateSandboxWarmPoolDetailsTypeManaged CreateSandboxWarmPoolDetailsTypeEnum = "MANAGED"
	CreateSandboxWarmPoolDetailsTypeCustom  CreateSandboxWarmPoolDetailsTypeEnum = "CUSTOM"
)

var mappingCreateSandboxWarmPoolDetailsTypeEnum = map[string]CreateSandboxWarmPoolDetailsTypeEnum{
	"MANAGED": CreateSandboxWarmPoolDetailsTypeManaged,
	"CUSTOM":  CreateSandboxWarmPoolDetailsTypeCustom,
}

var mappingCreateSandboxWarmPoolDetailsTypeEnumLowerCase = map[string]CreateSandboxWarmPoolDetailsTypeEnum{
	"managed": CreateSandboxWarmPoolDetailsTypeManaged,
	"custom":  CreateSandboxWarmPoolDetailsTypeCustom,
}

// GetCreateSandboxWarmPoolDetailsTypeEnumValues Enumerates the set of values for CreateSandboxWarmPoolDetailsTypeEnum
func GetCreateSandboxWarmPoolDetailsTypeEnumValues() []CreateSandboxWarmPoolDetailsTypeEnum {
	values := make([]CreateSandboxWarmPoolDetailsTypeEnum, 0)
	for _, v := range mappingCreateSandboxWarmPoolDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSandboxWarmPoolDetailsTypeEnumStringValues Enumerates the set of values in String for CreateSandboxWarmPoolDetailsTypeEnum
func GetCreateSandboxWarmPoolDetailsTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingCreateSandboxWarmPoolDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSandboxWarmPoolDetailsTypeEnum(val string) (CreateSandboxWarmPoolDetailsTypeEnum, bool) {
	enum, ok := mappingCreateSandboxWarmPoolDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateSandboxWarmPoolDetailsShapeEnum Enum with underlying type: string
type CreateSandboxWarmPoolDetailsShapeEnum string

// Set of constants representing the allowable values for CreateSandboxWarmPoolDetailsShapeEnum
const (
	CreateSandboxWarmPoolDetailsShapeSmall  CreateSandboxWarmPoolDetailsShapeEnum = "SMALL"
	CreateSandboxWarmPoolDetailsShapeMedium CreateSandboxWarmPoolDetailsShapeEnum = "MEDIUM"
	CreateSandboxWarmPoolDetailsShapeLarge  CreateSandboxWarmPoolDetailsShapeEnum = "LARGE"
)

var mappingCreateSandboxWarmPoolDetailsShapeEnum = map[string]CreateSandboxWarmPoolDetailsShapeEnum{
	"SMALL":  CreateSandboxWarmPoolDetailsShapeSmall,
	"MEDIUM": CreateSandboxWarmPoolDetailsShapeMedium,
	"LARGE":  CreateSandboxWarmPoolDetailsShapeLarge,
}

var mappingCreateSandboxWarmPoolDetailsShapeEnumLowerCase = map[string]CreateSandboxWarmPoolDetailsShapeEnum{
	"small":  CreateSandboxWarmPoolDetailsShapeSmall,
	"medium": CreateSandboxWarmPoolDetailsShapeMedium,
	"large":  CreateSandboxWarmPoolDetailsShapeLarge,
}

// GetCreateSandboxWarmPoolDetailsShapeEnumValues Enumerates the set of values for CreateSandboxWarmPoolDetailsShapeEnum
func GetCreateSandboxWarmPoolDetailsShapeEnumValues() []CreateSandboxWarmPoolDetailsShapeEnum {
	values := make([]CreateSandboxWarmPoolDetailsShapeEnum, 0)
	for _, v := range mappingCreateSandboxWarmPoolDetailsShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSandboxWarmPoolDetailsShapeEnumStringValues Enumerates the set of values in String for CreateSandboxWarmPoolDetailsShapeEnum
func GetCreateSandboxWarmPoolDetailsShapeEnumStringValues() []string {
	return []string{
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
}

// GetMappingCreateSandboxWarmPoolDetailsShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSandboxWarmPoolDetailsShapeEnum(val string) (CreateSandboxWarmPoolDetailsShapeEnum, bool) {
	enum, ok := mappingCreateSandboxWarmPoolDetailsShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
