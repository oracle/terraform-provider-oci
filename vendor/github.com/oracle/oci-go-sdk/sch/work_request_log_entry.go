// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
//

package sch

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestLogEntry A log message from the execution of a work request.
type WorkRequestLogEntry struct {

	// Human-readable log message.
	Message *string `mandatory:"true" json:"message"`

	// The date and time when the log message was written.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestLogEntry) String() string {
	return common.PointerString(m)
}
