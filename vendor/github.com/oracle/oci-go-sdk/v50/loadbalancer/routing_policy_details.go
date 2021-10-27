// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// RoutingPolicyDetails An ordered list of routing rules.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type RoutingPolicyDetails struct {

	// The list of routing rules.
	Rules []RoutingRule `mandatory:"true" json:"rules"`
}

func (m RoutingPolicyDetails) String() string {
	return common.PointerString(m)
}
