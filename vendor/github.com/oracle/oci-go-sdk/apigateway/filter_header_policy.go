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

// FilterHeaderPolicy Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations,
// so any headers set or renamed must also be listed here when using an ALLOW type policy.
type FilterHeaderPolicy struct {

	// BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW
	// permits only the headers in the list and removes all others, so it acts as an inclusion list.
	Type FilterHeaderPolicyTypeEnum `mandatory:"true" json:"type"`

	// The list of headers.
	Items []FilterHeaderPolicyItem `mandatory:"true" json:"items"`
}

func (m FilterHeaderPolicy) String() string {
	return common.PointerString(m)
}

// FilterHeaderPolicyTypeEnum Enum with underlying type: string
type FilterHeaderPolicyTypeEnum string

// Set of constants representing the allowable values for FilterHeaderPolicyTypeEnum
const (
	FilterHeaderPolicyTypeAllow FilterHeaderPolicyTypeEnum = "ALLOW"
	FilterHeaderPolicyTypeBlock FilterHeaderPolicyTypeEnum = "BLOCK"
)

var mappingFilterHeaderPolicyType = map[string]FilterHeaderPolicyTypeEnum{
	"ALLOW": FilterHeaderPolicyTypeAllow,
	"BLOCK": FilterHeaderPolicyTypeBlock,
}

// GetFilterHeaderPolicyTypeEnumValues Enumerates the set of values for FilterHeaderPolicyTypeEnum
func GetFilterHeaderPolicyTypeEnumValues() []FilterHeaderPolicyTypeEnum {
	values := make([]FilterHeaderPolicyTypeEnum, 0)
	for _, v := range mappingFilterHeaderPolicyType {
		values = append(values, v)
	}
	return values
}
