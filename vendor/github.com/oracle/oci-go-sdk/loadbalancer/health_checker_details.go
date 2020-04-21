// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// HealthCheckerDetails The health check policy's configuration details.
type HealthCheckerDetails struct {

	// The protocol the health check must use; either HTTP or TCP.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol"`

	// The path against which to run the health check.
	// Example: `/healthcheck`
	UrlPath *string `mandatory:"false" json:"urlPath"`

	// The backend server port against which to run the health check. If the port is not specified, the load balancer uses the
	// port information from the `Backend` object.
	// Example: `8080`
	Port *int `mandatory:"false" json:"port"`

	// The status code a healthy backend server should return.
	// Example: `200`
	ReturnCode *int `mandatory:"false" json:"returnCode"`

	// The number of retries to attempt before a backend server is considered "unhealthy". This number also applies
	// when recovering a server to the "healthy" state.
	// Example: `3`
	Retries *int `mandatory:"false" json:"retries"`

	// The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply
	// returns within this timeout period.
	// Example: `3000`
	TimeoutInMillis *int `mandatory:"false" json:"timeoutInMillis"`

	// The interval between health checks, in milliseconds.
	// Example: `10000`
	IntervalInMillis *int `mandatory:"false" json:"intervalInMillis"`

	// A regular expression for parsing the response body from the backend server.
	// Example: `^((?!false).|\s)*$`
	ResponseBodyRegex *string `mandatory:"false" json:"responseBodyRegex"`
}

func (m HealthCheckerDetails) String() string {
	return common.PointerString(m)
}
