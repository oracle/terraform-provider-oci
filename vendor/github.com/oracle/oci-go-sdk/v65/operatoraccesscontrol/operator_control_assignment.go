// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// OperatorControlAssignment An Operator Control Assignment identifies the target resource that is placed under the governance of an Operator Control. Creating an Operator Control Assignment Assignment with a time duration ensures that
// human accesses to the target resource will be governed by Operator Control for the duration specified.
type OperatorControlAssignment struct {

	// The OCID of the operator control assignment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the operator control.
	OperatorControlId *string `mandatory:"true" json:"operatorControlId"`

	// The OCID of the target resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Name of the target resource.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The OCID of the compartment that contains the target resource.
	ResourceCompartmentId *string `mandatory:"false" json:"resourceCompartmentId"`

	// The OCID of the comparment that contains the operator control assignment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// resourceType for which the OperatorControlAssignment is applicable
	ResourceType ResourceTypesEnum `mandatory:"false" json:"resourceType,omitempty"`

	// The time at which the target resource will be brought under the governance of the operator control expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: '2020-05-22T21:10:29.600Z'
	TimeAssignmentFrom *common.SDKTime `mandatory:"false" json:"timeAssignmentFrom"`

	// The time at which the target resource will leave the governance of the operator control expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: '2020-05-22T21:10:29.600Z'
	TimeAssignmentTo *common.SDKTime `mandatory:"false" json:"timeAssignmentTo"`

	// If set, then the target resource is always governed by the operator control.
	IsEnforcedAlways *bool `mandatory:"false" json:"isEnforcedAlways"`

	// The current lifcycle state of the OperatorControl.
	LifecycleState OperatorControlAssignmentLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// More in detail about the lifeCycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the user who created this operator control assignment.
	AssignerId *string `mandatory:"false" json:"assignerId"`

	// Time when the operator control assignment is created in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfAssignment *common.SDKTime `mandatory:"false" json:"timeOfAssignment"`

	// Comment about the assignment of the operator control to this target resource.
	Comment *string `mandatory:"false" json:"comment"`

	// User id who released the operatorControl.
	UnassignerId *string `mandatory:"false" json:"unassignerId"`

	// Time on which the operator control assignment was deleted in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// description containing reason for releasing of OperatorControl.
	DetachmentDescription *string `mandatory:"false" json:"detachmentDescription"`

	// If set indicates that the audit logs are being forwarded to the relevant remote logging server
	IsLogForwarded *bool `mandatory:"false" json:"isLogForwarded"`

	// The address of the remote syslog server where the audit logs are being forwarded to. Address in host or IP format.
	RemoteSyslogServerAddress *string `mandatory:"false" json:"remoteSyslogServerAddress"`

	// The listening port of the remote syslog server. The port range is 0 - 65535. Only TCP supported.
	RemoteSyslogServerPort *int `mandatory:"false" json:"remoteSyslogServerPort"`

	// The CA certificate of the remote syslog server.
	RemoteSyslogServerCACert *string `mandatory:"false" json:"remoteSyslogServerCACert"`

	// If set, then the hypervisor audit logs will be forwarded to the relevant remote syslog server
	IsHypervisorLogForwarded *bool `mandatory:"false" json:"isHypervisorLogForwarded"`

	// Name of the operator control name associated.
	OpControlName *string `mandatory:"false" json:"opControlName"`

	// The boolean if true would autoApprove during maintenance.
	IsAutoApproveDuringMaintenance *bool `mandatory:"false" json:"isAutoApproveDuringMaintenance"`

	// The code identifying the error occurred during Assignment operation.
	ErrorCode *int `mandatory:"false" json:"errorCode"`

	// The message describing the error occurred during Assignment operation.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// Whether the assignment is a default assignment.
	IsDefaultAssignment *bool `mandatory:"false" json:"isDefaultAssignment"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OperatorControlAssignment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OperatorControlAssignment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceTypesEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetResourceTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOperatorControlAssignmentLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOperatorControlAssignmentLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
