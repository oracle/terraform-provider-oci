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

// HostGroupConfiguration Host group configuration
type HostGroupConfiguration struct {

	// Either the platform name or compute shape that the configuration is targeting
	Target *string `mandatory:"false" json:"target"`

	// The OCID for firmware bundle
	FirmwareBundleId *string `mandatory:"false" json:"firmwareBundleId"`

	// Preferred recycle level for hosts associated with the reservation config.
	// * `SKIP_RECYCLE` - Skips host wipe.
	// * `FULL_RECYCLE` - Does not skip host wipe. This is the default behavior.
	RecycleLevel HostGroupConfigurationRecycleLevelEnum `mandatory:"false" json:"recycleLevel,omitempty"`
}

func (m HostGroupConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostGroupConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHostGroupConfigurationRecycleLevelEnum(string(m.RecycleLevel)); !ok && m.RecycleLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecycleLevel: %s. Supported values are: %s.", m.RecycleLevel, strings.Join(GetHostGroupConfigurationRecycleLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostGroupConfigurationRecycleLevelEnum Enum with underlying type: string
type HostGroupConfigurationRecycleLevelEnum string

// Set of constants representing the allowable values for HostGroupConfigurationRecycleLevelEnum
const (
	HostGroupConfigurationRecycleLevelSkipRecycle HostGroupConfigurationRecycleLevelEnum = "SKIP_RECYCLE"
	HostGroupConfigurationRecycleLevelFullRecycle HostGroupConfigurationRecycleLevelEnum = "FULL_RECYCLE"
)

var mappingHostGroupConfigurationRecycleLevelEnum = map[string]HostGroupConfigurationRecycleLevelEnum{
	"SKIP_RECYCLE": HostGroupConfigurationRecycleLevelSkipRecycle,
	"FULL_RECYCLE": HostGroupConfigurationRecycleLevelFullRecycle,
}

var mappingHostGroupConfigurationRecycleLevelEnumLowerCase = map[string]HostGroupConfigurationRecycleLevelEnum{
	"skip_recycle": HostGroupConfigurationRecycleLevelSkipRecycle,
	"full_recycle": HostGroupConfigurationRecycleLevelFullRecycle,
}

// GetHostGroupConfigurationRecycleLevelEnumValues Enumerates the set of values for HostGroupConfigurationRecycleLevelEnum
func GetHostGroupConfigurationRecycleLevelEnumValues() []HostGroupConfigurationRecycleLevelEnum {
	values := make([]HostGroupConfigurationRecycleLevelEnum, 0)
	for _, v := range mappingHostGroupConfigurationRecycleLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetHostGroupConfigurationRecycleLevelEnumStringValues Enumerates the set of values in String for HostGroupConfigurationRecycleLevelEnum
func GetHostGroupConfigurationRecycleLevelEnumStringValues() []string {
	return []string{
		"SKIP_RECYCLE",
		"FULL_RECYCLE",
	}
}

// GetMappingHostGroupConfigurationRecycleLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostGroupConfigurationRecycleLevelEnum(val string) (HostGroupConfigurationRecycleLevelEnum, bool) {
	enum, ok := mappingHostGroupConfigurationRecycleLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
