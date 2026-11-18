// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateHealthCheckerDetails The health checker's configuration details.
type UpdateHealthCheckerDetails struct {

	// The protocol the health check must use; either HTTP or TCP.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol"`

	// The backend server port against which to run the health check.
	// Example: `8080`
	Port *int `mandatory:"true" json:"port"`

	// The status code a healthy backend server should return.
	// Example: `200`
	ReturnCode *int `mandatory:"true" json:"returnCode"`

	// The number of retries to attempt before a backend server is considered "unhealthy". This number also applies
	// when recovering a server to the "healthy" state.
	// Example: `3`
	Retries *int `mandatory:"true" json:"retries"`

	// The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply
	// returns within this timeout period.
	// Example: `3000`
	TimeoutInMillis *int `mandatory:"true" json:"timeoutInMillis"`

	// The interval between health checks, in milliseconds.
	// Example: `10000`
	IntervalInMillis *int `mandatory:"true" json:"intervalInMillis"`

	// A regular expression for parsing the response body from the backend server.
	// Example: `^((?!false).|\s)*$`
	ResponseBodyRegex *string `mandatory:"true" json:"responseBodyRegex"`

	// The path against which to run the health check.
	// Example: `/healthcheck`
	UrlPath *string `mandatory:"false" json:"urlPath"`

	// You can enable dynamic backend weights for HTTP health checks.
	// When enabled, the load balancer reads the `X-Dynamic-Backend-Weight`
	// header from each backend's HTTP health check response.
	// The load balancer uses that value as the backend's effective weight
	// when distributing traffic.
	// A valid header value is an integer from `0` through `100`.
	// If the backend response does not include this header, or includes an invalid
	// value, the load balancer uses `defaultWeight`.
	// Update semantics:
	// * omitted: preserve existing value.
	// * false: disable dynamic backend weight.
	// * true: enable dynamic backend weight.
	// For TCP health checks, this field is accepted and stored, but it has no runtime effect.
	IsDynamicBackendWeightEnabled *bool `mandatory:"false" json:"isDynamicBackendWeightEnabled"`

	// You can configure this fallback backend weight when dynamic backend weights
	// are enabled and the backend's HTTP health check response does not include a valid
	// `X-Dynamic-Backend-Weight` header value.
	// This value must be an integer from `1` through `100`.
	// Update semantics:
	// * provided: replace the existing value.
	// * omitted or null: preserve the existing configured value.
	// If dynamic backend weights are enabled and no existing `defaultWeight` is
	// present, the service uses `1`.
	// Setting or updating `defaultWeight` does not enable dynamic backend weights.
	// To enable dynamic backend weights, set `isDynamicBackendWeightEnabled` to
	// `true`. If your request includes `defaultWeight`, you must also set
	// `isDynamicBackendWeightEnabled` to `true`; otherwise, the request is rejected.
	// For TCP health checks, this field is accepted and stored, but it has no runtime effect.
	// Example: `1`
	DefaultWeight *int `mandatory:"false" json:"defaultWeight"`

	// Specifies if health checks should always be done using plain text instead of depending on
	// whether or not the associated backend set is using SSL.
	// If "true", health checks will be done using plain text even if the associated backend set is configured
	// to use SSL.
	// If "false", health checks will be done using SSL encryption if the associated backend set is configured
	// to use SSL. If the backend set is not so configured the health checks will be done using plain text.
	// Example: `true`
	IsForcePlainText *bool `mandatory:"false" json:"isForcePlainText"`
}

func (m UpdateHealthCheckerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateHealthCheckerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
