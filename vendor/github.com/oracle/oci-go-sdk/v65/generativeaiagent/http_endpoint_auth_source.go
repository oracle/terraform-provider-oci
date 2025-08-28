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

// HttpEndpointAuthSource A credential source and configuration for a specific scope to HTTP Endpoint tools.
type HttpEndpointAuthSource struct {

	// Specifies the level from which credentials should be resolved.
	HttpEndpointAuthScope HttpEndpointAuthSourceHttpEndpointAuthScopeEnum `mandatory:"true" json:"httpEndpointAuthScope"`

	HttpEndpointAuthScopeConfig HttpEndpointAuthScopeConfig `mandatory:"true" json:"httpEndpointAuthScopeConfig"`
}

func (m HttpEndpointAuthSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpEndpointAuthSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHttpEndpointAuthSourceHttpEndpointAuthScopeEnum(string(m.HttpEndpointAuthScope)); !ok && m.HttpEndpointAuthScope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HttpEndpointAuthScope: %s. Supported values are: %s.", m.HttpEndpointAuthScope, strings.Join(GetHttpEndpointAuthSourceHttpEndpointAuthScopeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *HttpEndpointAuthSource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		HttpEndpointAuthScope       HttpEndpointAuthSourceHttpEndpointAuthScopeEnum `json:"httpEndpointAuthScope"`
		HttpEndpointAuthScopeConfig httpendpointauthscopeconfig                     `json:"httpEndpointAuthScopeConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.HttpEndpointAuthScope = model.HttpEndpointAuthScope

	nn, e = model.HttpEndpointAuthScopeConfig.UnmarshalPolymorphicJSON(model.HttpEndpointAuthScopeConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.HttpEndpointAuthScopeConfig = nn.(HttpEndpointAuthScopeConfig)
	} else {
		m.HttpEndpointAuthScopeConfig = nil
	}

	return
}

// HttpEndpointAuthSourceHttpEndpointAuthScopeEnum Enum with underlying type: string
type HttpEndpointAuthSourceHttpEndpointAuthScopeEnum string

// Set of constants representing the allowable values for HttpEndpointAuthSourceHttpEndpointAuthScopeEnum
const (
	HttpEndpointAuthSourceHttpEndpointAuthScopeAgent HttpEndpointAuthSourceHttpEndpointAuthScopeEnum = "AGENT"
)

var mappingHttpEndpointAuthSourceHttpEndpointAuthScopeEnum = map[string]HttpEndpointAuthSourceHttpEndpointAuthScopeEnum{
	"AGENT": HttpEndpointAuthSourceHttpEndpointAuthScopeAgent,
}

var mappingHttpEndpointAuthSourceHttpEndpointAuthScopeEnumLowerCase = map[string]HttpEndpointAuthSourceHttpEndpointAuthScopeEnum{
	"agent": HttpEndpointAuthSourceHttpEndpointAuthScopeAgent,
}

// GetHttpEndpointAuthSourceHttpEndpointAuthScopeEnumValues Enumerates the set of values for HttpEndpointAuthSourceHttpEndpointAuthScopeEnum
func GetHttpEndpointAuthSourceHttpEndpointAuthScopeEnumValues() []HttpEndpointAuthSourceHttpEndpointAuthScopeEnum {
	values := make([]HttpEndpointAuthSourceHttpEndpointAuthScopeEnum, 0)
	for _, v := range mappingHttpEndpointAuthSourceHttpEndpointAuthScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpEndpointAuthSourceHttpEndpointAuthScopeEnumStringValues Enumerates the set of values in String for HttpEndpointAuthSourceHttpEndpointAuthScopeEnum
func GetHttpEndpointAuthSourceHttpEndpointAuthScopeEnumStringValues() []string {
	return []string{
		"AGENT",
	}
}

// GetMappingHttpEndpointAuthSourceHttpEndpointAuthScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpEndpointAuthSourceHttpEndpointAuthScopeEnum(val string) (HttpEndpointAuthSourceHttpEndpointAuthScopeEnum, bool) {
	enum, ok := mappingHttpEndpointAuthSourceHttpEndpointAuthScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
