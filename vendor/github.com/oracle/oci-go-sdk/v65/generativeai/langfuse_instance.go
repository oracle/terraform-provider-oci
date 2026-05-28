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

// LangfuseInstance An OCI Managed Langfuse instance.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives OCI resource access to users. See
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type LangfuseInstance struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that uniquely identifies a Langfuse instance.
	Id *string `mandatory:"true" json:"id"`

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Langfuse instance.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time that the Langfuse instance was created. Format is defined by RFC3339 (https://www.rfc-editor.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the Langfuse instance.
	LifecycleState LangfuseInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The HTTPS URL of the identity domain used for single sign-on. The first DNS label of this URL is used as the identity domain stripe ID.
	IdentityDomainUrl *string `mandatory:"true" json:"identityDomainUrl"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the Langfuse instance.
	Description *string `mandatory:"false" json:"description"`

	// The date and time that the Langfuse instance was updated. Format is defined by RFC3339 (https://www.rfc-editor.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The HTTPS endpoint for the Langfuse instance.
	Endpoint *string `mandatory:"false" json:"endpoint"`

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

func (m LangfuseInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LangfuseInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLangfuseInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLangfuseInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LangfuseInstanceLifecycleStateEnum Enum with underlying type: string
type LangfuseInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for LangfuseInstanceLifecycleStateEnum
const (
	LangfuseInstanceLifecycleStateActive   LangfuseInstanceLifecycleStateEnum = "ACTIVE"
	LangfuseInstanceLifecycleStateCreating LangfuseInstanceLifecycleStateEnum = "CREATING"
	LangfuseInstanceLifecycleStateUpdating LangfuseInstanceLifecycleStateEnum = "UPDATING"
	LangfuseInstanceLifecycleStateDeleting LangfuseInstanceLifecycleStateEnum = "DELETING"
	LangfuseInstanceLifecycleStateDeleted  LangfuseInstanceLifecycleStateEnum = "DELETED"
	LangfuseInstanceLifecycleStateFailed   LangfuseInstanceLifecycleStateEnum = "FAILED"
)

var mappingLangfuseInstanceLifecycleStateEnum = map[string]LangfuseInstanceLifecycleStateEnum{
	"ACTIVE":   LangfuseInstanceLifecycleStateActive,
	"CREATING": LangfuseInstanceLifecycleStateCreating,
	"UPDATING": LangfuseInstanceLifecycleStateUpdating,
	"DELETING": LangfuseInstanceLifecycleStateDeleting,
	"DELETED":  LangfuseInstanceLifecycleStateDeleted,
	"FAILED":   LangfuseInstanceLifecycleStateFailed,
}

var mappingLangfuseInstanceLifecycleStateEnumLowerCase = map[string]LangfuseInstanceLifecycleStateEnum{
	"active":   LangfuseInstanceLifecycleStateActive,
	"creating": LangfuseInstanceLifecycleStateCreating,
	"updating": LangfuseInstanceLifecycleStateUpdating,
	"deleting": LangfuseInstanceLifecycleStateDeleting,
	"deleted":  LangfuseInstanceLifecycleStateDeleted,
	"failed":   LangfuseInstanceLifecycleStateFailed,
}

// GetLangfuseInstanceLifecycleStateEnumValues Enumerates the set of values for LangfuseInstanceLifecycleStateEnum
func GetLangfuseInstanceLifecycleStateEnumValues() []LangfuseInstanceLifecycleStateEnum {
	values := make([]LangfuseInstanceLifecycleStateEnum, 0)
	for _, v := range mappingLangfuseInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLangfuseInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for LangfuseInstanceLifecycleStateEnum
func GetLangfuseInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingLangfuseInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLangfuseInstanceLifecycleStateEnum(val string) (LangfuseInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingLangfuseInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
