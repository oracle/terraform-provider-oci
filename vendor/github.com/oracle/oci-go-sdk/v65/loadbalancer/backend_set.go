// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackendSet The configuration of a load balancer backend set.
// For more information on backend set configuration, see
// Managing Backend Sets (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingbackendsets.htm).
// **Note:** The `sessionPersistenceConfiguration` (application cookie stickiness) and `lbCookieSessionPersistenceConfiguration`
// (LB cookie stickiness) attributes are mutually exclusive. To avoid returning an error, configure only one of these two
// attributes per backend set.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type BackendSet struct {

	// A friendly name for the backend set. It must be unique and it cannot be changed.
	// Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot
	// contain spaces. Avoid entering confidential information.
	// Example: `example_backend_set`
	Name *string `mandatory:"true" json:"name"`

	// The load balancer policy for the backend set. To get a list of available policies, use the
	// ListPolicies operation.
	// Example: `LEAST_CONNECTIONS`
	Policy *string `mandatory:"true" json:"policy"`

	Backends []Backend `mandatory:"true" json:"backends"`

	HealthChecker *HealthChecker `mandatory:"true" json:"healthChecker"`

	// The maximum number of simultaneous connections the load balancer can make to any backend
	// in the backend set unless the backend has its own maxConnections setting. If this is not
	// set then the number of simultaneous connections the load balancer can make to any backend
	// in the backend set unless the backend has its own maxConnections setting is unlimited.
	// Example: `300`
	BackendMaxConnections *int `mandatory:"false" json:"backendMaxConnections"`

	SslConfiguration *SslConfiguration `mandatory:"false" json:"sslConfiguration"`

	SessionPersistenceConfiguration *SessionPersistenceConfigurationDetails `mandatory:"false" json:"sessionPersistenceConfiguration"`

	LbCookieSessionPersistenceConfiguration *LbCookieSessionPersistenceConfigurationDetails `mandatory:"false" json:"lbCookieSessionPersistenceConfiguration"`
}

func (m BackendSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackendSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
