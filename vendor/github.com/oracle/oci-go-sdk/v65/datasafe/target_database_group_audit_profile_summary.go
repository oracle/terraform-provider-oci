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

// TargetDatabaseGroupAuditProfileSummary Summary of an target group audit profile.
type TargetDatabaseGroupAuditProfileSummary struct {

	// The OCID of the target group audit profile.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the target group audit profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the target group audit profile.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the target group audit profile was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the target group audit profile was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the Data Safe target group for which the target group audit profile is created.
	TargetDatabaseGroupId *string `mandatory:"true" json:"targetDatabaseGroupId"`

	// The current state of the target group audit profile.
	LifecycleState TargetDatabaseGroupAuditProfileLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database,
	// potentially incurring additional charges. The default value is inherited from the global settings.
	// You can change at the global level or at the target level.
	IsPaidUsageEnabled *bool `mandatory:"true" json:"isPaidUsageEnabled"`

	// Indicates the number of months the audit records will be stored online in Oracle Data Safe audit repository for immediate reporting and analysis.
	// Minimum: 1; Maximum: 12 months
	OnlineMonths *int `mandatory:"true" json:"onlineMonths"`

	// Indicates the number of months the audit records will be stored offline in the Data Safe audit archive.
	// Minimum: 0; Maximum: 72 months.
	// If you have a requirement to store the audit data even longer (up to 168 months) in archive, please contact the Oracle Support.
	OfflineMonths *int `mandatory:"true" json:"offlineMonths"`

	// The description of audit profile.
	Description *string `mandatory:"false" json:"description"`

	// Details about the current state of the target group audit profile in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

func (m TargetDatabaseGroupAuditProfileSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetDatabaseGroupAuditProfileSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetDatabaseGroupAuditProfileLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetDatabaseGroupAuditProfileLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
