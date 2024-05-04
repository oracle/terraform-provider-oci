// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleMigrationObjectCollection Database objects to migrate.
type OracleMigrationObjectCollection struct {

	// Database objects to exclude/include from migration
	Items []OracleDatabaseObjectSummary `mandatory:"true" json:"items"`

	// Specifies the database objects to be excluded from the migration in bulk.
	// The definition accepts input in a CSV format, newline separated for each entry.
	// More details can be found in the documentation.
	BulkIncludeExcludeData *string `mandatory:"false" json:"bulkIncludeExcludeData"`
}

func (m OracleMigrationObjectCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleMigrationObjectCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleMigrationObjectCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleMigrationObjectCollection OracleMigrationObjectCollection
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeOracleMigrationObjectCollection
	}{
		"ORACLE",
		(MarshalTypeOracleMigrationObjectCollection)(m),
	}

	return json.Marshal(&s)
}
