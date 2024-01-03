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

// RateLimitingPolicy Limit the number of requests that should be handled for the specified window using a specfic key.
type RateLimitingPolicy struct {

	// The maximum number of requests per second to allow.
	RateInRequestsPerSecond *int `mandatory:"true" json:"rateInRequestsPerSecond"`

	// The key used to group requests together.
	RateKey RateLimitingPolicyRateKeyEnum `mandatory:"true" json:"rateKey"`
}

func (m RateLimitingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RateLimitingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRateLimitingPolicyRateKeyEnum(string(m.RateKey)); !ok && m.RateKey != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RateKey: %s. Supported values are: %s.", m.RateKey, strings.Join(GetRateLimitingPolicyRateKeyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RateLimitingPolicyRateKeyEnum Enum with underlying type: string
type RateLimitingPolicyRateKeyEnum string

// Set of constants representing the allowable values for RateLimitingPolicyRateKeyEnum
const (
	RateLimitingPolicyRateKeyClientIp RateLimitingPolicyRateKeyEnum = "CLIENT_IP"
	RateLimitingPolicyRateKeyTotal    RateLimitingPolicyRateKeyEnum = "TOTAL"
)

var mappingRateLimitingPolicyRateKeyEnum = map[string]RateLimitingPolicyRateKeyEnum{
	"CLIENT_IP": RateLimitingPolicyRateKeyClientIp,
	"TOTAL":     RateLimitingPolicyRateKeyTotal,
}

var mappingRateLimitingPolicyRateKeyEnumLowerCase = map[string]RateLimitingPolicyRateKeyEnum{
	"client_ip": RateLimitingPolicyRateKeyClientIp,
	"total":     RateLimitingPolicyRateKeyTotal,
}

// GetRateLimitingPolicyRateKeyEnumValues Enumerates the set of values for RateLimitingPolicyRateKeyEnum
func GetRateLimitingPolicyRateKeyEnumValues() []RateLimitingPolicyRateKeyEnum {
	values := make([]RateLimitingPolicyRateKeyEnum, 0)
	for _, v := range mappingRateLimitingPolicyRateKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetRateLimitingPolicyRateKeyEnumStringValues Enumerates the set of values in String for RateLimitingPolicyRateKeyEnum
func GetRateLimitingPolicyRateKeyEnumStringValues() []string {
	return []string{
		"CLIENT_IP",
		"TOTAL",
	}
}

// GetMappingRateLimitingPolicyRateKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRateLimitingPolicyRateKeyEnum(val string) (RateLimitingPolicyRateKeyEnum, bool) {
	enum, ok := mappingRateLimitingPolicyRateKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
