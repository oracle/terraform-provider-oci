// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DesktopAction Provides information about actions performed on a desktop, including type and time.
type DesktopAction struct {

	// An action performed on a desktop.
	Action DesktopActionActionEnum `mandatory:"true" json:"action"`

	// The time of an action performed on a desktop.
	TimeApplied *common.SDKTime `mandatory:"true" json:"timeApplied"`
}

func (m DesktopAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DesktopAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDesktopActionActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDesktopActionActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DesktopActionActionEnum Enum with underlying type: string
type DesktopActionActionEnum string

// Set of constants representing the allowable values for DesktopActionActionEnum
const (
	DesktopActionActionApiStart           DesktopActionActionEnum = "API_START"
	DesktopActionActionApiStop            DesktopActionActionEnum = "API_STOP"
	DesktopActionActionAvailabilityStart  DesktopActionActionEnum = "AVAILABILITY_START"
	DesktopActionActionAvailabilityStop   DesktopActionActionEnum = "AVAILABILITY_STOP"
	DesktopActionActionDisconnectedStop   DesktopActionActionEnum = "DISCONNECTED_STOP"
	DesktopActionActionDisconnectedDelete DesktopActionActionEnum = "DISCONNECTED_DELETE"
	DesktopActionActionScheduledStart     DesktopActionActionEnum = "SCHEDULED_START"
	DesktopActionActionScheduledStop      DesktopActionActionEnum = "SCHEDULED_STOP"
	DesktopActionActionSynchronize        DesktopActionActionEnum = "SYNCHRONIZE"
	DesktopActionActionConnected          DesktopActionActionEnum = "CONNECTED"
	DesktopActionActionDisconnected       DesktopActionActionEnum = "DISCONNECTED"
)

var mappingDesktopActionActionEnum = map[string]DesktopActionActionEnum{
	"API_START":           DesktopActionActionApiStart,
	"API_STOP":            DesktopActionActionApiStop,
	"AVAILABILITY_START":  DesktopActionActionAvailabilityStart,
	"AVAILABILITY_STOP":   DesktopActionActionAvailabilityStop,
	"DISCONNECTED_STOP":   DesktopActionActionDisconnectedStop,
	"DISCONNECTED_DELETE": DesktopActionActionDisconnectedDelete,
	"SCHEDULED_START":     DesktopActionActionScheduledStart,
	"SCHEDULED_STOP":      DesktopActionActionScheduledStop,
	"SYNCHRONIZE":         DesktopActionActionSynchronize,
	"CONNECTED":           DesktopActionActionConnected,
	"DISCONNECTED":        DesktopActionActionDisconnected,
}

var mappingDesktopActionActionEnumLowerCase = map[string]DesktopActionActionEnum{
	"api_start":           DesktopActionActionApiStart,
	"api_stop":            DesktopActionActionApiStop,
	"availability_start":  DesktopActionActionAvailabilityStart,
	"availability_stop":   DesktopActionActionAvailabilityStop,
	"disconnected_stop":   DesktopActionActionDisconnectedStop,
	"disconnected_delete": DesktopActionActionDisconnectedDelete,
	"scheduled_start":     DesktopActionActionScheduledStart,
	"scheduled_stop":      DesktopActionActionScheduledStop,
	"synchronize":         DesktopActionActionSynchronize,
	"connected":           DesktopActionActionConnected,
	"disconnected":        DesktopActionActionDisconnected,
}

// GetDesktopActionActionEnumValues Enumerates the set of values for DesktopActionActionEnum
func GetDesktopActionActionEnumValues() []DesktopActionActionEnum {
	values := make([]DesktopActionActionEnum, 0)
	for _, v := range mappingDesktopActionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDesktopActionActionEnumStringValues Enumerates the set of values in String for DesktopActionActionEnum
func GetDesktopActionActionEnumStringValues() []string {
	return []string{
		"API_START",
		"API_STOP",
		"AVAILABILITY_START",
		"AVAILABILITY_STOP",
		"DISCONNECTED_STOP",
		"DISCONNECTED_DELETE",
		"SCHEDULED_START",
		"SCHEDULED_STOP",
		"SYNCHRONIZE",
		"CONNECTED",
		"DISCONNECTED",
	}
}

// GetMappingDesktopActionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDesktopActionActionEnum(val string) (DesktopActionActionEnum, bool) {
	enum, ok := mappingDesktopActionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
