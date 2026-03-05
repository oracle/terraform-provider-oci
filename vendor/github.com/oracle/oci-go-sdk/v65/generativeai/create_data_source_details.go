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

// CreateDataSourceDetails Defines the data source that the semantic model connects to.
type CreateDataSourceDetails interface {
}

type createdatasourcedetails struct {
	JsonData       []byte
	ConnectionType string `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *createdatasourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatasourcedetails createdatasourcedetails
	s := struct {
		Model Unmarshalercreatedatasourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatasourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "DATABASE_TOOLS_CONNECTION":
		mm := CreateDataSourceDatabaseToolsConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDataSourceDetails: %s.", m.ConnectionType)
		return *m, nil
	}
}

func (m createdatasourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatasourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDataSourceDetailsConnectionTypeEnum Enum with underlying type: string
type CreateDataSourceDetailsConnectionTypeEnum string

// Set of constants representing the allowable values for CreateDataSourceDetailsConnectionTypeEnum
const (
	CreateDataSourceDetailsConnectionTypeDatabaseToolsConnection CreateDataSourceDetailsConnectionTypeEnum = "DATABASE_TOOLS_CONNECTION"
)

var mappingCreateDataSourceDetailsConnectionTypeEnum = map[string]CreateDataSourceDetailsConnectionTypeEnum{
	"DATABASE_TOOLS_CONNECTION": CreateDataSourceDetailsConnectionTypeDatabaseToolsConnection,
}

var mappingCreateDataSourceDetailsConnectionTypeEnumLowerCase = map[string]CreateDataSourceDetailsConnectionTypeEnum{
	"database_tools_connection": CreateDataSourceDetailsConnectionTypeDatabaseToolsConnection,
}

// GetCreateDataSourceDetailsConnectionTypeEnumValues Enumerates the set of values for CreateDataSourceDetailsConnectionTypeEnum
func GetCreateDataSourceDetailsConnectionTypeEnumValues() []CreateDataSourceDetailsConnectionTypeEnum {
	values := make([]CreateDataSourceDetailsConnectionTypeEnum, 0)
	for _, v := range mappingCreateDataSourceDetailsConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDataSourceDetailsConnectionTypeEnumStringValues Enumerates the set of values in String for CreateDataSourceDetailsConnectionTypeEnum
func GetCreateDataSourceDetailsConnectionTypeEnumStringValues() []string {
	return []string{
		"DATABASE_TOOLS_CONNECTION",
	}
}

// GetMappingCreateDataSourceDetailsConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDataSourceDetailsConnectionTypeEnum(val string) (CreateDataSourceDetailsConnectionTypeEnum, bool) {
	enum, ok := mappingCreateDataSourceDetailsConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
