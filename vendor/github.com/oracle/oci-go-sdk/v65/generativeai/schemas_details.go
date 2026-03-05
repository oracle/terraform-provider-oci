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

// SchemasDetails Array of database schemas or other database objects to include in enrichment pipeline.
type SchemasDetails interface {
}

type schemasdetails struct {
	JsonData       []byte
	ConnectionType string `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *schemasdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerschemasdetails schemasdetails
	s := struct {
		Model Unmarshalerschemasdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *schemasdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "DATABASE_TOOLS_CONNECTION":
		mm := SchemasDatabaseToolsConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SchemasDetails: %s.", m.ConnectionType)
		return *m, nil
	}
}

func (m schemasdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m schemasdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchemasDetailsConnectionTypeEnum Enum with underlying type: string
type SchemasDetailsConnectionTypeEnum string

// Set of constants representing the allowable values for SchemasDetailsConnectionTypeEnum
const (
	SchemasDetailsConnectionTypeDatabaseToolsConnection SchemasDetailsConnectionTypeEnum = "DATABASE_TOOLS_CONNECTION"
)

var mappingSchemasDetailsConnectionTypeEnum = map[string]SchemasDetailsConnectionTypeEnum{
	"DATABASE_TOOLS_CONNECTION": SchemasDetailsConnectionTypeDatabaseToolsConnection,
}

var mappingSchemasDetailsConnectionTypeEnumLowerCase = map[string]SchemasDetailsConnectionTypeEnum{
	"database_tools_connection": SchemasDetailsConnectionTypeDatabaseToolsConnection,
}

// GetSchemasDetailsConnectionTypeEnumValues Enumerates the set of values for SchemasDetailsConnectionTypeEnum
func GetSchemasDetailsConnectionTypeEnumValues() []SchemasDetailsConnectionTypeEnum {
	values := make([]SchemasDetailsConnectionTypeEnum, 0)
	for _, v := range mappingSchemasDetailsConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemasDetailsConnectionTypeEnumStringValues Enumerates the set of values in String for SchemasDetailsConnectionTypeEnum
func GetSchemasDetailsConnectionTypeEnumStringValues() []string {
	return []string{
		"DATABASE_TOOLS_CONNECTION",
	}
}

// GetMappingSchemasDetailsConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemasDetailsConnectionTypeEnum(val string) (SchemasDetailsConnectionTypeEnum, bool) {
	enum, ok := mappingSchemasDetailsConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
