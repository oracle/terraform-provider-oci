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

// CreateOracleOfflineLogicalMigrationSettings Create settings for an offline logical Oracle migration.
type CreateOracleOfflineLogicalMigrationSettings struct {
	DataTransferMediumDetails CreateOracleDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	InitialLoadSettings *CreateOracleInitialLoadSettings `mandatory:"false" json:"initialLoadSettings"`

	AdvisorSettings *CreateOracleAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	// List of Migration Parameter objects.
	AdvancedParameters []MigrationParameterDetails `mandatory:"false" json:"advancedParameters"`

	// Database objects to exclude from migration, cannot be specified alongside 'includeObjects'
	ExcludeObjects []OracleDatabaseObject `mandatory:"false" json:"excludeObjects"`

	// Database objects to include from migration, cannot be specified alongside 'excludeObjects'
	IncludeObjects []OracleDatabaseObject `mandatory:"false" json:"includeObjects"`

	// Specifies the database objects to be excluded from the migration in bulk.
	// The definition accepts input in a CSV format, newline separated for each entry.
	// More details can be found in the documentation.
	BulkIncludeExcludeData *string `mandatory:"false" json:"bulkIncludeExcludeData"`
}

func (m CreateOracleOfflineLogicalMigrationSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOracleOfflineLogicalMigrationSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOracleOfflineLogicalMigrationSettings) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOracleOfflineLogicalMigrationSettings CreateOracleOfflineLogicalMigrationSettings
	s := struct {
		DiscriminatorParam string `json:"migrationMethod"`
		MarshalTypeCreateOracleOfflineLogicalMigrationSettings
	}{
		"OFFLINE_LOGICAL",
		(MarshalTypeCreateOracleOfflineLogicalMigrationSettings)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateOracleOfflineLogicalMigrationSettings) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataTransferMediumDetails createoracledatatransfermediumdetails `json:"dataTransferMediumDetails"`
		InitialLoadSettings       *CreateOracleInitialLoadSettings      `json:"initialLoadSettings"`
		AdvisorSettings           *CreateOracleAdvisorSettings          `json:"advisorSettings"`
		AdvancedParameters        []MigrationParameterDetails           `json:"advancedParameters"`
		ExcludeObjects            []OracleDatabaseObject                `json:"excludeObjects"`
		IncludeObjects            []OracleDatabaseObject                `json:"includeObjects"`
		BulkIncludeExcludeData    *string                               `json:"bulkIncludeExcludeData"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.DataTransferMediumDetails.UnmarshalPolymorphicJSON(model.DataTransferMediumDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataTransferMediumDetails = nn.(CreateOracleDataTransferMediumDetails)
	} else {
		m.DataTransferMediumDetails = nil
	}

	m.InitialLoadSettings = model.InitialLoadSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.AdvancedParameters = make([]MigrationParameterDetails, len(model.AdvancedParameters))
	copy(m.AdvancedParameters, model.AdvancedParameters)
	m.ExcludeObjects = make([]OracleDatabaseObject, len(model.ExcludeObjects))
	copy(m.ExcludeObjects, model.ExcludeObjects)
	m.IncludeObjects = make([]OracleDatabaseObject, len(model.IncludeObjects))
	copy(m.IncludeObjects, model.IncludeObjects)
	m.BulkIncludeExcludeData = model.BulkIncludeExcludeData

	return
}
