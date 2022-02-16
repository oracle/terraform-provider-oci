// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// AuditArchiveRetrieval Represents the archive retrieve request for the audit data. You can retrieve audit data for a target database from the archive and store it online.
// For more information, see Retrieve Audit Data for a Target Database from the Archive (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/security-assessment-overview.html).
type AuditArchiveRetrieval struct {

	// The OCID of the archive retrieval.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains archive retrieval.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the archive retrieval. The name does not have to be unique, and is changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Start month of the archive retrieval, in the format defined by RFC3339.
	StartDate *common.SDKTime `mandatory:"true" json:"startDate"`

	// End month of the archive retrieval, in the format defined by RFC3339.
	EndDate *common.SDKTime `mandatory:"true" json:"endDate"`

	// The OCID of the target associated with the archive retrieval.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The current state of the archive retrieval.
	LifecycleState AuditArchiveRetrievalLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the archive retrieval.
	Description *string `mandatory:"false" json:"description"`

	// The date time when archive retrieval was requested, in the format defined by RFC3339.
	TimeRequested *common.SDKTime `mandatory:"false" json:"timeRequested"`

	// The date time when archive retrieval request was fulfilled, in the format defined by RFC3339.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The date time when retrieved archive data will be deleted from Data Safe and unloaded back into archival.
	TimeOfExpiry *common.SDKTime `mandatory:"false" json:"timeOfExpiry"`

	// Total count of audit events to be retrieved from the archive for the specified date range.
	AuditEventCount *int64 `mandatory:"false" json:"auditEventCount"`

	// The Error details of a failed archive retrieval.
	ErrorInfo *string `mandatory:"false" json:"errorInfo"`

	// Details about the current state of the archive retrieval.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AuditArchiveRetrieval) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditArchiveRetrieval) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditArchiveRetrievalLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAuditArchiveRetrievalLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
