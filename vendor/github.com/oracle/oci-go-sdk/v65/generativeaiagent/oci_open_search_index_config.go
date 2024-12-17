// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// OciOpenSearchIndexConfig The details of customer managed OCI OpenSearch.
type OciOpenSearchIndexConfig struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OpenSearch Cluster.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	SecretDetail SecretDetail `mandatory:"true" json:"secretDetail"`

	// Index configuration for open search.
	Indexes []Index `mandatory:"true" json:"indexes"`
}

func (m OciOpenSearchIndexConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciOpenSearchIndexConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciOpenSearchIndexConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciOpenSearchIndexConfig OciOpenSearchIndexConfig
	s := struct {
		DiscriminatorParam string `json:"indexConfigType"`
		MarshalTypeOciOpenSearchIndexConfig
	}{
		"OCI_OPEN_SEARCH_INDEX_CONFIG",
		(MarshalTypeOciOpenSearchIndexConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OciOpenSearchIndexConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ClusterId    *string      `json:"clusterId"`
		SecretDetail secretdetail `json:"secretDetail"`
		Indexes      []Index      `json:"indexes"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ClusterId = model.ClusterId

	nn, e = model.SecretDetail.UnmarshalPolymorphicJSON(model.SecretDetail.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecretDetail = nn.(SecretDetail)
	} else {
		m.SecretDetail = nil
	}

	m.Indexes = make([]Index, len(model.Indexes))
	copy(m.Indexes, model.Indexes)
	return
}
