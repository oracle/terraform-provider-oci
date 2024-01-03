// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityContext Security context for container.
type SecurityContext interface {
}

type securitycontext struct {
	JsonData            []byte
	SecurityContextType string `json:"securityContextType"`
}

// UnmarshalJSON unmarshals json
func (m *securitycontext) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecuritycontext securitycontext
	s := struct {
		Model Unmarshalersecuritycontext
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SecurityContextType = s.Model.SecurityContextType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *securitycontext) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SecurityContextType {
	case "LINUX":
		mm := LinuxSecurityContext{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SecurityContext: %s.", m.SecurityContextType)
		return *m, nil
	}
}

func (m securitycontext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m securitycontext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityContextSecurityContextTypeEnum Enum with underlying type: string
type SecurityContextSecurityContextTypeEnum string

// Set of constants representing the allowable values for SecurityContextSecurityContextTypeEnum
const (
	SecurityContextSecurityContextTypeLinux SecurityContextSecurityContextTypeEnum = "LINUX"
)

var mappingSecurityContextSecurityContextTypeEnum = map[string]SecurityContextSecurityContextTypeEnum{
	"LINUX": SecurityContextSecurityContextTypeLinux,
}

var mappingSecurityContextSecurityContextTypeEnumLowerCase = map[string]SecurityContextSecurityContextTypeEnum{
	"linux": SecurityContextSecurityContextTypeLinux,
}

// GetSecurityContextSecurityContextTypeEnumValues Enumerates the set of values for SecurityContextSecurityContextTypeEnum
func GetSecurityContextSecurityContextTypeEnumValues() []SecurityContextSecurityContextTypeEnum {
	values := make([]SecurityContextSecurityContextTypeEnum, 0)
	for _, v := range mappingSecurityContextSecurityContextTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityContextSecurityContextTypeEnumStringValues Enumerates the set of values in String for SecurityContextSecurityContextTypeEnum
func GetSecurityContextSecurityContextTypeEnumStringValues() []string {
	return []string{
		"LINUX",
	}
}

// GetMappingSecurityContextSecurityContextTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityContextSecurityContextTypeEnum(val string) (SecurityContextSecurityContextTypeEnum, bool) {
	enum, ok := mappingSecurityContextSecurityContextTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
