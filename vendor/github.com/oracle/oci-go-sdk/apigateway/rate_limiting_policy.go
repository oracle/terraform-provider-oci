// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
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

// RateLimitingPolicyRateKeyEnum Enum with underlying type: string
type RateLimitingPolicyRateKeyEnum string

// Set of constants representing the allowable values for RateLimitingPolicyRateKeyEnum
const (
	RateLimitingPolicyRateKeyClientIp RateLimitingPolicyRateKeyEnum = "CLIENT_IP"
	RateLimitingPolicyRateKeyTotal    RateLimitingPolicyRateKeyEnum = "TOTAL"
)

var mappingRateLimitingPolicyRateKey = map[string]RateLimitingPolicyRateKeyEnum{
	"CLIENT_IP": RateLimitingPolicyRateKeyClientIp,
	"TOTAL":     RateLimitingPolicyRateKeyTotal,
}

// GetRateLimitingPolicyRateKeyEnumValues Enumerates the set of values for RateLimitingPolicyRateKeyEnum
func GetRateLimitingPolicyRateKeyEnumValues() []RateLimitingPolicyRateKeyEnum {
	values := make([]RateLimitingPolicyRateKeyEnum, 0)
	for _, v := range mappingRateLimitingPolicyRateKey {
		values = append(values, v)
	}
	return values
}
