// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseConnectionCredentialsByName Existing named credential used to connect to the database.
type DatabaseConnectionCredentialsByName struct {

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

func (m DatabaseConnectionCredentialsByName) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseConnectionCredentialsByName) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseConnectionCredentialsByName) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseConnectionCredentialsByName DatabaseConnectionCredentialsByName
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeDatabaseConnectionCredentialsByName
	}{
		"NAME_REFERENCE",
		(MarshalTypeDatabaseConnectionCredentialsByName)(m),
	}

	return json.Marshal(&s)
}
