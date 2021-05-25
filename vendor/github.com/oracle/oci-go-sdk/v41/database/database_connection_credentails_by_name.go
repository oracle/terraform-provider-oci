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
	"github.com/oracle/oci-go-sdk/v41/common"
)

// DatabaseConnectionCredentailsByName Existing named credential used to connect to the database.
type DatabaseConnectionCredentailsByName struct {

	// The name of the credential information that used to connect to the database. The name should be in "x.y" format, where
	// the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters.
	// The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for
	// the "." character that separates the "x" and "y" portions of the name.
	// *IMPORTANT* - The name must be unique within the OCI region the credential is being created in. If you specify a name
	// that duplicates the name of another credential within the same OCI region, you may overwrite or corrupt the credential that is already
	// using the name.
	// For example: inventorydb.abc112233445566778899
	CredentialName *string `mandatory:"true" json:"credentialName"`
}

func (m DatabaseConnectionCredentailsByName) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DatabaseConnectionCredentailsByName) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseConnectionCredentailsByName DatabaseConnectionCredentailsByName
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeDatabaseConnectionCredentailsByName
	}{
		"NAME_REFERENCE",
		(MarshalTypeDatabaseConnectionCredentailsByName)(m),
	}

	return json.Marshal(&s)
}
