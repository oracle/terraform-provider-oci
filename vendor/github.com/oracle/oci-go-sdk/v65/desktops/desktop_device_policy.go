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

// DesktopDevicePolicy Provides the settings for desktop and client device options, such as audio in and out, client drive mapping, and clipboard access.
type DesktopDevicePolicy struct {

	// The clipboard mode.
	// NONE: No access to the local clipboard is permitted.
	// TODESKTOP: The clipboard can be used to transfer data to the desktop only.
	// FROMDESKTOP: The clipboard can be used to transfer data from the desktop only.
	// FULL: The clipboard can be used to transfer data to and from the desktop.
	ClipboardMode DesktopDevicePolicyClipboardModeEnum `mandatory:"true" json:"clipboardMode"`

	// The audio mode.
	// NONE: No access to the local audio devices is permitted.
	// TODESKTOP: The user may record audio on their desktop.
	// FROMDESKTOP: The user may play audio on their desktop.
	// FULL: The user may play and record audio on their desktop.
	AudioMode DesktopDevicePolicyAudioModeEnum `mandatory:"true" json:"audioMode"`

	// The client local drive access mode.
	// NONE: No access to local drives permitted.
	// READONLY: The user may read from local drives on their desktop.
	// FULL: The user may read from and write to their local drives on their desktop.
	CdmMode DesktopDevicePolicyCdmModeEnum `mandatory:"true" json:"cdmMode"`

	// Indicates whether printing is enabled.
	IsPrintingEnabled *bool `mandatory:"true" json:"isPrintingEnabled"`

	// Indicates whether the pointer is enabled.
	IsPointerEnabled *bool `mandatory:"true" json:"isPointerEnabled"`

	// Indicates whether the keyboard is enabled.
	IsKeyboardEnabled *bool `mandatory:"true" json:"isKeyboardEnabled"`

	// Indicates whether the display is enabled.
	IsDisplayEnabled *bool `mandatory:"true" json:"isDisplayEnabled"`
}

func (m DesktopDevicePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DesktopDevicePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDesktopDevicePolicyClipboardModeEnum(string(m.ClipboardMode)); !ok && m.ClipboardMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClipboardMode: %s. Supported values are: %s.", m.ClipboardMode, strings.Join(GetDesktopDevicePolicyClipboardModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDesktopDevicePolicyAudioModeEnum(string(m.AudioMode)); !ok && m.AudioMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AudioMode: %s. Supported values are: %s.", m.AudioMode, strings.Join(GetDesktopDevicePolicyAudioModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDesktopDevicePolicyCdmModeEnum(string(m.CdmMode)); !ok && m.CdmMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CdmMode: %s. Supported values are: %s.", m.CdmMode, strings.Join(GetDesktopDevicePolicyCdmModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DesktopDevicePolicyClipboardModeEnum Enum with underlying type: string
type DesktopDevicePolicyClipboardModeEnum string

// Set of constants representing the allowable values for DesktopDevicePolicyClipboardModeEnum
const (
	DesktopDevicePolicyClipboardModeNone        DesktopDevicePolicyClipboardModeEnum = "NONE"
	DesktopDevicePolicyClipboardModeTodesktop   DesktopDevicePolicyClipboardModeEnum = "TODESKTOP"
	DesktopDevicePolicyClipboardModeFromdesktop DesktopDevicePolicyClipboardModeEnum = "FROMDESKTOP"
	DesktopDevicePolicyClipboardModeFull        DesktopDevicePolicyClipboardModeEnum = "FULL"
)

var mappingDesktopDevicePolicyClipboardModeEnum = map[string]DesktopDevicePolicyClipboardModeEnum{
	"NONE":        DesktopDevicePolicyClipboardModeNone,
	"TODESKTOP":   DesktopDevicePolicyClipboardModeTodesktop,
	"FROMDESKTOP": DesktopDevicePolicyClipboardModeFromdesktop,
	"FULL":        DesktopDevicePolicyClipboardModeFull,
}

var mappingDesktopDevicePolicyClipboardModeEnumLowerCase = map[string]DesktopDevicePolicyClipboardModeEnum{
	"none":        DesktopDevicePolicyClipboardModeNone,
	"todesktop":   DesktopDevicePolicyClipboardModeTodesktop,
	"fromdesktop": DesktopDevicePolicyClipboardModeFromdesktop,
	"full":        DesktopDevicePolicyClipboardModeFull,
}

// GetDesktopDevicePolicyClipboardModeEnumValues Enumerates the set of values for DesktopDevicePolicyClipboardModeEnum
func GetDesktopDevicePolicyClipboardModeEnumValues() []DesktopDevicePolicyClipboardModeEnum {
	values := make([]DesktopDevicePolicyClipboardModeEnum, 0)
	for _, v := range mappingDesktopDevicePolicyClipboardModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDesktopDevicePolicyClipboardModeEnumStringValues Enumerates the set of values in String for DesktopDevicePolicyClipboardModeEnum
func GetDesktopDevicePolicyClipboardModeEnumStringValues() []string {
	return []string{
		"NONE",
		"TODESKTOP",
		"FROMDESKTOP",
		"FULL",
	}
}

// GetMappingDesktopDevicePolicyClipboardModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDesktopDevicePolicyClipboardModeEnum(val string) (DesktopDevicePolicyClipboardModeEnum, bool) {
	enum, ok := mappingDesktopDevicePolicyClipboardModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DesktopDevicePolicyAudioModeEnum Enum with underlying type: string
type DesktopDevicePolicyAudioModeEnum string

// Set of constants representing the allowable values for DesktopDevicePolicyAudioModeEnum
const (
	DesktopDevicePolicyAudioModeNone        DesktopDevicePolicyAudioModeEnum = "NONE"
	DesktopDevicePolicyAudioModeTodesktop   DesktopDevicePolicyAudioModeEnum = "TODESKTOP"
	DesktopDevicePolicyAudioModeFromdesktop DesktopDevicePolicyAudioModeEnum = "FROMDESKTOP"
	DesktopDevicePolicyAudioModeFull        DesktopDevicePolicyAudioModeEnum = "FULL"
)

var mappingDesktopDevicePolicyAudioModeEnum = map[string]DesktopDevicePolicyAudioModeEnum{
	"NONE":        DesktopDevicePolicyAudioModeNone,
	"TODESKTOP":   DesktopDevicePolicyAudioModeTodesktop,
	"FROMDESKTOP": DesktopDevicePolicyAudioModeFromdesktop,
	"FULL":        DesktopDevicePolicyAudioModeFull,
}

var mappingDesktopDevicePolicyAudioModeEnumLowerCase = map[string]DesktopDevicePolicyAudioModeEnum{
	"none":        DesktopDevicePolicyAudioModeNone,
	"todesktop":   DesktopDevicePolicyAudioModeTodesktop,
	"fromdesktop": DesktopDevicePolicyAudioModeFromdesktop,
	"full":        DesktopDevicePolicyAudioModeFull,
}

// GetDesktopDevicePolicyAudioModeEnumValues Enumerates the set of values for DesktopDevicePolicyAudioModeEnum
func GetDesktopDevicePolicyAudioModeEnumValues() []DesktopDevicePolicyAudioModeEnum {
	values := make([]DesktopDevicePolicyAudioModeEnum, 0)
	for _, v := range mappingDesktopDevicePolicyAudioModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDesktopDevicePolicyAudioModeEnumStringValues Enumerates the set of values in String for DesktopDevicePolicyAudioModeEnum
func GetDesktopDevicePolicyAudioModeEnumStringValues() []string {
	return []string{
		"NONE",
		"TODESKTOP",
		"FROMDESKTOP",
		"FULL",
	}
}

// GetMappingDesktopDevicePolicyAudioModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDesktopDevicePolicyAudioModeEnum(val string) (DesktopDevicePolicyAudioModeEnum, bool) {
	enum, ok := mappingDesktopDevicePolicyAudioModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DesktopDevicePolicyCdmModeEnum Enum with underlying type: string
type DesktopDevicePolicyCdmModeEnum string

// Set of constants representing the allowable values for DesktopDevicePolicyCdmModeEnum
const (
	DesktopDevicePolicyCdmModeNone     DesktopDevicePolicyCdmModeEnum = "NONE"
	DesktopDevicePolicyCdmModeReadonly DesktopDevicePolicyCdmModeEnum = "READONLY"
	DesktopDevicePolicyCdmModeFull     DesktopDevicePolicyCdmModeEnum = "FULL"
)

var mappingDesktopDevicePolicyCdmModeEnum = map[string]DesktopDevicePolicyCdmModeEnum{
	"NONE":     DesktopDevicePolicyCdmModeNone,
	"READONLY": DesktopDevicePolicyCdmModeReadonly,
	"FULL":     DesktopDevicePolicyCdmModeFull,
}

var mappingDesktopDevicePolicyCdmModeEnumLowerCase = map[string]DesktopDevicePolicyCdmModeEnum{
	"none":     DesktopDevicePolicyCdmModeNone,
	"readonly": DesktopDevicePolicyCdmModeReadonly,
	"full":     DesktopDevicePolicyCdmModeFull,
}

// GetDesktopDevicePolicyCdmModeEnumValues Enumerates the set of values for DesktopDevicePolicyCdmModeEnum
func GetDesktopDevicePolicyCdmModeEnumValues() []DesktopDevicePolicyCdmModeEnum {
	values := make([]DesktopDevicePolicyCdmModeEnum, 0)
	for _, v := range mappingDesktopDevicePolicyCdmModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDesktopDevicePolicyCdmModeEnumStringValues Enumerates the set of values in String for DesktopDevicePolicyCdmModeEnum
func GetDesktopDevicePolicyCdmModeEnumStringValues() []string {
	return []string{
		"NONE",
		"READONLY",
		"FULL",
	}
}

// GetMappingDesktopDevicePolicyCdmModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDesktopDevicePolicyCdmModeEnum(val string) (DesktopDevicePolicyCdmModeEnum, bool) {
	enum, ok := mappingDesktopDevicePolicyCdmModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
