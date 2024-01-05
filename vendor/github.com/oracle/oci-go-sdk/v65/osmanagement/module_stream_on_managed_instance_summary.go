// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModuleStreamOnManagedInstanceSummary Summary information pertaining to a module stream on a managed instance
type ModuleStreamOnManagedInstanceSummary struct {

	// The name of the module that contains the stream.
	ModuleName *string `mandatory:"true" json:"moduleName"`

	// The name of the stream.
	StreamName *string `mandatory:"true" json:"streamName"`

	// The status of the stream
	// A stream with the "ENABLED" status can be used as a source for installing
	// profiles.  Streams with this status are also "ACTIVE".
	// A stream with the "DISABLED" status cannot be the source for installing
	// profiles.  To install profiles and packages from this stream, it must be
	// enabled.
	// A stream with the "ACTIVE" status can be used as a source for installing
	// profiles.  The packages that comprise the stream are also used when a
	// matching package is installed directly.  In general, a stream can have
	// this status if it is the default stream for the module and no stream has
	// been explicitly enabled.
	Status ModuleStreamOnManagedInstanceSummaryStatusEnum `mandatory:"true" json:"status"`

	// The set of profiles that the module stream contains.
	Profiles []ModuleStreamProfileOnManagedInstanceSummary `mandatory:"false" json:"profiles"`

	// The OCID of the software source that provides this module stream.
	SoftwareSourceId *string `mandatory:"false" json:"softwareSourceId"`

	// The date and time of the last status change for this profile, as
	// described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`
}

func (m ModuleStreamOnManagedInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModuleStreamOnManagedInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModuleStreamOnManagedInstanceSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetModuleStreamOnManagedInstanceSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModuleStreamOnManagedInstanceSummaryStatusEnum Enum with underlying type: string
type ModuleStreamOnManagedInstanceSummaryStatusEnum string

// Set of constants representing the allowable values for ModuleStreamOnManagedInstanceSummaryStatusEnum
const (
	ModuleStreamOnManagedInstanceSummaryStatusEnabled  ModuleStreamOnManagedInstanceSummaryStatusEnum = "ENABLED"
	ModuleStreamOnManagedInstanceSummaryStatusDisabled ModuleStreamOnManagedInstanceSummaryStatusEnum = "DISABLED"
	ModuleStreamOnManagedInstanceSummaryStatusActive   ModuleStreamOnManagedInstanceSummaryStatusEnum = "ACTIVE"
)

var mappingModuleStreamOnManagedInstanceSummaryStatusEnum = map[string]ModuleStreamOnManagedInstanceSummaryStatusEnum{
	"ENABLED":  ModuleStreamOnManagedInstanceSummaryStatusEnabled,
	"DISABLED": ModuleStreamOnManagedInstanceSummaryStatusDisabled,
	"ACTIVE":   ModuleStreamOnManagedInstanceSummaryStatusActive,
}

var mappingModuleStreamOnManagedInstanceSummaryStatusEnumLowerCase = map[string]ModuleStreamOnManagedInstanceSummaryStatusEnum{
	"enabled":  ModuleStreamOnManagedInstanceSummaryStatusEnabled,
	"disabled": ModuleStreamOnManagedInstanceSummaryStatusDisabled,
	"active":   ModuleStreamOnManagedInstanceSummaryStatusActive,
}

// GetModuleStreamOnManagedInstanceSummaryStatusEnumValues Enumerates the set of values for ModuleStreamOnManagedInstanceSummaryStatusEnum
func GetModuleStreamOnManagedInstanceSummaryStatusEnumValues() []ModuleStreamOnManagedInstanceSummaryStatusEnum {
	values := make([]ModuleStreamOnManagedInstanceSummaryStatusEnum, 0)
	for _, v := range mappingModuleStreamOnManagedInstanceSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetModuleStreamOnManagedInstanceSummaryStatusEnumStringValues Enumerates the set of values in String for ModuleStreamOnManagedInstanceSummaryStatusEnum
func GetModuleStreamOnManagedInstanceSummaryStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"ACTIVE",
	}
}

// GetMappingModuleStreamOnManagedInstanceSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModuleStreamOnManagedInstanceSummaryStatusEnum(val string) (ModuleStreamOnManagedInstanceSummaryStatusEnum, bool) {
	enum, ok := mappingModuleStreamOnManagedInstanceSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
