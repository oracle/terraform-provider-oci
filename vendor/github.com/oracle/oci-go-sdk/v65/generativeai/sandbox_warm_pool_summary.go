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

// SandboxWarmPoolSummary CP-owned configuration and management lifecycle summary. List does not perform per-pool runtime reads and does not return observed ready capacity.
type SandboxWarmPoolSummary struct {

	// Owning Project OCID.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Unique immutable name within the Project.
	Name *string `mandatory:"true" json:"name"`

	// MANAGED uses the service runtime; CUSTOM uses a customer image.
	Type SandboxWarmPoolSummaryTypeEnum `mandatory:"true" json:"type"`

	// Compute shape for each idle Sandbox member.
	Shape SandboxWarmPoolSummaryShapeEnum `mandatory:"true" json:"shape"`

	// Requested idle target, not observed capacity.
	DesiredCapacity *int `mandatory:"true" json:"desiredCapacity"`

	// CP management lifecycle. ACTIVE means the template and requested pool configuration were applied and verified, independently of ready capacity. Scheduling, image-pull and runtime capacity failures do not by themselves change ACTIVE to FAILED.
	LifecycleState SandboxWarmPoolLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Time the desired configuration was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time the configuration or management lifecycle last changed; not runtime readiness.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Digest-pinned customer image for CUSTOM; omitted for MANAGED.
	ContainerImage *string `mandatory:"false" json:"containerImage"`

	// Additional information about the current CP management operation or failure.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m SandboxWarmPoolSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxWarmPoolSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSandboxWarmPoolSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSandboxWarmPoolSummaryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSandboxWarmPoolSummaryShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetSandboxWarmPoolSummaryShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSandboxWarmPoolLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSandboxWarmPoolLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxWarmPoolSummaryTypeEnum Enum with underlying type: string
type SandboxWarmPoolSummaryTypeEnum string

// Set of constants representing the allowable values for SandboxWarmPoolSummaryTypeEnum
const (
	SandboxWarmPoolSummaryTypeManaged SandboxWarmPoolSummaryTypeEnum = "MANAGED"
	SandboxWarmPoolSummaryTypeCustom  SandboxWarmPoolSummaryTypeEnum = "CUSTOM"
)

var mappingSandboxWarmPoolSummaryTypeEnum = map[string]SandboxWarmPoolSummaryTypeEnum{
	"MANAGED": SandboxWarmPoolSummaryTypeManaged,
	"CUSTOM":  SandboxWarmPoolSummaryTypeCustom,
}

var mappingSandboxWarmPoolSummaryTypeEnumLowerCase = map[string]SandboxWarmPoolSummaryTypeEnum{
	"managed": SandboxWarmPoolSummaryTypeManaged,
	"custom":  SandboxWarmPoolSummaryTypeCustom,
}

// GetSandboxWarmPoolSummaryTypeEnumValues Enumerates the set of values for SandboxWarmPoolSummaryTypeEnum
func GetSandboxWarmPoolSummaryTypeEnumValues() []SandboxWarmPoolSummaryTypeEnum {
	values := make([]SandboxWarmPoolSummaryTypeEnum, 0)
	for _, v := range mappingSandboxWarmPoolSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxWarmPoolSummaryTypeEnumStringValues Enumerates the set of values in String for SandboxWarmPoolSummaryTypeEnum
func GetSandboxWarmPoolSummaryTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingSandboxWarmPoolSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxWarmPoolSummaryTypeEnum(val string) (SandboxWarmPoolSummaryTypeEnum, bool) {
	enum, ok := mappingSandboxWarmPoolSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SandboxWarmPoolSummaryShapeEnum Enum with underlying type: string
type SandboxWarmPoolSummaryShapeEnum string

// Set of constants representing the allowable values for SandboxWarmPoolSummaryShapeEnum
const (
	SandboxWarmPoolSummaryShapeSmall  SandboxWarmPoolSummaryShapeEnum = "SMALL"
	SandboxWarmPoolSummaryShapeMedium SandboxWarmPoolSummaryShapeEnum = "MEDIUM"
	SandboxWarmPoolSummaryShapeLarge  SandboxWarmPoolSummaryShapeEnum = "LARGE"
)

var mappingSandboxWarmPoolSummaryShapeEnum = map[string]SandboxWarmPoolSummaryShapeEnum{
	"SMALL":  SandboxWarmPoolSummaryShapeSmall,
	"MEDIUM": SandboxWarmPoolSummaryShapeMedium,
	"LARGE":  SandboxWarmPoolSummaryShapeLarge,
}

var mappingSandboxWarmPoolSummaryShapeEnumLowerCase = map[string]SandboxWarmPoolSummaryShapeEnum{
	"small":  SandboxWarmPoolSummaryShapeSmall,
	"medium": SandboxWarmPoolSummaryShapeMedium,
	"large":  SandboxWarmPoolSummaryShapeLarge,
}

// GetSandboxWarmPoolSummaryShapeEnumValues Enumerates the set of values for SandboxWarmPoolSummaryShapeEnum
func GetSandboxWarmPoolSummaryShapeEnumValues() []SandboxWarmPoolSummaryShapeEnum {
	values := make([]SandboxWarmPoolSummaryShapeEnum, 0)
	for _, v := range mappingSandboxWarmPoolSummaryShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxWarmPoolSummaryShapeEnumStringValues Enumerates the set of values in String for SandboxWarmPoolSummaryShapeEnum
func GetSandboxWarmPoolSummaryShapeEnumStringValues() []string {
	return []string{
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
}

// GetMappingSandboxWarmPoolSummaryShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxWarmPoolSummaryShapeEnum(val string) (SandboxWarmPoolSummaryShapeEnum, bool) {
	enum, ok := mappingSandboxWarmPoolSummaryShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
