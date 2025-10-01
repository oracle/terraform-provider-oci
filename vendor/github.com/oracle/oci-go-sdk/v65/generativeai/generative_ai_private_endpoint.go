// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GenerativeAiPrivateEndpoint Generative AI private endpoint.
type GenerativeAiPrivateEndpoint struct {

	// The OCID of a private endpoint.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the private endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the subnet that the private endpoint belongs to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The current state of the Generative AI Private Endpoint.
	LifecycleState GenerativeAiPrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Fully qualified domain name the customer will use for access (for eg: xyz.oraclecloud.com)
	Fqdn *string `mandatory:"true" json:"fqdn"`

	// The date and time that the Generative AI private endpoint was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time that the Generative AI private endpoint was updated expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A user friendly name. It doesn't have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of this private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The detailed messages about the lifecycle state
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A list of the OCIDs of the network security groups that the private endpoint's VNIC belongs to.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private IP address (in the customer's VCN) that represents the access point for the associated endpoint service.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	PreviousState *GenerativeAiPrivateEndpoint `mandatory:"false" json:"previousState"`

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

func (m GenerativeAiPrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerativeAiPrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerativeAiPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGenerativeAiPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerativeAiPrivateEndpointLifecycleStateEnum Enum with underlying type: string
type GenerativeAiPrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for GenerativeAiPrivateEndpointLifecycleStateEnum
const (
	GenerativeAiPrivateEndpointLifecycleStateCreating GenerativeAiPrivateEndpointLifecycleStateEnum = "CREATING"
	GenerativeAiPrivateEndpointLifecycleStateActive   GenerativeAiPrivateEndpointLifecycleStateEnum = "ACTIVE"
	GenerativeAiPrivateEndpointLifecycleStateUpdating GenerativeAiPrivateEndpointLifecycleStateEnum = "UPDATING"
	GenerativeAiPrivateEndpointLifecycleStateDeleting GenerativeAiPrivateEndpointLifecycleStateEnum = "DELETING"
	GenerativeAiPrivateEndpointLifecycleStateDeleted  GenerativeAiPrivateEndpointLifecycleStateEnum = "DELETED"
	GenerativeAiPrivateEndpointLifecycleStateFailed   GenerativeAiPrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingGenerativeAiPrivateEndpointLifecycleStateEnum = map[string]GenerativeAiPrivateEndpointLifecycleStateEnum{
	"CREATING": GenerativeAiPrivateEndpointLifecycleStateCreating,
	"ACTIVE":   GenerativeAiPrivateEndpointLifecycleStateActive,
	"UPDATING": GenerativeAiPrivateEndpointLifecycleStateUpdating,
	"DELETING": GenerativeAiPrivateEndpointLifecycleStateDeleting,
	"DELETED":  GenerativeAiPrivateEndpointLifecycleStateDeleted,
	"FAILED":   GenerativeAiPrivateEndpointLifecycleStateFailed,
}

var mappingGenerativeAiPrivateEndpointLifecycleStateEnumLowerCase = map[string]GenerativeAiPrivateEndpointLifecycleStateEnum{
	"creating": GenerativeAiPrivateEndpointLifecycleStateCreating,
	"active":   GenerativeAiPrivateEndpointLifecycleStateActive,
	"updating": GenerativeAiPrivateEndpointLifecycleStateUpdating,
	"deleting": GenerativeAiPrivateEndpointLifecycleStateDeleting,
	"deleted":  GenerativeAiPrivateEndpointLifecycleStateDeleted,
	"failed":   GenerativeAiPrivateEndpointLifecycleStateFailed,
}

// GetGenerativeAiPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for GenerativeAiPrivateEndpointLifecycleStateEnum
func GetGenerativeAiPrivateEndpointLifecycleStateEnumValues() []GenerativeAiPrivateEndpointLifecycleStateEnum {
	values := make([]GenerativeAiPrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingGenerativeAiPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerativeAiPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for GenerativeAiPrivateEndpointLifecycleStateEnum
func GetGenerativeAiPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingGenerativeAiPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerativeAiPrivateEndpointLifecycleStateEnum(val string) (GenerativeAiPrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingGenerativeAiPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
