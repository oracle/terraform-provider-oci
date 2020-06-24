// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestSummary The status of an asynchronous task in the system.
type WorkRequestSummary struct {

	// The id of the work request.
	Id *string `mandatory:"true" json:"id"`

	// the original operation ID requested
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Current status of the work request
	Status WorkRequestOperationStatusEnum `mandatory:"true" json:"status"`

	// The ocid of the compartment that contains the work request. Work
	// requests should be scoped to the same compartment as the resource
	// the work request affects. If the work request affects multiple
	// resources, and those resources are not in the same compartment, it
	// is up to the service team to pick the primary resource whose
	// compartment should be used
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the object was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}
