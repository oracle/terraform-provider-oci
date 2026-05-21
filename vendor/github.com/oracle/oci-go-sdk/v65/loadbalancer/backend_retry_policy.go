// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackendRetryPolicy The representation of BackendRetryPolicy
type BackendRetryPolicy struct {

	// The backend retry policy supports the following modes:
	// "OFF" — If the initial connection attempt to the first backend fails, no retries are attempted on other backends.
	// "CUSTOM" — Retries are attempted on the number of backends specified in the limit field, until a successful connection is established.
	// "ALL" (default) — Retries are attempted on all available backends until a successful connection is established.
	Mode BackendRetryPolicyModeEnum `mandatory:"true" json:"mode"`

	// When the mode is "Custom", use this field to specify how many backends the Load Balancer will try to establish
	// a successful connection with.
	// The minimum value is 1, which means that if the initial connection attempt to the first backend fails,
	// no retries are attempted on other backends. This is the same behavior as setting the mode to "Off".
	// The maximum value is the total number of backends.
	// ** Notes **
	// * If you set the limit to the total number of backends and later add another backend, you will need to increase
	// the limit to ensure retries cover all backends. Alternatively, you can set the mode to All.
	// * The limit field should only be set when the mode is "Custom". You will receive a 400 error otherwise.
	// Example: `3`
	Limit *int `mandatory:"false" json:"limit"`
}

func (m BackendRetryPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackendRetryPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackendRetryPolicyModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetBackendRetryPolicyModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackendRetryPolicyModeEnum Enum with underlying type: string
type BackendRetryPolicyModeEnum string

// Set of constants representing the allowable values for BackendRetryPolicyModeEnum
const (
	BackendRetryPolicyModeOff    BackendRetryPolicyModeEnum = "OFF"
	BackendRetryPolicyModeCustom BackendRetryPolicyModeEnum = "CUSTOM"
	BackendRetryPolicyModeAll    BackendRetryPolicyModeEnum = "ALL"
)

var mappingBackendRetryPolicyModeEnum = map[string]BackendRetryPolicyModeEnum{
	"OFF":    BackendRetryPolicyModeOff,
	"CUSTOM": BackendRetryPolicyModeCustom,
	"ALL":    BackendRetryPolicyModeAll,
}

var mappingBackendRetryPolicyModeEnumLowerCase = map[string]BackendRetryPolicyModeEnum{
	"off":    BackendRetryPolicyModeOff,
	"custom": BackendRetryPolicyModeCustom,
	"all":    BackendRetryPolicyModeAll,
}

// GetBackendRetryPolicyModeEnumValues Enumerates the set of values for BackendRetryPolicyModeEnum
func GetBackendRetryPolicyModeEnumValues() []BackendRetryPolicyModeEnum {
	values := make([]BackendRetryPolicyModeEnum, 0)
	for _, v := range mappingBackendRetryPolicyModeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackendRetryPolicyModeEnumStringValues Enumerates the set of values in String for BackendRetryPolicyModeEnum
func GetBackendRetryPolicyModeEnumStringValues() []string {
	return []string{
		"OFF",
		"CUSTOM",
		"ALL",
	}
}

// GetMappingBackendRetryPolicyModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackendRetryPolicyModeEnum(val string) (BackendRetryPolicyModeEnum, bool) {
	enum, ok := mappingBackendRetryPolicyModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
