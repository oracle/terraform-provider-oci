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

// OlvmSerialNumber Serial number object in OLVM
type OlvmSerialNumber struct {

	// Type representing the policy of a Serial Number
	Policy OlvmSerialNumberPolicyEnum `mandatory:"false" json:"policy,omitempty"`

	// Value of the serial number policy.
	Value *string `mandatory:"false" json:"value"`
}

func (m OlvmSerialNumber) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmSerialNumber) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmSerialNumberPolicyEnum(string(m.Policy)); !ok && m.Policy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Policy: %s. Supported values are: %s.", m.Policy, strings.Join(GetOlvmSerialNumberPolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmSerialNumberPolicyEnum Enum with underlying type: string
type OlvmSerialNumberPolicyEnum string

// Set of constants representing the allowable values for OlvmSerialNumberPolicyEnum
const (
	OlvmSerialNumberPolicyCustom OlvmSerialNumberPolicyEnum = "CUSTOM"
	OlvmSerialNumberPolicyHost   OlvmSerialNumberPolicyEnum = "HOST"
	OlvmSerialNumberPolicyNone   OlvmSerialNumberPolicyEnum = "NONE"
	OlvmSerialNumberPolicyVm     OlvmSerialNumberPolicyEnum = "VM"
)

var mappingOlvmSerialNumberPolicyEnum = map[string]OlvmSerialNumberPolicyEnum{
	"CUSTOM": OlvmSerialNumberPolicyCustom,
	"HOST":   OlvmSerialNumberPolicyHost,
	"NONE":   OlvmSerialNumberPolicyNone,
	"VM":     OlvmSerialNumberPolicyVm,
}

var mappingOlvmSerialNumberPolicyEnumLowerCase = map[string]OlvmSerialNumberPolicyEnum{
	"custom": OlvmSerialNumberPolicyCustom,
	"host":   OlvmSerialNumberPolicyHost,
	"none":   OlvmSerialNumberPolicyNone,
	"vm":     OlvmSerialNumberPolicyVm,
}

// GetOlvmSerialNumberPolicyEnumValues Enumerates the set of values for OlvmSerialNumberPolicyEnum
func GetOlvmSerialNumberPolicyEnumValues() []OlvmSerialNumberPolicyEnum {
	values := make([]OlvmSerialNumberPolicyEnum, 0)
	for _, v := range mappingOlvmSerialNumberPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmSerialNumberPolicyEnumStringValues Enumerates the set of values in String for OlvmSerialNumberPolicyEnum
func GetOlvmSerialNumberPolicyEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"HOST",
		"NONE",
		"VM",
	}
}

// GetMappingOlvmSerialNumberPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmSerialNumberPolicyEnum(val string) (OlvmSerialNumberPolicyEnum, bool) {
	enum, ok := mappingOlvmSerialNumberPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
