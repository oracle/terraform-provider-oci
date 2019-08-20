// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ArchiverError An error related to a stream archiver.
type ArchiverError struct {

	// A short error code that defines the error, meant for programmatic parsing.
	Code *string `mandatory:"false" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"false" json:"message"`
}

func (m ArchiverError) String() string {
	return common.PointerString(m)
}
