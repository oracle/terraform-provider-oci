// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v32/common"
)

// DatabaseUpgradeWithDatabaseSoftwareImageDetails Details of Database Software Image for upgrading a database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DatabaseUpgradeWithDatabaseSoftwareImageDetails struct {

	// the database software id used for upgrading the database.
	DatabaseSoftwareImageId *string `mandatory:"true" json:"databaseSoftwareImageId"`
}

func (m DatabaseUpgradeWithDatabaseSoftwareImageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DatabaseUpgradeWithDatabaseSoftwareImageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseUpgradeWithDatabaseSoftwareImageDetails DatabaseUpgradeWithDatabaseSoftwareImageDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDatabaseUpgradeWithDatabaseSoftwareImageDetails
	}{
		"DB_SOFTWARE_IMAGE",
		(MarshalTypeDatabaseUpgradeWithDatabaseSoftwareImageDetails)(m),
	}

	return json.Marshal(&s)
}
