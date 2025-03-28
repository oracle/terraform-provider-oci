// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestLogEntry A log message from the execution of a work request.
type WorkRequestLogEntry struct {

	// Human-readable log message.
	Message *string `mandatory:"true" json:"message"`

	// The time the log message was written. An RFC3339 formatted datetime string
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
