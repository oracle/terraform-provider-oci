// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AuditProfile The resource represents audit profile settings and audit configurations for the database target, and helps evaluate the initial audit data volume for configuring collection in Data Safe. The resource is also responsible for auto-discovery of audit trails in the database target during target's registration.
type AuditProfile struct {

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

	// The current state of the audit profile.
	LifecycleState AuditProfileLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the target database or target database group for which the audit profile is created.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database,
	// potentially incurring additional charges. The default value is inherited from the global settings.
	// You can change at the global level or at the target level.
	IsPaidUsageEnabled *bool `mandatory:"true" json:"isPaidUsageEnabled"`

	// Number of months the audit records will be stored online in the audit repository for immediate reporting and analysis.
	// Minimum: 1; Maximum: 12 months
	OnlineMonths *int `mandatory:"true" json:"onlineMonths"`

	// Number of months the audit records will be stored offline in the offline archive.
	// Minimum: 0; Maximum: 72 months.
	// If you have a requirement to store the audit data even longer in the offline archive, please contact the Oracle Support.
	OfflineMonths *int `mandatory:"true" json:"offlineMonths"`

	// Indicates whether audit retention settings like online and offline months set at the
	// target level override both the global settings and the target group level audit retention settings.
	IsOverrideGlobalRetentionSetting *bool `mandatory:"true" json:"isOverrideGlobalRetentionSetting"`

	// Details about the current state of the audit profile in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The description of the audit profile.
	Description *string `mandatory:"false" json:"description"`

	// Contains the list of available audit trails on the target database.
	AuditTrails []AuditTrail `mandatory:"false" json:"auditTrails"`

	// Number of audit records collected in the current calendar month.
	// Audit records for the Data Safe service account are excluded and are not counted towards your monthly free limit.
	AuditCollectedVolume *int64 `mandatory:"false" json:"auditCollectedVolume"`

	// Indicates whether audit paid usage settings specified at the target database level override both the global settings and the target group level paid usage settings.
	// Enabling paid usage continues the collection of audit records beyond the free limit of one million audit records per month per target database,
	// potentially incurring additional charges. For more information, see Data Safe Price List (https://www.oracle.com/cloud/price-list/#data-safe).
	IsOverrideGlobalPaidUsage *bool `mandatory:"false" json:"isOverrideGlobalPaidUsage"`

	// The name or the OCID of the resource from which the online month retention setting is sourced. For example, a global setting or a target database group OCID.
	OnlineMonthsSource *string `mandatory:"false" json:"onlineMonthsSource"`

	// The name or the OCID of the resource from which the offline month retention setting is sourced. For example, a global setting or a target database group OCID.
	OfflineMonthsSource *string `mandatory:"false" json:"offlineMonthsSource"`

	// The name or the OCID of the resource from which the paid usage setting is sourced. For example, a global setting or a target database group OCID.
	PaidUsageSource *string `mandatory:"false" json:"paidUsageSource"`

	// The resource type that is represented by the audit profile.
	TargetType AuditProfileTargetTypeEnum `mandatory:"false" json:"targetType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AuditProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditProfile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditProfileLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAuditProfileLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAuditProfileTargetTypeEnum(string(m.TargetType)); !ok && m.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", m.TargetType, strings.Join(GetAuditProfileTargetTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
