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

// Flow Flow is a graph-based execution workflow to represent predefined and user-defined agent flows, offering flexible, composable, and debuggable automation at scale.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type Flow struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Flow.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Flow.
	LifecycleState FlowLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Flow was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated agent's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the agent that this Flow is attached to.
	AgentId *string `mandatory:"true" json:"agentId"`

	FlowConfig FlowConfig `mandatory:"true" json:"flowConfig"`

	// The date and time the Flow was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the Flow.
	Description *string `mandatory:"false" json:"description"`

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

func (m Flow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Flow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFlowLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFlowLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Flow) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeUpdated    *common.SDKTime                   `json:"timeUpdated"`
		DisplayName    *string                           `json:"displayName"`
		Description    *string                           `json:"description"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		SystemTags     map[string]map[string]interface{} `json:"systemTags"`
		Id             *string                           `json:"id"`
		LifecycleState FlowLifecycleStateEnum            `json:"lifecycleState"`
		TimeCreated    *common.SDKTime                   `json:"timeCreated"`
		CompartmentId  *string                           `json:"compartmentId"`
		AgentId        *string                           `json:"agentId"`
		FlowConfig     flowconfig                        `json:"flowConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeUpdated = model.TimeUpdated

	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.CompartmentId = model.CompartmentId

	m.AgentId = model.AgentId

	nn, e = model.FlowConfig.UnmarshalPolymorphicJSON(model.FlowConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FlowConfig = nn.(FlowConfig)
	} else {
		m.FlowConfig = nil
	}

	return
}

// FlowLifecycleStateEnum Enum with underlying type: string
type FlowLifecycleStateEnum string

// Set of constants representing the allowable values for FlowLifecycleStateEnum
const (
	FlowLifecycleStateCreating FlowLifecycleStateEnum = "CREATING"
	FlowLifecycleStateUpdating FlowLifecycleStateEnum = "UPDATING"
	FlowLifecycleStateActive   FlowLifecycleStateEnum = "ACTIVE"
	FlowLifecycleStateDeleting FlowLifecycleStateEnum = "DELETING"
	FlowLifecycleStateDeleted  FlowLifecycleStateEnum = "DELETED"
	FlowLifecycleStateFailed   FlowLifecycleStateEnum = "FAILED"
)

var mappingFlowLifecycleStateEnum = map[string]FlowLifecycleStateEnum{
	"CREATING": FlowLifecycleStateCreating,
	"UPDATING": FlowLifecycleStateUpdating,
	"ACTIVE":   FlowLifecycleStateActive,
	"DELETING": FlowLifecycleStateDeleting,
	"DELETED":  FlowLifecycleStateDeleted,
	"FAILED":   FlowLifecycleStateFailed,
}

var mappingFlowLifecycleStateEnumLowerCase = map[string]FlowLifecycleStateEnum{
	"creating": FlowLifecycleStateCreating,
	"updating": FlowLifecycleStateUpdating,
	"active":   FlowLifecycleStateActive,
	"deleting": FlowLifecycleStateDeleting,
	"deleted":  FlowLifecycleStateDeleted,
	"failed":   FlowLifecycleStateFailed,
}

// GetFlowLifecycleStateEnumValues Enumerates the set of values for FlowLifecycleStateEnum
func GetFlowLifecycleStateEnumValues() []FlowLifecycleStateEnum {
	values := make([]FlowLifecycleStateEnum, 0)
	for _, v := range mappingFlowLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFlowLifecycleStateEnumStringValues Enumerates the set of values in String for FlowLifecycleStateEnum
func GetFlowLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFlowLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlowLifecycleStateEnum(val string) (FlowLifecycleStateEnum, bool) {
	enum, ok := mappingFlowLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
