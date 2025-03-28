// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// Agent An agent is an LLM-based autonomous system that understands and generates human-like text, enabling natural-language processing interactions. OCI Generative AI Agents supports retrieval-augmented generation (RAG) agents. A RAG agent connects to a data source, retrieves data, and augments model responses with the information from the data sources to generate more relevant responses.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type Agent struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the agent was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the agent.
	LifecycleState AgentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description about the agent.
	Description *string `mandatory:"false" json:"description"`

	// List of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledgeBases associated with agent. This field is deprecated and will be removed after March 26 2026.
	KnowledgeBaseIds []string `mandatory:"false" json:"knowledgeBaseIds"`

	// Details about purpose and responsibility of the agent
	WelcomeMessage *string `mandatory:"false" json:"welcomeMessage"`

	LlmConfig *LlmConfig `mandatory:"false" json:"llmConfig"`

	// The date and time the agent was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the agent in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
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

func (m Agent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Agent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAgentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAgentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AgentLifecycleStateEnum Enum with underlying type: string
type AgentLifecycleStateEnum string

// Set of constants representing the allowable values for AgentLifecycleStateEnum
const (
	AgentLifecycleStateCreating AgentLifecycleStateEnum = "CREATING"
	AgentLifecycleStateUpdating AgentLifecycleStateEnum = "UPDATING"
	AgentLifecycleStateActive   AgentLifecycleStateEnum = "ACTIVE"
	AgentLifecycleStateDeleting AgentLifecycleStateEnum = "DELETING"
	AgentLifecycleStateDeleted  AgentLifecycleStateEnum = "DELETED"
	AgentLifecycleStateFailed   AgentLifecycleStateEnum = "FAILED"
)

var mappingAgentLifecycleStateEnum = map[string]AgentLifecycleStateEnum{
	"CREATING": AgentLifecycleStateCreating,
	"UPDATING": AgentLifecycleStateUpdating,
	"ACTIVE":   AgentLifecycleStateActive,
	"DELETING": AgentLifecycleStateDeleting,
	"DELETED":  AgentLifecycleStateDeleted,
	"FAILED":   AgentLifecycleStateFailed,
}

var mappingAgentLifecycleStateEnumLowerCase = map[string]AgentLifecycleStateEnum{
	"creating": AgentLifecycleStateCreating,
	"updating": AgentLifecycleStateUpdating,
	"active":   AgentLifecycleStateActive,
	"deleting": AgentLifecycleStateDeleting,
	"deleted":  AgentLifecycleStateDeleted,
	"failed":   AgentLifecycleStateFailed,
}

// GetAgentLifecycleStateEnumValues Enumerates the set of values for AgentLifecycleStateEnum
func GetAgentLifecycleStateEnumValues() []AgentLifecycleStateEnum {
	values := make([]AgentLifecycleStateEnum, 0)
	for _, v := range mappingAgentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAgentLifecycleStateEnumStringValues Enumerates the set of values in String for AgentLifecycleStateEnum
func GetAgentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAgentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgentLifecycleStateEnum(val string) (AgentLifecycleStateEnum, bool) {
	enum, ok := mappingAgentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
