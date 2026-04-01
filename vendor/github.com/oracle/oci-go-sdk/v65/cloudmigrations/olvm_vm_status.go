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

// OlvmVmStatus The current status of the virtual machine.
type OlvmVmStatus struct {

	// Type representing a status of a virtual machine.
	Status OlvmVmStatusStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m OlvmVmStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmVmStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmVmStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOlvmVmStatusStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmVmStatusStatusEnum Enum with underlying type: string
type OlvmVmStatusStatusEnum string

// Set of constants representing the allowable values for OlvmVmStatusStatusEnum
const (
	OlvmVmStatusStatusDown             OlvmVmStatusStatusEnum = "DOWN"
	OlvmVmStatusStatusImageLocked      OlvmVmStatusStatusEnum = "IMAGE_LOCKED"
	OlvmVmStatusStatusMigrating        OlvmVmStatusStatusEnum = "MIGRATING"
	OlvmVmStatusStatusNotResponding    OlvmVmStatusStatusEnum = "NOT_RESPONDING"
	OlvmVmStatusStatusPaused           OlvmVmStatusStatusEnum = "PAUSED"
	OlvmVmStatusStatusPoweringDown     OlvmVmStatusStatusEnum = "POWERING_DOWN"
	OlvmVmStatusStatusPoweringUp       OlvmVmStatusStatusEnum = "POWERING_UP"
	OlvmVmStatusStatusRebootInProgress OlvmVmStatusStatusEnum = "REBOOT_IN_PROGRESS"
	OlvmVmStatusStatusRestoringState   OlvmVmStatusStatusEnum = "RESTORING_STATE"
	OlvmVmStatusStatusSavingState      OlvmVmStatusStatusEnum = "SAVING_STATE"
	OlvmVmStatusStatusSuspended        OlvmVmStatusStatusEnum = "SUSPENDED"
	OlvmVmStatusStatusUnassigned       OlvmVmStatusStatusEnum = "UNASSIGNED"
	OlvmVmStatusStatusUnknown          OlvmVmStatusStatusEnum = "UNKNOWN"
	OlvmVmStatusStatusUp               OlvmVmStatusStatusEnum = "UP"
	OlvmVmStatusStatusWaitForLaunch    OlvmVmStatusStatusEnum = "WAIT_FOR_LAUNCH"
)

var mappingOlvmVmStatusStatusEnum = map[string]OlvmVmStatusStatusEnum{
	"DOWN":               OlvmVmStatusStatusDown,
	"IMAGE_LOCKED":       OlvmVmStatusStatusImageLocked,
	"MIGRATING":          OlvmVmStatusStatusMigrating,
	"NOT_RESPONDING":     OlvmVmStatusStatusNotResponding,
	"PAUSED":             OlvmVmStatusStatusPaused,
	"POWERING_DOWN":      OlvmVmStatusStatusPoweringDown,
	"POWERING_UP":        OlvmVmStatusStatusPoweringUp,
	"REBOOT_IN_PROGRESS": OlvmVmStatusStatusRebootInProgress,
	"RESTORING_STATE":    OlvmVmStatusStatusRestoringState,
	"SAVING_STATE":       OlvmVmStatusStatusSavingState,
	"SUSPENDED":          OlvmVmStatusStatusSuspended,
	"UNASSIGNED":         OlvmVmStatusStatusUnassigned,
	"UNKNOWN":            OlvmVmStatusStatusUnknown,
	"UP":                 OlvmVmStatusStatusUp,
	"WAIT_FOR_LAUNCH":    OlvmVmStatusStatusWaitForLaunch,
}

var mappingOlvmVmStatusStatusEnumLowerCase = map[string]OlvmVmStatusStatusEnum{
	"down":               OlvmVmStatusStatusDown,
	"image_locked":       OlvmVmStatusStatusImageLocked,
	"migrating":          OlvmVmStatusStatusMigrating,
	"not_responding":     OlvmVmStatusStatusNotResponding,
	"paused":             OlvmVmStatusStatusPaused,
	"powering_down":      OlvmVmStatusStatusPoweringDown,
	"powering_up":        OlvmVmStatusStatusPoweringUp,
	"reboot_in_progress": OlvmVmStatusStatusRebootInProgress,
	"restoring_state":    OlvmVmStatusStatusRestoringState,
	"saving_state":       OlvmVmStatusStatusSavingState,
	"suspended":          OlvmVmStatusStatusSuspended,
	"unassigned":         OlvmVmStatusStatusUnassigned,
	"unknown":            OlvmVmStatusStatusUnknown,
	"up":                 OlvmVmStatusStatusUp,
	"wait_for_launch":    OlvmVmStatusStatusWaitForLaunch,
}

// GetOlvmVmStatusStatusEnumValues Enumerates the set of values for OlvmVmStatusStatusEnum
func GetOlvmVmStatusStatusEnumValues() []OlvmVmStatusStatusEnum {
	values := make([]OlvmVmStatusStatusEnum, 0)
	for _, v := range mappingOlvmVmStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmVmStatusStatusEnumStringValues Enumerates the set of values in String for OlvmVmStatusStatusEnum
func GetOlvmVmStatusStatusEnumStringValues() []string {
	return []string{
		"DOWN",
		"IMAGE_LOCKED",
		"MIGRATING",
		"NOT_RESPONDING",
		"PAUSED",
		"POWERING_DOWN",
		"POWERING_UP",
		"REBOOT_IN_PROGRESS",
		"RESTORING_STATE",
		"SAVING_STATE",
		"SUSPENDED",
		"UNASSIGNED",
		"UNKNOWN",
		"UP",
		"WAIT_FOR_LAUNCH",
	}
}

// GetMappingOlvmVmStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmVmStatusStatusEnum(val string) (OlvmVmStatusStatusEnum, bool) {
	enum, ok := mappingOlvmVmStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
