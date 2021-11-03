// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// IamWorkRequestErrorSummary An error encountered while executing an operation that is tracked by a IAM work request.
type IamWorkRequestErrorSummary struct {

	// A machine-usable code for the error that occured.
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the error occurred.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m IamWorkRequestErrorSummary) String() string {
	return common.PointerString(m)
}
