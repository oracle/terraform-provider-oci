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

// OlvmUsb Configuration of USB devices for this virtual machine (count, type).
type OlvmUsb struct {

	// Determines whether the USB device should be included or not.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Type of USB device redirection
	UsbType OlvmUsbUsbTypeEnum `mandatory:"false" json:"usbType,omitempty"`
}

func (m OlvmUsb) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmUsb) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmUsbUsbTypeEnum(string(m.UsbType)); !ok && m.UsbType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsbType: %s. Supported values are: %s.", m.UsbType, strings.Join(GetOlvmUsbUsbTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmUsbUsbTypeEnum Enum with underlying type: string
type OlvmUsbUsbTypeEnum string

// Set of constants representing the allowable values for OlvmUsbUsbTypeEnum
const (
	OlvmUsbUsbTypeLegacy OlvmUsbUsbTypeEnum = "LEGACY"
	OlvmUsbUsbTypeNative OlvmUsbUsbTypeEnum = "NATIVE"
)

var mappingOlvmUsbUsbTypeEnum = map[string]OlvmUsbUsbTypeEnum{
	"LEGACY": OlvmUsbUsbTypeLegacy,
	"NATIVE": OlvmUsbUsbTypeNative,
}

var mappingOlvmUsbUsbTypeEnumLowerCase = map[string]OlvmUsbUsbTypeEnum{
	"legacy": OlvmUsbUsbTypeLegacy,
	"native": OlvmUsbUsbTypeNative,
}

// GetOlvmUsbUsbTypeEnumValues Enumerates the set of values for OlvmUsbUsbTypeEnum
func GetOlvmUsbUsbTypeEnumValues() []OlvmUsbUsbTypeEnum {
	values := make([]OlvmUsbUsbTypeEnum, 0)
	for _, v := range mappingOlvmUsbUsbTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmUsbUsbTypeEnumStringValues Enumerates the set of values in String for OlvmUsbUsbTypeEnum
func GetOlvmUsbUsbTypeEnumStringValues() []string {
	return []string{
		"LEGACY",
		"NATIVE",
	}
}

// GetMappingOlvmUsbUsbTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmUsbUsbTypeEnum(val string) (OlvmUsbUsbTypeEnum, bool) {
	enum, ok := mappingOlvmUsbUsbTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
