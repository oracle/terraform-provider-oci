// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// **Generative AI Agents API**
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

// IndexConfig **IndexConfig**
// The index configuration of Knowledge bases.
type IndexConfig interface {
}

type indexconfig struct {
	JsonData        []byte
	IndexConfigType string `json:"indexConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *indexconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerindexconfig indexconfig
	s := struct {
		Model Unmarshalerindexconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IndexConfigType = s.Model.IndexConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *indexconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.IndexConfigType {
	case "DEFAULT_INDEX_CONFIG":
		mm := DefaultIndexConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE_CONFIG":
		mm := OciDatabaseConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_OPEN_SEARCH_INDEX_CONFIG":
		mm := OciOpenSearchIndexConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for IndexConfig: %s.", m.IndexConfigType)
		return *m, nil
	}
}

func (m indexconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m indexconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IndexConfigIndexConfigTypeEnum Enum with underlying type: string
type IndexConfigIndexConfigTypeEnum string

// Set of constants representing the allowable values for IndexConfigIndexConfigTypeEnum
const (
	IndexConfigIndexConfigTypeDefaultIndexConfig       IndexConfigIndexConfigTypeEnum = "DEFAULT_INDEX_CONFIG"
	IndexConfigIndexConfigTypeOciOpenSearchIndexConfig IndexConfigIndexConfigTypeEnum = "OCI_OPEN_SEARCH_INDEX_CONFIG"
	IndexConfigIndexConfigTypeOciDatabaseConfig        IndexConfigIndexConfigTypeEnum = "OCI_DATABASE_CONFIG"
)

var mappingIndexConfigIndexConfigTypeEnum = map[string]IndexConfigIndexConfigTypeEnum{
	"DEFAULT_INDEX_CONFIG":         IndexConfigIndexConfigTypeDefaultIndexConfig,
	"OCI_OPEN_SEARCH_INDEX_CONFIG": IndexConfigIndexConfigTypeOciOpenSearchIndexConfig,
	"OCI_DATABASE_CONFIG":          IndexConfigIndexConfigTypeOciDatabaseConfig,
}

var mappingIndexConfigIndexConfigTypeEnumLowerCase = map[string]IndexConfigIndexConfigTypeEnum{
	"default_index_config":         IndexConfigIndexConfigTypeDefaultIndexConfig,
	"oci_open_search_index_config": IndexConfigIndexConfigTypeOciOpenSearchIndexConfig,
	"oci_database_config":          IndexConfigIndexConfigTypeOciDatabaseConfig,
}

// GetIndexConfigIndexConfigTypeEnumValues Enumerates the set of values for IndexConfigIndexConfigTypeEnum
func GetIndexConfigIndexConfigTypeEnumValues() []IndexConfigIndexConfigTypeEnum {
	values := make([]IndexConfigIndexConfigTypeEnum, 0)
	for _, v := range mappingIndexConfigIndexConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIndexConfigIndexConfigTypeEnumStringValues Enumerates the set of values in String for IndexConfigIndexConfigTypeEnum
func GetIndexConfigIndexConfigTypeEnumStringValues() []string {
	return []string{
		"DEFAULT_INDEX_CONFIG",
		"OCI_OPEN_SEARCH_INDEX_CONFIG",
		"OCI_DATABASE_CONFIG",
	}
}

// GetMappingIndexConfigIndexConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIndexConfigIndexConfigTypeEnum(val string) (IndexConfigIndexConfigTypeEnum, bool) {
	enum, ok := mappingIndexConfigIndexConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
