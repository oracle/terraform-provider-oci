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

// OperatorControl Operator Access Control enables you to grant, audit, or revoke the access Oracle has to your Exadata Cloud@Customer infrastructure, and obtain audit reports of all actions taken by a human operator, in a near real-time manner.
type OperatorControl struct {

	// The OCID of the operator control.
	Id *string `mandatory:"true" json:"id"`

	// Name of the operator control. The name must be unique.
	OperatorControlName *string `mandatory:"true" json:"operatorControlName"`

	// The OCID of the compartment that contains the operator control.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Description of operator control.
	Description *string `mandatory:"false" json:"description"`

	// List of users who can approve an access request associated with a target resource under the governance of this operator control.
	ApproversList []string `mandatory:"false" json:"approversList"`

	// List of user groups who can approve an access request associated with a target resource under the governance of this operator control.
	ApproverGroupsList []string `mandatory:"false" json:"approverGroupsList"`

	// List of pre-approved operator actions. Access requests associated with a resource governed by this operator control will be
	// automatically approved if the access request only contain operator actions in the pre-approved list.
	PreApprovedOpActionList []string `mandatory:"false" json:"preApprovedOpActionList"`

	// List of operator actions that need explicit approval. Any operator action not in the pre-approved list will require explicit
	// approval. Access requests associated with a resource governed by this operator control will be
	// require explicit approval if the access request contains any operator action in this list.
	ApprovalRequiredOpActionList []string `mandatory:"false" json:"approvalRequiredOpActionList"`

	// Whether all the operator actions have been pre-approved. If yes, all access requests associated with a resource governed by this operator control
	// will be auto-approved.
	IsFullyPreApproved *bool `mandatory:"false" json:"isFullyPreApproved"`

	// List of emailId.
	EmailIdList []string `mandatory:"false" json:"emailIdList"`

	// resourceType for which the OperatorControl is applicable
	ResourceType ResourceTypesEnum `mandatory:"false" json:"resourceType,omitempty"`

	// System message that would be displayed to the operator users on accessing the target resource under the governance of this operator control.
	SystemMessage *string `mandatory:"false" json:"systemMessage"`

	// The current lifecycle state of the operator control.
	LifecycleState OperatorControlLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Time when the operator control was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfCreation *common.SDKTime `mandatory:"false" json:"timeOfCreation"`

	// Time when the operator control was last modified expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfModification *common.SDKTime `mandatory:"false" json:"timeOfModification"`

	// Time when deleted expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'.
	// Note a deleted operator control still stays in the system, so that you can still audit operator actions associated with access requests
	// raised on target resources governed by the deleted operator control.
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// Description associated with the latest modification of the operator control.
	LastModifiedInfo *string `mandatory:"false" json:"lastModifiedInfo"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OperatorControl) String() string {
	return common.PointerString(m)
}
