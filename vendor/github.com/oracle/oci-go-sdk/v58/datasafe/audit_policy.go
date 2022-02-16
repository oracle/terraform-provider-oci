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

// AuditPolicy The resource represents all available audit policies relevant for the target database with their corresponding audit conditions.
// The audit policies could be in any one of the following 3 states in the target database
// 1) Created and enabled
// 2) Created but not enabled
// 3) Not created
// For more details on available audit policies, refer to documentation (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/audit-policies.html#GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827).
type AuditPolicy struct {

	// The OCID of the audit policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the audit policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the audit policy.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the the audit policy was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the audit policy.
	LifecycleState AuditPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the target for which the audit policy is created.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Option provided to users at the target to indicate whether the Data Safe service account has to be excluded while provisioning the audit policies.
	IsDataSafeServiceAccountExcluded *bool `mandatory:"true" json:"isDataSafeServiceAccountExcluded"`

	// Description of the audit policy.
	Description *string `mandatory:"false" json:"description"`

	// The last date and time the audit policy was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Details about the current state of the audit policy in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Indicates the last provisioning time of audit policies on the target, in the format defined by RFC3339.
	TimeLastProvisioned *common.SDKTime `mandatory:"false" json:"timeLastProvisioned"`

	// The time when the audit policies was last retrieved from this target, in the format defined by RFC3339.
	TimeLastRetrieved *common.SDKTime `mandatory:"false" json:"timeLastRetrieved"`

	// Represents all available audit policy specifications relevant for the target database. For more details on available audit polcies, refer to documentation (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/audit-policies.html#GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827).
	AuditSpecifications []AuditSpecification `mandatory:"false" json:"auditSpecifications"`

	// Lists the audit policy provisioning conditions for the target database.
	AuditConditions []AuditConditions `mandatory:"false" json:"auditConditions"`

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

func (m AuditPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAuditPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
