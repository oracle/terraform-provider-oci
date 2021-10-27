// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// Action An entity that represents an action to apply for a routing rule.
type Action interface {
}

type action struct {
	JsonData []byte
	Name     string `json:"name"`
}

// UnmarshalJSON unmarshals json
func (m *action) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraction action
	s := struct {
		Model Unmarshaleraction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *action) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Name {
	case "FORWARD_TO_BACKENDSET":
		mm := ForwardToBackendSet{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m action) String() string {
	return common.PointerString(m)
}

// ActionNameEnum Enum with underlying type: string
type ActionNameEnum string

// Set of constants representing the allowable values for ActionNameEnum
const (
	ActionNameForwardToBackendset ActionNameEnum = "FORWARD_TO_BACKENDSET"
)

var mappingActionName = map[string]ActionNameEnum{
	"FORWARD_TO_BACKENDSET": ActionNameForwardToBackendset,
}

// GetActionNameEnumValues Enumerates the set of values for ActionNameEnum
func GetActionNameEnumValues() []ActionNameEnum {
	values := make([]ActionNameEnum, 0)
	for _, v := range mappingActionName {
		values = append(values, v)
	}
	return values
}
