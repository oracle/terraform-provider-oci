// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// SecretDetail The details of configured security configuration on OpenSearch.
type SecretDetail interface {
}

type secretdetail struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *secretdetail) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecretdetail secretdetail
	s := struct {
		Model Unmarshalersecretdetail
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *secretdetail) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "IDCS_SECRET":
		mm := IdcsSecret{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BASIC_AUTH_SECRET":
		mm := BasicAuthSecret{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SecretDetail: %s.", m.Type)
		return *m, nil
	}
}

func (m secretdetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m secretdetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecretDetailTypeEnum Enum with underlying type: string
type SecretDetailTypeEnum string

// Set of constants representing the allowable values for SecretDetailTypeEnum
const (
	SecretDetailTypeIdcsSecret      SecretDetailTypeEnum = "IDCS_SECRET"
	SecretDetailTypeBasicAuthSecret SecretDetailTypeEnum = "BASIC_AUTH_SECRET"
)

var mappingSecretDetailTypeEnum = map[string]SecretDetailTypeEnum{
	"IDCS_SECRET":       SecretDetailTypeIdcsSecret,
	"BASIC_AUTH_SECRET": SecretDetailTypeBasicAuthSecret,
}

var mappingSecretDetailTypeEnumLowerCase = map[string]SecretDetailTypeEnum{
	"idcs_secret":       SecretDetailTypeIdcsSecret,
	"basic_auth_secret": SecretDetailTypeBasicAuthSecret,
}

// GetSecretDetailTypeEnumValues Enumerates the set of values for SecretDetailTypeEnum
func GetSecretDetailTypeEnumValues() []SecretDetailTypeEnum {
	values := make([]SecretDetailTypeEnum, 0)
	for _, v := range mappingSecretDetailTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretDetailTypeEnumStringValues Enumerates the set of values in String for SecretDetailTypeEnum
func GetSecretDetailTypeEnumStringValues() []string {
	return []string{
		"IDCS_SECRET",
		"BASIC_AUTH_SECRET",
	}
}

// GetMappingSecretDetailTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretDetailTypeEnum(val string) (SecretDetailTypeEnum, bool) {
	enum, ok := mappingSecretDetailTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
