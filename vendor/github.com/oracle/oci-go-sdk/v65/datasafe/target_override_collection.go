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

// TargetOverrideCollection Collection of target override summary.
type TargetOverrideCollection struct {

	// Number of target databases within the target database group.
	TargetsCount *int `mandatory:"true" json:"targetsCount"`

	// Number of target databases within the target database group that override the audit profile of the target database group.
	TargetsOverridingCount *int `mandatory:"true" json:"targetsOverridingCount"`

	// Number of target databases within the target database group that conform with the audit profile of the target database group.
	TargetsConformingCount *int `mandatory:"true" json:"targetsConformingCount"`

	// Number of target databases within the group that override the paid usage setting of the audit profile for the target database group.
	TargetsOverridingPaidUsageCount *int `mandatory:"true" json:"targetsOverridingPaidUsageCount"`

	// Number of target databases within the group that override the online retention setting of the audit profile for the target database group.
	TargetsOverridingOnlineMonthsCount *int `mandatory:"true" json:"targetsOverridingOnlineMonthsCount"`

	// Number of target databases within the group that override the offline retention setting of the audit profile for the target database group.
	TargetsOverridingOfflineMonthsCount *int `mandatory:"true" json:"targetsOverridingOfflineMonthsCount"`

	// Array of target database override summary.
	Items []TargetOverrideSummary `mandatory:"true" json:"items"`
}

func (m TargetOverrideCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetOverrideCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
