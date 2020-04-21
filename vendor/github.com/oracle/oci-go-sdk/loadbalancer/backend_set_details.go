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

// BackendSetDetails The configuration details for a load balancer backend set.
// For more information on backend set configuration, see
// Managing Backend Sets (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingbackendsets.htm).
// **Note:** The `sessionPersistenceConfiguration` (application cookie stickiness) and `lbCookieSessionPersistenceConfiguration`
// (LB cookie stickiness) attributes are mutually exclusive. To avoid returning an error, configure only one of these two
// attributes per backend set.
type BackendSetDetails struct {

	// The load balancer policy for the backend set. To get a list of available policies, use the
	// ListPolicies operation.
	// Example: `LEAST_CONNECTIONS`
	Policy *string `mandatory:"true" json:"policy"`

	HealthChecker *HealthCheckerDetails `mandatory:"true" json:"healthChecker"`

	Backends []BackendDetails `mandatory:"false" json:"backends"`

	SslConfiguration *SslConfigurationDetails `mandatory:"false" json:"sslConfiguration"`

	SessionPersistenceConfiguration *SessionPersistenceConfigurationDetails `mandatory:"false" json:"sessionPersistenceConfiguration"`

	LbCookieSessionPersistenceConfiguration *LbCookieSessionPersistenceConfigurationDetails `mandatory:"false" json:"lbCookieSessionPersistenceConfiguration"`
}

func (m BackendSetDetails) String() string {
	return common.PointerString(m)
}
