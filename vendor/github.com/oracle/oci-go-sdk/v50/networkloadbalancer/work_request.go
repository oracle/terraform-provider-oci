// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// WorkRequest A description of work request status.
type WorkRequest struct {

	// Type of work request.
	OperationType OperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of current work request.
	Status OperationStatusEnum `mandatory:"true" json:"status"`

	// The identifier of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the work request. Work requests are scoped to
	// the same compartment as the resource that the work request affects. If the work request affects multiple resources,
	// and those resources are not in the same compartment, then the service team must choose the primary
	// resource whose compartment is to be used.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time that the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time that the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time that the object was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}
