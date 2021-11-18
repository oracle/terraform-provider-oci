// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateWaitCriteriaDetails Specifies wait criteria for wait stage.
type CreateWaitCriteriaDetails interface {
}

type createwaitcriteriadetails struct {
	JsonData []byte
	WaitType string `json:"waitType"`
}

// UnmarshalJSON unmarshals json
func (m *createwaitcriteriadetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatewaitcriteriadetails createwaitcriteriadetails
	s := struct {
		Model Unmarshalercreatewaitcriteriadetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WaitType = s.Model.WaitType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createwaitcriteriadetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.WaitType {
	case "ABSOLUTE_WAIT":
		mm := CreateAbsoluteWaitCriteriaDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createwaitcriteriadetails) String() string {
	return common.PointerString(m)
}

// CreateWaitCriteriaDetailsWaitTypeEnum Enum with underlying type: string
type CreateWaitCriteriaDetailsWaitTypeEnum string

// Set of constants representing the allowable values for CreateWaitCriteriaDetailsWaitTypeEnum
const (
	CreateWaitCriteriaDetailsWaitTypeAbsoluteWait CreateWaitCriteriaDetailsWaitTypeEnum = "ABSOLUTE_WAIT"
)

var mappingCreateWaitCriteriaDetailsWaitType = map[string]CreateWaitCriteriaDetailsWaitTypeEnum{
	"ABSOLUTE_WAIT": CreateWaitCriteriaDetailsWaitTypeAbsoluteWait,
}

// GetCreateWaitCriteriaDetailsWaitTypeEnumValues Enumerates the set of values for CreateWaitCriteriaDetailsWaitTypeEnum
func GetCreateWaitCriteriaDetailsWaitTypeEnumValues() []CreateWaitCriteriaDetailsWaitTypeEnum {
	values := make([]CreateWaitCriteriaDetailsWaitTypeEnum, 0)
	for _, v := range mappingCreateWaitCriteriaDetailsWaitType {
		values = append(values, v)
	}
	return values
}
