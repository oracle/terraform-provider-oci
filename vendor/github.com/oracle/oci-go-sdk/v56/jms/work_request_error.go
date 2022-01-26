// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkRequestError An error encountered while executing an operation that is tracked by a work request.
type WorkRequestError struct {

	// A machine-usable code for the error that occured. Error codes are listed at
	// API Errors (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm).
	Code *string `mandatory:"true" json:"code"`

	// A human readable description of the issue encountered.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the error occured (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
