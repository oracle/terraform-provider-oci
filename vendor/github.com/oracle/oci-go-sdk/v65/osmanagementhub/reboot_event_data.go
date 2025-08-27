// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RebootEventData Provides additional information for a reboot event.
type RebootEventData struct {

	// Reboot status for the current event
	RebootStatus RebootEventDataRebootStatusEnum `mandatory:"true" json:"rebootStatus"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m RebootEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RebootEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRebootEventDataRebootStatusEnum(string(m.RebootStatus)); !ok && m.RebootStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RebootStatus: %s. Supported values are: %s.", m.RebootStatus, strings.Join(GetRebootEventDataRebootStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RebootEventDataRebootStatusEnum Enum with underlying type: string
type RebootEventDataRebootStatusEnum string

// Set of constants representing the allowable values for RebootEventDataRebootStatusEnum
const (
	RebootEventDataRebootStatusStarted               RebootEventDataRebootStatusEnum = "REBOOT_STARTED"
	RebootEventDataRebootStatusSucceeded             RebootEventDataRebootStatusEnum = "REBOOT_SUCCEEDED"
	RebootEventDataRebootStatusFailed                RebootEventDataRebootStatusEnum = "REBOOT_FAILED"
	RebootEventDataRebootStatusSucceededAfterTimeout RebootEventDataRebootStatusEnum = "REBOOT_SUCCEEDED_AFTER_TIMEOUT"
)

var mappingRebootEventDataRebootStatusEnum = map[string]RebootEventDataRebootStatusEnum{
	"REBOOT_STARTED":                 RebootEventDataRebootStatusStarted,
	"REBOOT_SUCCEEDED":               RebootEventDataRebootStatusSucceeded,
	"REBOOT_FAILED":                  RebootEventDataRebootStatusFailed,
	"REBOOT_SUCCEEDED_AFTER_TIMEOUT": RebootEventDataRebootStatusSucceededAfterTimeout,
}

var mappingRebootEventDataRebootStatusEnumLowerCase = map[string]RebootEventDataRebootStatusEnum{
	"reboot_started":                 RebootEventDataRebootStatusStarted,
	"reboot_succeeded":               RebootEventDataRebootStatusSucceeded,
	"reboot_failed":                  RebootEventDataRebootStatusFailed,
	"reboot_succeeded_after_timeout": RebootEventDataRebootStatusSucceededAfterTimeout,
}

// GetRebootEventDataRebootStatusEnumValues Enumerates the set of values for RebootEventDataRebootStatusEnum
func GetRebootEventDataRebootStatusEnumValues() []RebootEventDataRebootStatusEnum {
	values := make([]RebootEventDataRebootStatusEnum, 0)
	for _, v := range mappingRebootEventDataRebootStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRebootEventDataRebootStatusEnumStringValues Enumerates the set of values in String for RebootEventDataRebootStatusEnum
func GetRebootEventDataRebootStatusEnumStringValues() []string {
	return []string{
		"REBOOT_STARTED",
		"REBOOT_SUCCEEDED",
		"REBOOT_FAILED",
		"REBOOT_SUCCEEDED_AFTER_TIMEOUT",
	}
}

// GetMappingRebootEventDataRebootStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRebootEventDataRebootStatusEnum(val string) (RebootEventDataRebootStatusEnum, bool) {
	enum, ok := mappingRebootEventDataRebootStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
