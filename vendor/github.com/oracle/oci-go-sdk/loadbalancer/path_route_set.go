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

// PathRouteSet A named set of path route rules. For more information, see
// Managing Request Routing (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingrequest.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type PathRouteSet struct {

	// The unique name for this set of path route rules. Avoid entering confidential information.
	// Example: `example_path_route_set`
	Name *string `mandatory:"true" json:"name"`

	// The set of path route rules.
	PathRoutes []PathRoute `mandatory:"true" json:"pathRoutes"`
}

func (m PathRouteSet) String() string {
	return common.PointerString(m)
}
