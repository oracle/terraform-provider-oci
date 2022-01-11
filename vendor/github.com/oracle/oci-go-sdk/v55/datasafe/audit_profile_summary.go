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
	"github.com/oracle/oci-go-sdk/v55/common"
	"strings"
)

// AuditProfileSummary Summary of an audit profile.
type AuditProfileSummary struct {

	// The OCID of the audit profile.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the audit profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the audit profile.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the audit profile was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the audit profile was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the Data Safe target for which the audit profile is created.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The current state of the audit profile.
	LifecycleState AuditProfileLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Indicates if you want to continue audit record collection beyond the free limit
	// of one million audit records per month per target database, incurring additional charges.
	// The default is to continue collection beyond a million audit records.
	// If you do not want to pay after it reaches the one million audit records per target per
	// month limit, please configure the service to stop collecting.
	IsPaidUsageEnabled *bool `mandatory:"true" json:"isPaidUsageEnabled"`

	// Indicates the number of months the audit data will be stored in the Oracle Data
	// Safe audit repository which is online storage and is available for immediate reporting and analysis.
	// You can store up to 12 months of audit data online. The minimum online retention period is 1 month.
	OnlineMonths *int `mandatory:"true" json:"onlineMonths"`

	// Indicates the number of months the audit data will stored in the Object Store which is
	// offline/ archive storage.The minimum archive period you can set is zero months and the maximum you can
	// set is 72 months (6 years). If there is a business need to store the audit data beyond 72 months in
	// archive, contact Oracle Support to increase archive retention to maximum of 168 months (14 years).
	OfflineMonths *int `mandatory:"true" json:"offlineMonths"`

	// Indicates whether audit retention settings like online and offline months is set at the
	// target level overriding the global audit retention settings.
	IsOverrideGlobalRetentionSetting *bool `mandatory:"true" json:"isOverrideGlobalRetentionSetting"`

	// The description of audit profile.
	Description *string `mandatory:"false" json:"description"`

	// Details about the current state of the audit profile in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Indicates count of audit records collected by Data Safe from the target which is eligible
	// for the current month's billing cycle. Audit records for actions performed by Data Safe service
	// account on the target is excluded.
	AuditCollectedVolume *int64 `mandatory:"false" json:"auditCollectedVolume"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AuditProfileSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditProfileSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingAuditProfileLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAuditProfileLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
