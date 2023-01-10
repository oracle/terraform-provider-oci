// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// PdbConversionToNewDatabaseDetails Details of the new container database in which the converted pluggable database will be located.
type PdbConversionToNewDatabaseDetails struct {

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 8 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	CdbName *string `mandatory:"true" json:"cdbName"`

	// A strong password for SYS, SYSTEM, and the plugbable database ADMIN user of the container database after conversion. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	CdbAdminPassword *string `mandatory:"true" json:"cdbAdminPassword"`

	// The existing TDE wallet password of the non-container database.
	NonCdbTdeWalletPassword *string `mandatory:"true" json:"nonCdbTdeWalletPassword"`

	// A strong password for plugbable database ADMIN user of the container database after conversion. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	PdbAdminPassword *string `mandatory:"false" json:"pdbAdminPassword"`

	// The password to open the TDE wallet of the container database after conversion. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	CdbTdeWalletPassword *string `mandatory:"false" json:"cdbTdeWalletPassword"`

	// Additional container database parameters.
	// Example: "_pdb_name_case_sensitive=true"
	AdditionalCdbParams *string `mandatory:"false" json:"additionalCdbParams"`
}

func (m PdbConversionToNewDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PdbConversionToNewDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PdbConversionToNewDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePdbConversionToNewDatabaseDetails PdbConversionToNewDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"target"`
		MarshalTypePdbConversionToNewDatabaseDetails
	}{
		"NEW_DATABASE",
		(MarshalTypePdbConversionToNewDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
