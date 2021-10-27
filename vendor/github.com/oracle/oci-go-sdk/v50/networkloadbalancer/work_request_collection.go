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

// WorkRequestCollection Wrapper object for an array of WorkRequest objects.
type WorkRequestCollection struct {

	// An array of WorkRequest objects.
	Items []WorkRequestSummary `mandatory:"false" json:"items"`
}

func (m WorkRequestCollection) String() string {
	return common.PointerString(m)
}
