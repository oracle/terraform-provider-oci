// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComponentProperties The properties of the task.
type ComponentProperties struct {

	// The action to be taken in case of task failure.
	ActionOnFailure ComponentPropertiesActionOnFailureEnum `mandatory:"true" json:"actionOnFailure"`

	// The hosts to execute on.
	RunOn *string `mandatory:"false" json:"runOn"`

	// The condition in which the task is to be executed.
	Condition *string `mandatory:"false" json:"condition"`
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
