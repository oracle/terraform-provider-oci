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

// UpdateOperatorControlAssignmentDetails Details for modifying the Operator Control assignment.
type UpdateOperatorControlAssignmentDetails struct {

	// If true, then the target resource is always governed by the operator control. Otherwise governance is time-based as specified by timeAssignmentTo and timeAssignmentFrom.
	IsEnforcedAlways *bool `mandatory:"true" json:"isEnforcedAlways"`

	// The time at which the target resource will be brought under the governance of the operator control in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeAssignmentFrom *common.SDKTime `mandatory:"false" json:"timeAssignmentFrom"`

	// The time at which the target resource will leave the governance of the operator control in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeAssignmentTo *common.SDKTime `mandatory:"false" json:"timeAssignmentTo"`

	// Comment about the modification of the operator control assignment.
	Comment *string `mandatory:"false" json:"comment"`

	// If set, then the audit logs will be forwarded to the relevant remote logging server
	IsLogForwarded *bool `mandatory:"false" json:"isLogForwarded"`

	// The address of the remote syslog server where the audit logs will be forwarded to. Address in host or IP format.
	RemoteSyslogServerAddress *string `mandatory:"false" json:"remoteSyslogServerAddress"`

	// The listening port of the remote syslog server. The port range is 0 - 65535. Only TCP supported.
	RemoteSyslogServerPort *int `mandatory:"false" json:"remoteSyslogServerPort"`

	// The CA certificate of the remote syslog server. Identity of the remote syslog server will be asserted based on this certificate.
	RemoteSyslogServerCACert *string `mandatory:"false" json:"remoteSyslogServerCACert"`

	// The boolean if true would autoApprove during maintenance.
	IsAutoApproveDuringMaintenance *bool `mandatory:"false" json:"isAutoApproveDuringMaintenance"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOperatorControlAssignmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOperatorControlAssignmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
