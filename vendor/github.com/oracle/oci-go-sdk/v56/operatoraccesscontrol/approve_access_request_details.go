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

// ApproveAccessRequestDetails Details of the access request approval.
type ApproveAccessRequestDetails struct {

	// Comment by the approver during approval.
	ApproverComment *string `mandatory:"false" json:"approverComment"`

	// Specifies the type of auditing to be enabled. There are two levels of auditing: command-level and keystroke-level.
	// By default, auditing is enabled at the command level i.e., each command issued by the operator is audited. When keystroke-level is chosen,
	// in addition to command level logging, key strokes are also logged.
	AuditType []string `mandatory:"false" json:"auditType"`

	// Message that needs to be displayed to the Ops User.
	AdditionalMessage *string `mandatory:"false" json:"additionalMessage"`

	// The time when access request is scheduled to be approved in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfUserCreation *common.SDKTime `mandatory:"false" json:"timeOfUserCreation"`
}

func (m ApproveAccessRequestDetails) String() string {
	return common.PointerString(m)
}
