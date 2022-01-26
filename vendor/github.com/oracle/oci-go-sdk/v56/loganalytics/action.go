// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Action Action for scheduled task.
type Action interface {
}

type action struct {
	JsonData []byte
	Type     string `json:"type"`
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
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *action) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "PURGE":
		mm := PurgeAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STREAM":
		mm := StreamAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m action) String() string {
	return common.PointerString(m)
}

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeStream ActionTypeEnum = "STREAM"
	ActionTypePurge  ActionTypeEnum = "PURGE"
)

var mappingActionType = map[string]ActionTypeEnum{
	"STREAM": ActionTypeStream,
	"PURGE":  ActionTypePurge,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionType {
		values = append(values, v)
	}
	return values
}
