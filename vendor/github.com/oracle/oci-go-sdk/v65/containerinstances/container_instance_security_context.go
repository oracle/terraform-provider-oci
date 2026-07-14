// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerInstanceSecurityContext Security context for all containers in a container instance.
type ContainerInstanceSecurityContext interface {
}

type containerinstancesecuritycontext struct {
	JsonData            []byte
	SecurityContextType string `json:"securityContextType"`
}

// UnmarshalJSON unmarshals json
func (m *containerinstancesecuritycontext) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercontainerinstancesecuritycontext containerinstancesecuritycontext
	s := struct {
		Model Unmarshalercontainerinstancesecuritycontext
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SecurityContextType = s.Model.SecurityContextType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *containerinstancesecuritycontext) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SecurityContextType {
	case "LINUX":
		mm := LinuxContainerInstanceSecurityContext{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ContainerInstanceSecurityContext: %s.", m.SecurityContextType)
		return *m, nil
	}
}

func (m containerinstancesecuritycontext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m containerinstancesecuritycontext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContainerInstanceSecurityContextSecurityContextTypeEnum Enum with underlying type: string
type ContainerInstanceSecurityContextSecurityContextTypeEnum string

// Set of constants representing the allowable values for ContainerInstanceSecurityContextSecurityContextTypeEnum
const (
	ContainerInstanceSecurityContextSecurityContextTypeLinux ContainerInstanceSecurityContextSecurityContextTypeEnum = "LINUX"
)

var mappingContainerInstanceSecurityContextSecurityContextTypeEnum = map[string]ContainerInstanceSecurityContextSecurityContextTypeEnum{
	"LINUX": ContainerInstanceSecurityContextSecurityContextTypeLinux,
}

var mappingContainerInstanceSecurityContextSecurityContextTypeEnumLowerCase = map[string]ContainerInstanceSecurityContextSecurityContextTypeEnum{
	"linux": ContainerInstanceSecurityContextSecurityContextTypeLinux,
}

// GetContainerInstanceSecurityContextSecurityContextTypeEnumValues Enumerates the set of values for ContainerInstanceSecurityContextSecurityContextTypeEnum
func GetContainerInstanceSecurityContextSecurityContextTypeEnumValues() []ContainerInstanceSecurityContextSecurityContextTypeEnum {
	values := make([]ContainerInstanceSecurityContextSecurityContextTypeEnum, 0)
	for _, v := range mappingContainerInstanceSecurityContextSecurityContextTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerInstanceSecurityContextSecurityContextTypeEnumStringValues Enumerates the set of values in String for ContainerInstanceSecurityContextSecurityContextTypeEnum
func GetContainerInstanceSecurityContextSecurityContextTypeEnumStringValues() []string {
	return []string{
		"LINUX",
	}
}

// GetMappingContainerInstanceSecurityContextSecurityContextTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerInstanceSecurityContextSecurityContextTypeEnum(val string) (ContainerInstanceSecurityContextSecurityContextTypeEnum, bool) {
	enum, ok := mappingContainerInstanceSecurityContextSecurityContextTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
