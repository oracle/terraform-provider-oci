// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestLogEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
