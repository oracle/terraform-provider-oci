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

// HealthCheckResult Information about a single backend server health check result reported by a network load balancer.
type HealthCheckResult struct {

	// The date and time the data was retrieved, in the format defined by RFC3339.
	// Example: `2020-05-01T18:28:11+00:00`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// The result of the most recent health check.
	HealthCheckStatus HealthCheckResultHealthCheckStatusEnum `mandatory:"true" json:"healthCheckStatus"`
}

func (m HealthCheckResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HealthCheckResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHealthCheckResultHealthCheckStatusEnum(string(m.HealthCheckStatus)); !ok && m.HealthCheckStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HealthCheckStatus: %s. Supported values are: %s.", m.HealthCheckStatus, strings.Join(GetHealthCheckResultHealthCheckStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HealthCheckResultHealthCheckStatusEnum Enum with underlying type: string
type HealthCheckResultHealthCheckStatusEnum string

// Set of constants representing the allowable values for HealthCheckResultHealthCheckStatusEnum
const (
	HealthCheckResultHealthCheckStatusOk                    HealthCheckResultHealthCheckStatusEnum = "OK"
	HealthCheckResultHealthCheckStatusInvalidStatusCode     HealthCheckResultHealthCheckStatusEnum = "INVALID_STATUS_CODE"
	HealthCheckResultHealthCheckStatusTimedOut              HealthCheckResultHealthCheckStatusEnum = "TIMED_OUT"
	HealthCheckResultHealthCheckStatusHealthPayloadMismatch HealthCheckResultHealthCheckStatusEnum = "HEALTH_PAYLOAD_MISMATCH"
	HealthCheckResultHealthCheckStatusConnectFailed         HealthCheckResultHealthCheckStatusEnum = "CONNECT_FAILED"
	HealthCheckResultHealthCheckStatusUnknown               HealthCheckResultHealthCheckStatusEnum = "UNKNOWN"
)

var mappingHealthCheckResultHealthCheckStatusEnum = map[string]HealthCheckResultHealthCheckStatusEnum{
	"OK":                      HealthCheckResultHealthCheckStatusOk,
	"INVALID_STATUS_CODE":     HealthCheckResultHealthCheckStatusInvalidStatusCode,
	"TIMED_OUT":               HealthCheckResultHealthCheckStatusTimedOut,
	"HEALTH_PAYLOAD_MISMATCH": HealthCheckResultHealthCheckStatusHealthPayloadMismatch,
	"CONNECT_FAILED":          HealthCheckResultHealthCheckStatusConnectFailed,
	"UNKNOWN":                 HealthCheckResultHealthCheckStatusUnknown,
}

var mappingHealthCheckResultHealthCheckStatusEnumLowerCase = map[string]HealthCheckResultHealthCheckStatusEnum{
	"ok":                      HealthCheckResultHealthCheckStatusOk,
	"invalid_status_code":     HealthCheckResultHealthCheckStatusInvalidStatusCode,
	"timed_out":               HealthCheckResultHealthCheckStatusTimedOut,
	"health_payload_mismatch": HealthCheckResultHealthCheckStatusHealthPayloadMismatch,
	"connect_failed":          HealthCheckResultHealthCheckStatusConnectFailed,
	"unknown":                 HealthCheckResultHealthCheckStatusUnknown,
}

// GetHealthCheckResultHealthCheckStatusEnumValues Enumerates the set of values for HealthCheckResultHealthCheckStatusEnum
func GetHealthCheckResultHealthCheckStatusEnumValues() []HealthCheckResultHealthCheckStatusEnum {
	values := make([]HealthCheckResultHealthCheckStatusEnum, 0)
	for _, v := range mappingHealthCheckResultHealthCheckStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthCheckResultHealthCheckStatusEnumStringValues Enumerates the set of values in String for HealthCheckResultHealthCheckStatusEnum
func GetHealthCheckResultHealthCheckStatusEnumStringValues() []string {
	return []string{
		"OK",
		"INVALID_STATUS_CODE",
		"TIMED_OUT",
		"HEALTH_PAYLOAD_MISMATCH",
		"CONNECT_FAILED",
		"UNKNOWN",
	}
}

// GetMappingHealthCheckResultHealthCheckStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthCheckResultHealthCheckStatusEnum(val string) (HealthCheckResultHealthCheckStatusEnum, bool) {
	enum, ok := mappingHealthCheckResultHealthCheckStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
