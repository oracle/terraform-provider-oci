// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SnapUpdateEventData Provides additional information for a snap update event.
type SnapUpdateEventData struct {

	// Type of snap update operation.
	OperationType SnapUpdateEventDataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the snap update.
	Status EventStatusEnum `mandatory:"true" json:"status"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m SnapUpdateEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SnapUpdateEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSnapUpdateEventDataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetSnapUpdateEventDataOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SnapUpdateEventDataOperationTypeEnum Enum with underlying type: string
type SnapUpdateEventDataOperationTypeEnum string

// Set of constants representing the allowable values for SnapUpdateEventDataOperationTypeEnum
const (
	SnapUpdateEventDataOperationTypeListSnaps         SnapUpdateEventDataOperationTypeEnum = "LIST_SNAPS"
	SnapUpdateEventDataOperationTypeInstallSnaps      SnapUpdateEventDataOperationTypeEnum = "INSTALL_SNAPS"
	SnapUpdateEventDataOperationTypeRemoveSnaps       SnapUpdateEventDataOperationTypeEnum = "REMOVE_SNAPS"
	SnapUpdateEventDataOperationTypeSwitchSnapChannel SnapUpdateEventDataOperationTypeEnum = "SWITCH_SNAP_CHANNEL"
)

var mappingSnapUpdateEventDataOperationTypeEnum = map[string]SnapUpdateEventDataOperationTypeEnum{
	"LIST_SNAPS":          SnapUpdateEventDataOperationTypeListSnaps,
	"INSTALL_SNAPS":       SnapUpdateEventDataOperationTypeInstallSnaps,
	"REMOVE_SNAPS":        SnapUpdateEventDataOperationTypeRemoveSnaps,
	"SWITCH_SNAP_CHANNEL": SnapUpdateEventDataOperationTypeSwitchSnapChannel,
}

var mappingSnapUpdateEventDataOperationTypeEnumLowerCase = map[string]SnapUpdateEventDataOperationTypeEnum{
	"list_snaps":          SnapUpdateEventDataOperationTypeListSnaps,
	"install_snaps":       SnapUpdateEventDataOperationTypeInstallSnaps,
	"remove_snaps":        SnapUpdateEventDataOperationTypeRemoveSnaps,
	"switch_snap_channel": SnapUpdateEventDataOperationTypeSwitchSnapChannel,
}

// GetSnapUpdateEventDataOperationTypeEnumValues Enumerates the set of values for SnapUpdateEventDataOperationTypeEnum
func GetSnapUpdateEventDataOperationTypeEnumValues() []SnapUpdateEventDataOperationTypeEnum {
	values := make([]SnapUpdateEventDataOperationTypeEnum, 0)
	for _, v := range mappingSnapUpdateEventDataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapUpdateEventDataOperationTypeEnumStringValues Enumerates the set of values in String for SnapUpdateEventDataOperationTypeEnum
func GetSnapUpdateEventDataOperationTypeEnumStringValues() []string {
	return []string{
		"LIST_SNAPS",
		"INSTALL_SNAPS",
		"REMOVE_SNAPS",
		"SWITCH_SNAP_CHANNEL",
	}
}

// GetMappingSnapUpdateEventDataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapUpdateEventDataOperationTypeEnum(val string) (SnapUpdateEventDataOperationTypeEnum, bool) {
	enum, ok := mappingSnapUpdateEventDataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
