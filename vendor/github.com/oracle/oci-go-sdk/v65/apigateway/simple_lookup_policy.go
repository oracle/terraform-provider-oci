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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SimpleLookupPolicy Provides ability to vary the cache key using context expressions.
type SimpleLookupPolicy struct {

	// Whether this policy is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Set true to allow caching responses where the request has an Authorization header. Ensure you have configured your
	// cache key additions to get the level of isolation across authenticated requests that you require.
	// When false, any request with an Authorization header will not be stored in the Response Cache.
	// If using the CustomAuthenticationPolicy then the tokenHeader/tokenQueryParam are also subject to this check.
	IsPrivateCachingEnabled *bool `mandatory:"false" json:"isPrivateCachingEnabled"`

	// A list of context expressions whose values will be added to the base cache key. Values should contain an expression enclosed within
	// ${} delimiters. Only the request context is available.
	CacheKeyAdditions []string `mandatory:"false" json:"cacheKeyAdditions"`
}

// GetIsEnabled returns IsEnabled
func (m SimpleLookupPolicy) GetIsEnabled() *bool {
	return m.IsEnabled
}

// GetIsPrivateCachingEnabled returns IsPrivateCachingEnabled
func (m SimpleLookupPolicy) GetIsPrivateCachingEnabled() *bool {
	return m.IsPrivateCachingEnabled
}

func (m SimpleLookupPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SimpleLookupPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SimpleLookupPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSimpleLookupPolicy SimpleLookupPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSimpleLookupPolicy
	}{
		"SIMPLE_LOOKUP_POLICY",
		(MarshalTypeSimpleLookupPolicy)(m),
	}

	return json.Marshal(&s)
}
