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

// RateLimit Rate-limiting policy for a usage plan.
type RateLimit struct {

	// The number of requests that can be made per time period.
	Value *int `mandatory:"true" json:"value"`

	// The unit of time over which rate limits are calculated.
	// Example: `SECOND`
	Unit RateLimitUnitEnum `mandatory:"true" json:"unit"`
}

func (m RateLimit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RateLimit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRateLimitUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetRateLimitUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RateLimitUnitEnum Enum with underlying type: string
type RateLimitUnitEnum string

// Set of constants representing the allowable values for RateLimitUnitEnum
const (
	RateLimitUnitSecond RateLimitUnitEnum = "SECOND"
)

var mappingRateLimitUnitEnum = map[string]RateLimitUnitEnum{
	"SECOND": RateLimitUnitSecond,
}

var mappingRateLimitUnitEnumLowerCase = map[string]RateLimitUnitEnum{
	"second": RateLimitUnitSecond,
}

// GetRateLimitUnitEnumValues Enumerates the set of values for RateLimitUnitEnum
func GetRateLimitUnitEnumValues() []RateLimitUnitEnum {
	values := make([]RateLimitUnitEnum, 0)
	for _, v := range mappingRateLimitUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetRateLimitUnitEnumStringValues Enumerates the set of values in String for RateLimitUnitEnum
func GetRateLimitUnitEnumStringValues() []string {
	return []string{
		"SECOND",
	}
}

// GetMappingRateLimitUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRateLimitUnitEnum(val string) (RateLimitUnitEnum, bool) {
	enum, ok := mappingRateLimitUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
