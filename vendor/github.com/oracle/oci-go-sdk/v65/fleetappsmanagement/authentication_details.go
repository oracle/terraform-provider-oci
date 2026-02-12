// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuthenticationDetails Authentication details for api request.
type AuthenticationDetails interface {
}

type authenticationdetails struct {
	JsonData           []byte
	AuthenticationType string `json:"authenticationType"`
}

// UnmarshalJSON unmarshals json
func (m *authenticationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerauthenticationdetails authenticationdetails
	s := struct {
		Model Unmarshalerauthenticationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AuthenticationType = s.Model.AuthenticationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *authenticationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AuthenticationType {
	case "RESOURCE_PRINCIPAL":
		mm := ResourcePrincipalAuthenticationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBO_TOKEN":
		mm := OboTokenAuthenticationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AuthenticationDetails: %s.", m.AuthenticationType)
		return *m, nil
	}
}

func (m authenticationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m authenticationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthenticationDetailsAuthenticationTypeEnum Enum with underlying type: string
type AuthenticationDetailsAuthenticationTypeEnum string

// Set of constants representing the allowable values for AuthenticationDetailsAuthenticationTypeEnum
const (
	AuthenticationDetailsAuthenticationTypeOboToken          AuthenticationDetailsAuthenticationTypeEnum = "OBO_TOKEN"
	AuthenticationDetailsAuthenticationTypeResourcePrincipal AuthenticationDetailsAuthenticationTypeEnum = "RESOURCE_PRINCIPAL"
)

var mappingAuthenticationDetailsAuthenticationTypeEnum = map[string]AuthenticationDetailsAuthenticationTypeEnum{
	"OBO_TOKEN":          AuthenticationDetailsAuthenticationTypeOboToken,
	"RESOURCE_PRINCIPAL": AuthenticationDetailsAuthenticationTypeResourcePrincipal,
}

var mappingAuthenticationDetailsAuthenticationTypeEnumLowerCase = map[string]AuthenticationDetailsAuthenticationTypeEnum{
	"obo_token":          AuthenticationDetailsAuthenticationTypeOboToken,
	"resource_principal": AuthenticationDetailsAuthenticationTypeResourcePrincipal,
}

// GetAuthenticationDetailsAuthenticationTypeEnumValues Enumerates the set of values for AuthenticationDetailsAuthenticationTypeEnum
func GetAuthenticationDetailsAuthenticationTypeEnumValues() []AuthenticationDetailsAuthenticationTypeEnum {
	values := make([]AuthenticationDetailsAuthenticationTypeEnum, 0)
	for _, v := range mappingAuthenticationDetailsAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationDetailsAuthenticationTypeEnumStringValues Enumerates the set of values in String for AuthenticationDetailsAuthenticationTypeEnum
func GetAuthenticationDetailsAuthenticationTypeEnumStringValues() []string {
	return []string{
		"OBO_TOKEN",
		"RESOURCE_PRINCIPAL",
	}
}

// GetMappingAuthenticationDetailsAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationDetailsAuthenticationTypeEnum(val string) (AuthenticationDetailsAuthenticationTypeEnum, bool) {
	enum, ok := mappingAuthenticationDetailsAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
