// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration (WAA) API
//
// API for the Web Application Acceleration service.
// Use this API to manage regional Web App Acceleration policies such as Caching and Compression
// for accelerating HTTP services.
//

package waa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResponseCachingPolicy An object that specifies an HTTP response caching policy.
type ResponseCachingPolicy struct {

	// When false, responses will not be cached by the backend based on response headers.
	// When true, responses that contain one of the supported cache control headers will be cached according to the
	// values specified in the cache control headers.
	// The "X-Accel-Expires" header field sets caching time of a response in seconds. The zero value disables
	// caching for a response. If the value starts with the @ prefix, it sets an absolute time in seconds since
	// Epoch, up to which the response may be cached.
	// If the header does not include the "X-Accel-Expires" field, parameters of caching may be set in the header
	// fields "Expires" or "Cache-Control".
	// If the header includes the "Set-Cookie" field, such a response will not be cached.
	// If the header includes the "Vary" field with the special value "*", such a response will not be cached. If the
	// header includes the "Vary" field with another value, such a response will be cached taking into account the
	// corresponding request header fields.
	IsResponseHeaderBasedCachingEnabled *bool `mandatory:"false" json:"isResponseHeaderBasedCachingEnabled"`
}

func (m ResponseCachingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResponseCachingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
