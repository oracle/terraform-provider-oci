// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDbHomeWithDbSystemIdFromDatabaseDetails Note that a valid `dbSystemId` value must be supplied for the `CreateDbHomeWithDbSystemIdFromDatabase` API operation to successfully complete.
type CreateDbHomeWithDbSystemIdFromDatabaseDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	Database *CreateDatabaseFromAnotherDatabaseDetails `mandatory:"true" json:"database"`

	// The user-provided name of the Database Home.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The database software image OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

//GetFreeformTags returns FreeformTags
func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithDbSystemIdFromDatabaseDetails CreateDbHomeWithDbSystemIdFromDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithDbSystemIdFromDatabaseDetails
	}{
		"DATABASE",
		(MarshalTypeCreateDbHomeWithDbSystemIdFromDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
