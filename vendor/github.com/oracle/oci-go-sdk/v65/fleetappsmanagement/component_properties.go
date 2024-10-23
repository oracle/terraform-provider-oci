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

// ComponentProperties The properties of the component.
type ComponentProperties struct {

	// The action to be taken in case of a failure.
	ActionOnFailure ComponentPropertiesActionOnFailureEnum `mandatory:"true" json:"actionOnFailure"`

	// The runOn condition for the task/group/container.
	// Build task execution conditions if applicable to product and product-specific components.
	// This condition is relevant when handling product stack workflows.
	// Example: target.product.name = Oracle WebLogic Server OR target.product.name = Oracle HTTP Server
	RunOn *string `mandatory:"false" json:"runOn"`

	// Build control flow conditions that determine the relevance of the task execution.
	Condition *string `mandatory:"false" json:"condition"`

	PauseDetails PauseDetails `mandatory:"false" json:"pauseDetails"`

	NotificationPreferences *TaskNotificationPreferences `mandatory:"false" json:"notificationPreferences"`
}

func (m ComponentProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComponentProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComponentPropertiesActionOnFailureEnum(string(m.ActionOnFailure)); !ok && m.ActionOnFailure != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionOnFailure: %s. Supported values are: %s.", m.ActionOnFailure, strings.Join(GetComponentPropertiesActionOnFailureEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ComponentProperties) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		RunOn                   *string                                `json:"runOn"`
		Condition               *string                                `json:"condition"`
		PauseDetails            pausedetails                           `json:"pauseDetails"`
		NotificationPreferences *TaskNotificationPreferences           `json:"notificationPreferences"`
		ActionOnFailure         ComponentPropertiesActionOnFailureEnum `json:"actionOnFailure"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.RunOn = model.RunOn

	m.Condition = model.Condition

	nn, e = model.PauseDetails.UnmarshalPolymorphicJSON(model.PauseDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PauseDetails = nn.(PauseDetails)
	} else {
		m.PauseDetails = nil
	}

	m.NotificationPreferences = model.NotificationPreferences

	m.ActionOnFailure = model.ActionOnFailure

	return
}

// ComponentPropertiesActionOnFailureEnum Enum with underlying type: string
type ComponentPropertiesActionOnFailureEnum string

// Set of constants representing the allowable values for ComponentPropertiesActionOnFailureEnum
const (
	ComponentPropertiesActionOnFailureAbort    ComponentPropertiesActionOnFailureEnum = "ABORT"
	ComponentPropertiesActionOnFailureContinue ComponentPropertiesActionOnFailureEnum = "CONTINUE"
	ComponentPropertiesActionOnFailureRollback ComponentPropertiesActionOnFailureEnum = "ROLLBACK"
)

var mappingComponentPropertiesActionOnFailureEnum = map[string]ComponentPropertiesActionOnFailureEnum{
	"ABORT":    ComponentPropertiesActionOnFailureAbort,
	"CONTINUE": ComponentPropertiesActionOnFailureContinue,
	"ROLLBACK": ComponentPropertiesActionOnFailureRollback,
}

var mappingComponentPropertiesActionOnFailureEnumLowerCase = map[string]ComponentPropertiesActionOnFailureEnum{
	"abort":    ComponentPropertiesActionOnFailureAbort,
	"continue": ComponentPropertiesActionOnFailureContinue,
	"rollback": ComponentPropertiesActionOnFailureRollback,
}

// GetComponentPropertiesActionOnFailureEnumValues Enumerates the set of values for ComponentPropertiesActionOnFailureEnum
func GetComponentPropertiesActionOnFailureEnumValues() []ComponentPropertiesActionOnFailureEnum {
	values := make([]ComponentPropertiesActionOnFailureEnum, 0)
	for _, v := range mappingComponentPropertiesActionOnFailureEnum {
		values = append(values, v)
	}
	return values
}

// GetComponentPropertiesActionOnFailureEnumStringValues Enumerates the set of values in String for ComponentPropertiesActionOnFailureEnum
func GetComponentPropertiesActionOnFailureEnumStringValues() []string {
	return []string{
		"ABORT",
		"CONTINUE",
		"ROLLBACK",
	}
}

// GetMappingComponentPropertiesActionOnFailureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComponentPropertiesActionOnFailureEnum(val string) (ComponentPropertiesActionOnFailureEnum, bool) {
	enum, ok := mappingComponentPropertiesActionOnFailureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
