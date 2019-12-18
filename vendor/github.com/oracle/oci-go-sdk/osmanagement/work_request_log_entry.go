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

// WorkRequestLogEntry Human readable log message describing what the work request is doing
type WorkRequestLogEntry struct {

	// A human readable log message.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the error happened, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestLogEntry) String() string {
	return common.PointerString(m)
}
