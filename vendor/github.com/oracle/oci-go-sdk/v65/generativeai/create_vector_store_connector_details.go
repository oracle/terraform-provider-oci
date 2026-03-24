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

// CreateVectorStoreConnectorDetails The data to create a VectorStoreConnector.
type CreateVectorStoreConnectorDetails struct {

	// Owning compartment OCID for a ConversationStore.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// An OCID that identifies the Vector Store to which this connector is connected.
	VectorStoreId *string `mandatory:"true" json:"vectorStoreId"`

	// A user-friendly name for the VectorStoreConnector.
	DisplayName *string `mandatory:"true" json:"displayName"`

	Configuration ConnectorConfiguration `mandatory:"true" json:"configuration"`

	// An optional description of the VectorStoreConnector.
	Description *string `mandatory:"false" json:"description"`

	ScheduleConfig ScheduleConfig `mandatory:"false" json:"scheduleConfig"`

	// An optional customer Encryption Key stored in OCI Vault that can be used to decrypt the data downloaded from the data source.
	VaultSecretId *string `mandatory:"false" json:"vaultSecretId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateVectorStoreConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVectorStoreConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateVectorStoreConnectorDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description    *string                           `json:"description"`
		ScheduleConfig scheduleconfig                    `json:"scheduleConfig"`
		VaultSecretId  *string                           `json:"vaultSecretId"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId  *string                           `json:"compartmentId"`
		VectorStoreId  *string                           `json:"vectorStoreId"`
		DisplayName    *string                           `json:"displayName"`
		Configuration  connectorconfiguration            `json:"configuration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	nn, e = model.ScheduleConfig.UnmarshalPolymorphicJSON(model.ScheduleConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ScheduleConfig = nn.(ScheduleConfig)
	} else {
		m.ScheduleConfig = nil
	}

	m.VaultSecretId = model.VaultSecretId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.VectorStoreId = model.VectorStoreId

	m.DisplayName = model.DisplayName

	nn, e = model.Configuration.UnmarshalPolymorphicJSON(model.Configuration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Configuration = nn.(ConnectorConfiguration)
	} else {
		m.Configuration = nil
	}

	return
}
