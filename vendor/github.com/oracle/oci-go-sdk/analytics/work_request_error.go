// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
