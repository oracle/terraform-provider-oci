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

// ApiKey ApiKeys are resources used to access GenAI models.
// You must be authorized through an IAM policy to use any API operations. If you're not authorized, contact an administrator who manages OCI resource access. See
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and Getting Access to Generative AI Resources (https://docs.oracle.com/iaas/Content/generative-ai/iam-policies.htm).
type ApiKey struct {

	// the ApiKey id.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The compartment OCID to create the apiKey in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time that the ApiKey was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The list of keys.
	Keys []ApiKeyItem `mandatory:"true" json:"keys"`

	// The current state of the API key.
	LifecycleState ApiKeyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An optional description of the Api key.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the ApiKey was updated, in the format defined by RFC 3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state with detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ApiKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiKeyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApiKeyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiKeyLifecycleStateEnum Enum with underlying type: string
type ApiKeyLifecycleStateEnum string

// Set of constants representing the allowable values for ApiKeyLifecycleStateEnum
const (
	ApiKeyLifecycleStateCreating ApiKeyLifecycleStateEnum = "CREATING"
	ApiKeyLifecycleStateActive   ApiKeyLifecycleStateEnum = "ACTIVE"
	ApiKeyLifecycleStateUpdating ApiKeyLifecycleStateEnum = "UPDATING"
	ApiKeyLifecycleStateDeleting ApiKeyLifecycleStateEnum = "DELETING"
	ApiKeyLifecycleStateDeleted  ApiKeyLifecycleStateEnum = "DELETED"
	ApiKeyLifecycleStateFailed   ApiKeyLifecycleStateEnum = "FAILED"
)

var mappingApiKeyLifecycleStateEnum = map[string]ApiKeyLifecycleStateEnum{
	"CREATING": ApiKeyLifecycleStateCreating,
	"ACTIVE":   ApiKeyLifecycleStateActive,
	"UPDATING": ApiKeyLifecycleStateUpdating,
	"DELETING": ApiKeyLifecycleStateDeleting,
	"DELETED":  ApiKeyLifecycleStateDeleted,
	"FAILED":   ApiKeyLifecycleStateFailed,
}

var mappingApiKeyLifecycleStateEnumLowerCase = map[string]ApiKeyLifecycleStateEnum{
	"creating": ApiKeyLifecycleStateCreating,
	"active":   ApiKeyLifecycleStateActive,
	"updating": ApiKeyLifecycleStateUpdating,
	"deleting": ApiKeyLifecycleStateDeleting,
	"deleted":  ApiKeyLifecycleStateDeleted,
	"failed":   ApiKeyLifecycleStateFailed,
}

// GetApiKeyLifecycleStateEnumValues Enumerates the set of values for ApiKeyLifecycleStateEnum
func GetApiKeyLifecycleStateEnumValues() []ApiKeyLifecycleStateEnum {
	values := make([]ApiKeyLifecycleStateEnum, 0)
	for _, v := range mappingApiKeyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApiKeyLifecycleStateEnumStringValues Enumerates the set of values in String for ApiKeyLifecycleStateEnum
func GetApiKeyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingApiKeyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiKeyLifecycleStateEnum(val string) (ApiKeyLifecycleStateEnum, bool) {
	enum, ok := mappingApiKeyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
