// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ActionDetails Object used to create an action.
type ActionDetails interface {

	// Whether or not this action is currently enabled.
	// Example: `true`
	GetIsEnabled() *bool

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	GetDescription() *string
}

type actiondetails struct {
	JsonData    []byte
	IsEnabled   *bool   `mandatory:"true" json:"isEnabled"`
	Description *string `mandatory:"false" json:"description"`
	ActionType  string  `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *actiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleractiondetails actiondetails
	s := struct {
		Model Unmarshaleractiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsEnabled = s.Model.IsEnabled
	m.Description = s.Model.Description
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *actiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "OSS":
		mm := CreateStreamingServiceActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAAS":
		mm := CreateFaaSActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONS":
		mm := CreateNotificationServiceActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetIsEnabled returns IsEnabled
func (m actiondetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m actiondetails) GetDescription() *string {
	return m.Description
}

func (m actiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m actiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ActionDetailsActionTypeEnum Enum with underlying type: string
type ActionDetailsActionTypeEnum string

// Set of constants representing the allowable values for ActionDetailsActionTypeEnum
const (
	ActionDetailsActionTypeOns  ActionDetailsActionTypeEnum = "ONS"
	ActionDetailsActionTypeOss  ActionDetailsActionTypeEnum = "OSS"
	ActionDetailsActionTypeFaas ActionDetailsActionTypeEnum = "FAAS"
)

var mappingActionDetailsActionTypeEnum = map[string]ActionDetailsActionTypeEnum{
	"ONS":  ActionDetailsActionTypeOns,
	"OSS":  ActionDetailsActionTypeOss,
	"FAAS": ActionDetailsActionTypeFaas,
}

// GetActionDetailsActionTypeEnumValues Enumerates the set of values for ActionDetailsActionTypeEnum
func GetActionDetailsActionTypeEnumValues() []ActionDetailsActionTypeEnum {
	values := make([]ActionDetailsActionTypeEnum, 0)
	for _, v := range mappingActionDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionDetailsActionTypeEnumStringValues Enumerates the set of values in String for ActionDetailsActionTypeEnum
func GetActionDetailsActionTypeEnumStringValues() []string {
	return []string{
		"ONS",
		"OSS",
		"FAAS",
	}
}

// GetMappingActionDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionDetailsActionTypeEnum(val string) (ActionDetailsActionTypeEnum, bool) {
	mappingActionDetailsActionTypeEnumIgnoreCase := make(map[string]ActionDetailsActionTypeEnum)
	for k, v := range mappingActionDetailsActionTypeEnum {
		mappingActionDetailsActionTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingActionDetailsActionTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
