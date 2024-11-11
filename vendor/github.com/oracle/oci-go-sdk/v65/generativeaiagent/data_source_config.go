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

// DataSourceConfig **DataSourceConfig**
// The details of data source.
type DataSourceConfig interface {
}

type datasourceconfig struct {
	JsonData             []byte
	DataSourceConfigType string `json:"dataSourceConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *datasourceconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatasourceconfig datasourceconfig
	s := struct {
		Model Unmarshalerdatasourceconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DataSourceConfigType = s.Model.DataSourceConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datasourceconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DataSourceConfigType {
	case "OCI_OBJECT_STORAGE":
		mm := OciObjectStorageDataSourceConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DataSourceConfig: %s.", m.DataSourceConfigType)
		return *m, nil
	}
}

func (m datasourceconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m datasourceconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataSourceConfigDataSourceConfigTypeEnum Enum with underlying type: string
type DataSourceConfigDataSourceConfigTypeEnum string

// Set of constants representing the allowable values for DataSourceConfigDataSourceConfigTypeEnum
const (
	DataSourceConfigDataSourceConfigTypeOciObjectStorage DataSourceConfigDataSourceConfigTypeEnum = "OCI_OBJECT_STORAGE"
)

var mappingDataSourceConfigDataSourceConfigTypeEnum = map[string]DataSourceConfigDataSourceConfigTypeEnum{
	"OCI_OBJECT_STORAGE": DataSourceConfigDataSourceConfigTypeOciObjectStorage,
}

var mappingDataSourceConfigDataSourceConfigTypeEnumLowerCase = map[string]DataSourceConfigDataSourceConfigTypeEnum{
	"oci_object_storage": DataSourceConfigDataSourceConfigTypeOciObjectStorage,
}

// GetDataSourceConfigDataSourceConfigTypeEnumValues Enumerates the set of values for DataSourceConfigDataSourceConfigTypeEnum
func GetDataSourceConfigDataSourceConfigTypeEnumValues() []DataSourceConfigDataSourceConfigTypeEnum {
	values := make([]DataSourceConfigDataSourceConfigTypeEnum, 0)
	for _, v := range mappingDataSourceConfigDataSourceConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataSourceConfigDataSourceConfigTypeEnumStringValues Enumerates the set of values in String for DataSourceConfigDataSourceConfigTypeEnum
func GetDataSourceConfigDataSourceConfigTypeEnumStringValues() []string {
	return []string{
		"OCI_OBJECT_STORAGE",
	}
}

// GetMappingDataSourceConfigDataSourceConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataSourceConfigDataSourceConfigTypeEnum(val string) (DataSourceConfigDataSourceConfigTypeEnum, bool) {
	enum, ok := mappingDataSourceConfigDataSourceConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
