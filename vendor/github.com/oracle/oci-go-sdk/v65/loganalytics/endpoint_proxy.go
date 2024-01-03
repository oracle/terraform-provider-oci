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

// EndpointProxy An object containing the endpoint proxy details.
type EndpointProxy struct {

	// The proxy URL.
	Url *string `mandatory:"true" json:"url"`

	// The named credential name on the management agent, containing the proxy credentials.
	CredentialName *string `mandatory:"false" json:"credentialName"`

	// The credential type. NONE indicates credentials are not needed to access the proxy.
	// BASIC_AUTH represents a username and password based model. TOKEN represents a token based model.
	CredentialType EndpointProxyCredentialTypeEnum `mandatory:"false" json:"credentialType,omitempty"`
}

func (m EndpointProxy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EndpointProxy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEndpointProxyCredentialTypeEnum(string(m.CredentialType)); !ok && m.CredentialType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CredentialType: %s. Supported values are: %s.", m.CredentialType, strings.Join(GetEndpointProxyCredentialTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EndpointProxyCredentialTypeEnum Enum with underlying type: string
type EndpointProxyCredentialTypeEnum string

// Set of constants representing the allowable values for EndpointProxyCredentialTypeEnum
const (
	EndpointProxyCredentialTypeNone      EndpointProxyCredentialTypeEnum = "NONE"
	EndpointProxyCredentialTypeBasicAuth EndpointProxyCredentialTypeEnum = "BASIC_AUTH"
	EndpointProxyCredentialTypeToken     EndpointProxyCredentialTypeEnum = "TOKEN"
)

var mappingEndpointProxyCredentialTypeEnum = map[string]EndpointProxyCredentialTypeEnum{
	"NONE":       EndpointProxyCredentialTypeNone,
	"BASIC_AUTH": EndpointProxyCredentialTypeBasicAuth,
	"TOKEN":      EndpointProxyCredentialTypeToken,
}

var mappingEndpointProxyCredentialTypeEnumLowerCase = map[string]EndpointProxyCredentialTypeEnum{
	"none":       EndpointProxyCredentialTypeNone,
	"basic_auth": EndpointProxyCredentialTypeBasicAuth,
	"token":      EndpointProxyCredentialTypeToken,
}

// GetEndpointProxyCredentialTypeEnumValues Enumerates the set of values for EndpointProxyCredentialTypeEnum
func GetEndpointProxyCredentialTypeEnumValues() []EndpointProxyCredentialTypeEnum {
	values := make([]EndpointProxyCredentialTypeEnum, 0)
	for _, v := range mappingEndpointProxyCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointProxyCredentialTypeEnumStringValues Enumerates the set of values in String for EndpointProxyCredentialTypeEnum
func GetEndpointProxyCredentialTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"BASIC_AUTH",
		"TOKEN",
	}
}

// GetMappingEndpointProxyCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointProxyCredentialTypeEnum(val string) (EndpointProxyCredentialTypeEnum, bool) {
	enum, ok := mappingEndpointProxyCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
