// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WorkRequest An asynchronous work request. See Work Requests (https://docs.cloud.oracle.com/Content/General/Concepts/workrequestoverview.htm).
type WorkRequest struct {

	// The asynchronous operation tracked by this work request.
	OperationType OperationTypeEnum `mandatory:"true" json:"operationType"`

	// The status of the work request.
	Status OperationStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the work request. Work requests should be scoped to
	// the same compartment as the resource the work request affects. If the work request affects multiple resources,
	// and those resources are not in the same compartment, it is up to the service team to pick the primary
	// resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources that are affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// The percentage complete of the operation tracked by this work request.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request transitioned from _ACCEPTED_ to _IN_PROGRESS_ (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request reached a terminal state, either _FAILED_ or _SUCCEEDED_ (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}
