// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CompartmentMetadata The representation of CompartmentMetadata
type CompartmentMetadata struct {

	// The compartment id.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The access level.
	AccessLevel CompartmentMetadataAccessLevelEnum `mandatory:"true" json:"accessLevel"`
}

func (m CompartmentMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompartmentMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCompartmentMetadataAccessLevelEnum(string(m.AccessLevel)); !ok && m.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", m.AccessLevel, strings.Join(GetCompartmentMetadataAccessLevelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CompartmentMetadataAccessLevelEnum Enum with underlying type: string
type CompartmentMetadataAccessLevelEnum string

// Set of constants representing the allowable values for CompartmentMetadataAccessLevelEnum
const (
	CompartmentMetadataAccessLevelAccessible   CompartmentMetadataAccessLevelEnum = "accessible"
	CompartmentMetadataAccessLevelVisible      CompartmentMetadataAccessLevelEnum = "visible"
	CompartmentMetadataAccessLevelInaccessible CompartmentMetadataAccessLevelEnum = "inaccessible"
)

var mappingCompartmentMetadataAccessLevelEnum = map[string]CompartmentMetadataAccessLevelEnum{
	"accessible":   CompartmentMetadataAccessLevelAccessible,
	"visible":      CompartmentMetadataAccessLevelVisible,
	"inaccessible": CompartmentMetadataAccessLevelInaccessible,
}

// GetCompartmentMetadataAccessLevelEnumValues Enumerates the set of values for CompartmentMetadataAccessLevelEnum
func GetCompartmentMetadataAccessLevelEnumValues() []CompartmentMetadataAccessLevelEnum {
	values := make([]CompartmentMetadataAccessLevelEnum, 0)
	for _, v := range mappingCompartmentMetadataAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetCompartmentMetadataAccessLevelEnumStringValues Enumerates the set of values in String for CompartmentMetadataAccessLevelEnum
func GetCompartmentMetadataAccessLevelEnumStringValues() []string {
	return []string{
		"accessible",
		"visible",
		"inaccessible",
	}
}

// GetMappingCompartmentMetadataAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompartmentMetadataAccessLevelEnum(val string) (CompartmentMetadataAccessLevelEnum, bool) {
	mappingCompartmentMetadataAccessLevelEnumIgnoreCase := make(map[string]CompartmentMetadataAccessLevelEnum)
	for k, v := range mappingCompartmentMetadataAccessLevelEnum {
		mappingCompartmentMetadataAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCompartmentMetadataAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
