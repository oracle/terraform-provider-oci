// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestError Error encountered during the execution of a work request.
type WorkRequestError struct {

	// A short error code that defines the error, meant for programmatic parsing.
	Code *string `mandatory:"true" json:"code"`

	// Error message.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the error occured, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
