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

// HttpEndpointToolConfig The configuration for HTTP endpoint Tool.
type HttpEndpointToolConfig struct {
	ApiSchema ApiSchemaInputLocation `mandatory:"true" json:"apiSchema"`

	// The subnet ID from agent developer tenancy through which the egress is going to be routed.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	HttpEndpointAuthConfig *HttpEndpointAuthConfig `mandatory:"true" json:"httpEndpointAuthConfig"`
}

func (m HttpEndpointToolConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpEndpointToolConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpEndpointToolConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpEndpointToolConfig HttpEndpointToolConfig
	s := struct {
		DiscriminatorParam string `json:"toolConfigType"`
		MarshalTypeHttpEndpointToolConfig
	}{
		"HTTP_ENDPOINT_TOOL_CONFIG",
		(MarshalTypeHttpEndpointToolConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *HttpEndpointToolConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ApiSchema              apischemainputlocation  `json:"apiSchema"`
		SubnetId               *string                 `json:"subnetId"`
		HttpEndpointAuthConfig *HttpEndpointAuthConfig `json:"httpEndpointAuthConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ApiSchema.UnmarshalPolymorphicJSON(model.ApiSchema.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ApiSchema = nn.(ApiSchemaInputLocation)
	} else {
		m.ApiSchema = nil
	}

	m.SubnetId = model.SubnetId

	m.HttpEndpointAuthConfig = model.HttpEndpointAuthConfig

	return
}
