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

// OlvmSelinux Represents SELinux in the system
type OlvmSelinux struct {

	// Represents an SELinux enforcement mode.
	SeLinuxMode OlvmSelinuxSeLinuxModeEnum `mandatory:"false" json:"seLinuxMode,omitempty"`
}

func (m OlvmSelinux) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmSelinux) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmSelinuxSeLinuxModeEnum(string(m.SeLinuxMode)); !ok && m.SeLinuxMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SeLinuxMode: %s. Supported values are: %s.", m.SeLinuxMode, strings.Join(GetOlvmSelinuxSeLinuxModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmSelinuxSeLinuxModeEnum Enum with underlying type: string
type OlvmSelinuxSeLinuxModeEnum string

// Set of constants representing the allowable values for OlvmSelinuxSeLinuxModeEnum
const (
	OlvmSelinuxSeLinuxModeDisabled   OlvmSelinuxSeLinuxModeEnum = "DISABLED"
	OlvmSelinuxSeLinuxModeEnforcing  OlvmSelinuxSeLinuxModeEnum = "ENFORCING"
	OlvmSelinuxSeLinuxModePermissive OlvmSelinuxSeLinuxModeEnum = "PERMISSIVE"
)

var mappingOlvmSelinuxSeLinuxModeEnum = map[string]OlvmSelinuxSeLinuxModeEnum{
	"DISABLED":   OlvmSelinuxSeLinuxModeDisabled,
	"ENFORCING":  OlvmSelinuxSeLinuxModeEnforcing,
	"PERMISSIVE": OlvmSelinuxSeLinuxModePermissive,
}

var mappingOlvmSelinuxSeLinuxModeEnumLowerCase = map[string]OlvmSelinuxSeLinuxModeEnum{
	"disabled":   OlvmSelinuxSeLinuxModeDisabled,
	"enforcing":  OlvmSelinuxSeLinuxModeEnforcing,
	"permissive": OlvmSelinuxSeLinuxModePermissive,
}

// GetOlvmSelinuxSeLinuxModeEnumValues Enumerates the set of values for OlvmSelinuxSeLinuxModeEnum
func GetOlvmSelinuxSeLinuxModeEnumValues() []OlvmSelinuxSeLinuxModeEnum {
	values := make([]OlvmSelinuxSeLinuxModeEnum, 0)
	for _, v := range mappingOlvmSelinuxSeLinuxModeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmSelinuxSeLinuxModeEnumStringValues Enumerates the set of values in String for OlvmSelinuxSeLinuxModeEnum
func GetOlvmSelinuxSeLinuxModeEnumStringValues() []string {
	return []string{
		"DISABLED",
		"ENFORCING",
		"PERMISSIVE",
	}
}

// GetMappingOlvmSelinuxSeLinuxModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmSelinuxSeLinuxModeEnum(val string) (OlvmSelinuxSeLinuxModeEnum, bool) {
	enum, ok := mappingOlvmSelinuxSeLinuxModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
