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

// OciObjectStorageDataSourceConfig The details of OCI Search with OpenSearch data source information.
type OciObjectStorageDataSourceConfig struct {

	// Flag to enable or disable multi modality such as image processing while ingestion of data. True enable the processing and false exclude the multi modality contents during ingestion.
	ShouldEnableMultiModality *bool `mandatory:"false" json:"shouldEnableMultiModality"`

	// The locations of data items in Object Storage, can either be an object (File) or a prefix (folder).
	ObjectStoragePrefixes []ObjectStoragePrefix `mandatory:"false" json:"objectStoragePrefixes"`
}

// GetShouldEnableMultiModality returns ShouldEnableMultiModality
func (m OciObjectStorageDataSourceConfig) GetShouldEnableMultiModality() *bool {
	return m.ShouldEnableMultiModality
}

func (m OciObjectStorageDataSourceConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciObjectStorageDataSourceConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciObjectStorageDataSourceConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciObjectStorageDataSourceConfig OciObjectStorageDataSourceConfig
	s := struct {
		DiscriminatorParam string `json:"dataSourceConfigType"`
		MarshalTypeOciObjectStorageDataSourceConfig
	}{
		"OCI_OBJECT_STORAGE",
		(MarshalTypeOciObjectStorageDataSourceConfig)(m),
	}

	return json.Marshal(&s)
}
