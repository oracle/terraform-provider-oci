// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AvailableOverrideMaintenanceStartTimeSummary Information about the list of available start times on a particular date.
// User can choose their preferred date and start time for creating request/input for Override operation.
type AvailableOverrideMaintenanceStartTimeSummary struct {

	// The date corresponding to the list of start times available.
	// Example: `2024-04-25T00:00:00.000Z`
	TimeDateAvailable *common.SDKTime `mandatory:"true" json:"timeDateAvailable"`

	// List of available start times. Each array item is of the format `HH:mm`
	StartTimes []string `mandatory:"true" json:"startTimes"`
}

func (m AvailableOverrideMaintenanceStartTimeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AvailableOverrideMaintenanceStartTimeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
