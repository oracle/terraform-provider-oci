// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProvisionedCapacity A provisioned capacity is a resource pool of Genrative AI DACs with properties like AgentRuntimeVersion, unit size, and a tool name to DAC ID mapping.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type ProvisionedCapacity struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the provisioned capacity.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Provisioned Capacity Unit corresponds to the amount of characters processed per minute.
	NumberOfUnits *int `mandatory:"true" json:"numberOfUnits"`

	// The date and time the provisioned capacity was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the provisioned capacity.
	LifecycleState ProvisionedCapacityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the provisioned capacity.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the provisioned capacity.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the provisioned capacity was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

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

func (m ProvisionedCapacity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProvisionedCapacity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProvisionedCapacityLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProvisionedCapacityLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProvisionedCapacityLifecycleStateEnum Enum with underlying type: string
type ProvisionedCapacityLifecycleStateEnum string

// Set of constants representing the allowable values for ProvisionedCapacityLifecycleStateEnum
const (
	ProvisionedCapacityLifecycleStateCreating ProvisionedCapacityLifecycleStateEnum = "CREATING"
	ProvisionedCapacityLifecycleStateUpdating ProvisionedCapacityLifecycleStateEnum = "UPDATING"
	ProvisionedCapacityLifecycleStateActive   ProvisionedCapacityLifecycleStateEnum = "ACTIVE"
	ProvisionedCapacityLifecycleStateDeleting ProvisionedCapacityLifecycleStateEnum = "DELETING"
	ProvisionedCapacityLifecycleStateDeleted  ProvisionedCapacityLifecycleStateEnum = "DELETED"
	ProvisionedCapacityLifecycleStateFailed   ProvisionedCapacityLifecycleStateEnum = "FAILED"
)

var mappingProvisionedCapacityLifecycleStateEnum = map[string]ProvisionedCapacityLifecycleStateEnum{
	"CREATING": ProvisionedCapacityLifecycleStateCreating,
	"UPDATING": ProvisionedCapacityLifecycleStateUpdating,
	"ACTIVE":   ProvisionedCapacityLifecycleStateActive,
	"DELETING": ProvisionedCapacityLifecycleStateDeleting,
	"DELETED":  ProvisionedCapacityLifecycleStateDeleted,
	"FAILED":   ProvisionedCapacityLifecycleStateFailed,
}

var mappingProvisionedCapacityLifecycleStateEnumLowerCase = map[string]ProvisionedCapacityLifecycleStateEnum{
	"creating": ProvisionedCapacityLifecycleStateCreating,
	"updating": ProvisionedCapacityLifecycleStateUpdating,
	"active":   ProvisionedCapacityLifecycleStateActive,
	"deleting": ProvisionedCapacityLifecycleStateDeleting,
	"deleted":  ProvisionedCapacityLifecycleStateDeleted,
	"failed":   ProvisionedCapacityLifecycleStateFailed,
}

// GetProvisionedCapacityLifecycleStateEnumValues Enumerates the set of values for ProvisionedCapacityLifecycleStateEnum
func GetProvisionedCapacityLifecycleStateEnumValues() []ProvisionedCapacityLifecycleStateEnum {
	values := make([]ProvisionedCapacityLifecycleStateEnum, 0)
	for _, v := range mappingProvisionedCapacityLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProvisionedCapacityLifecycleStateEnumStringValues Enumerates the set of values in String for ProvisionedCapacityLifecycleStateEnum
func GetProvisionedCapacityLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingProvisionedCapacityLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProvisionedCapacityLifecycleStateEnum(val string) (ProvisionedCapacityLifecycleStateEnum, bool) {
	enum, ok := mappingProvisionedCapacityLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
