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

// WorkRequestError An error encountered while running a work request.
type WorkRequestError struct {

	// A machine-usable code for the error that occured. Error codes are listed here:
	// (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm)
	Code *string `mandatory:"true" json:"code"`

	// A human-readable description of the issue encountered.
	Message *string `mandatory:"true" json:"message"`

	// The time the error occured in the form of an RFC3339 formatted date-time string.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
