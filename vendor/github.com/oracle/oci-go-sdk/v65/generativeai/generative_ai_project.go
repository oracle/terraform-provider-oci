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

// GenerativeAiProject A GenerativeAiProject is a logical container that stores conversation, file and containers.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives OCI resource access to users. See
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and Getting Access to Generative AI Resources (https://docs.oracle.com/iaas/Content/generative-ai/iam-policies.htm).
type GenerativeAiProject struct {

	// An OCID that uniquely identifies a GenerativeAiProject.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Owning compartment OCID for a GenerativeAiProject.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time that the generativeAiProject was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of a GenerativeAiProject.
	LifecycleState GenerativeAiProjectLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// An optional description of the GenerativeAiProject.
	Description *string `mandatory:"false" json:"description"`

	// The date and time that the generativeAiProject was updated in the format of an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	ConversationConfig *ConversationConfig `mandatory:"false" json:"conversationConfig"`

	LongTermMemoryConfig *LongTermMemoryConfig `mandatory:"false" json:"longTermMemoryConfig"`

	ShortTermMemoryOptimizationConfig *ShortTermMemoryOptimizationConfig `mandatory:"false" json:"shortTermMemoryOptimizationConfig"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m GenerativeAiProject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerativeAiProject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerativeAiProjectLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGenerativeAiProjectLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerativeAiProjectLifecycleStateEnum Enum with underlying type: string
type GenerativeAiProjectLifecycleStateEnum string

// Set of constants representing the allowable values for GenerativeAiProjectLifecycleStateEnum
const (
	GenerativeAiProjectLifecycleStateActive   GenerativeAiProjectLifecycleStateEnum = "ACTIVE"
	GenerativeAiProjectLifecycleStateCreating GenerativeAiProjectLifecycleStateEnum = "CREATING"
	GenerativeAiProjectLifecycleStateUpdating GenerativeAiProjectLifecycleStateEnum = "UPDATING"
	GenerativeAiProjectLifecycleStateDeleting GenerativeAiProjectLifecycleStateEnum = "DELETING"
	GenerativeAiProjectLifecycleStateDeleted  GenerativeAiProjectLifecycleStateEnum = "DELETED"
	GenerativeAiProjectLifecycleStateFailed   GenerativeAiProjectLifecycleStateEnum = "FAILED"
)

var mappingGenerativeAiProjectLifecycleStateEnum = map[string]GenerativeAiProjectLifecycleStateEnum{
	"ACTIVE":   GenerativeAiProjectLifecycleStateActive,
	"CREATING": GenerativeAiProjectLifecycleStateCreating,
	"UPDATING": GenerativeAiProjectLifecycleStateUpdating,
	"DELETING": GenerativeAiProjectLifecycleStateDeleting,
	"DELETED":  GenerativeAiProjectLifecycleStateDeleted,
	"FAILED":   GenerativeAiProjectLifecycleStateFailed,
}

var mappingGenerativeAiProjectLifecycleStateEnumLowerCase = map[string]GenerativeAiProjectLifecycleStateEnum{
	"active":   GenerativeAiProjectLifecycleStateActive,
	"creating": GenerativeAiProjectLifecycleStateCreating,
	"updating": GenerativeAiProjectLifecycleStateUpdating,
	"deleting": GenerativeAiProjectLifecycleStateDeleting,
	"deleted":  GenerativeAiProjectLifecycleStateDeleted,
	"failed":   GenerativeAiProjectLifecycleStateFailed,
}

// GetGenerativeAiProjectLifecycleStateEnumValues Enumerates the set of values for GenerativeAiProjectLifecycleStateEnum
func GetGenerativeAiProjectLifecycleStateEnumValues() []GenerativeAiProjectLifecycleStateEnum {
	values := make([]GenerativeAiProjectLifecycleStateEnum, 0)
	for _, v := range mappingGenerativeAiProjectLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerativeAiProjectLifecycleStateEnumStringValues Enumerates the set of values in String for GenerativeAiProjectLifecycleStateEnum
func GetGenerativeAiProjectLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingGenerativeAiProjectLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerativeAiProjectLifecycleStateEnum(val string) (GenerativeAiProjectLifecycleStateEnum, bool) {
	enum, ok := mappingGenerativeAiProjectLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
