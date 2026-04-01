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

// OlvmVnicProfileProperties OLVM Vnic Profile properties.
type OlvmVnicProfileProperties struct {

	// A human-readable name in plain text
	VnicProfileName *string `mandatory:"true" json:"vnicProfileName"`

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	// Custom properties applied to the vNIC profile.
	CustomProperties []OlvmCustomProperty `mandatory:"false" json:"customProperties"`

	// A human-readable description in plain text.
	Description *string `mandatory:"false" json:"description"`

	// Indicates whether passThrough NIC is migratable or not.
	IsMigratable *bool `mandatory:"false" json:"isMigratable"`

	// Describes whether the vNIC is to be implemented as a pass-through device or a virtual one.
	PassThrough OlvmVnicProfilePropertiesPassThroughEnum `mandatory:"false" json:"passThrough,omitempty"`

	// Indicates if port mirroring is enabled.
	IsPortMirroring *bool `mandatory:"false" json:"isPortMirroring"`
}

func (m OlvmVnicProfileProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmVnicProfileProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmVnicProfilePropertiesPassThroughEnum(string(m.PassThrough)); !ok && m.PassThrough != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PassThrough: %s. Supported values are: %s.", m.PassThrough, strings.Join(GetOlvmVnicProfilePropertiesPassThroughEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmVnicProfilePropertiesPassThroughEnum Enum with underlying type: string
type OlvmVnicProfilePropertiesPassThroughEnum string

// Set of constants representing the allowable values for OlvmVnicProfilePropertiesPassThroughEnum
const (
	OlvmVnicProfilePropertiesPassThroughDisabled OlvmVnicProfilePropertiesPassThroughEnum = "DISABLED"
	OlvmVnicProfilePropertiesPassThroughEnabled  OlvmVnicProfilePropertiesPassThroughEnum = "ENABLED"
)

var mappingOlvmVnicProfilePropertiesPassThroughEnum = map[string]OlvmVnicProfilePropertiesPassThroughEnum{
	"DISABLED": OlvmVnicProfilePropertiesPassThroughDisabled,
	"ENABLED":  OlvmVnicProfilePropertiesPassThroughEnabled,
}

var mappingOlvmVnicProfilePropertiesPassThroughEnumLowerCase = map[string]OlvmVnicProfilePropertiesPassThroughEnum{
	"disabled": OlvmVnicProfilePropertiesPassThroughDisabled,
	"enabled":  OlvmVnicProfilePropertiesPassThroughEnabled,
}

// GetOlvmVnicProfilePropertiesPassThroughEnumValues Enumerates the set of values for OlvmVnicProfilePropertiesPassThroughEnum
func GetOlvmVnicProfilePropertiesPassThroughEnumValues() []OlvmVnicProfilePropertiesPassThroughEnum {
	values := make([]OlvmVnicProfilePropertiesPassThroughEnum, 0)
	for _, v := range mappingOlvmVnicProfilePropertiesPassThroughEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmVnicProfilePropertiesPassThroughEnumStringValues Enumerates the set of values in String for OlvmVnicProfilePropertiesPassThroughEnum
func GetOlvmVnicProfilePropertiesPassThroughEnumStringValues() []string {
	return []string{
		"DISABLED",
		"ENABLED",
	}
}

// GetMappingOlvmVnicProfilePropertiesPassThroughEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmVnicProfilePropertiesPassThroughEnum(val string) (OlvmVnicProfilePropertiesPassThroughEnum, bool) {
	enum, ok := mappingOlvmVnicProfilePropertiesPassThroughEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
