// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// HealthCheckerDetails The health check policy configuration.
// For more information, see Editing Health Check Policies (https://docs.cloud.oracle.com/Content/Balance/Tasks/editinghealthcheck.htm).
type HealthCheckerDetails struct {

	// The protocol the health check must use; either HTTP or HTTPS, or UDP or TCP.
	// Example: `HTTP`
	Protocol HealthCheckProtocolsEnum `mandatory:"true" json:"protocol"`

	// The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the
	// port information from the `Backend` object. The port must be specified if the backend port is 0.
	// Example: `8080`
	Port *int `mandatory:"false" json:"port"`

	// The number of retries to attempt before a backend server is considered "unhealthy". This number also applies
	// when recovering a server to the "healthy" state. The default value is 3.
	// Example: `3`
	Retries *int `mandatory:"false" json:"retries"`

	// The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply
	// returns within this timeout period. The default value is 3000 (3 seconds).
	// Example: `3000`
	TimeoutInMillis *int `mandatory:"false" json:"timeoutInMillis"`

	// The interval between health checks, in milliseconds. The default value is 10000 (10 seconds).
	// Example: `10000`
	IntervalInMillis *int `mandatory:"false" json:"intervalInMillis"`

	// The path against which to run the health check.
	// Example: `/healthcheck`
	UrlPath *string `mandatory:"false" json:"urlPath"`

	// A regular expression for parsing the response body from the backend server.
	// Example: `^((?!false).|\s)*$`
	ResponseBodyRegex *string `mandatory:"false" json:"responseBodyRegex"`

	// The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol,
	// then you can use common HTTP status codes such as "200".
	// Example: `200`
	ReturnCode *int `mandatory:"false" json:"returnCode"`

	// Base64 encoded pattern to be sent as UDP or TCP health check probe.
	RequestData []byte `mandatory:"false" json:"requestData"`

	// Base64 encoded pattern to be validated as UDP or TCP health check probe response.
	ResponseData []byte `mandatory:"false" json:"responseData"`
}

func (m HealthCheckerDetails) String() string {
	return common.PointerString(m)
}
