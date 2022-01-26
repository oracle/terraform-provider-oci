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

// UpdateWaitCriteriaDetails Specifies wait criteria for the Wait stage.
type UpdateWaitCriteriaDetails interface {
}

type updatewaitcriteriadetails struct {
	JsonData []byte
	WaitType string `json:"waitType"`
}

// UnmarshalJSON unmarshals json
func (m *updatewaitcriteriadetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatewaitcriteriadetails updatewaitcriteriadetails
	s := struct {
		Model Unmarshalerupdatewaitcriteriadetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WaitType = s.Model.WaitType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatewaitcriteriadetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.WaitType {
	case "ABSOLUTE_WAIT":
		mm := UpdateAbsoluteWaitCriteriaDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m updatewaitcriteriadetails) String() string {
	return common.PointerString(m)
}

// UpdateWaitCriteriaDetailsWaitTypeEnum Enum with underlying type: string
type UpdateWaitCriteriaDetailsWaitTypeEnum string

// Set of constants representing the allowable values for UpdateWaitCriteriaDetailsWaitTypeEnum
const (
	UpdateWaitCriteriaDetailsWaitTypeAbsoluteWait UpdateWaitCriteriaDetailsWaitTypeEnum = "ABSOLUTE_WAIT"
)

var mappingUpdateWaitCriteriaDetailsWaitType = map[string]UpdateWaitCriteriaDetailsWaitTypeEnum{
	"ABSOLUTE_WAIT": UpdateWaitCriteriaDetailsWaitTypeAbsoluteWait,
}

// GetUpdateWaitCriteriaDetailsWaitTypeEnumValues Enumerates the set of values for UpdateWaitCriteriaDetailsWaitTypeEnum
func GetUpdateWaitCriteriaDetailsWaitTypeEnumValues() []UpdateWaitCriteriaDetailsWaitTypeEnum {
	values := make([]UpdateWaitCriteriaDetailsWaitTypeEnum, 0)
	for _, v := range mappingUpdateWaitCriteriaDetailsWaitType {
		values = append(values, v)
	}
	return values
}
