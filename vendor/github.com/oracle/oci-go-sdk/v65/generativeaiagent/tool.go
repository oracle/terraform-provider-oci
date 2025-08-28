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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Tool The description of Tool.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type Tool struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Tool.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Tool.
	LifecycleState ToolLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Tool was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Description of the Tool.
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the agent that this Tool is attached to.
	AgentId *string `mandatory:"true" json:"agentId"`

	ToolConfig ToolConfig `mandatory:"true" json:"toolConfig"`

	// The date and time the Tool was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Key-value pairs to allow additional configurations.
	Metadata map[string]string `mandatory:"false" json:"metadata"`

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

func (m Tool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Tool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingToolLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetToolLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Tool) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeUpdated    *common.SDKTime                   `json:"timeUpdated"`
		DisplayName    *string                           `json:"displayName"`
		Metadata       map[string]string                 `json:"metadata"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		SystemTags     map[string]map[string]interface{} `json:"systemTags"`
		Id             *string                           `json:"id"`
		LifecycleState ToolLifecycleStateEnum            `json:"lifecycleState"`
		TimeCreated    *common.SDKTime                   `json:"timeCreated"`
		Description    *string                           `json:"description"`
		CompartmentId  *string                           `json:"compartmentId"`
		AgentId        *string                           `json:"agentId"`
		ToolConfig     toolconfig                        `json:"toolConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeUpdated = model.TimeUpdated

	m.DisplayName = model.DisplayName

	m.Metadata = model.Metadata

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.Description = model.Description

	m.CompartmentId = model.CompartmentId

	m.AgentId = model.AgentId

	nn, e = model.ToolConfig.UnmarshalPolymorphicJSON(model.ToolConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ToolConfig = nn.(ToolConfig)
	} else {
		m.ToolConfig = nil
	}

	return
}

// ToolLifecycleStateEnum Enum with underlying type: string
type ToolLifecycleStateEnum string

// Set of constants representing the allowable values for ToolLifecycleStateEnum
const (
	ToolLifecycleStateCreating ToolLifecycleStateEnum = "CREATING"
	ToolLifecycleStateUpdating ToolLifecycleStateEnum = "UPDATING"
	ToolLifecycleStateActive   ToolLifecycleStateEnum = "ACTIVE"
	ToolLifecycleStateDeleting ToolLifecycleStateEnum = "DELETING"
	ToolLifecycleStateDeleted  ToolLifecycleStateEnum = "DELETED"
	ToolLifecycleStateFailed   ToolLifecycleStateEnum = "FAILED"
)

var mappingToolLifecycleStateEnum = map[string]ToolLifecycleStateEnum{
	"CREATING": ToolLifecycleStateCreating,
	"UPDATING": ToolLifecycleStateUpdating,
	"ACTIVE":   ToolLifecycleStateActive,
	"DELETING": ToolLifecycleStateDeleting,
	"DELETED":  ToolLifecycleStateDeleted,
	"FAILED":   ToolLifecycleStateFailed,
}

var mappingToolLifecycleStateEnumLowerCase = map[string]ToolLifecycleStateEnum{
	"creating": ToolLifecycleStateCreating,
	"updating": ToolLifecycleStateUpdating,
	"active":   ToolLifecycleStateActive,
	"deleting": ToolLifecycleStateDeleting,
	"deleted":  ToolLifecycleStateDeleted,
	"failed":   ToolLifecycleStateFailed,
}

// GetToolLifecycleStateEnumValues Enumerates the set of values for ToolLifecycleStateEnum
func GetToolLifecycleStateEnumValues() []ToolLifecycleStateEnum {
	values := make([]ToolLifecycleStateEnum, 0)
	for _, v := range mappingToolLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetToolLifecycleStateEnumStringValues Enumerates the set of values in String for ToolLifecycleStateEnum
func GetToolLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingToolLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingToolLifecycleStateEnum(val string) (ToolLifecycleStateEnum, bool) {
	enum, ok := mappingToolLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
