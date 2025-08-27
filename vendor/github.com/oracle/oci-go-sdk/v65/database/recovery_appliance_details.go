// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecoveryApplianceDetails Information about the recovery appliance configuration associated with the Autonomous Container Database.
type RecoveryApplianceDetails struct {

	// The storage size of the backup destination allocated for an Autonomous Container Database to store backups on the recovery appliance, in GBs, rounded to the nearest integer.
	AllocatedStorageSizeInGBs *int `mandatory:"false" json:"allocatedStorageSizeInGBs"`

	// Number of days between the current and earliest point of recoverability covered by automatic backups.
	RecoveryWindowInDays *int `mandatory:"false" json:"recoveryWindowInDays"`

	// The time when the recovery appliance details are updated.
	TimeRecoveryApplianceDetailsUpdated *common.SDKTime `mandatory:"false" json:"timeRecoveryApplianceDetailsUpdated"`
}

func (m RecoveryApplianceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecoveryApplianceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
