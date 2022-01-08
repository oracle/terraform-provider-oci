// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"strings"
)

// ShapeMemoryEncryptionTechnologyOptions Configuration options for the underlying memory encryption technology.
type ShapeMemoryEncryptionTechnologyOptions struct {

	// The supported values for this platform configuration property.
	AllowedValues []ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum `mandatory:"false" json:"allowedValues,omitempty"`

	// The default memory encryption technology.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`
}

func (m ShapeMemoryEncryptionTechnologyOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShapeMemoryEncryptionTechnologyOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.AllowedValues {
		if _, ok := mappingShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum[string(val)]; !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AllowedValues: %s. Supported values are: %s.", val, strings.Join(GetShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum Enum with underlying type: string
type ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum string

// Set of constants representing the allowable values for ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum
const (
	ShapeMemoryEncryptionTechnologyOptionsAllowedValuesTsme ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum = "TSME"
	ShapeMemoryEncryptionTechnologyOptionsAllowedValuesSmee ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum = "SMEE"
)

var mappingShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum = map[string]ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum{
	"TSME": ShapeMemoryEncryptionTechnologyOptionsAllowedValuesTsme,
	"SMEE": ShapeMemoryEncryptionTechnologyOptionsAllowedValuesSmee,
}

// GetShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnumValues Enumerates the set of values for ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum
func GetShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnumValues() []ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum {
	values := make([]ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum, 0)
	for _, v := range mappingShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum {
		values = append(values, v)
	}
	return values
}

// GetShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnumStringValues Enumerates the set of values in String for ShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnum
func GetShapeMemoryEncryptionTechnologyOptionsAllowedValuesEnumStringValues() []string {
	return []string{
		"TSME",
		"SMEE",
	}
}
