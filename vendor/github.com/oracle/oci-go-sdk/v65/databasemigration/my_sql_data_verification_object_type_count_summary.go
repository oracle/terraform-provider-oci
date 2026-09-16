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

// MySqlDataVerificationObjectTypeCountSummary Object type count comparison summary line for MySQL migrations.
type MySqlDataVerificationObjectTypeCountSummary struct {

	// Number of objects in source.
	SourceObjectCount *int `mandatory:"true" json:"sourceObjectCount"`

	// Number of objects in target.
	TargetObjectCount *int `mandatory:"true" json:"targetObjectCount"`

	// Percentage delta between source and target counts.
	// This value is expected to be non-negative.
	// If a negative value is encountered, it indicates that the target has objects not present in the source,
	// which should be treated as an inconsistency (i.e., a failed comparison/report result).
	DeltaPercent *float64 `mandatory:"true" json:"deltaPercent"`

	// Number of invalid objects in source.
	SourceInvalidCount *int `mandatory:"true" json:"sourceInvalidCount"`

	// Number of invalid objects in target.
	TargetInvalidCount *int `mandatory:"true" json:"targetInvalidCount"`

	// Schema/owner name.
	// This field is omitted (null/empty depending on backend serialization) for database-wide object types
	// that are not schema-scoped.
	// Oracle non-schema-scoped object types: USER, ROLE, PROFILE, TABLESPACE, DATABASE_LINK, CONTROLFILE,
	// DATAFILE, REDO_LOG, DIRECTORY, LIBRARY, CONTEXT.
	// MySQL non-schema-scoped object types: USER, ROLE, SERVER, TABLESPACE, LOGFILE_GROUP.
	SchemaName *string `mandatory:"false" json:"schemaName"`

	// MySQL database object type.
	ObjectType MySqlDatabaseObjectTypesEnum `mandatory:"true" json:"objectType"`
}

// GetSchemaName returns SchemaName
func (m MySqlDataVerificationObjectTypeCountSummary) GetSchemaName() *string {
	return m.SchemaName
}

// GetSourceObjectCount returns SourceObjectCount
func (m MySqlDataVerificationObjectTypeCountSummary) GetSourceObjectCount() *int {
	return m.SourceObjectCount
}

// GetTargetObjectCount returns TargetObjectCount
func (m MySqlDataVerificationObjectTypeCountSummary) GetTargetObjectCount() *int {
	return m.TargetObjectCount
}

// GetDeltaPercent returns DeltaPercent
func (m MySqlDataVerificationObjectTypeCountSummary) GetDeltaPercent() *float64 {
	return m.DeltaPercent
}

func (m MySqlDataVerificationObjectTypeCountSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlDataVerificationObjectTypeCountSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMySqlDatabaseObjectTypesEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetMySqlDatabaseObjectTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MySqlDataVerificationObjectTypeCountSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlDataVerificationObjectTypeCountSummary MySqlDataVerificationObjectTypeCountSummary
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlDataVerificationObjectTypeCountSummary
	}{
		"MYSQL",
		(MarshalTypeMySqlDataVerificationObjectTypeCountSummary)(m),
	}

	return json.Marshal(&s)
}
