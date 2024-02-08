// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateProtectionPolicyDetails Describes the parameters required to create a custom protection policy.
type CreateProtectionPolicyDetails struct {

	// A user provided name for the protection policy. The 'displayName' does not have to be unique, and it can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The maximum number of days to retain backups for a protected database.
	BackupRetentionPeriodInDays *int `mandatory:"true" json:"backupRetentionPeriodInDays"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// An RFC3339 formatted datetime string that specifies the exact date and time for the retention lock to take effect and permanently lock the retention period defined in the policy.
	// * The retention lock feature controls whether Recovery Service strictly preserves backups for the duration defined in a policy. Retention lock is useful to enforce recovery window compliance and to prevent unintentional modifications to protected database backups.
	// * Recovery Service enforces a 14-day delay before the retention lock set for a policy can take effect. Therefore, you must set policyLockedDateTime  to a date that occurs 14 days after the current date.
	// * For example, assuming that the current date is Aug 1, 2023 9 pm, you can set policyLockedDateTime  to '2023-08-15T21:00:00.600Z' (Aug 15, 2023, 9:00 pm), or greater.
	// * During the 14-day delay period, you can either increase or decrease the retention period in the policy.
	// * However, you are only allowed to increase the retention period on or after the retention lock date.
	// * You cannot change the value of policyLockedDateTime if the retention lock is already in effect.
	PolicyLockedDateTime *string `mandatory:"false" json:"policyLockedDateTime"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. For more information, see Resource Tags (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm)
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateProtectionPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateProtectionPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
