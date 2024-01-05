// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Quota Quota policy for a usage plan.
type Quota struct {

	// The number of requests that can be made per time period.
	Value *int `mandatory:"true" json:"value"`

	// The unit of time over which quotas are calculated.
	// Example: `MINUTE` or `MONTH`
	Unit QuotaUnitEnum `mandatory:"true" json:"unit"`

	// The policy that controls when quotas will reset.
	// Example: `CALENDAR`
	ResetPolicy QuotaResetPolicyEnum `mandatory:"true" json:"resetPolicy"`

	// What the usage plan will do when a quota is breached:
	// `REJECT` will allow no further requests
	// `ALLOW` will continue to allow further requests
	OperationOnBreach QuotaOperationOnBreachEnum `mandatory:"true" json:"operationOnBreach"`
}

func (m Quota) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Quota) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQuotaUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetQuotaUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingQuotaResetPolicyEnum(string(m.ResetPolicy)); !ok && m.ResetPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResetPolicy: %s. Supported values are: %s.", m.ResetPolicy, strings.Join(GetQuotaResetPolicyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingQuotaOperationOnBreachEnum(string(m.OperationOnBreach)); !ok && m.OperationOnBreach != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationOnBreach: %s. Supported values are: %s.", m.OperationOnBreach, strings.Join(GetQuotaOperationOnBreachEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QuotaUnitEnum Enum with underlying type: string
type QuotaUnitEnum string

// Set of constants representing the allowable values for QuotaUnitEnum
const (
	QuotaUnitMinute QuotaUnitEnum = "MINUTE"
	QuotaUnitHour   QuotaUnitEnum = "HOUR"
	QuotaUnitDay    QuotaUnitEnum = "DAY"
	QuotaUnitWeek   QuotaUnitEnum = "WEEK"
	QuotaUnitMonth  QuotaUnitEnum = "MONTH"
)

var mappingQuotaUnitEnum = map[string]QuotaUnitEnum{
	"MINUTE": QuotaUnitMinute,
	"HOUR":   QuotaUnitHour,
	"DAY":    QuotaUnitDay,
	"WEEK":   QuotaUnitWeek,
	"MONTH":  QuotaUnitMonth,
}

var mappingQuotaUnitEnumLowerCase = map[string]QuotaUnitEnum{
	"minute": QuotaUnitMinute,
	"hour":   QuotaUnitHour,
	"day":    QuotaUnitDay,
	"week":   QuotaUnitWeek,
	"month":  QuotaUnitMonth,
}

// GetQuotaUnitEnumValues Enumerates the set of values for QuotaUnitEnum
func GetQuotaUnitEnumValues() []QuotaUnitEnum {
	values := make([]QuotaUnitEnum, 0)
	for _, v := range mappingQuotaUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetQuotaUnitEnumStringValues Enumerates the set of values in String for QuotaUnitEnum
func GetQuotaUnitEnumStringValues() []string {
	return []string{
		"MINUTE",
		"HOUR",
		"DAY",
		"WEEK",
		"MONTH",
	}
}

// GetMappingQuotaUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuotaUnitEnum(val string) (QuotaUnitEnum, bool) {
	enum, ok := mappingQuotaUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// QuotaResetPolicyEnum Enum with underlying type: string
type QuotaResetPolicyEnum string

// Set of constants representing the allowable values for QuotaResetPolicyEnum
const (
	QuotaResetPolicyCalendar QuotaResetPolicyEnum = "CALENDAR"
)

var mappingQuotaResetPolicyEnum = map[string]QuotaResetPolicyEnum{
	"CALENDAR": QuotaResetPolicyCalendar,
}

var mappingQuotaResetPolicyEnumLowerCase = map[string]QuotaResetPolicyEnum{
	"calendar": QuotaResetPolicyCalendar,
}

// GetQuotaResetPolicyEnumValues Enumerates the set of values for QuotaResetPolicyEnum
func GetQuotaResetPolicyEnumValues() []QuotaResetPolicyEnum {
	values := make([]QuotaResetPolicyEnum, 0)
	for _, v := range mappingQuotaResetPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetQuotaResetPolicyEnumStringValues Enumerates the set of values in String for QuotaResetPolicyEnum
func GetQuotaResetPolicyEnumStringValues() []string {
	return []string{
		"CALENDAR",
	}
}

// GetMappingQuotaResetPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuotaResetPolicyEnum(val string) (QuotaResetPolicyEnum, bool) {
	enum, ok := mappingQuotaResetPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// QuotaOperationOnBreachEnum Enum with underlying type: string
type QuotaOperationOnBreachEnum string

// Set of constants representing the allowable values for QuotaOperationOnBreachEnum
const (
	QuotaOperationOnBreachReject QuotaOperationOnBreachEnum = "REJECT"
	QuotaOperationOnBreachAllow  QuotaOperationOnBreachEnum = "ALLOW"
)

var mappingQuotaOperationOnBreachEnum = map[string]QuotaOperationOnBreachEnum{
	"REJECT": QuotaOperationOnBreachReject,
	"ALLOW":  QuotaOperationOnBreachAllow,
}

var mappingQuotaOperationOnBreachEnumLowerCase = map[string]QuotaOperationOnBreachEnum{
	"reject": QuotaOperationOnBreachReject,
	"allow":  QuotaOperationOnBreachAllow,
}

// GetQuotaOperationOnBreachEnumValues Enumerates the set of values for QuotaOperationOnBreachEnum
func GetQuotaOperationOnBreachEnumValues() []QuotaOperationOnBreachEnum {
	values := make([]QuotaOperationOnBreachEnum, 0)
	for _, v := range mappingQuotaOperationOnBreachEnum {
		values = append(values, v)
	}
	return values
}

// GetQuotaOperationOnBreachEnumStringValues Enumerates the set of values in String for QuotaOperationOnBreachEnum
func GetQuotaOperationOnBreachEnumStringValues() []string {
	return []string{
		"REJECT",
		"ALLOW",
	}
}

// GetMappingQuotaOperationOnBreachEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuotaOperationOnBreachEnum(val string) (QuotaOperationOnBreachEnum, bool) {
	enum, ok := mappingQuotaOperationOnBreachEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
