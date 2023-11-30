// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// MySqlMigrationObjectCollection Database objects to migrate.
type MySqlMigrationObjectCollection struct {

	// Database objects to exclude/include from migration
	Items []MySqlDatabaseObjectSummary `mandatory:"true" json:"items"`

	// Database objects to exclude/include from migration in CSV format. The excludeObjects and includeObjects
	// fields will be ignored if this field is not null.
	BulkIncludeExcludeData *string `mandatory:"false" json:"bulkIncludeExcludeData"`
}

func (m MySqlMigrationObjectCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlMigrationObjectCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MySqlMigrationObjectCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlMigrationObjectCollection MySqlMigrationObjectCollection
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlMigrationObjectCollection
	}{
		"MYSQL",
		(MarshalTypeMySqlMigrationObjectCollection)(m),
	}

	return json.Marshal(&s)
}
