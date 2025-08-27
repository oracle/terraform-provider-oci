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

// ComputeHostConfigurationCheckDetails Compute Host Group Configuration Details Check
type ComputeHostConfigurationCheckDetails struct {

	// The type of configuration
	Type ComputeHostConfigurationCheckDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The current state of the host configuration. The Host is either |
	// CONFORMANT - current state matches the desired configuration
	// NON_CONFORMANT - current state does not match the desired configuration
	// PRE_APPLYING, APPLYING, CHECKING- transitional states
	// UNKNOWN - current state is unknown
	ConfigurationState ConfigurationStateEnum `mandatory:"false" json:"configurationState,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique firmware bundle associated with the Host Configuration.
	FirmwareBundleId *string `mandatory:"false" json:"firmwareBundleId"`

	// Preferred recycle level for hosts associated with the reservation config.
	// * `SKIP_RECYCLE` - Skips host wipe.
	// * `FULL_RECYCLE` - Does not skip host wipe. This is the default behavior.
	RecycleLevel ComputeHostConfigurationCheckDetailsRecycleLevelEnum `mandatory:"false" json:"recycleLevel,omitempty"`
}

func (m ComputeHostConfigurationCheckDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeHostConfigurationCheckDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputeHostConfigurationCheckDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetComputeHostConfigurationCheckDetailsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigurationStateEnum(string(m.ConfigurationState)); !ok && m.ConfigurationState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigurationState: %s. Supported values are: %s.", m.ConfigurationState, strings.Join(GetConfigurationStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingComputeHostConfigurationCheckDetailsRecycleLevelEnum(string(m.RecycleLevel)); !ok && m.RecycleLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecycleLevel: %s. Supported values are: %s.", m.RecycleLevel, strings.Join(GetComputeHostConfigurationCheckDetailsRecycleLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeHostConfigurationCheckDetailsTypeEnum Enum with underlying type: string
type ComputeHostConfigurationCheckDetailsTypeEnum string

// Set of constants representing the allowable values for ComputeHostConfigurationCheckDetailsTypeEnum
const (
	ComputeHostConfigurationCheckDetailsTypeFirmware ComputeHostConfigurationCheckDetailsTypeEnum = "FIRMWARE"
	ComputeHostConfigurationCheckDetailsTypeRecycle  ComputeHostConfigurationCheckDetailsTypeEnum = "RECYCLE"
)

var mappingComputeHostConfigurationCheckDetailsTypeEnum = map[string]ComputeHostConfigurationCheckDetailsTypeEnum{
	"FIRMWARE": ComputeHostConfigurationCheckDetailsTypeFirmware,
	"RECYCLE":  ComputeHostConfigurationCheckDetailsTypeRecycle,
}

var mappingComputeHostConfigurationCheckDetailsTypeEnumLowerCase = map[string]ComputeHostConfigurationCheckDetailsTypeEnum{
	"firmware": ComputeHostConfigurationCheckDetailsTypeFirmware,
	"recycle":  ComputeHostConfigurationCheckDetailsTypeRecycle,
}

// GetComputeHostConfigurationCheckDetailsTypeEnumValues Enumerates the set of values for ComputeHostConfigurationCheckDetailsTypeEnum
func GetComputeHostConfigurationCheckDetailsTypeEnumValues() []ComputeHostConfigurationCheckDetailsTypeEnum {
	values := make([]ComputeHostConfigurationCheckDetailsTypeEnum, 0)
	for _, v := range mappingComputeHostConfigurationCheckDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeHostConfigurationCheckDetailsTypeEnumStringValues Enumerates the set of values in String for ComputeHostConfigurationCheckDetailsTypeEnum
func GetComputeHostConfigurationCheckDetailsTypeEnumStringValues() []string {
	return []string{
		"FIRMWARE",
		"RECYCLE",
	}
}

// GetMappingComputeHostConfigurationCheckDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeHostConfigurationCheckDetailsTypeEnum(val string) (ComputeHostConfigurationCheckDetailsTypeEnum, bool) {
	enum, ok := mappingComputeHostConfigurationCheckDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ComputeHostConfigurationCheckDetailsRecycleLevelEnum Enum with underlying type: string
type ComputeHostConfigurationCheckDetailsRecycleLevelEnum string

// Set of constants representing the allowable values for ComputeHostConfigurationCheckDetailsRecycleLevelEnum
const (
	ComputeHostConfigurationCheckDetailsRecycleLevelSkipRecycle ComputeHostConfigurationCheckDetailsRecycleLevelEnum = "SKIP_RECYCLE"
	ComputeHostConfigurationCheckDetailsRecycleLevelFullRecycle ComputeHostConfigurationCheckDetailsRecycleLevelEnum = "FULL_RECYCLE"
)

var mappingComputeHostConfigurationCheckDetailsRecycleLevelEnum = map[string]ComputeHostConfigurationCheckDetailsRecycleLevelEnum{
	"SKIP_RECYCLE": ComputeHostConfigurationCheckDetailsRecycleLevelSkipRecycle,
	"FULL_RECYCLE": ComputeHostConfigurationCheckDetailsRecycleLevelFullRecycle,
}

var mappingComputeHostConfigurationCheckDetailsRecycleLevelEnumLowerCase = map[string]ComputeHostConfigurationCheckDetailsRecycleLevelEnum{
	"skip_recycle": ComputeHostConfigurationCheckDetailsRecycleLevelSkipRecycle,
	"full_recycle": ComputeHostConfigurationCheckDetailsRecycleLevelFullRecycle,
}

// GetComputeHostConfigurationCheckDetailsRecycleLevelEnumValues Enumerates the set of values for ComputeHostConfigurationCheckDetailsRecycleLevelEnum
func GetComputeHostConfigurationCheckDetailsRecycleLevelEnumValues() []ComputeHostConfigurationCheckDetailsRecycleLevelEnum {
	values := make([]ComputeHostConfigurationCheckDetailsRecycleLevelEnum, 0)
	for _, v := range mappingComputeHostConfigurationCheckDetailsRecycleLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeHostConfigurationCheckDetailsRecycleLevelEnumStringValues Enumerates the set of values in String for ComputeHostConfigurationCheckDetailsRecycleLevelEnum
func GetComputeHostConfigurationCheckDetailsRecycleLevelEnumStringValues() []string {
	return []string{
		"SKIP_RECYCLE",
		"FULL_RECYCLE",
	}
}

// GetMappingComputeHostConfigurationCheckDetailsRecycleLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeHostConfigurationCheckDetailsRecycleLevelEnum(val string) (ComputeHostConfigurationCheckDetailsRecycleLevelEnum, bool) {
	enum, ok := mappingComputeHostConfigurationCheckDetailsRecycleLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
