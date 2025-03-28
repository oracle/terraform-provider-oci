// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccUpgradeInformation Upgrade information that relates to a Compute Cloud@Customer infrastructure. This information
// cannot be updated.
type CccUpgradeInformation struct {

	// The current version of software installed on the Compute Cloud@Customer infrastructure.
	CurrentVersion *string `mandatory:"false" json:"currentVersion"`

	// Compute Cloud@Customer infrastructure next upgrade time. The rack might have performance
	// impacts during this time.
	TimeOfScheduledUpgrade *common.SDKTime `mandatory:"false" json:"timeOfScheduledUpgrade"`

	// Expected duration of Compute Cloud@Customer infrastructure scheduled upgrade. The actual
	// upgrade time might be longer or shorter than this duration depending on rack activity, this
	// is only an estimate.
	ScheduledUpgradeDuration *string `mandatory:"false" json:"scheduledUpgradeDuration"`

	// Indication that the Compute Cloud@Customer infrastructure is in the process of
	// an upgrade or an upgrade activity (such as preloading upgrade images).
	IsActive *bool `mandatory:"false" json:"isActive"`
}

func (m CccUpgradeInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccUpgradeInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
