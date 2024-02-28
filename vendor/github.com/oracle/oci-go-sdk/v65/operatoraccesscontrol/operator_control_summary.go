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

// OperatorControlSummary Summary of the OperatorControl.
type OperatorControlSummary struct {

	// The OCID of the operator control.
	Id *string `mandatory:"true" json:"id"`

	// Name of the operator control.
	OperatorControlName *string `mandatory:"true" json:"operatorControlName"`

	// The OCID of the compartment that contains the operator control.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Whether all operator actions are pre-approved. If yes, an access request associated with a resource governed by the operator control will be automatically approved by the system.
	IsFullyPreApproved *bool `mandatory:"false" json:"isFullyPreApproved"`

	// resourceType for which the OperatorControl is applicable
	ResourceType ResourceTypesEnum `mandatory:"false" json:"resourceType,omitempty"`

	// Number of approvers required to approve an access request.
	NumberOfApprovers *int `mandatory:"false" json:"numberOfApprovers"`

	// Time when the operator control was created, expressed in RFC 3339  (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfCreation *common.SDKTime `mandatory:"false" json:"timeOfCreation"`

	// Time when the operator control was last modified, expressed in RFC 3339  (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfModification *common.SDKTime `mandatory:"false" json:"timeOfModification"`

	// Time when the operator control was deleted, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// The current lifecycle state of the operator control.
	LifecycleState OperatorControlLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OperatorControlSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OperatorControlSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceTypesEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetResourceTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOperatorControlLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOperatorControlLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
