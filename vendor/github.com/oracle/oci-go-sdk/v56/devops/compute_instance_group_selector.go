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
	"github.com/oracle/oci-go-sdk/v56/common"
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

// ComputeInstanceGroupSelectorSelectorTypeEnum Enum with underlying type: string
type ComputeInstanceGroupSelectorSelectorTypeEnum string

// Set of constants representing the allowable values for ComputeInstanceGroupSelectorSelectorTypeEnum
const (
	ComputeInstanceGroupSelectorSelectorTypeIds   ComputeInstanceGroupSelectorSelectorTypeEnum = "INSTANCE_IDS"
	ComputeInstanceGroupSelectorSelectorTypeQuery ComputeInstanceGroupSelectorSelectorTypeEnum = "INSTANCE_QUERY"
)

var mappingComputeInstanceGroupSelectorSelectorType = map[string]ComputeInstanceGroupSelectorSelectorTypeEnum{
	"INSTANCE_IDS":   ComputeInstanceGroupSelectorSelectorTypeIds,
	"INSTANCE_QUERY": ComputeInstanceGroupSelectorSelectorTypeQuery,
}

// GetComputeInstanceGroupSelectorSelectorTypeEnumValues Enumerates the set of values for ComputeInstanceGroupSelectorSelectorTypeEnum
func GetComputeInstanceGroupSelectorSelectorTypeEnumValues() []ComputeInstanceGroupSelectorSelectorTypeEnum {
	values := make([]ComputeInstanceGroupSelectorSelectorTypeEnum, 0)
	for _, v := range mappingComputeInstanceGroupSelectorSelectorType {
		values = append(values, v)
	}
	return values
}
