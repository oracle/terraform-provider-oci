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

// OlvmVmType Determines whether the virtual machine is optimized for desktop or server.
type OlvmVmType struct {

	// Type representing what the virtual machine is optimized for.
	VmType OlvmVmTypeVmTypeEnum `mandatory:"false" json:"vmType,omitempty"`
}

func (m OlvmVmType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmVmType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmVmTypeVmTypeEnum(string(m.VmType)); !ok && m.VmType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VmType: %s. Supported values are: %s.", m.VmType, strings.Join(GetOlvmVmTypeVmTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmVmTypeVmTypeEnum Enum with underlying type: string
type OlvmVmTypeVmTypeEnum string

// Set of constants representing the allowable values for OlvmVmTypeVmTypeEnum
const (
	OlvmVmTypeVmTypeDesktop         OlvmVmTypeVmTypeEnum = "DESKTOP"
	OlvmVmTypeVmTypeHighPerformance OlvmVmTypeVmTypeEnum = "HIGH_PERFORMANCE"
	OlvmVmTypeVmTypeServer          OlvmVmTypeVmTypeEnum = "SERVER"
)

var mappingOlvmVmTypeVmTypeEnum = map[string]OlvmVmTypeVmTypeEnum{
	"DESKTOP":          OlvmVmTypeVmTypeDesktop,
	"HIGH_PERFORMANCE": OlvmVmTypeVmTypeHighPerformance,
	"SERVER":           OlvmVmTypeVmTypeServer,
}

var mappingOlvmVmTypeVmTypeEnumLowerCase = map[string]OlvmVmTypeVmTypeEnum{
	"desktop":          OlvmVmTypeVmTypeDesktop,
	"high_performance": OlvmVmTypeVmTypeHighPerformance,
	"server":           OlvmVmTypeVmTypeServer,
}

// GetOlvmVmTypeVmTypeEnumValues Enumerates the set of values for OlvmVmTypeVmTypeEnum
func GetOlvmVmTypeVmTypeEnumValues() []OlvmVmTypeVmTypeEnum {
	values := make([]OlvmVmTypeVmTypeEnum, 0)
	for _, v := range mappingOlvmVmTypeVmTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmVmTypeVmTypeEnumStringValues Enumerates the set of values in String for OlvmVmTypeVmTypeEnum
func GetOlvmVmTypeVmTypeEnumStringValues() []string {
	return []string{
		"DESKTOP",
		"HIGH_PERFORMANCE",
		"SERVER",
	}
}

// GetMappingOlvmVmTypeVmTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmVmTypeVmTypeEnum(val string) (OlvmVmTypeVmTypeEnum, bool) {
	enum, ok := mappingOlvmVmTypeVmTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
