// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApproverDetail details of Approver Detail.
type ApproverDetail struct {

	// The userId of the approver.
	ApproverId *string `mandatory:"false" json:"approverId"`

	// The action done by the approver.
	ApprovalAction *string `mandatory:"false" json:"approvalAction"`

	// Comment specified by the approver of the request.
	ApprovalComment *string `mandatory:"false" json:"approvalComment"`

	// Additional message specified by the approver of the request.
	ApprovalAdditionalMessage *string `mandatory:"false" json:"approvalAdditionalMessage"`

	// Time when the access request was authorized by the customer in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfAuthorization *common.SDKTime `mandatory:"false" json:"timeOfAuthorization"`

	// Time for when the access request should start that is authorized by the customer in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeApprovedForAccess *common.SDKTime `mandatory:"false" json:"timeApprovedForAccess"`
}

func (m ApproverDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApproverDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
