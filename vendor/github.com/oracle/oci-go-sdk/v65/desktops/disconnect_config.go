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

// DisconnectConfig Action and grace period for disconnect
type DisconnectConfig struct {

	// a disconnect action to be triggered
	Action DisconnectConfigActionEnum `mandatory:"true" json:"action"`

	// The period of time (in minutes) after disconnect before any action occurs.
	// If the value is not provided, a default value is used.
	GracePeriodInMinutes *int `mandatory:"false" json:"gracePeriodInMinutes"`
}

func (m DisconnectConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisconnectConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDisconnectConfigActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDisconnectConfigActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DisconnectConfigActionEnum Enum with underlying type: string
type DisconnectConfigActionEnum string

// Set of constants representing the allowable values for DisconnectConfigActionEnum
const (
	DisconnectConfigActionNone DisconnectConfigActionEnum = "NONE"
	DisconnectConfigActionStop DisconnectConfigActionEnum = "STOP"
)

var mappingDisconnectConfigActionEnum = map[string]DisconnectConfigActionEnum{
	"NONE": DisconnectConfigActionNone,
	"STOP": DisconnectConfigActionStop,
}

var mappingDisconnectConfigActionEnumLowerCase = map[string]DisconnectConfigActionEnum{
	"none": DisconnectConfigActionNone,
	"stop": DisconnectConfigActionStop,
}

// GetDisconnectConfigActionEnumValues Enumerates the set of values for DisconnectConfigActionEnum
func GetDisconnectConfigActionEnumValues() []DisconnectConfigActionEnum {
	values := make([]DisconnectConfigActionEnum, 0)
	for _, v := range mappingDisconnectConfigActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDisconnectConfigActionEnumStringValues Enumerates the set of values in String for DisconnectConfigActionEnum
func GetDisconnectConfigActionEnumStringValues() []string {
	return []string{
		"NONE",
		"STOP",
	}
}

// GetMappingDisconnectConfigActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisconnectConfigActionEnum(val string) (DisconnectConfigActionEnum, bool) {
	enum, ok := mappingDisconnectConfigActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
