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

// FlowConfig Configuration for an agent set for node based workflow
type FlowConfig interface {
}

type flowconfig struct {
	JsonData       []byte
	FlowConfigType string `json:"flowConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *flowconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerflowconfig flowconfig
	s := struct {
		Model Unmarshalerflowconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FlowConfigType = s.Model.FlowConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *flowconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FlowConfigType {
	case "INLINE":
		mm := FlowConfigInline{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for FlowConfig: %s.", m.FlowConfigType)
		return *m, nil
	}
}

func (m flowconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m flowconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FlowConfigFlowConfigTypeEnum Enum with underlying type: string
type FlowConfigFlowConfigTypeEnum string

// Set of constants representing the allowable values for FlowConfigFlowConfigTypeEnum
const (
	FlowConfigFlowConfigTypeInline FlowConfigFlowConfigTypeEnum = "INLINE"
)

var mappingFlowConfigFlowConfigTypeEnum = map[string]FlowConfigFlowConfigTypeEnum{
	"INLINE": FlowConfigFlowConfigTypeInline,
}

var mappingFlowConfigFlowConfigTypeEnumLowerCase = map[string]FlowConfigFlowConfigTypeEnum{
	"inline": FlowConfigFlowConfigTypeInline,
}

// GetFlowConfigFlowConfigTypeEnumValues Enumerates the set of values for FlowConfigFlowConfigTypeEnum
func GetFlowConfigFlowConfigTypeEnumValues() []FlowConfigFlowConfigTypeEnum {
	values := make([]FlowConfigFlowConfigTypeEnum, 0)
	for _, v := range mappingFlowConfigFlowConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFlowConfigFlowConfigTypeEnumStringValues Enumerates the set of values in String for FlowConfigFlowConfigTypeEnum
func GetFlowConfigFlowConfigTypeEnumStringValues() []string {
	return []string{
		"INLINE",
	}
}

// GetMappingFlowConfigFlowConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlowConfigFlowConfigTypeEnum(val string) (FlowConfigFlowConfigTypeEnum, bool) {
	enum, ok := mappingFlowConfigFlowConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
