// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmDisplay Display object in OLVM
type OlvmDisplay struct {

	// The IP address of the guest to connect the graphic console client to.
	Address *string `mandatory:"false" json:"address"`

	// Indicates if to override the display address per host.
	IsAllowOverride *bool `mandatory:"false" json:"isAllowOverride"`

	Certificate *OlvmCertificate `mandatory:"false" json:"certificate"`

	// Indicates whether a user is able to copy and paste content from an external host into the graphic console.
	IsCopyPasteEnabled *bool `mandatory:"false" json:"isCopyPasteEnabled"`

	// Returns the action that will take place when the graphic console is disconnected.
	DisconnectAction *string `mandatory:"false" json:"disconnectAction"`

	// Delay (in minutes) before the graphic console disconnect action is carried out.
	DisconnectActionDelayInMinutes *int `mandatory:"false" json:"disconnectActionDelayInMinutes"`

	// Indicates if a user is able to drag and drop files from an external host into the graphic console.
	IsFileTransferEnabled *bool `mandatory:"false" json:"isFileTransferEnabled"`

	// The keyboard layout to use with this graphic console.
	KeyboardLayout *string `mandatory:"false" json:"keyboardLayout"`

	// The number of monitors opened for this graphic console.
	Monitors *int `mandatory:"false" json:"monitors"`

	// The port address on the guest to connect the graphic console client to
	Port *int `mandatory:"false" json:"port"`

	// The proxy IP which will be used by the graphic console client to connect to the guest.
	Proxy *string `mandatory:"false" json:"proxy"`

	// The secured port address on the guest, in case of using TLS, to connect the graphic console client to.
	SecurePort *int `mandatory:"false" json:"securePort"`

	// The engine now sets it automatically according to the operating system.
	IsSingleQxlPci *bool `mandatory:"false" json:"isSingleQxlPci"`

	// The graphic console protocol type.
	DisplayType OlvmDisplayDisplayTypeEnum `mandatory:"false" json:"displayType,omitempty"`
}

func (m OlvmDisplay) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmDisplay) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmDisplayDisplayTypeEnum(string(m.DisplayType)); !ok && m.DisplayType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisplayType: %s. Supported values are: %s.", m.DisplayType, strings.Join(GetOlvmDisplayDisplayTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmDisplayDisplayTypeEnum Enum with underlying type: string
type OlvmDisplayDisplayTypeEnum string

// Set of constants representing the allowable values for OlvmDisplayDisplayTypeEnum
const (
	OlvmDisplayDisplayTypeSpice OlvmDisplayDisplayTypeEnum = "SPICE"
	OlvmDisplayDisplayTypeVnc   OlvmDisplayDisplayTypeEnum = "VNC"
)

var mappingOlvmDisplayDisplayTypeEnum = map[string]OlvmDisplayDisplayTypeEnum{
	"SPICE": OlvmDisplayDisplayTypeSpice,
	"VNC":   OlvmDisplayDisplayTypeVnc,
}

var mappingOlvmDisplayDisplayTypeEnumLowerCase = map[string]OlvmDisplayDisplayTypeEnum{
	"spice": OlvmDisplayDisplayTypeSpice,
	"vnc":   OlvmDisplayDisplayTypeVnc,
}

// GetOlvmDisplayDisplayTypeEnumValues Enumerates the set of values for OlvmDisplayDisplayTypeEnum
func GetOlvmDisplayDisplayTypeEnumValues() []OlvmDisplayDisplayTypeEnum {
	values := make([]OlvmDisplayDisplayTypeEnum, 0)
	for _, v := range mappingOlvmDisplayDisplayTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmDisplayDisplayTypeEnumStringValues Enumerates the set of values in String for OlvmDisplayDisplayTypeEnum
func GetOlvmDisplayDisplayTypeEnumStringValues() []string {
	return []string{
		"SPICE",
		"VNC",
	}
}

// GetMappingOlvmDisplayDisplayTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmDisplayDisplayTypeEnum(val string) (OlvmDisplayDisplayTypeEnum, bool) {
	enum, ok := mappingOlvmDisplayDisplayTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
