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

// AuditLogReport The auditLog report details.
type AuditLogReport struct {

	// auditReportStatus for the accessRequestId
	AuditReportStatus AuditReportStatusEnum `mandatory:"true" json:"auditReportStatus"`

	// Contains the report data.
	Report *string `mandatory:"false" json:"report"`

	// Contains the process tree data
	ProcessTree *string `mandatory:"false" json:"processTree"`

	// Time when the audit report was generated RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfReportGeneration *common.SDKTime `mandatory:"false" json:"timeOfReportGeneration"`
}

func (m AuditLogReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditLogReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditReportStatusEnum(string(m.AuditReportStatus)); !ok && m.AuditReportStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditReportStatus: %s. Supported values are: %s.", m.AuditReportStatus, strings.Join(GetAuditReportStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
