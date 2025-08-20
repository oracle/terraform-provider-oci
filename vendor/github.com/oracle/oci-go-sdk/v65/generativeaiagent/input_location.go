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

// InputLocation The input location definition.
type InputLocation interface {
}

type inputlocation struct {
	JsonData          []byte
	InputLocationType string `json:"inputLocationType"`
}

// UnmarshalJSON unmarshals json
func (m *inputlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinputlocation inputlocation
	s := struct {
		Model Unmarshalerinputlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InputLocationType = s.Model.InputLocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.InputLocationType {
	case "OBJECT_STORAGE_PREFIX":
		mm := ObjectStorageInputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INLINE":
		mm := InlineInputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for InputLocation: %s.", m.InputLocationType)
		return *m, nil
	}
}

func (m inputlocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m inputlocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InputLocationInputLocationTypeEnum Enum with underlying type: string
type InputLocationInputLocationTypeEnum string

// Set of constants representing the allowable values for InputLocationInputLocationTypeEnum
const (
	InputLocationInputLocationTypeInline              InputLocationInputLocationTypeEnum = "INLINE"
	InputLocationInputLocationTypeObjectStoragePrefix InputLocationInputLocationTypeEnum = "OBJECT_STORAGE_PREFIX"
)

var mappingInputLocationInputLocationTypeEnum = map[string]InputLocationInputLocationTypeEnum{
	"INLINE":                InputLocationInputLocationTypeInline,
	"OBJECT_STORAGE_PREFIX": InputLocationInputLocationTypeObjectStoragePrefix,
}

var mappingInputLocationInputLocationTypeEnumLowerCase = map[string]InputLocationInputLocationTypeEnum{
	"inline":                InputLocationInputLocationTypeInline,
	"object_storage_prefix": InputLocationInputLocationTypeObjectStoragePrefix,
}

// GetInputLocationInputLocationTypeEnumValues Enumerates the set of values for InputLocationInputLocationTypeEnum
func GetInputLocationInputLocationTypeEnumValues() []InputLocationInputLocationTypeEnum {
	values := make([]InputLocationInputLocationTypeEnum, 0)
	for _, v := range mappingInputLocationInputLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInputLocationInputLocationTypeEnumStringValues Enumerates the set of values in String for InputLocationInputLocationTypeEnum
func GetInputLocationInputLocationTypeEnumStringValues() []string {
	return []string{
		"INLINE",
		"OBJECT_STORAGE_PREFIX",
	}
}

// GetMappingInputLocationInputLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputLocationInputLocationTypeEnum(val string) (InputLocationInputLocationTypeEnum, bool) {
	enum, ok := mappingInputLocationInputLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
