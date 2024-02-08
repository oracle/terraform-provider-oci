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

// UpdateOperatorControlDetails Information about the new operator control.
type UpdateOperatorControlDetails struct {

	// Name of the operator control.
	OperatorControlName *string `mandatory:"true" json:"operatorControlName"`

	// List of user groups who can approve an access request associated with a target resource under the governance of this operator control.
	ApproverGroupsList []string `mandatory:"true" json:"approverGroupsList"`

	// Whether all the operator actions have been pre-approved. If yes, all access requests associated with a resource governed by this operator control
	// will be auto-approved.
	IsFullyPreApproved *bool `mandatory:"true" json:"isFullyPreApproved"`

	// Description of the operator control.
	Description *string `mandatory:"false" json:"description"`

	// List of users who can approve an access request associated with a target resource under the governance of this operator control.
	ApproversList []string `mandatory:"false" json:"approversList"`

	// List of pre-approved operator actions. Access requests associated with a resource governed by this operator control will be
	// automatically approved if the access request only contain operator actions in the pre-approved list.
	PreApprovedOpActionList []string `mandatory:"false" json:"preApprovedOpActionList"`

	// List of emailId.
	EmailIdList []string `mandatory:"false" json:"emailIdList"`

	// System message that would be displayed to the operator users on accessing the target resource under the governance of this operator control.
	SystemMessage *string `mandatory:"false" json:"systemMessage"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOperatorControlDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOperatorControlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
