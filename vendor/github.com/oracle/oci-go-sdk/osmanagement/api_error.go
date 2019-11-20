// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApiError Error Information
type ApiError struct {

	// A short error code that defines the error, meant for programmatic
	// parsing.
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`
}

func (m ApiError) String() string {
	return common.PointerString(m)
}
