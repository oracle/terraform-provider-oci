// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// CompartmentMetadataAccessLevelEnum Enum with underlying type: string
type CompartmentMetadataAccessLevelEnum string

// Set of constants representing the allowable values for CompartmentMetadataAccessLevelEnum
const (
	CompartmentMetadataAccessLevelAccessible   CompartmentMetadataAccessLevelEnum = "accessible"
	CompartmentMetadataAccessLevelVisible      CompartmentMetadataAccessLevelEnum = "visible"
	CompartmentMetadataAccessLevelInaccessible CompartmentMetadataAccessLevelEnum = "inaccessible"
)

var mappingCompartmentMetadataAccessLevel = map[string]CompartmentMetadataAccessLevelEnum{
	"accessible":   CompartmentMetadataAccessLevelAccessible,
	"visible":      CompartmentMetadataAccessLevelVisible,
	"inaccessible": CompartmentMetadataAccessLevelInaccessible,
}

// GetCompartmentMetadataAccessLevelEnumValues Enumerates the set of values for CompartmentMetadataAccessLevelEnum
func GetCompartmentMetadataAccessLevelEnumValues() []CompartmentMetadataAccessLevelEnum {
	values := make([]CompartmentMetadataAccessLevelEnum, 0)
	for _, v := range mappingCompartmentMetadataAccessLevel {
		values = append(values, v)
	}
	return values
}
