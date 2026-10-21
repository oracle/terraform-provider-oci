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

// SandboxWarmPool CP-owned configuration and management lifecycle of a Project-owned Sandbox warm pool. Get can include best-effort runtime observations independently of lifecycle.
type SandboxWarmPool struct {

	// Owning Project OCID.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Unique immutable name within the Project.
	Name *string `mandatory:"true" json:"name"`

	// MANAGED uses the service runtime; CUSTOM uses a customer image.
	Type SandboxWarmPoolTypeEnum `mandatory:"true" json:"type"`

	// Compute shape for each idle Sandbox member.
	Shape SandboxWarmPoolShapeEnum `mandatory:"true" json:"shape"`

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

	RuntimeStatus *SandboxWarmPoolRuntimeStatus `mandatory:"false" json:"runtimeStatus"`
}

func (m SandboxWarmPool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SandboxWarmPool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSandboxWarmPoolTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSandboxWarmPoolTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSandboxWarmPoolShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetSandboxWarmPoolShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSandboxWarmPoolLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSandboxWarmPoolLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SandboxWarmPoolTypeEnum Enum with underlying type: string
type SandboxWarmPoolTypeEnum string

// Set of constants representing the allowable values for SandboxWarmPoolTypeEnum
const (
	SandboxWarmPoolTypeManaged SandboxWarmPoolTypeEnum = "MANAGED"
	SandboxWarmPoolTypeCustom  SandboxWarmPoolTypeEnum = "CUSTOM"
)

var mappingSandboxWarmPoolTypeEnum = map[string]SandboxWarmPoolTypeEnum{
	"MANAGED": SandboxWarmPoolTypeManaged,
	"CUSTOM":  SandboxWarmPoolTypeCustom,
}

var mappingSandboxWarmPoolTypeEnumLowerCase = map[string]SandboxWarmPoolTypeEnum{
	"managed": SandboxWarmPoolTypeManaged,
	"custom":  SandboxWarmPoolTypeCustom,
}

// GetSandboxWarmPoolTypeEnumValues Enumerates the set of values for SandboxWarmPoolTypeEnum
func GetSandboxWarmPoolTypeEnumValues() []SandboxWarmPoolTypeEnum {
	values := make([]SandboxWarmPoolTypeEnum, 0)
	for _, v := range mappingSandboxWarmPoolTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxWarmPoolTypeEnumStringValues Enumerates the set of values in String for SandboxWarmPoolTypeEnum
func GetSandboxWarmPoolTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingSandboxWarmPoolTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxWarmPoolTypeEnum(val string) (SandboxWarmPoolTypeEnum, bool) {
	enum, ok := mappingSandboxWarmPoolTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SandboxWarmPoolShapeEnum Enum with underlying type: string
type SandboxWarmPoolShapeEnum string

// Set of constants representing the allowable values for SandboxWarmPoolShapeEnum
const (
	SandboxWarmPoolShapeSmall  SandboxWarmPoolShapeEnum = "SMALL"
	SandboxWarmPoolShapeMedium SandboxWarmPoolShapeEnum = "MEDIUM"
	SandboxWarmPoolShapeLarge  SandboxWarmPoolShapeEnum = "LARGE"
)

var mappingSandboxWarmPoolShapeEnum = map[string]SandboxWarmPoolShapeEnum{
	"SMALL":  SandboxWarmPoolShapeSmall,
	"MEDIUM": SandboxWarmPoolShapeMedium,
	"LARGE":  SandboxWarmPoolShapeLarge,
}

var mappingSandboxWarmPoolShapeEnumLowerCase = map[string]SandboxWarmPoolShapeEnum{
	"small":  SandboxWarmPoolShapeSmall,
	"medium": SandboxWarmPoolShapeMedium,
	"large":  SandboxWarmPoolShapeLarge,
}

// GetSandboxWarmPoolShapeEnumValues Enumerates the set of values for SandboxWarmPoolShapeEnum
func GetSandboxWarmPoolShapeEnumValues() []SandboxWarmPoolShapeEnum {
	values := make([]SandboxWarmPoolShapeEnum, 0)
	for _, v := range mappingSandboxWarmPoolShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxWarmPoolShapeEnumStringValues Enumerates the set of values in String for SandboxWarmPoolShapeEnum
func GetSandboxWarmPoolShapeEnumStringValues() []string {
	return []string{
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
}

// GetMappingSandboxWarmPoolShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxWarmPoolShapeEnum(val string) (SandboxWarmPoolShapeEnum, bool) {
	enum, ok := mappingSandboxWarmPoolShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SandboxWarmPoolLifecycleStateEnum Enum with underlying type: string
type SandboxWarmPoolLifecycleStateEnum string

// Set of constants representing the allowable values for SandboxWarmPoolLifecycleStateEnum
const (
	SandboxWarmPoolLifecycleStateCreating SandboxWarmPoolLifecycleStateEnum = "CREATING"
	SandboxWarmPoolLifecycleStateActive   SandboxWarmPoolLifecycleStateEnum = "ACTIVE"
	SandboxWarmPoolLifecycleStateUpdating SandboxWarmPoolLifecycleStateEnum = "UPDATING"
	SandboxWarmPoolLifecycleStateFailed   SandboxWarmPoolLifecycleStateEnum = "FAILED"
	SandboxWarmPoolLifecycleStateDeleting SandboxWarmPoolLifecycleStateEnum = "DELETING"
	SandboxWarmPoolLifecycleStateDeleted  SandboxWarmPoolLifecycleStateEnum = "DELETED"
)

var mappingSandboxWarmPoolLifecycleStateEnum = map[string]SandboxWarmPoolLifecycleStateEnum{
	"CREATING": SandboxWarmPoolLifecycleStateCreating,
	"ACTIVE":   SandboxWarmPoolLifecycleStateActive,
	"UPDATING": SandboxWarmPoolLifecycleStateUpdating,
	"FAILED":   SandboxWarmPoolLifecycleStateFailed,
	"DELETING": SandboxWarmPoolLifecycleStateDeleting,
	"DELETED":  SandboxWarmPoolLifecycleStateDeleted,
}

var mappingSandboxWarmPoolLifecycleStateEnumLowerCase = map[string]SandboxWarmPoolLifecycleStateEnum{
	"creating": SandboxWarmPoolLifecycleStateCreating,
	"active":   SandboxWarmPoolLifecycleStateActive,
	"updating": SandboxWarmPoolLifecycleStateUpdating,
	"failed":   SandboxWarmPoolLifecycleStateFailed,
	"deleting": SandboxWarmPoolLifecycleStateDeleting,
	"deleted":  SandboxWarmPoolLifecycleStateDeleted,
}

// GetSandboxWarmPoolLifecycleStateEnumValues Enumerates the set of values for SandboxWarmPoolLifecycleStateEnum
func GetSandboxWarmPoolLifecycleStateEnumValues() []SandboxWarmPoolLifecycleStateEnum {
	values := make([]SandboxWarmPoolLifecycleStateEnum, 0)
	for _, v := range mappingSandboxWarmPoolLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSandboxWarmPoolLifecycleStateEnumStringValues Enumerates the set of values in String for SandboxWarmPoolLifecycleStateEnum
func GetSandboxWarmPoolLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSandboxWarmPoolLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSandboxWarmPoolLifecycleStateEnum(val string) (SandboxWarmPoolLifecycleStateEnum, bool) {
	enum, ok := mappingSandboxWarmPoolLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
