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

// AuditTrail An audit trail represents the source of audit records that provides documentary evidence of
// the sequence of activities in the target database. Configuring audit trails in Data Safe, and enabling
// audit data collection on the audit trails copies the audit records from the target database's audit trail
// into the Data Safe repository.
type AuditTrail struct {

	// The OCID of the audit trail.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the  parent audit.
	AuditProfileId *string `mandatory:"true" json:"auditProfileId"`

	// The OCID of the Data Safe target for which the audit trail is created.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The display name of the audit trail.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the audit trail was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the audit trail was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the audit trail.
	LifecycleState AuditTrailLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current sub-state of the audit trail.
	Status AuditTrailStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the compartment that contains the audit trail and is the same as the compartment of the audit profile resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Details about the current state of the audit trail in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// An audit trail location represents the source of audit records that provides documentary
	// evidence of the sequence of activities in the target database.
	TrailLocation *string `mandatory:"false" json:"trailLocation"`

	// The description of the audit trail.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if auto purge is enabled on the target database, which helps delete audit data in the
	// target database every seven days so that the database's audit trail does not become too large.
	IsAutoPurgeEnabled *bool `mandatory:"false" json:"isAutoPurgeEnabled"`

	// The date from which the audit trail must start collecting data, in the format defined by RFC3339.
	AuditCollectionStartTime *common.SDKTime `mandatory:"false" json:"auditCollectionStartTime"`

	// The OCID of the workrequest for audit trail which collects audit records.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// The date and time until when the audit events were collected from the target database by the Data Safe audit trail
	// collection process, in the format defined by RFC3339.
	TimeLastCollected *common.SDKTime `mandatory:"false" json:"timeLastCollected"`

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

func (m AuditTrail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditTrail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditTrailLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAuditTrailLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuditTrailStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAuditTrailStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
