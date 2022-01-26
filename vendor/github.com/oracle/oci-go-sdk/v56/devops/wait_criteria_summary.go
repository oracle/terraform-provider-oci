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

// WaitCriteriaSummary Specifies wait criteria for the Wait stage.
type WaitCriteriaSummary interface {
}

type waitcriteriasummary struct {
	JsonData []byte
	WaitType string `json:"waitType"`
}

// UnmarshalJSON unmarshals json
func (m *waitcriteriasummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerwaitcriteriasummary waitcriteriasummary
	s := struct {
		Model Unmarshalerwaitcriteriasummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WaitType = s.Model.WaitType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *waitcriteriasummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.WaitType {
	case "ABSOLUTE_WAIT":
		mm := AbsoluteWaitCriteriaSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m waitcriteriasummary) String() string {
	return common.PointerString(m)
}

// WaitCriteriaSummaryWaitTypeEnum Enum with underlying type: string
type WaitCriteriaSummaryWaitTypeEnum string

// Set of constants representing the allowable values for WaitCriteriaSummaryWaitTypeEnum
const (
	WaitCriteriaSummaryWaitTypeAbsoluteWait WaitCriteriaSummaryWaitTypeEnum = "ABSOLUTE_WAIT"
)

var mappingWaitCriteriaSummaryWaitType = map[string]WaitCriteriaSummaryWaitTypeEnum{
	"ABSOLUTE_WAIT": WaitCriteriaSummaryWaitTypeAbsoluteWait,
}

// GetWaitCriteriaSummaryWaitTypeEnumValues Enumerates the set of values for WaitCriteriaSummaryWaitTypeEnum
func GetWaitCriteriaSummaryWaitTypeEnumValues() []WaitCriteriaSummaryWaitTypeEnum {
	values := make([]WaitCriteriaSummaryWaitTypeEnum, 0)
	for _, v := range mappingWaitCriteriaSummaryWaitType {
		values = append(values, v)
	}
	return values
}
