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

// UpdateModuleStreamDetails Information detailing the state of a module stream
type UpdateModuleStreamDetails struct {

	// The name of the stream of the parent module
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
	Status UpdateModuleStreamDetailsStatusEnum `mandatory:"true" json:"status"`

	// The date and time of the last status change for this object, as
	// described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeModified *common.SDKTime `mandatory:"true" json:"timeModified"`

	// The name of the software source that publishes this stream.
	SoftwareSourceName *string `mandatory:"false" json:"softwareSourceName"`

	// The URL of the software source that publishes this stream.
	SoftwareSourceUrl *string `mandatory:"false" json:"softwareSourceUrl"`

	// Indicates if the module stream is the default
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// The profiles of the stream
	Profiles []UpdateModuleStreamProfileDetails `mandatory:"false" json:"profiles"`
}

func (m UpdateModuleStreamDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateModuleStreamDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateModuleStreamDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateModuleStreamDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateModuleStreamDetailsStatusEnum Enum with underlying type: string
type UpdateModuleStreamDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateModuleStreamDetailsStatusEnum
const (
	UpdateModuleStreamDetailsStatusEnabled  UpdateModuleStreamDetailsStatusEnum = "ENABLED"
	UpdateModuleStreamDetailsStatusDisabled UpdateModuleStreamDetailsStatusEnum = "DISABLED"
	UpdateModuleStreamDetailsStatusActive   UpdateModuleStreamDetailsStatusEnum = "ACTIVE"
)

var mappingUpdateModuleStreamDetailsStatusEnum = map[string]UpdateModuleStreamDetailsStatusEnum{
	"ENABLED":  UpdateModuleStreamDetailsStatusEnabled,
	"DISABLED": UpdateModuleStreamDetailsStatusDisabled,
	"ACTIVE":   UpdateModuleStreamDetailsStatusActive,
}

var mappingUpdateModuleStreamDetailsStatusEnumLowerCase = map[string]UpdateModuleStreamDetailsStatusEnum{
	"enabled":  UpdateModuleStreamDetailsStatusEnabled,
	"disabled": UpdateModuleStreamDetailsStatusDisabled,
	"active":   UpdateModuleStreamDetailsStatusActive,
}

// GetUpdateModuleStreamDetailsStatusEnumValues Enumerates the set of values for UpdateModuleStreamDetailsStatusEnum
func GetUpdateModuleStreamDetailsStatusEnumValues() []UpdateModuleStreamDetailsStatusEnum {
	values := make([]UpdateModuleStreamDetailsStatusEnum, 0)
	for _, v := range mappingUpdateModuleStreamDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateModuleStreamDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateModuleStreamDetailsStatusEnum
func GetUpdateModuleStreamDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"ACTIVE",
	}
}

// GetMappingUpdateModuleStreamDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateModuleStreamDetailsStatusEnum(val string) (UpdateModuleStreamDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateModuleStreamDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
