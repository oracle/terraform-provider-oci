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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkRequest A description of the work request status.
type WorkRequest struct {

	// The ID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the work request. Work requests should be scoped to
	// the same compartment as the resource the work request affects. If the work request affects multiple resources that are not in the same compartment,
	// then the system picks the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The status of the current work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The completed percentage of the operation tracked by the work request.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the work request was accepted, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	// The precision for this time object in milliseconds.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// The date and time the work request transitioned from ACCEPTED to IN_PROGRESS, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	// The precision for this time object is in milliseconds.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request reached a terminal state, either FAILED or SUCCEEDED, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	// The precision for this time object is in milliseconds.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}
