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

// OciDatabaseConfig The details of the customer Database Connection.
type OciDatabaseConfig struct {
	DatabaseConnection DatabaseConnection `mandatory:"true" json:"databaseConnection"`

	// Array of Database functions to be used.
	DatabaseFunctions []DatabaseFunction `mandatory:"true" json:"databaseFunctions"`
}

func (m OciDatabaseConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciDatabaseConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciDatabaseConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciDatabaseConfig OciDatabaseConfig
	s := struct {
		DiscriminatorParam string `json:"indexConfigType"`
		MarshalTypeOciDatabaseConfig
	}{
		"OCI_DATABASE_CONFIG",
		(MarshalTypeOciDatabaseConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OciDatabaseConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DatabaseConnection databaseconnection `json:"databaseConnection"`
		DatabaseFunctions  []DatabaseFunction `json:"databaseFunctions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.DatabaseConnection.UnmarshalPolymorphicJSON(model.DatabaseConnection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseConnection = nn.(DatabaseConnection)
	} else {
		m.DatabaseConnection = nil
	}

	m.DatabaseFunctions = make([]DatabaseFunction, len(model.DatabaseFunctions))
	copy(m.DatabaseFunctions, model.DatabaseFunctions)
	return
}
