// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InactivityConfig Action and grace period for inactivity
type InactivityConfig struct {

	// an inactivity action to be triggered
	Action InactivityConfigActionEnum `mandatory:"true" json:"action"`

	// The period of time (in minutes) during which the session must remain inactive before any action occurs.
	// If the value is not provided, a default value is used.
	GracePeriodInMinutes *int `mandatory:"false" json:"gracePeriodInMinutes"`
}

func (m InactivityConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InactivityConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInactivityConfigActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetInactivityConfigActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InactivityConfigActionEnum Enum with underlying type: string
type InactivityConfigActionEnum string

// Set of constants representing the allowable values for InactivityConfigActionEnum
const (
	InactivityConfigActionNone       InactivityConfigActionEnum = "NONE"
	InactivityConfigActionDisconnect InactivityConfigActionEnum = "DISCONNECT"
)

var mappingInactivityConfigActionEnum = map[string]InactivityConfigActionEnum{
	"NONE":       InactivityConfigActionNone,
	"DISCONNECT": InactivityConfigActionDisconnect,
}

var mappingInactivityConfigActionEnumLowerCase = map[string]InactivityConfigActionEnum{
	"none":       InactivityConfigActionNone,
	"disconnect": InactivityConfigActionDisconnect,
}

// GetInactivityConfigActionEnumValues Enumerates the set of values for InactivityConfigActionEnum
func GetInactivityConfigActionEnumValues() []InactivityConfigActionEnum {
	values := make([]InactivityConfigActionEnum, 0)
	for _, v := range mappingInactivityConfigActionEnum {
		values = append(values, v)
	}
	return values
}

// GetInactivityConfigActionEnumStringValues Enumerates the set of values in String for InactivityConfigActionEnum
func GetInactivityConfigActionEnumStringValues() []string {
	return []string{
		"NONE",
		"DISCONNECT",
	}
}

// GetMappingInactivityConfigActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInactivityConfigActionEnum(val string) (InactivityConfigActionEnum, bool) {
	enum, ok := mappingInactivityConfigActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
