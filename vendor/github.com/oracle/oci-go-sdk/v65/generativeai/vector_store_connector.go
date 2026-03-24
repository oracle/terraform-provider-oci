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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VectorStoreConnector A VectorStore Connector offers a lightweight and configurable mechanism to continuously synchronize data
// from external systems into the VectorStore at scale. It captures the configuration of the datasource for data
// ingestion.
type VectorStoreConnector struct {

	// An OCID that uniquely identifies a VectorStoreConnector
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// An OCID that identifies the Vector Store to which this connector is connected.
	VectorStoreId *string `mandatory:"true" json:"vectorStoreId"`

	// Owning compartment OCID for a VectorStoreConnector.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Owning tenant OCID for a VectorStoreConnector
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The date and time that the VectorStoreConnector was created in the format of an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of a VectorStoreConnector.
	LifecycleState VectorStoreConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Configuration ConnectorConfiguration `mandatory:"true" json:"configuration"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// An optional description of the VectorStoreConnector
	Description *string `mandatory:"false" json:"description"`

	// The date and time that the VectorStoreConnector was updated in the format of an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An optional customer Encryption Key stored in OCI Vault that can be used to decrypt the data downloaded from the data source.
	VaultSecretId *string `mandatory:"false" json:"vaultSecretId"`

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	ScheduleConfig ScheduleConfig `mandatory:"false" json:"scheduleConfig"`
}

func (m VectorStoreConnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VectorStoreConnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVectorStoreConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVectorStoreConnectorLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *VectorStoreConnector) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                                `json:"description"`
		TimeUpdated      *common.SDKTime                        `json:"timeUpdated"`
		VaultSecretId    *string                                `json:"vaultSecretId"`
		LifecycleDetails *string                                `json:"lifecycleDetails"`
		ScheduleConfig   scheduleconfig                         `json:"scheduleConfig"`
		Id               *string                                `json:"id"`
		DisplayName      *string                                `json:"displayName"`
		VectorStoreId    *string                                `json:"vectorStoreId"`
		CompartmentId    *string                                `json:"compartmentId"`
		TenantId         *string                                `json:"tenantId"`
		TimeCreated      *common.SDKTime                        `json:"timeCreated"`
		LifecycleState   VectorStoreConnectorLifecycleStateEnum `json:"lifecycleState"`
		Configuration    connectorconfiguration                 `json:"configuration"`
		FreeformTags     map[string]string                      `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{}      `json:"definedTags"`
		SystemTags       map[string]map[string]interface{}      `json:"systemTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.VaultSecretId = model.VaultSecretId

	m.LifecycleDetails = model.LifecycleDetails

	nn, e = model.ScheduleConfig.UnmarshalPolymorphicJSON(model.ScheduleConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ScheduleConfig = nn.(ScheduleConfig)
	} else {
		m.ScheduleConfig = nil
	}

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.VectorStoreId = model.VectorStoreId

	m.CompartmentId = model.CompartmentId

	m.TenantId = model.TenantId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	nn, e = model.Configuration.UnmarshalPolymorphicJSON(model.Configuration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Configuration = nn.(ConnectorConfiguration)
	} else {
		m.Configuration = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	return
}

// VectorStoreConnectorLifecycleStateEnum Enum with underlying type: string
type VectorStoreConnectorLifecycleStateEnum string

// Set of constants representing the allowable values for VectorStoreConnectorLifecycleStateEnum
const (
	VectorStoreConnectorLifecycleStateActive   VectorStoreConnectorLifecycleStateEnum = "ACTIVE"
	VectorStoreConnectorLifecycleStateCreating VectorStoreConnectorLifecycleStateEnum = "CREATING"
	VectorStoreConnectorLifecycleStateUpdating VectorStoreConnectorLifecycleStateEnum = "UPDATING"
	VectorStoreConnectorLifecycleStateDeleting VectorStoreConnectorLifecycleStateEnum = "DELETING"
	VectorStoreConnectorLifecycleStateDeleted  VectorStoreConnectorLifecycleStateEnum = "DELETED"
	VectorStoreConnectorLifecycleStateFailed   VectorStoreConnectorLifecycleStateEnum = "FAILED"
)

var mappingVectorStoreConnectorLifecycleStateEnum = map[string]VectorStoreConnectorLifecycleStateEnum{
	"ACTIVE":   VectorStoreConnectorLifecycleStateActive,
	"CREATING": VectorStoreConnectorLifecycleStateCreating,
	"UPDATING": VectorStoreConnectorLifecycleStateUpdating,
	"DELETING": VectorStoreConnectorLifecycleStateDeleting,
	"DELETED":  VectorStoreConnectorLifecycleStateDeleted,
	"FAILED":   VectorStoreConnectorLifecycleStateFailed,
}

var mappingVectorStoreConnectorLifecycleStateEnumLowerCase = map[string]VectorStoreConnectorLifecycleStateEnum{
	"active":   VectorStoreConnectorLifecycleStateActive,
	"creating": VectorStoreConnectorLifecycleStateCreating,
	"updating": VectorStoreConnectorLifecycleStateUpdating,
	"deleting": VectorStoreConnectorLifecycleStateDeleting,
	"deleted":  VectorStoreConnectorLifecycleStateDeleted,
	"failed":   VectorStoreConnectorLifecycleStateFailed,
}

// GetVectorStoreConnectorLifecycleStateEnumValues Enumerates the set of values for VectorStoreConnectorLifecycleStateEnum
func GetVectorStoreConnectorLifecycleStateEnumValues() []VectorStoreConnectorLifecycleStateEnum {
	values := make([]VectorStoreConnectorLifecycleStateEnum, 0)
	for _, v := range mappingVectorStoreConnectorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVectorStoreConnectorLifecycleStateEnumStringValues Enumerates the set of values in String for VectorStoreConnectorLifecycleStateEnum
func GetVectorStoreConnectorLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingVectorStoreConnectorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVectorStoreConnectorLifecycleStateEnum(val string) (VectorStoreConnectorLifecycleStateEnum, bool) {
	enum, ok := mappingVectorStoreConnectorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
