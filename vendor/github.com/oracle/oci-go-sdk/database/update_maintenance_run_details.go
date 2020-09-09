// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateMaintenanceRunDetails Describes the modification parameters for the maintenance run.
type UpdateMaintenanceRunDetails struct {

	// If `FALSE`, skips the maintenance run.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The scheduled date and time of the maintenance run to update.
	TimeScheduled *common.SDKTime `mandatory:"false" json:"timeScheduled"`

	// If set to `TRUE`, starts patching immediately.
	IsPatchNowEnabled *bool `mandatory:"false" json:"isPatchNowEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
	PatchId *string `mandatory:"false" json:"patchId"`
}

func (m UpdateMaintenanceRunDetails) String() string {
	return common.PointerString(m)
}
