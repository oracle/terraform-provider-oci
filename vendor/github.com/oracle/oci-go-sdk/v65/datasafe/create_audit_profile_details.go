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

// CreateAuditProfileDetails The details used to create a new audit profile.
type CreateAuditProfileDetails struct {

	// The OCID of the compartment where you want to create the audit profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the target database or target database group for which the audit profile is created.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The resource type that is represented by the audit profile.
	TargetType AuditProfileTargetTypeEnum `mandatory:"true" json:"targetType"`

	// The display name of the audit profile. The name does not have to be unique, and it's updatable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the audit profile.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database,
	// potentially incurring additional charges. The default value is inherited from the global settings.
	// You can change at the global level or at the target level.
	IsPaidUsageEnabled *bool `mandatory:"false" json:"isPaidUsageEnabled"`

	// Number of months the audit records will be stored online in the audit repository for immediate reporting and analysis.
	// Minimum: 1; Maximum: 12 months
	OnlineMonths *int `mandatory:"false" json:"onlineMonths"`

	// Number of months the audit records will be stored offline in the offline archive.
	// Minimum: 0; Maximum: 72 months.
	// If you have a requirement to store the audit data even longer in the offline archive, please contact the Oracle Support.
	OfflineMonths *int `mandatory:"false" json:"offlineMonths"`

	// Indicates whether audit paid usage settings specified at the target database level override both the global and the target database group level paid usage settings.
	// Enabling paid usage continues the collection of audit records beyond the free limit of one million audit records per month per target database,
	// potentially incurring additional charges. For more information, see Data Safe Price List (https://www.oracle.com/cloud/price-list/#data-safe).
	IsOverrideGlobalPaidUsage *bool `mandatory:"false" json:"isOverrideGlobalPaidUsage"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAuditProfileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAuditProfileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditProfileTargetTypeEnum(string(m.TargetType)); !ok && m.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", m.TargetType, strings.Join(GetAuditProfileTargetTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
