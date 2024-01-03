// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FindingsChangeAuditLogSummary Summary of audit log of risk updates of findings of specified security assessment.
type FindingsChangeAuditLogSummary struct {

	// The unique key that identifies the finding risk change.
	Key *string `mandatory:"true" json:"key"`

	// The unique key that identifies the finding.
	FindingKey *string `mandatory:"true" json:"findingKey"`

	// The short title for the finding whose risk is being modified.
	FindingTitle *string `mandatory:"true" json:"findingTitle"`

	// The OCID of the latest security assessment.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// The OCID of the target database.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The date and time the risk level of finding was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The original severity / risk level of the finding as determined by security assessment.
	Severity FindingSeverityEnum `mandatory:"true" json:"severity"`

	// The severity of the finding as determined by security assessment by Oracle.
	OracleDefinedSeverity FindingSeverityEnum `mandatory:"true" json:"oracleDefinedSeverity"`

	// Determines if the user has deferred the risk level of this finding when he is ok with it
	// and does not plan to do anything about it.
	IsRiskDeferred *bool `mandatory:"true" json:"isRiskDeferred"`

	// If the risk level is changed more than once, the previous modified value.
	PreviousSeverity FindingSeverityEnum `mandatory:"true" json:"previousSeverity"`

	// The justification given by the user for accepting or modifying the risk level.
	Justification *string `mandatory:"true" json:"justification"`

	// The user who initiated change of risk level of the finding
	ModifiedBy *string `mandatory:"true" json:"modifiedBy"`

	// The date and time, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339),
	// the risk level change as updated by user is valid until. After this date passes, the risk level
	// will be that of what is determined by the latest security assessment.
	TimeValidUntil *common.SDKTime `mandatory:"false" json:"timeValidUntil"`
}

func (m FindingsChangeAuditLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FindingsChangeAuditLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFindingSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetFindingSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFindingSeverityEnum(string(m.OracleDefinedSeverity)); !ok && m.OracleDefinedSeverity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OracleDefinedSeverity: %s. Supported values are: %s.", m.OracleDefinedSeverity, strings.Join(GetFindingSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFindingSeverityEnum(string(m.PreviousSeverity)); !ok && m.PreviousSeverity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreviousSeverity: %s. Supported values are: %s.", m.PreviousSeverity, strings.Join(GetFindingSeverityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
