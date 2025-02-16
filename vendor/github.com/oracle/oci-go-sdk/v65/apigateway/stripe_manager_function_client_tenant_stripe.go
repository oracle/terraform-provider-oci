// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StripeManagerFunctionClientTenantStripe Stripe manager function details
type StripeManagerFunctionClientTenantStripe struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource.
	FunctionId *string `mandatory:"true" json:"functionId"`

	// The duration for which parameters should be cached before it is fetched again.
	MaxCacheDurationInHours *int `mandatory:"true" json:"maxCacheDurationInHours"`

	// A map where key is a user defined string and value is a context expressions whose values will be sent to the custom auth function. Values should contain an expression.
	// Example: `{"foo": "request.header[abc]"}`
	Parameters map[string]string `mandatory:"false" json:"parameters"`

	// A list of keys from "parameters" attribute value whose values will be added to the cache key.
	CacheKey []string `mandatory:"false" json:"cacheKey"`
}

func (m StripeManagerFunctionClientTenantStripe) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StripeManagerFunctionClientTenantStripe) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StripeManagerFunctionClientTenantStripe) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStripeManagerFunctionClientTenantStripe StripeManagerFunctionClientTenantStripe
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStripeManagerFunctionClientTenantStripe
	}{
		"STRIPE_MANAGER_FUNCTION",
		(MarshalTypeStripeManagerFunctionClientTenantStripe)(m),
	}

	return json.Marshal(&s)
}
