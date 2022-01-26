// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// QueryWorkRequestSummary High level summary of query job work request.
type QueryWorkRequestSummary struct {

	// Unique OCID identifier to reference this query job work Request with.
	Id *string `mandatory:"true" json:"id"`

	// When the work request started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// Current execution mode for the job.
	Mode JobModeEnum `mandatory:"true" json:"mode"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// When the work request was accepted. Should match timeStarted in all cases.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// When the work request finished execution.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// When the work request will expire.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// Percentage progress completion of the query.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// Work request status.
	Status WorkRequestStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Asynchronous action name.
	OperationType QueryOperationTypeEnum `mandatory:"false" json:"operationType,omitempty"`
}

func (m QueryWorkRequestSummary) String() string {
	return common.PointerString(m)
}
