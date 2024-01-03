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

// AssignmentValidationStatus Summary of assignment Validation status.
type AssignmentValidationStatus struct {

	// Id of the unique execution.
	ExecutionId *string `mandatory:"true" json:"executionId"`

	// Status of the execution Success or Failure.
	ExecutionStatus AssignmentValidationLifecycleStatesEnum `mandatory:"true" json:"executionStatus"`

	// Id of the user who triggered the Assignment Validation.
	UserId *string `mandatory:"false" json:"userId"`

	// Id of the accessRequest which got created as part of Assignment Validation.
	AccessRequestId *string `mandatory:"false" json:"accessRequestId"`

	// any errorMessage during validation.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// List of execution detail for the validate assignment.
	StatusDetail []StatusDetail `mandatory:"false" json:"statusDetail"`

	// Time when the respective action happened in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'.
	TimeOfValidationStart *common.SDKTime `mandatory:"false" json:"timeOfValidationStart"`

	// Time when the respective action happened in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'.
	TimeOfValidationFinish *common.SDKTime `mandatory:"false" json:"timeOfValidationFinish"`
}

func (m AssignmentValidationStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssignmentValidationStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssignmentValidationLifecycleStatesEnum(string(m.ExecutionStatus)); !ok && m.ExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionStatus: %s. Supported values are: %s.", m.ExecutionStatus, strings.Join(GetAssignmentValidationLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
