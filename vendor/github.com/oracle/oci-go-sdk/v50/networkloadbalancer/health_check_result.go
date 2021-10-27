// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
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

var mappingHealthCheckResultHealthCheckStatus = map[string]HealthCheckResultHealthCheckStatusEnum{
	"OK":                      HealthCheckResultHealthCheckStatusOk,
	"INVALID_STATUS_CODE":     HealthCheckResultHealthCheckStatusInvalidStatusCode,
	"TIMED_OUT":               HealthCheckResultHealthCheckStatusTimedOut,
	"HEALTH_PAYLOAD_MISMATCH": HealthCheckResultHealthCheckStatusHealthPayloadMismatch,
	"CONNECT_FAILED":          HealthCheckResultHealthCheckStatusConnectFailed,
	"UNKNOWN":                 HealthCheckResultHealthCheckStatusUnknown,
}

// GetHealthCheckResultHealthCheckStatusEnumValues Enumerates the set of values for HealthCheckResultHealthCheckStatusEnum
func GetHealthCheckResultHealthCheckStatusEnumValues() []HealthCheckResultHealthCheckStatusEnum {
	values := make([]HealthCheckResultHealthCheckStatusEnum, 0)
	for _, v := range mappingHealthCheckResultHealthCheckStatus {
		values = append(values, v)
	}
	return values
}
