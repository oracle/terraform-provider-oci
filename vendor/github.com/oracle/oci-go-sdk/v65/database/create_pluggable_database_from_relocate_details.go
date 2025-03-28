// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePluggableDatabaseFromRelocateDetails Specifies the creation type Relocate.
// Additional input 'dblinkUsername` and `dblinkUserPassword` can be provided for Relocate Operation.
// If not provided, Backend will create a temporary user to perform Relocate operation.
type CreatePluggableDatabaseFromRelocateDetails struct {

	// The OCID of the Source Pluggable Database.
	SourcePluggableDatabaseId *string `mandatory:"true" json:"sourcePluggableDatabaseId"`

	// The DB system administrator password of the source Container Database.
	SourceContainerDatabaseAdminPassword *string `mandatory:"true" json:"sourceContainerDatabaseAdminPassword"`

	// The name of the DB link user.
	DblinkUsername *string `mandatory:"false" json:"dblinkUsername"`

	// The DB link user password.
	DblinkUserPassword *string `mandatory:"false" json:"dblinkUserPassword"`
}

func (m CreatePluggableDatabaseFromRelocateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePluggableDatabaseFromRelocateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreatePluggableDatabaseFromRelocateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreatePluggableDatabaseFromRelocateDetails CreatePluggableDatabaseFromRelocateDetails
	s := struct {
		DiscriminatorParam string `json:"creationType"`
		MarshalTypeCreatePluggableDatabaseFromRelocateDetails
	}{
		"RELOCATE_PDB",
		(MarshalTypeCreatePluggableDatabaseFromRelocateDetails)(m),
	}

	return json.Marshal(&s)
}
