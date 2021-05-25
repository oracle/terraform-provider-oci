// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// UpdateDatabaseDetails Details to update a database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateDatabaseDetails struct {
	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// A new strong password for SYS, SYSTEM, and the plugbable database ADMIN user. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	NewAdminPassword *string `mandatory:"false" json:"newAdminPassword"`

	// The existing TDE wallet password. You must provide the existing password in order to set a new TDE wallet password.
	OldTdeWalletPassword *string `mandatory:"false" json:"oldTdeWalletPassword"`

	// The new password to open the TDE wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	NewTdeWalletPassword *string `mandatory:"false" json:"newTdeWalletPassword"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateDatabaseDetails) String() string {
	return common.PointerString(m)
}
