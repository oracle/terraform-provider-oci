// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ComputeInstanceGroupSelector Defines how the instances in a instance group environment is selected.
type ComputeInstanceGroupSelector interface {
}

type computeinstancegroupselector struct {
	JsonData     []byte
	SelectorType string `json:"selectorType"`
}

// UnmarshalJSON unmarshals json
func (m *computeinstancegroupselector) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercomputeinstancegroupselector computeinstancegroupselector
	s := struct {
		Model Unmarshalercomputeinstancegroupselector
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SelectorType = s.Model.SelectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *computeinstancegroupselector) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SelectorType {
	case "INSTANCE_IDS":
		mm := ComputeInstanceGroupByIdsSelector{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INSTANCE_QUERY":
		mm := ComputeInstanceGroupByQuerySelector{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m computeinstancegroupselector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m computeinstancegroupselector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeInstanceGroupSelectorSelectorTypeEnum Enum with underlying type: string
type ComputeInstanceGroupSelectorSelectorTypeEnum string

// Set of constants representing the allowable values for ComputeInstanceGroupSelectorSelectorTypeEnum
const (
	ComputeInstanceGroupSelectorSelectorTypeIds   ComputeInstanceGroupSelectorSelectorTypeEnum = "INSTANCE_IDS"
	ComputeInstanceGroupSelectorSelectorTypeQuery ComputeInstanceGroupSelectorSelectorTypeEnum = "INSTANCE_QUERY"
)

var mappingComputeInstanceGroupSelectorSelectorTypeEnum = map[string]ComputeInstanceGroupSelectorSelectorTypeEnum{
	"INSTANCE_IDS":   ComputeInstanceGroupSelectorSelectorTypeIds,
	"INSTANCE_QUERY": ComputeInstanceGroupSelectorSelectorTypeQuery,
}

// GetComputeInstanceGroupSelectorSelectorTypeEnumValues Enumerates the set of values for ComputeInstanceGroupSelectorSelectorTypeEnum
func GetComputeInstanceGroupSelectorSelectorTypeEnumValues() []ComputeInstanceGroupSelectorSelectorTypeEnum {
	values := make([]ComputeInstanceGroupSelectorSelectorTypeEnum, 0)
	for _, v := range mappingComputeInstanceGroupSelectorSelectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeInstanceGroupSelectorSelectorTypeEnumStringValues Enumerates the set of values in String for ComputeInstanceGroupSelectorSelectorTypeEnum
func GetComputeInstanceGroupSelectorSelectorTypeEnumStringValues() []string {
	return []string{
		"INSTANCE_IDS",
		"INSTANCE_QUERY",
	}
}

// GetMappingComputeInstanceGroupSelectorSelectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeInstanceGroupSelectorSelectorTypeEnum(val string) (ComputeInstanceGroupSelectorSelectorTypeEnum, bool) {
	mappingComputeInstanceGroupSelectorSelectorTypeEnumIgnoreCase := make(map[string]ComputeInstanceGroupSelectorSelectorTypeEnum)
	for k, v := range mappingComputeInstanceGroupSelectorSelectorTypeEnum {
		mappingComputeInstanceGroupSelectorSelectorTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingComputeInstanceGroupSelectorSelectorTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
