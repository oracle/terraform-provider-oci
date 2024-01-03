// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EndpointCredentials An object containing credential details to authenticate/authorize a REST request.
type EndpointCredentials struct {

	// The credential type. NONE indicates credentials are not needed to access the endpoint.
	// BASIC_AUTH represents a username and password based model. TOKEN could be static or dynamic.
	// In case of dynamic tokens, also specify the endpoint from which the token must be fetched.
	CredentialType EndpointCredentialsCredentialTypeEnum `mandatory:"false" json:"credentialType,omitempty"`

	// The named credential name on the management agent.
	CredentialName *string `mandatory:"false" json:"credentialName"`

	CredentialEndpoint *CredentialEndpoint `mandatory:"false" json:"credentialEndpoint"`
}

func (m EndpointCredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EndpointCredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEndpointCredentialsCredentialTypeEnum(string(m.CredentialType)); !ok && m.CredentialType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CredentialType: %s. Supported values are: %s.", m.CredentialType, strings.Join(GetEndpointCredentialsCredentialTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EndpointCredentialsCredentialTypeEnum Enum with underlying type: string
type EndpointCredentialsCredentialTypeEnum string

// Set of constants representing the allowable values for EndpointCredentialsCredentialTypeEnum
const (
	EndpointCredentialsCredentialTypeNone         EndpointCredentialsCredentialTypeEnum = "NONE"
	EndpointCredentialsCredentialTypeBasicAuth    EndpointCredentialsCredentialTypeEnum = "BASIC_AUTH"
	EndpointCredentialsCredentialTypeStaticToken  EndpointCredentialsCredentialTypeEnum = "STATIC_TOKEN"
	EndpointCredentialsCredentialTypeDynamicToken EndpointCredentialsCredentialTypeEnum = "DYNAMIC_TOKEN"
)

var mappingEndpointCredentialsCredentialTypeEnum = map[string]EndpointCredentialsCredentialTypeEnum{
	"NONE":          EndpointCredentialsCredentialTypeNone,
	"BASIC_AUTH":    EndpointCredentialsCredentialTypeBasicAuth,
	"STATIC_TOKEN":  EndpointCredentialsCredentialTypeStaticToken,
	"DYNAMIC_TOKEN": EndpointCredentialsCredentialTypeDynamicToken,
}

var mappingEndpointCredentialsCredentialTypeEnumLowerCase = map[string]EndpointCredentialsCredentialTypeEnum{
	"none":          EndpointCredentialsCredentialTypeNone,
	"basic_auth":    EndpointCredentialsCredentialTypeBasicAuth,
	"static_token":  EndpointCredentialsCredentialTypeStaticToken,
	"dynamic_token": EndpointCredentialsCredentialTypeDynamicToken,
}

// GetEndpointCredentialsCredentialTypeEnumValues Enumerates the set of values for EndpointCredentialsCredentialTypeEnum
func GetEndpointCredentialsCredentialTypeEnumValues() []EndpointCredentialsCredentialTypeEnum {
	values := make([]EndpointCredentialsCredentialTypeEnum, 0)
	for _, v := range mappingEndpointCredentialsCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointCredentialsCredentialTypeEnumStringValues Enumerates the set of values in String for EndpointCredentialsCredentialTypeEnum
func GetEndpointCredentialsCredentialTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"BASIC_AUTH",
		"STATIC_TOKEN",
		"DYNAMIC_TOKEN",
	}
}

// GetMappingEndpointCredentialsCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointCredentialsCredentialTypeEnum(val string) (EndpointCredentialsCredentialTypeEnum, bool) {
	enum, ok := mappingEndpointCredentialsCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
