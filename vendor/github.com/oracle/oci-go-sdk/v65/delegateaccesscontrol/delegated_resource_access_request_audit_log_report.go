// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DelegatedResourceAccessRequestAuditLogReport The audit log report details.
type DelegatedResourceAccessRequestAuditLogReport struct {

	// Status of the audit report
	AuditReportStatus AuditReportStatusEnum `mandatory:"true" json:"auditReportStatus"`

	// Audit log report.
	Report *string `mandatory:"false" json:"report"`

	// The process tree data
	ProcessTree *string `mandatory:"false" json:"processTree"`

	// Time when the audit report was generated RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeReportGenerated *common.SDKTime `mandatory:"false" json:"timeReportGenerated"`
}

func (m DelegatedResourceAccessRequestAuditLogReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DelegatedResourceAccessRequestAuditLogReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditReportStatusEnum(string(m.AuditReportStatus)); !ok && m.AuditReportStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditReportStatus: %s. Supported values are: %s.", m.AuditReportStatus, strings.Join(GetAuditReportStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
