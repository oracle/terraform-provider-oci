// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkRequestError An error encountered while executing a work request.
type WorkRequestError struct {

	// A machine-usable code for the error that occurred. Error codes are listed on
	// (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm)
	Code *string `mandatory:"true" json:"code"`

	// A human readable description of the issue encountered.
	Message *string `mandatory:"true" json:"message"`

	// The time the error occurred. An RFC3339 formatted datetime string.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
