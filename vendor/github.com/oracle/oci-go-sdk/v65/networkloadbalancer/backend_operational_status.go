// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackendOperationalStatus The operational status of the specified backend server.
type BackendOperationalStatus struct {

	// The operational status.
	Status BackendOperationalStatusStatusEnum `mandatory:"true" json:"status"`
}

func (m BackendOperationalStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackendOperationalStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackendOperationalStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetBackendOperationalStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackendOperationalStatusStatusEnum Enum with underlying type: string
type BackendOperationalStatusStatusEnum string

// Set of constants representing the allowable values for BackendOperationalStatusStatusEnum
const (
	BackendOperationalStatusStatusActive  BackendOperationalStatusStatusEnum = "ACTIVE"
	BackendOperationalStatusStatusStandby BackendOperationalStatusStatusEnum = "STANDBY"
	BackendOperationalStatusStatusUnknown BackendOperationalStatusStatusEnum = "UNKNOWN"
)

var mappingBackendOperationalStatusStatusEnum = map[string]BackendOperationalStatusStatusEnum{
	"ACTIVE":  BackendOperationalStatusStatusActive,
	"STANDBY": BackendOperationalStatusStatusStandby,
	"UNKNOWN": BackendOperationalStatusStatusUnknown,
}

var mappingBackendOperationalStatusStatusEnumLowerCase = map[string]BackendOperationalStatusStatusEnum{
	"active":  BackendOperationalStatusStatusActive,
	"standby": BackendOperationalStatusStatusStandby,
	"unknown": BackendOperationalStatusStatusUnknown,
}

// GetBackendOperationalStatusStatusEnumValues Enumerates the set of values for BackendOperationalStatusStatusEnum
func GetBackendOperationalStatusStatusEnumValues() []BackendOperationalStatusStatusEnum {
	values := make([]BackendOperationalStatusStatusEnum, 0)
	for _, v := range mappingBackendOperationalStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBackendOperationalStatusStatusEnumStringValues Enumerates the set of values in String for BackendOperationalStatusStatusEnum
func GetBackendOperationalStatusStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"STANDBY",
		"UNKNOWN",
	}
}

// GetMappingBackendOperationalStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackendOperationalStatusStatusEnum(val string) (BackendOperationalStatusStatusEnum, bool) {
	enum, ok := mappingBackendOperationalStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
