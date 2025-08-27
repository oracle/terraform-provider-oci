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

// OutputLocation Location of the output.
type OutputLocation interface {
}

type outputlocation struct {
	JsonData           []byte
	OutputLocationType string `json:"outputLocationType"`
}

// UnmarshalJSON unmarshals json
func (m *outputlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroutputlocation outputlocation
	s := struct {
		Model Unmarshaleroutputlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OutputLocationType = s.Model.OutputLocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *outputlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputLocationType {
	case "OBJECT_STORAGE_PREFIX":
		mm := ObjectStoragePrefixOutputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for OutputLocation: %s.", m.OutputLocationType)
		return *m, nil
	}
}

func (m outputlocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m outputlocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OutputLocationOutputLocationTypeEnum Enum with underlying type: string
type OutputLocationOutputLocationTypeEnum string

// Set of constants representing the allowable values for OutputLocationOutputLocationTypeEnum
const (
	OutputLocationOutputLocationTypeObjectStoragePrefix OutputLocationOutputLocationTypeEnum = "OBJECT_STORAGE_PREFIX"
)

var mappingOutputLocationOutputLocationTypeEnum = map[string]OutputLocationOutputLocationTypeEnum{
	"OBJECT_STORAGE_PREFIX": OutputLocationOutputLocationTypeObjectStoragePrefix,
}

var mappingOutputLocationOutputLocationTypeEnumLowerCase = map[string]OutputLocationOutputLocationTypeEnum{
	"object_storage_prefix": OutputLocationOutputLocationTypeObjectStoragePrefix,
}

// GetOutputLocationOutputLocationTypeEnumValues Enumerates the set of values for OutputLocationOutputLocationTypeEnum
func GetOutputLocationOutputLocationTypeEnumValues() []OutputLocationOutputLocationTypeEnum {
	values := make([]OutputLocationOutputLocationTypeEnum, 0)
	for _, v := range mappingOutputLocationOutputLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOutputLocationOutputLocationTypeEnumStringValues Enumerates the set of values in String for OutputLocationOutputLocationTypeEnum
func GetOutputLocationOutputLocationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_PREFIX",
	}
}

// GetMappingOutputLocationOutputLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOutputLocationOutputLocationTypeEnum(val string) (OutputLocationOutputLocationTypeEnum, bool) {
	enum, ok := mappingOutputLocationOutputLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
