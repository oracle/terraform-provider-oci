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

// PathRoute A "path route rule" to evaluate an incoming URI path, and then route a matching request to the specified backend set.
// Path route rules apply only to HTTP and HTTPS requests. They have no effect on TCP requests.
type PathRoute struct {

	// The path string to match against the incoming URI path.
	// *  Path strings are case-insensitive.
	// *  Asterisk (*) wildcards are not supported.
	// *  Regular expressions are not supported.
	// Example: `/example/video/123`
	Path *string `mandatory:"true" json:"path"`

	// The type of matching to apply to incoming URIs.
	PathMatchType *PathMatchType `mandatory:"true" json:"pathMatchType"`

	// The name of the target backend set for requests where the incoming URI matches the specified path.
	// Example: `example_backend_set`
	BackendSetName *string `mandatory:"true" json:"backendSetName"`
}

func (m PathRoute) String() string {
	return common.PointerString(m)
}
