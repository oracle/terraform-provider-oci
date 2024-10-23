// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PauseDetails Pause Details
type PauseDetails interface {
}

type pausedetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *pausedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpausedetails pausedetails
	s := struct {
		Model Unmarshalerpausedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pausedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "USER_ACTION":
		mm := UserActionBasedPauseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_BASED":
		mm := TimeBasedPauseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PauseDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m pausedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pausedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PauseDetailsKindEnum Enum with underlying type: string
type PauseDetailsKindEnum string

// Set of constants representing the allowable values for PauseDetailsKindEnum
const (
	PauseDetailsKindUserAction PauseDetailsKindEnum = "USER_ACTION"
	PauseDetailsKindTimeBased  PauseDetailsKindEnum = "TIME_BASED"
)

var mappingPauseDetailsKindEnum = map[string]PauseDetailsKindEnum{
	"USER_ACTION": PauseDetailsKindUserAction,
	"TIME_BASED":  PauseDetailsKindTimeBased,
}

var mappingPauseDetailsKindEnumLowerCase = map[string]PauseDetailsKindEnum{
	"user_action": PauseDetailsKindUserAction,
	"time_based":  PauseDetailsKindTimeBased,
}

// GetPauseDetailsKindEnumValues Enumerates the set of values for PauseDetailsKindEnum
func GetPauseDetailsKindEnumValues() []PauseDetailsKindEnum {
	values := make([]PauseDetailsKindEnum, 0)
	for _, v := range mappingPauseDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetPauseDetailsKindEnumStringValues Enumerates the set of values in String for PauseDetailsKindEnum
func GetPauseDetailsKindEnumStringValues() []string {
	return []string{
		"USER_ACTION",
		"TIME_BASED",
	}
}

// GetMappingPauseDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPauseDetailsKindEnum(val string) (PauseDetailsKindEnum, bool) {
	enum, ok := mappingPauseDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
