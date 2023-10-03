// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HttpRateLimitConfiguration HTTP rate limiting configuration is used to limit number of HTTP/GRPC requests in the specific time window.
type HttpRateLimitConfiguration struct {

	// The interval for which requests are limited.
	IntervalInMs *int64 `mandatory:"true" json:"intervalInMs"`

	// Number of requests to limit.
	RequestsPerTarget *int `mandatory:"true" json:"requestsPerTarget"`

	// Overrides the existing rate limit configurations.
	Overrides []HttpRateLimitConfigurationOverrides `mandatory:"false" json:"overrides"`

	// Response headers which should be added to every rate limited response.
	HeadersToAdd []HttpHeader `mandatory:"false" json:"headersToAdd"`
}

// GetIntervalInMs returns IntervalInMs
func (m HttpRateLimitConfiguration) GetIntervalInMs() *int64 {
	return m.IntervalInMs
}

func (m HttpRateLimitConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpRateLimitConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpRateLimitConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpRateLimitConfiguration HttpRateLimitConfiguration
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeHttpRateLimitConfiguration
	}{
		"HTTP",
		(MarshalTypeHttpRateLimitConfiguration)(m),
	}

	return json.Marshal(&s)
}
