// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AsmConnectionCredentailsByName The existing named credential used to connect to the ASM instance.
type AsmConnectionCredentailsByName struct {

	// The name of the credential information that used to connect to the DB system resource.
	// The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters,
	// and length of "y" has a maximum of 199 characters. The name strings can contain letters,
	// numbers and the underscore character only. Other characters are not valid, except for
	// the "." character that separates the "x" and "y" portions of the name.
	// *IMPORTANT* - The name must be unique within the OCI region the credential is being created in.
	// If you specify a name that duplicates the name of another credential within the same OCI region,
	// you may overwrite or corrupt the credential that is already using the name.
	// For example: inventorydb.abc112233445566778899
	CredentialName *string `mandatory:"true" json:"credentialName"`
}

func (m AsmConnectionCredentailsByName) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AsmConnectionCredentailsByName) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AsmConnectionCredentailsByName) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAsmConnectionCredentailsByName AsmConnectionCredentailsByName
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeAsmConnectionCredentailsByName
	}{
		"NAME_REFERENCE",
		(MarshalTypeAsmConnectionCredentailsByName)(m),
	}

	return json.Marshal(&s)
}
