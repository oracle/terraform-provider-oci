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

// AuditProfileDimensions Details of aggregation dimensions used for summarizing audit profiles.
type AuditProfileDimensions struct {

	// Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database,
	// potentially incurring additional charges. The default value is inherited from the global settings.
	// You can change at the global level or at the target level.
	IsPaidUsageEnabled *bool `mandatory:"false" json:"isPaidUsageEnabled"`

	// The resource type that is represented by the audit profile.
	TargetType *string `mandatory:"false" json:"targetType"`

	// The name or the OCID of the resource from which the online month retention setting is sourced. For example a target database group OCID or global.
	OnlineMonthsSource *string `mandatory:"false" json:"onlineMonthsSource"`

	// The name or the OCID of the resource from which the offline month retention setting is sourced. For example a target database group OCID or global.
	OfflineMonthsSource *string `mandatory:"false" json:"offlineMonthsSource"`

	// The name or the OCID of the resource from which the paid usage setting is sourced. For example a target database group OCID or global.
	PaidUsageSource *string `mandatory:"false" json:"paidUsageSource"`
}

func (m AuditProfileDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditProfileDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
