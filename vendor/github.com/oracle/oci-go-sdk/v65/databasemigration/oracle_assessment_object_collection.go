// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// OracleAssessmentObjectCollection List of affected database objects.
type OracleAssessmentObjectCollection struct {

	// An array of database objects.
	Items []OracleDatabaseObjectSummary `mandatory:"true" json:"items"`

	// Specifies the database objects to be excluded from the migration in bulk.
	// The definition accepts input in a CSV format, newline separated for each entry.
	// More details can be found in the documentation.
	BulkIncludeExcludeData *string `mandatory:"false" json:"bulkIncludeExcludeData"`
}

func (m OracleAssessmentObjectCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleAssessmentObjectCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleAssessmentObjectCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleAssessmentObjectCollection OracleAssessmentObjectCollection
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeOracleAssessmentObjectCollection
	}{
		"ORACLE",
		(MarshalTypeOracleAssessmentObjectCollection)(m),
	}

	return json.Marshal(&s)
}
