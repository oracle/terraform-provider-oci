// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerConfig Specifies the container configuration.
type ContainerConfig interface {
}

type containerconfig struct {
	JsonData            []byte
	ContainerConfigType string `json:"containerConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *containerconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercontainerconfig containerconfig
	s := struct {
		Model Unmarshalercontainerconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ContainerConfigType = s.Model.ContainerConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *containerconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ContainerConfigType {
	case "CONTAINER_INSTANCE_CONFIG":
		mm := ContainerInstanceConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ContainerConfig: %s.", m.ContainerConfigType)
		return *m, nil
	}
}

func (m containerconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m containerconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContainerConfigContainerConfigTypeEnum Enum with underlying type: string
type ContainerConfigContainerConfigTypeEnum string

// Set of constants representing the allowable values for ContainerConfigContainerConfigTypeEnum
const (
	ContainerConfigContainerConfigTypeContainerInstanceConfig ContainerConfigContainerConfigTypeEnum = "CONTAINER_INSTANCE_CONFIG"
)

var mappingContainerConfigContainerConfigTypeEnum = map[string]ContainerConfigContainerConfigTypeEnum{
	"CONTAINER_INSTANCE_CONFIG": ContainerConfigContainerConfigTypeContainerInstanceConfig,
}

var mappingContainerConfigContainerConfigTypeEnumLowerCase = map[string]ContainerConfigContainerConfigTypeEnum{
	"container_instance_config": ContainerConfigContainerConfigTypeContainerInstanceConfig,
}

// GetContainerConfigContainerConfigTypeEnumValues Enumerates the set of values for ContainerConfigContainerConfigTypeEnum
func GetContainerConfigContainerConfigTypeEnumValues() []ContainerConfigContainerConfigTypeEnum {
	values := make([]ContainerConfigContainerConfigTypeEnum, 0)
	for _, v := range mappingContainerConfigContainerConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerConfigContainerConfigTypeEnumStringValues Enumerates the set of values in String for ContainerConfigContainerConfigTypeEnum
func GetContainerConfigContainerConfigTypeEnumStringValues() []string {
	return []string{
		"CONTAINER_INSTANCE_CONFIG",
	}
}

// GetMappingContainerConfigContainerConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerConfigContainerConfigTypeEnum(val string) (ContainerConfigContainerConfigTypeEnum, bool) {
	enum, ok := mappingContainerConfigContainerConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
