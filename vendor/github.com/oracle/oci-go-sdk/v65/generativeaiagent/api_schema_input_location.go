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

// ApiSchemaInputLocation The input location definition for Api schema.
type ApiSchemaInputLocation interface {
}

type apischemainputlocation struct {
	JsonData                   []byte
	ApiSchemaInputLocationType string `json:"apiSchemaInputLocationType"`
}

// UnmarshalJSON unmarshals json
func (m *apischemainputlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerapischemainputlocation apischemainputlocation
	s := struct {
		Model Unmarshalerapischemainputlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ApiSchemaInputLocationType = s.Model.ApiSchemaInputLocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *apischemainputlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ApiSchemaInputLocationType {
	case "OBJECT_STORAGE_LOCATION":
		mm := ApiSchemaObjectStorageInputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INLINE":
		mm := ApiSchemaInlineInputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ApiSchemaInputLocation: %s.", m.ApiSchemaInputLocationType)
		return *m, nil
	}
}

func (m apischemainputlocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m apischemainputlocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiSchemaInputLocationApiSchemaInputLocationTypeEnum Enum with underlying type: string
type ApiSchemaInputLocationApiSchemaInputLocationTypeEnum string

// Set of constants representing the allowable values for ApiSchemaInputLocationApiSchemaInputLocationTypeEnum
const (
	ApiSchemaInputLocationApiSchemaInputLocationTypeInline                ApiSchemaInputLocationApiSchemaInputLocationTypeEnum = "INLINE"
	ApiSchemaInputLocationApiSchemaInputLocationTypeObjectStorageLocation ApiSchemaInputLocationApiSchemaInputLocationTypeEnum = "OBJECT_STORAGE_LOCATION"
)

var mappingApiSchemaInputLocationApiSchemaInputLocationTypeEnum = map[string]ApiSchemaInputLocationApiSchemaInputLocationTypeEnum{
	"INLINE":                  ApiSchemaInputLocationApiSchemaInputLocationTypeInline,
	"OBJECT_STORAGE_LOCATION": ApiSchemaInputLocationApiSchemaInputLocationTypeObjectStorageLocation,
}

var mappingApiSchemaInputLocationApiSchemaInputLocationTypeEnumLowerCase = map[string]ApiSchemaInputLocationApiSchemaInputLocationTypeEnum{
	"inline":                  ApiSchemaInputLocationApiSchemaInputLocationTypeInline,
	"object_storage_location": ApiSchemaInputLocationApiSchemaInputLocationTypeObjectStorageLocation,
}

// GetApiSchemaInputLocationApiSchemaInputLocationTypeEnumValues Enumerates the set of values for ApiSchemaInputLocationApiSchemaInputLocationTypeEnum
func GetApiSchemaInputLocationApiSchemaInputLocationTypeEnumValues() []ApiSchemaInputLocationApiSchemaInputLocationTypeEnum {
	values := make([]ApiSchemaInputLocationApiSchemaInputLocationTypeEnum, 0)
	for _, v := range mappingApiSchemaInputLocationApiSchemaInputLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetApiSchemaInputLocationApiSchemaInputLocationTypeEnumStringValues Enumerates the set of values in String for ApiSchemaInputLocationApiSchemaInputLocationTypeEnum
func GetApiSchemaInputLocationApiSchemaInputLocationTypeEnumStringValues() []string {
	return []string{
		"INLINE",
		"OBJECT_STORAGE_LOCATION",
	}
}

// GetMappingApiSchemaInputLocationApiSchemaInputLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiSchemaInputLocationApiSchemaInputLocationTypeEnum(val string) (ApiSchemaInputLocationApiSchemaInputLocationTypeEnum, bool) {
	enum, ok := mappingApiSchemaInputLocationApiSchemaInputLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
