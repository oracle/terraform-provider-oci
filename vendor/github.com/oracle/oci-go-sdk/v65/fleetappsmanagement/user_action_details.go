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

// UserActionDetails User action details.
// This can be performed on a failed/paused task or action group.
type UserActionDetails interface {

	// Action to be Performed.
	GetAction() UserActionDetailsActionEnum
}

type useractiondetails struct {
	JsonData []byte
	Action   UserActionDetailsActionEnum `mandatory:"true" json:"action"`
	Level    string                      `json:"level"`
}

// UnmarshalJSON unmarshals json
func (m *useractiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleruseractiondetails useractiondetails
	s := struct {
		Model Unmarshaleruseractiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Action = s.Model.Action
	m.Level = s.Model.Level

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *useractiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Level {
	case "STEP_NAME":
		mm := StepBasedUserActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ACTION_GROUP":
		mm := ActionGroupBasedUserActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UserActionDetails: %s.", m.Level)
		return *m, nil
	}
}

// GetAction returns Action
func (m useractiondetails) GetAction() UserActionDetailsActionEnum {
	return m.Action
}

func (m useractiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m useractiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserActionDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetUserActionDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserActionDetailsActionEnum Enum with underlying type: string
type UserActionDetailsActionEnum string

// Set of constants representing the allowable values for UserActionDetailsActionEnum
const (
	UserActionDetailsActionRetry  UserActionDetailsActionEnum = "RETRY"
	UserActionDetailsActionResume UserActionDetailsActionEnum = "RESUME"
)

var mappingUserActionDetailsActionEnum = map[string]UserActionDetailsActionEnum{
	"RETRY":  UserActionDetailsActionRetry,
	"RESUME": UserActionDetailsActionResume,
}

var mappingUserActionDetailsActionEnumLowerCase = map[string]UserActionDetailsActionEnum{
	"retry":  UserActionDetailsActionRetry,
	"resume": UserActionDetailsActionResume,
}

// GetUserActionDetailsActionEnumValues Enumerates the set of values for UserActionDetailsActionEnum
func GetUserActionDetailsActionEnumValues() []UserActionDetailsActionEnum {
	values := make([]UserActionDetailsActionEnum, 0)
	for _, v := range mappingUserActionDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUserActionDetailsActionEnumStringValues Enumerates the set of values in String for UserActionDetailsActionEnum
func GetUserActionDetailsActionEnumStringValues() []string {
	return []string{
		"RETRY",
		"RESUME",
	}
}

// GetMappingUserActionDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserActionDetailsActionEnum(val string) (UserActionDetailsActionEnum, bool) {
	enum, ok := mappingUserActionDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UserActionDetailsLevelEnum Enum with underlying type: string
type UserActionDetailsLevelEnum string

// Set of constants representing the allowable values for UserActionDetailsLevelEnum
const (
	UserActionDetailsLevelActionGroup UserActionDetailsLevelEnum = "ACTION_GROUP"
	UserActionDetailsLevelStepName    UserActionDetailsLevelEnum = "STEP_NAME"
)

var mappingUserActionDetailsLevelEnum = map[string]UserActionDetailsLevelEnum{
	"ACTION_GROUP": UserActionDetailsLevelActionGroup,
	"STEP_NAME":    UserActionDetailsLevelStepName,
}

var mappingUserActionDetailsLevelEnumLowerCase = map[string]UserActionDetailsLevelEnum{
	"action_group": UserActionDetailsLevelActionGroup,
	"step_name":    UserActionDetailsLevelStepName,
}

// GetUserActionDetailsLevelEnumValues Enumerates the set of values for UserActionDetailsLevelEnum
func GetUserActionDetailsLevelEnumValues() []UserActionDetailsLevelEnum {
	values := make([]UserActionDetailsLevelEnum, 0)
	for _, v := range mappingUserActionDetailsLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetUserActionDetailsLevelEnumStringValues Enumerates the set of values in String for UserActionDetailsLevelEnum
func GetUserActionDetailsLevelEnumStringValues() []string {
	return []string{
		"ACTION_GROUP",
		"STEP_NAME",
	}
}

// GetMappingUserActionDetailsLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserActionDetailsLevelEnum(val string) (UserActionDetailsLevelEnum, bool) {
	enum, ok := mappingUserActionDetailsLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
