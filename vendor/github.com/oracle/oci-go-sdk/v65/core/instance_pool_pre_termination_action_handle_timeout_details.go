// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstancePoolPreTerminationActionHandleTimeoutDetails Options to handle timeout for pre-termination action.
type InstancePoolPreTerminationActionHandleTimeoutDetails struct {

	// Whether the block volume should be preserved after termination.
	PreserveBlockVolumeMode InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum `mandatory:"true" json:"preserveBlockVolumeMode"`

	// Whether the boot volume should be preserved after termination.
	PreserveBootVolumeMode InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum `mandatory:"true" json:"preserveBootVolumeMode"`
}

func (m InstancePoolPreTerminationActionHandleTimeoutDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstancePoolPreTerminationActionHandleTimeoutDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum(string(m.PreserveBlockVolumeMode)); !ok && m.PreserveBlockVolumeMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreserveBlockVolumeMode: %s. Supported values are: %s.", m.PreserveBlockVolumeMode, strings.Join(GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum(string(m.PreserveBootVolumeMode)); !ok && m.PreserveBootVolumeMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreserveBootVolumeMode: %s. Supported values are: %s.", m.PreserveBootVolumeMode, strings.Join(GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum Enum with underlying type: string
type InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum string

// Set of constants representing the allowable values for InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum
const (
	InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModePreserveAlways    InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum = "PRESERVE_ALWAYS"
	InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModePreserveOnTimeout InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum = "PRESERVE_ON_TIMEOUT"
	InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeDeleteAlways      InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum = "DELETE_ALWAYS"
)

var mappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum = map[string]InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum{
	"PRESERVE_ALWAYS":     InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModePreserveAlways,
	"PRESERVE_ON_TIMEOUT": InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModePreserveOnTimeout,
	"DELETE_ALWAYS":       InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeDeleteAlways,
}

var mappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnumLowerCase = map[string]InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum{
	"preserve_always":     InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModePreserveAlways,
	"preserve_on_timeout": InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModePreserveOnTimeout,
	"delete_always":       InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeDeleteAlways,
}

// GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnumValues Enumerates the set of values for InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum
func GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnumValues() []InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum {
	values := make([]InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum, 0)
	for _, v := range mappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnumStringValues Enumerates the set of values in String for InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum
func GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnumStringValues() []string {
	return []string{
		"PRESERVE_ALWAYS",
		"PRESERVE_ON_TIMEOUT",
		"DELETE_ALWAYS",
	}
}

// GetMappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum(val string) (InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnum, bool) {
	enum, ok := mappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBlockVolumeModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum Enum with underlying type: string
type InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum string

// Set of constants representing the allowable values for InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum
const (
	InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModePreserveAlways    InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum = "PRESERVE_ALWAYS"
	InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModePreserveOnTimeout InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum = "PRESERVE_ON_TIMEOUT"
	InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeDeleteAlways      InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum = "DELETE_ALWAYS"
)

var mappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum = map[string]InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum{
	"PRESERVE_ALWAYS":     InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModePreserveAlways,
	"PRESERVE_ON_TIMEOUT": InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModePreserveOnTimeout,
	"DELETE_ALWAYS":       InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeDeleteAlways,
}

var mappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnumLowerCase = map[string]InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum{
	"preserve_always":     InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModePreserveAlways,
	"preserve_on_timeout": InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModePreserveOnTimeout,
	"delete_always":       InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeDeleteAlways,
}

// GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnumValues Enumerates the set of values for InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum
func GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnumValues() []InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum {
	values := make([]InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum, 0)
	for _, v := range mappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnumStringValues Enumerates the set of values in String for InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum
func GetInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnumStringValues() []string {
	return []string{
		"PRESERVE_ALWAYS",
		"PRESERVE_ON_TIMEOUT",
		"DELETE_ALWAYS",
	}
}

// GetMappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum(val string) (InstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnum, bool) {
	enum, ok := mappingInstancePoolPreTerminationActionHandleTimeoutDetailsPreserveBootVolumeModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
