// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Digital Assistant Control Plane API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestError Description of the unexpected error that prevented completion of the request.
type WorkRequestError struct {

	// A machine-usable code for the error that occurred. Error codes are listed at
	// (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm)
	Code *string `mandatory:"true" json:"code"`

	// A human-readable description of the issue.
	Message *string `mandatory:"true" json:"message"`

	// When the error occurred. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeStamp *common.SDKTime `mandatory:"true" json:"timeStamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
