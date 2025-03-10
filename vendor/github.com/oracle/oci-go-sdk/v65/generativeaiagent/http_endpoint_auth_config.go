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

// HttpEndpointAuthConfig Auth related information to be used when invoking external endpoint
type HttpEndpointAuthConfig interface {
}

type httpendpointauthconfig struct {
	JsonData                   []byte
	HttpEndpointAuthConfigType string `json:"httpEndpointAuthConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *httpendpointauthconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhttpendpointauthconfig httpendpointauthconfig
	s := struct {
		Model Unmarshalerhttpendpointauthconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.HttpEndpointAuthConfigType = s.Model.HttpEndpointAuthConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *httpendpointauthconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.HttpEndpointAuthConfigType {
	case "HTTP_ENDPOINT_IDCS_AUTH_CONFIG":
		mm := HttpEndpointIdcsAuthConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_DELEGATED_BEARER_AUTH_CONFIG":
		mm := HttpEndpointDelegatedBearerAuthConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_NO_AUTH_CONFIG":
		mm := HttpEndpointNoAuthConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_OCI_RESOURCE_PRINCIPAL_AUTH_CONFIG":
		mm := HttpEndpointOciResourcePrincipalAuthConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for HttpEndpointAuthConfig: %s.", m.HttpEndpointAuthConfigType)
		return *m, nil
	}
}

func (m httpendpointauthconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m httpendpointauthconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum Enum with underlying type: string
type HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum string

// Set of constants representing the allowable values for HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum
const (
	HttpEndpointAuthConfigHttpEndpointAuthConfigTypeNoAuthConfig                   HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum = "HTTP_ENDPOINT_NO_AUTH_CONFIG"
	HttpEndpointAuthConfigHttpEndpointAuthConfigTypeDelegatedBearerAuthConfig      HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum = "HTTP_ENDPOINT_DELEGATED_BEARER_AUTH_CONFIG"
	HttpEndpointAuthConfigHttpEndpointAuthConfigTypeOciResourcePrincipalAuthConfig HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum = "HTTP_ENDPOINT_OCI_RESOURCE_PRINCIPAL_AUTH_CONFIG"
	HttpEndpointAuthConfigHttpEndpointAuthConfigTypeIdcsAuthConfig                 HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum = "HTTP_ENDPOINT_IDCS_AUTH_CONFIG"
)

var mappingHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum = map[string]HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum{
	"HTTP_ENDPOINT_NO_AUTH_CONFIG":                     HttpEndpointAuthConfigHttpEndpointAuthConfigTypeNoAuthConfig,
	"HTTP_ENDPOINT_DELEGATED_BEARER_AUTH_CONFIG":       HttpEndpointAuthConfigHttpEndpointAuthConfigTypeDelegatedBearerAuthConfig,
	"HTTP_ENDPOINT_OCI_RESOURCE_PRINCIPAL_AUTH_CONFIG": HttpEndpointAuthConfigHttpEndpointAuthConfigTypeOciResourcePrincipalAuthConfig,
	"HTTP_ENDPOINT_IDCS_AUTH_CONFIG":                   HttpEndpointAuthConfigHttpEndpointAuthConfigTypeIdcsAuthConfig,
}

var mappingHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnumLowerCase = map[string]HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum{
	"http_endpoint_no_auth_config":                     HttpEndpointAuthConfigHttpEndpointAuthConfigTypeNoAuthConfig,
	"http_endpoint_delegated_bearer_auth_config":       HttpEndpointAuthConfigHttpEndpointAuthConfigTypeDelegatedBearerAuthConfig,
	"http_endpoint_oci_resource_principal_auth_config": HttpEndpointAuthConfigHttpEndpointAuthConfigTypeOciResourcePrincipalAuthConfig,
	"http_endpoint_idcs_auth_config":                   HttpEndpointAuthConfigHttpEndpointAuthConfigTypeIdcsAuthConfig,
}

// GetHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnumValues Enumerates the set of values for HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum
func GetHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnumValues() []HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum {
	values := make([]HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum, 0)
	for _, v := range mappingHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnumStringValues Enumerates the set of values in String for HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum
func GetHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnumStringValues() []string {
	return []string{
		"HTTP_ENDPOINT_NO_AUTH_CONFIG",
		"HTTP_ENDPOINT_DELEGATED_BEARER_AUTH_CONFIG",
		"HTTP_ENDPOINT_OCI_RESOURCE_PRINCIPAL_AUTH_CONFIG",
		"HTTP_ENDPOINT_IDCS_AUTH_CONFIG",
	}
}

// GetMappingHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum(val string) (HttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnum, bool) {
	enum, ok := mappingHttpEndpointAuthConfigHttpEndpointAuthConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
