// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.cloud.oracle.com/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Endpoint To host a custom model for inference, create an endpoint for that model on a dedicated AI cluster of type HOSTING.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives OCI resource access to users. See
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and Getting Access to Generative AI Resouces (https://docs.cloud.oracle.com/iaas/Content/generative-ai/iam-policies.htm).
type Endpoint struct {

	// An OCID that uniquely identifies this endpoint resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the model that's used to create this endpoint.
	ModelId *string `mandatory:"true" json:"modelId"`

	// The compartment OCID to create the endpoint in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the dedicated AI cluster on which the model will be deployed to.
	DedicatedAiClusterId *string `mandatory:"true" json:"dedicatedAiClusterId"`

	// The date and time that the endpoint was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the endpoint.
	LifecycleState EndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The date and time that the endpoint was updated in the format of an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state of the endpoint in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	ContentModerationConfig *ContentModerationConfig `mandatory:"false" json:"contentModerationConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Endpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Endpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EndpointLifecycleStateEnum Enum with underlying type: string
type EndpointLifecycleStateEnum string

// Set of constants representing the allowable values for EndpointLifecycleStateEnum
const (
	EndpointLifecycleStateActive   EndpointLifecycleStateEnum = "ACTIVE"
	EndpointLifecycleStateCreating EndpointLifecycleStateEnum = "CREATING"
	EndpointLifecycleStateUpdating EndpointLifecycleStateEnum = "UPDATING"
	EndpointLifecycleStateDeleting EndpointLifecycleStateEnum = "DELETING"
	EndpointLifecycleStateDeleted  EndpointLifecycleStateEnum = "DELETED"
	EndpointLifecycleStateFailed   EndpointLifecycleStateEnum = "FAILED"
)

var mappingEndpointLifecycleStateEnum = map[string]EndpointLifecycleStateEnum{
	"ACTIVE":   EndpointLifecycleStateActive,
	"CREATING": EndpointLifecycleStateCreating,
	"UPDATING": EndpointLifecycleStateUpdating,
	"DELETING": EndpointLifecycleStateDeleting,
	"DELETED":  EndpointLifecycleStateDeleted,
	"FAILED":   EndpointLifecycleStateFailed,
}

var mappingEndpointLifecycleStateEnumLowerCase = map[string]EndpointLifecycleStateEnum{
	"active":   EndpointLifecycleStateActive,
	"creating": EndpointLifecycleStateCreating,
	"updating": EndpointLifecycleStateUpdating,
	"deleting": EndpointLifecycleStateDeleting,
	"deleted":  EndpointLifecycleStateDeleted,
	"failed":   EndpointLifecycleStateFailed,
}

// GetEndpointLifecycleStateEnumValues Enumerates the set of values for EndpointLifecycleStateEnum
func GetEndpointLifecycleStateEnumValues() []EndpointLifecycleStateEnum {
	values := make([]EndpointLifecycleStateEnum, 0)
	for _, v := range mappingEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for EndpointLifecycleStateEnum
func GetEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointLifecycleStateEnum(val string) (EndpointLifecycleStateEnum, bool) {
	enum, ok := mappingEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
