// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSource A data source points to the source of your data. After you add a data source to a knowledge base, you must ingest the data source's data, so that agents using the knowledge base can refer to the data.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type DataSource struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data source.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent KnowledgeBase.
	KnowledgeBaseId *string `mandatory:"true" json:"knowledgeBaseId"`

	DataSourceConfig DataSourceConfig `mandatory:"true" json:"dataSourceConfig"`

	// The date and time the data source was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the data source.
	LifecycleState DataSourceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A description of the data source.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the data source was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the data source in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DataSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDataSourceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DataSource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		DisplayName      *string                           `json:"displayName"`
		CompartmentId    *string                           `json:"compartmentId"`
		KnowledgeBaseId  *string                           `json:"knowledgeBaseId"`
		DataSourceConfig datasourceconfig                  `json:"dataSourceConfig"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		LifecycleState   DataSourceLifecycleStateEnum      `json:"lifecycleState"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.KnowledgeBaseId = model.KnowledgeBaseId

	nn, e = model.DataSourceConfig.UnmarshalPolymorphicJSON(model.DataSourceConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataSourceConfig = nn.(DataSourceConfig)
	} else {
		m.DataSourceConfig = nil
	}

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// DataSourceLifecycleStateEnum Enum with underlying type: string
type DataSourceLifecycleStateEnum string

// Set of constants representing the allowable values for DataSourceLifecycleStateEnum
const (
	DataSourceLifecycleStateCreating DataSourceLifecycleStateEnum = "CREATING"
	DataSourceLifecycleStateActive   DataSourceLifecycleStateEnum = "ACTIVE"
	DataSourceLifecycleStateInactive DataSourceLifecycleStateEnum = "INACTIVE"
	DataSourceLifecycleStateUpdating DataSourceLifecycleStateEnum = "UPDATING"
	DataSourceLifecycleStateDeleting DataSourceLifecycleStateEnum = "DELETING"
	DataSourceLifecycleStateDeleted  DataSourceLifecycleStateEnum = "DELETED"
	DataSourceLifecycleStateFailed   DataSourceLifecycleStateEnum = "FAILED"
)

var mappingDataSourceLifecycleStateEnum = map[string]DataSourceLifecycleStateEnum{
	"CREATING": DataSourceLifecycleStateCreating,
	"ACTIVE":   DataSourceLifecycleStateActive,
	"INACTIVE": DataSourceLifecycleStateInactive,
	"UPDATING": DataSourceLifecycleStateUpdating,
	"DELETING": DataSourceLifecycleStateDeleting,
	"DELETED":  DataSourceLifecycleStateDeleted,
	"FAILED":   DataSourceLifecycleStateFailed,
}

var mappingDataSourceLifecycleStateEnumLowerCase = map[string]DataSourceLifecycleStateEnum{
	"creating": DataSourceLifecycleStateCreating,
	"active":   DataSourceLifecycleStateActive,
	"inactive": DataSourceLifecycleStateInactive,
	"updating": DataSourceLifecycleStateUpdating,
	"deleting": DataSourceLifecycleStateDeleting,
	"deleted":  DataSourceLifecycleStateDeleted,
	"failed":   DataSourceLifecycleStateFailed,
}

// GetDataSourceLifecycleStateEnumValues Enumerates the set of values for DataSourceLifecycleStateEnum
func GetDataSourceLifecycleStateEnumValues() []DataSourceLifecycleStateEnum {
	values := make([]DataSourceLifecycleStateEnum, 0)
	for _, v := range mappingDataSourceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataSourceLifecycleStateEnumStringValues Enumerates the set of values in String for DataSourceLifecycleStateEnum
func GetDataSourceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDataSourceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataSourceLifecycleStateEnum(val string) (DataSourceLifecycleStateEnum, bool) {
	enum, ok := mappingDataSourceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
