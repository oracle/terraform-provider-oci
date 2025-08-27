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

// TargetOverrideSummary Summary of the target database-specific profiles that override the audit profile of the target database group.
type TargetOverrideSummary struct {

	// The OCID of the target database that overrides from the audit profile of the target database group.
	TargetDatabaseId *string `mandatory:"true" json:"targetDatabaseId"`

	// Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database,
	// potentially incurring additional charges. The default value is inherited from the global settings.
	// You can change at the global level or at the target level.
	IsPaidUsageEnabled *bool `mandatory:"true" json:"isPaidUsageEnabled"`

	// Number of months the audit records will be stored online in the audit repository for immediate reporting and analysis.
	// Minimum: 1; Maximum: 12 months
	OnlineMonths *int `mandatory:"true" json:"onlineMonths"`

	// Number of months the audit records will be stored offline in the offline archive.
	// Minimum: 0; Maximum: 72 months.
	// If you have a requirement to store the audit data even longer (up to 168 months) in the offline archive, please contact the Oracle Support.
	OfflineMonths *int `mandatory:"true" json:"offlineMonths"`

	// The name or the OCID of the resource from which the online month retention setting is sourced. For example a target database group OCID or global.
	OnlineMonthsSource *string `mandatory:"false" json:"onlineMonthsSource"`

	// The name or the OCID of the resource from which the offline month retention setting is sourced. For example a target database group OCID or global.
	OfflineMonthsSource *string `mandatory:"false" json:"offlineMonthsSource"`

	// The name or the OCID of the resource from which the paid usage setting is sourced. For example a target database group OCID or global.
	PaidUsageSource *string `mandatory:"false" json:"paidUsageSource"`

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

func (m TargetOverrideSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetOverrideSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
