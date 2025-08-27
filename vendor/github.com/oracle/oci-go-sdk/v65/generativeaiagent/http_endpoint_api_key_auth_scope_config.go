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

// HttpEndpointApiKeyAuthScopeConfig Specifies authentication using an API key injected either as a header or query parameter.
// - If `authScope = AGENT`: The API key is retrieved from OCI Vault using the agent’s identity.
type HttpEndpointApiKeyAuthScopeConfig struct {

	// The name of the key parameter in the location.
	KeyName *string `mandatory:"true" json:"keyName"`

	// The OCID of the vault secret with API key.
	// Required when `authScope` is AGENT.
	VaultSecretId *string `mandatory:"false" json:"vaultSecretId"`

	// The location of the API key in the request.
	KeyLocation HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum `mandatory:"true" json:"keyLocation"`
}

func (m HttpEndpointApiKeyAuthScopeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpEndpointApiKeyAuthScopeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHttpEndpointApiKeyAuthScopeConfigKeyLocationEnum(string(m.KeyLocation)); !ok && m.KeyLocation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyLocation: %s. Supported values are: %s.", m.KeyLocation, strings.Join(GetHttpEndpointApiKeyAuthScopeConfigKeyLocationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpEndpointApiKeyAuthScopeConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpEndpointApiKeyAuthScopeConfig HttpEndpointApiKeyAuthScopeConfig
	s := struct {
		DiscriminatorParam string `json:"httpEndpointAuthScopeConfigType"`
		MarshalTypeHttpEndpointApiKeyAuthScopeConfig
	}{
		"HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG",
		(MarshalTypeHttpEndpointApiKeyAuthScopeConfig)(m),
	}

	return json.Marshal(&s)
}

// HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum Enum with underlying type: string
type HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum string

// Set of constants representing the allowable values for HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum
const (
	HttpEndpointApiKeyAuthScopeConfigKeyLocationHeader         HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum = "HEADER"
	HttpEndpointApiKeyAuthScopeConfigKeyLocationQueryParameter HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum = "QUERY_PARAMETER"
)

var mappingHttpEndpointApiKeyAuthScopeConfigKeyLocationEnum = map[string]HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum{
	"HEADER":          HttpEndpointApiKeyAuthScopeConfigKeyLocationHeader,
	"QUERY_PARAMETER": HttpEndpointApiKeyAuthScopeConfigKeyLocationQueryParameter,
}

var mappingHttpEndpointApiKeyAuthScopeConfigKeyLocationEnumLowerCase = map[string]HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum{
	"header":          HttpEndpointApiKeyAuthScopeConfigKeyLocationHeader,
	"query_parameter": HttpEndpointApiKeyAuthScopeConfigKeyLocationQueryParameter,
}

// GetHttpEndpointApiKeyAuthScopeConfigKeyLocationEnumValues Enumerates the set of values for HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum
func GetHttpEndpointApiKeyAuthScopeConfigKeyLocationEnumValues() []HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum {
	values := make([]HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum, 0)
	for _, v := range mappingHttpEndpointApiKeyAuthScopeConfigKeyLocationEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpEndpointApiKeyAuthScopeConfigKeyLocationEnumStringValues Enumerates the set of values in String for HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum
func GetHttpEndpointApiKeyAuthScopeConfigKeyLocationEnumStringValues() []string {
	return []string{
		"HEADER",
		"QUERY_PARAMETER",
	}
}

// GetMappingHttpEndpointApiKeyAuthScopeConfigKeyLocationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpEndpointApiKeyAuthScopeConfigKeyLocationEnum(val string) (HttpEndpointApiKeyAuthScopeConfigKeyLocationEnum, bool) {
	enum, ok := mappingHttpEndpointApiKeyAuthScopeConfigKeyLocationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
