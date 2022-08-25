// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, agent configurations, log data models,
// continuous queries, and managed continuous queries.
// For more information, see https://docs.oracle.com/en-us/iaas/Content/Logging/Concepts/loggingoverview.htm.
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogRuleStartPolicy Log rule start policy object
type LogRuleStartPolicy interface {
}

type logrulestartpolicy struct {
	JsonData        []byte
	StartPolicyType string `json:"startPolicyType"`
}

// UnmarshalJSON unmarshals json
func (m *logrulestartpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlogrulestartpolicy logrulestartpolicy
	s := struct {
		Model Unmarshalerlogrulestartpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StartPolicyType = s.Model.StartPolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *logrulestartpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StartPolicyType {
	case "NO_DELAY_START_POLICY":
		mm := LogRuleNoDelayStartPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ABSOLUTE_TIME_START_POLICY":
		mm := LogRuleAbsoluteTimeStartPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for LogRuleStartPolicy: %s.", m.StartPolicyType)
		return *m, nil
	}
}

func (m logrulestartpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m logrulestartpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogRuleStartPolicyStartPolicyTypeEnum Enum with underlying type: string
type LogRuleStartPolicyStartPolicyTypeEnum string

// Set of constants representing the allowable values for LogRuleStartPolicyStartPolicyTypeEnum
const (
	LogRuleStartPolicyStartPolicyTypeNoDelayStartPolicy      LogRuleStartPolicyStartPolicyTypeEnum = "NO_DELAY_START_POLICY"
	LogRuleStartPolicyStartPolicyTypeAbsoluteTimeStartPolicy LogRuleStartPolicyStartPolicyTypeEnum = "ABSOLUTE_TIME_START_POLICY"
)

var mappingLogRuleStartPolicyStartPolicyTypeEnum = map[string]LogRuleStartPolicyStartPolicyTypeEnum{
	"NO_DELAY_START_POLICY":      LogRuleStartPolicyStartPolicyTypeNoDelayStartPolicy,
	"ABSOLUTE_TIME_START_POLICY": LogRuleStartPolicyStartPolicyTypeAbsoluteTimeStartPolicy,
}

var mappingLogRuleStartPolicyStartPolicyTypeEnumLowerCase = map[string]LogRuleStartPolicyStartPolicyTypeEnum{
	"no_delay_start_policy":      LogRuleStartPolicyStartPolicyTypeNoDelayStartPolicy,
	"absolute_time_start_policy": LogRuleStartPolicyStartPolicyTypeAbsoluteTimeStartPolicy,
}

// GetLogRuleStartPolicyStartPolicyTypeEnumValues Enumerates the set of values for LogRuleStartPolicyStartPolicyTypeEnum
func GetLogRuleStartPolicyStartPolicyTypeEnumValues() []LogRuleStartPolicyStartPolicyTypeEnum {
	values := make([]LogRuleStartPolicyStartPolicyTypeEnum, 0)
	for _, v := range mappingLogRuleStartPolicyStartPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogRuleStartPolicyStartPolicyTypeEnumStringValues Enumerates the set of values in String for LogRuleStartPolicyStartPolicyTypeEnum
func GetLogRuleStartPolicyStartPolicyTypeEnumStringValues() []string {
	return []string{
		"NO_DELAY_START_POLICY",
		"ABSOLUTE_TIME_START_POLICY",
	}
}

// GetMappingLogRuleStartPolicyStartPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogRuleStartPolicyStartPolicyTypeEnum(val string) (LogRuleStartPolicyStartPolicyTypeEnum, bool) {
	enum, ok := mappingLogRuleStartPolicyStartPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
