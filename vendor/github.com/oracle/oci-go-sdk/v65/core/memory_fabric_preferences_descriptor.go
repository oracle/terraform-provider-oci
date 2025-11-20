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

// MemoryFabricPreferencesDescriptor The preference object specified by customer. Contains customerDesiredFirmwareBundleId, fabricRecycleLevel.
type MemoryFabricPreferencesDescriptor struct {

	// The desired firmware bundle id on the GPU memory fabric.
	CustomerDesiredFirmwareBundleId *string `mandatory:"false" json:"customerDesiredFirmwareBundleId"`

	// The recycle level of GPU memory fabric.
	FabricRecycleLevel MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum `mandatory:"false" json:"fabricRecycleLevel,omitempty"`
}

func (m MemoryFabricPreferencesDescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MemoryFabricPreferencesDescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMemoryFabricPreferencesDescriptorFabricRecycleLevelEnum(string(m.FabricRecycleLevel)); !ok && m.FabricRecycleLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FabricRecycleLevel: %s. Supported values are: %s.", m.FabricRecycleLevel, strings.Join(GetMemoryFabricPreferencesDescriptorFabricRecycleLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum Enum with underlying type: string
type MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum string

// Set of constants representing the allowable values for MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum
const (
	MemoryFabricPreferencesDescriptorFabricRecycleLevelFullRecycle              MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum = "FULL_RECYCLE"
	MemoryFabricPreferencesDescriptorFabricRecycleLevelSkipRecycle              MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum = "SKIP_RECYCLE"
	MemoryFabricPreferencesDescriptorFabricRecycleLevelOpportunisticFullRecycle MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum = "OPPORTUNISTIC_FULL_RECYCLE"
)

var mappingMemoryFabricPreferencesDescriptorFabricRecycleLevelEnum = map[string]MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum{
	"FULL_RECYCLE":               MemoryFabricPreferencesDescriptorFabricRecycleLevelFullRecycle,
	"SKIP_RECYCLE":               MemoryFabricPreferencesDescriptorFabricRecycleLevelSkipRecycle,
	"OPPORTUNISTIC_FULL_RECYCLE": MemoryFabricPreferencesDescriptorFabricRecycleLevelOpportunisticFullRecycle,
}

var mappingMemoryFabricPreferencesDescriptorFabricRecycleLevelEnumLowerCase = map[string]MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum{
	"full_recycle":               MemoryFabricPreferencesDescriptorFabricRecycleLevelFullRecycle,
	"skip_recycle":               MemoryFabricPreferencesDescriptorFabricRecycleLevelSkipRecycle,
	"opportunistic_full_recycle": MemoryFabricPreferencesDescriptorFabricRecycleLevelOpportunisticFullRecycle,
}

// GetMemoryFabricPreferencesDescriptorFabricRecycleLevelEnumValues Enumerates the set of values for MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum
func GetMemoryFabricPreferencesDescriptorFabricRecycleLevelEnumValues() []MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum {
	values := make([]MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum, 0)
	for _, v := range mappingMemoryFabricPreferencesDescriptorFabricRecycleLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetMemoryFabricPreferencesDescriptorFabricRecycleLevelEnumStringValues Enumerates the set of values in String for MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum
func GetMemoryFabricPreferencesDescriptorFabricRecycleLevelEnumStringValues() []string {
	return []string{
		"FULL_RECYCLE",
		"SKIP_RECYCLE",
		"OPPORTUNISTIC_FULL_RECYCLE",
	}
}

// GetMappingMemoryFabricPreferencesDescriptorFabricRecycleLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMemoryFabricPreferencesDescriptorFabricRecycleLevelEnum(val string) (MemoryFabricPreferencesDescriptorFabricRecycleLevelEnum, bool) {
	enum, ok := mappingMemoryFabricPreferencesDescriptorFabricRecycleLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
