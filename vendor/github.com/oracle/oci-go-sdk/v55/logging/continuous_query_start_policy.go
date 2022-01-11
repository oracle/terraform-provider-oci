// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v55/common"
	"strings"
)

// ContinuousQueryStartPolicy Continuous query start policy object
type ContinuousQueryStartPolicy interface {
}

type continuousquerystartpolicy struct {
	JsonData        []byte
	StartPolicyType string `json:"startPolicyType"`
}

// UnmarshalJSON unmarshals json
func (m *continuousquerystartpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercontinuousquerystartpolicy continuousquerystartpolicy
	s := struct {
		Model Unmarshalercontinuousquerystartpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StartPolicyType = s.Model.StartPolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *continuousquerystartpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StartPolicyType {
	case "ABSOLUTE_TIME_START_POLICY":
		mm := AbsoluteTimeStartPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NO_DELAY_START_POLICY":
		mm := NoDelayStartPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m continuousquerystartpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m continuousquerystartpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContinuousQueryStartPolicyStartPolicyTypeEnum Enum with underlying type: string
type ContinuousQueryStartPolicyStartPolicyTypeEnum string

// Set of constants representing the allowable values for ContinuousQueryStartPolicyStartPolicyTypeEnum
const (
	ContinuousQueryStartPolicyStartPolicyTypeNoDelayStartPolicy      ContinuousQueryStartPolicyStartPolicyTypeEnum = "NO_DELAY_START_POLICY"
	ContinuousQueryStartPolicyStartPolicyTypeAbsoluteTimeStartPolicy ContinuousQueryStartPolicyStartPolicyTypeEnum = "ABSOLUTE_TIME_START_POLICY"
)

var mappingContinuousQueryStartPolicyStartPolicyTypeEnum = map[string]ContinuousQueryStartPolicyStartPolicyTypeEnum{
	"NO_DELAY_START_POLICY":      ContinuousQueryStartPolicyStartPolicyTypeNoDelayStartPolicy,
	"ABSOLUTE_TIME_START_POLICY": ContinuousQueryStartPolicyStartPolicyTypeAbsoluteTimeStartPolicy,
}

// GetContinuousQueryStartPolicyStartPolicyTypeEnumValues Enumerates the set of values for ContinuousQueryStartPolicyStartPolicyTypeEnum
func GetContinuousQueryStartPolicyStartPolicyTypeEnumValues() []ContinuousQueryStartPolicyStartPolicyTypeEnum {
	values := make([]ContinuousQueryStartPolicyStartPolicyTypeEnum, 0)
	for _, v := range mappingContinuousQueryStartPolicyStartPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContinuousQueryStartPolicyStartPolicyTypeEnumStringValues Enumerates the set of values in String for ContinuousQueryStartPolicyStartPolicyTypeEnum
func GetContinuousQueryStartPolicyStartPolicyTypeEnumStringValues() []string {
	return []string{
		"NO_DELAY_START_POLICY",
		"ABSOLUTE_TIME_START_POLICY",
	}
}
