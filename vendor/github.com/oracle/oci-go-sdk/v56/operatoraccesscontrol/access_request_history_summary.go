// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AccessRequestHistorySummary Summary of access request status.
type AccessRequestHistorySummary struct {

	// The current state of the AccessRequest.
	LifecycleState AccessRequestLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Approver who modified the access request.
	UserId *string `mandatory:"false" json:"userId"`

	// Reason or description about the cause of change.
	Description *string `mandatory:"false" json:"description"`

	// Duration for approval of request or extension depending on the type of action.
	Duration *int `mandatory:"false" json:"duration"`

	// Whether the access request was automatically approved.
	IsAutoApproved *bool `mandatory:"false" json:"isAutoApproved"`

	// List of operator actions for which approvals were requested by the operator.
	ActionsList []string `mandatory:"false" json:"actionsList"`

	// Time when the respective action happened in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfAction *common.SDKTime `mandatory:"false" json:"timeOfAction"`
}

func (m AccessRequestHistorySummary) String() string {
	return common.PointerString(m)
}
