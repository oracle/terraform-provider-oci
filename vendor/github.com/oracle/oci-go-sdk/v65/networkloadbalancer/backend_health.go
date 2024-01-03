// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// BackendHealth The health status of the specified backend server.
type BackendHealth struct {

	// The general health status of the specified backend server.
	// *   **OK:**  All health check probes return `OK`
	// *   **WARNING:** At least one of the health check probes does not return `OK`
	// *   **CRITICAL:** None of the health check probes return `OK`.
	// *
	// *   **UNKNOWN:** One of the health checks probes return `UNKNOWN`,
	// *   or the system is unable to retrieve metrics at this time.
	Status BackendHealthStatusEnum `mandatory:"true" json:"status"`

	// A list of the most recent health check results returned for the specified backend server.
	HealthCheckResults []HealthCheckResult `mandatory:"true" json:"healthCheckResults"`
}

func (m BackendHealth) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackendHealth) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackendHealthStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetBackendHealthStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackendHealthStatusEnum Enum with underlying type: string
type BackendHealthStatusEnum string

// Set of constants representing the allowable values for BackendHealthStatusEnum
const (
	BackendHealthStatusOk       BackendHealthStatusEnum = "OK"
	BackendHealthStatusWarning  BackendHealthStatusEnum = "WARNING"
	BackendHealthStatusCritical BackendHealthStatusEnum = "CRITICAL"
	BackendHealthStatusUnknown  BackendHealthStatusEnum = "UNKNOWN"
)

var mappingBackendHealthStatusEnum = map[string]BackendHealthStatusEnum{
	"OK":       BackendHealthStatusOk,
	"WARNING":  BackendHealthStatusWarning,
	"CRITICAL": BackendHealthStatusCritical,
	"UNKNOWN":  BackendHealthStatusUnknown,
}

var mappingBackendHealthStatusEnumLowerCase = map[string]BackendHealthStatusEnum{
	"ok":       BackendHealthStatusOk,
	"warning":  BackendHealthStatusWarning,
	"critical": BackendHealthStatusCritical,
	"unknown":  BackendHealthStatusUnknown,
}

// GetBackendHealthStatusEnumValues Enumerates the set of values for BackendHealthStatusEnum
func GetBackendHealthStatusEnumValues() []BackendHealthStatusEnum {
	values := make([]BackendHealthStatusEnum, 0)
	for _, v := range mappingBackendHealthStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBackendHealthStatusEnumStringValues Enumerates the set of values in String for BackendHealthStatusEnum
func GetBackendHealthStatusEnumStringValues() []string {
	return []string{
		"OK",
		"WARNING",
		"CRITICAL",
		"UNKNOWN",
	}
}

// GetMappingBackendHealthStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackendHealthStatusEnum(val string) (BackendHealthStatusEnum, bool) {
	enum, ok := mappingBackendHealthStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
