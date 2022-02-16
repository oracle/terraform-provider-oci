// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// WorkRequestError An error encountered while executing a work request.
type WorkRequestError struct {

	// The identifier of the work request error.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the work request.
	WorkRequestId *string `mandatory:"true" json:"workRequestId"`

	// A machine-usable code for the error that occurred. Error codes are listed on
	// (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm).
	Code *string `mandatory:"true" json:"code"`

	// A human-readable description of the issue that occurred.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the error occurred as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339). The precision for the time object is in milliseconds.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Determines if the work request error can be reproduced and tried again.
	IsRetryable *bool `mandatory:"false" json:"isRetryable"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestError) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
