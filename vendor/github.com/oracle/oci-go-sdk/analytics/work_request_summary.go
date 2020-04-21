// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestSummary An asynchronous work request.
type WorkRequestSummary struct {

	// The resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// The operation performed by the work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The current status of the work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The compartment OCID of this work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources this work request affects.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// The time the work request was accepted, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The time the work request was started, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the work request was finished, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}
