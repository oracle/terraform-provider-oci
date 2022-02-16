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

// CreateWaitCriteriaDetails Specifies wait criteria for the Wait stage.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createwaitcriteriadetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateWaitCriteriaDetailsWaitTypeEnum Enum with underlying type: string
type CreateWaitCriteriaDetailsWaitTypeEnum string

// Set of constants representing the allowable values for CreateWaitCriteriaDetailsWaitTypeEnum
const (
	CreateWaitCriteriaDetailsWaitTypeAbsoluteWait CreateWaitCriteriaDetailsWaitTypeEnum = "ABSOLUTE_WAIT"
)

var mappingCreateWaitCriteriaDetailsWaitTypeEnum = map[string]CreateWaitCriteriaDetailsWaitTypeEnum{
	"ABSOLUTE_WAIT": CreateWaitCriteriaDetailsWaitTypeAbsoluteWait,
}

// GetCreateWaitCriteriaDetailsWaitTypeEnumValues Enumerates the set of values for CreateWaitCriteriaDetailsWaitTypeEnum
func GetCreateWaitCriteriaDetailsWaitTypeEnumValues() []CreateWaitCriteriaDetailsWaitTypeEnum {
	values := make([]CreateWaitCriteriaDetailsWaitTypeEnum, 0)
	for _, v := range mappingCreateWaitCriteriaDetailsWaitTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateWaitCriteriaDetailsWaitTypeEnumStringValues Enumerates the set of values in String for CreateWaitCriteriaDetailsWaitTypeEnum
func GetCreateWaitCriteriaDetailsWaitTypeEnumStringValues() []string {
	return []string{
		"ABSOLUTE_WAIT",
	}
}

// GetMappingCreateWaitCriteriaDetailsWaitTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateWaitCriteriaDetailsWaitTypeEnum(val string) (CreateWaitCriteriaDetailsWaitTypeEnum, bool) {
	mappingCreateWaitCriteriaDetailsWaitTypeEnumIgnoreCase := make(map[string]CreateWaitCriteriaDetailsWaitTypeEnum)
	for k, v := range mappingCreateWaitCriteriaDetailsWaitTypeEnum {
		mappingCreateWaitCriteriaDetailsWaitTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateWaitCriteriaDetailsWaitTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
