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

// CreateMigrationDetails Create Migration resource parameters.
type CreateMigrationDetails struct {

	// Migration type.
	Type MigrationTypesEnum `mandatory:"true" json:"type"`

	// OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Source Database Connection.
	SourceDatabaseConnectionId *string `mandatory:"true" json:"sourceDatabaseConnectionId"`

	// The OCID of the Target Database Connection.
	TargetDatabaseConnectionId *string `mandatory:"true" json:"targetDatabaseConnectionId"`

	// Migration Display Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the registered ODMS Agent. Only valid for Offline Logical Migrations.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The OCID of the Source Container Database Connection. Only used for Online migrations.
	// Only Connections of type Non-Autonomous can be used as source container databases.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	DataTransferMediumDetailsV2 DataTransferMediumDetailsV2 `mandatory:"false" json:"dataTransferMediumDetailsV2"`

	DataTransferMediumDetails *CreateDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	DumpTransferDetails *CreateDumpTransferDetails `mandatory:"false" json:"dumpTransferDetails"`

	DatapumpSettings *CreateDataPumpSettings `mandatory:"false" json:"datapumpSettings"`

	AdvisorSettings *CreateAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	// Database objects to exclude from migration, cannot be specified alongside 'includeObjects'
	ExcludeObjects []DatabaseObject `mandatory:"false" json:"excludeObjects"`

	// Database objects to include from migration, cannot be specified alongside 'excludeObjects'
	IncludeObjects []DatabaseObject `mandatory:"false" json:"includeObjects"`

	// Database objects to exclude/include from migration in CSV format. The excludeObjects and includeObjects fields will be ignored if this field is not null.
	CsvText *string `mandatory:"false" json:"csvText"`

	GoldenGateDetails *CreateGoldenGateDetails `mandatory:"false" json:"goldenGateDetails"`

	GoldenGateServiceDetails *CreateGoldenGateServiceDetails `mandatory:"false" json:"goldenGateServiceDetails"`

	VaultDetails *CreateVaultDetails `mandatory:"false" json:"vaultDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMigrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrationTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMigrationTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                         *string                           `json:"displayName"`
		AgentId                             *string                           `json:"agentId"`
		SourceContainerDatabaseConnectionId *string                           `json:"sourceContainerDatabaseConnectionId"`
		DataTransferMediumDetailsV2         datatransfermediumdetailsv2       `json:"dataTransferMediumDetailsV2"`
		DataTransferMediumDetails           *CreateDataTransferMediumDetails  `json:"dataTransferMediumDetails"`
		DumpTransferDetails                 *CreateDumpTransferDetails        `json:"dumpTransferDetails"`
		DatapumpSettings                    *CreateDataPumpSettings           `json:"datapumpSettings"`
		AdvisorSettings                     *CreateAdvisorSettings            `json:"advisorSettings"`
		ExcludeObjects                      []DatabaseObject                  `json:"excludeObjects"`
		IncludeObjects                      []DatabaseObject                  `json:"includeObjects"`
		CsvText                             *string                           `json:"csvText"`
		GoldenGateDetails                   *CreateGoldenGateDetails          `json:"goldenGateDetails"`
		GoldenGateServiceDetails            *CreateGoldenGateServiceDetails   `json:"goldenGateServiceDetails"`
		VaultDetails                        *CreateVaultDetails               `json:"vaultDetails"`
		FreeformTags                        map[string]string                 `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{} `json:"definedTags"`
		Type                                MigrationTypesEnum                `json:"type"`
		CompartmentId                       *string                           `json:"compartmentId"`
		SourceDatabaseConnectionId          *string                           `json:"sourceDatabaseConnectionId"`
		TargetDatabaseConnectionId          *string                           `json:"targetDatabaseConnectionId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.AgentId = model.AgentId

	m.SourceContainerDatabaseConnectionId = model.SourceContainerDatabaseConnectionId

	nn, e = model.DataTransferMediumDetailsV2.UnmarshalPolymorphicJSON(model.DataTransferMediumDetailsV2.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataTransferMediumDetailsV2 = nn.(DataTransferMediumDetailsV2)
	} else {
		m.DataTransferMediumDetailsV2 = nil
	}

	m.DataTransferMediumDetails = model.DataTransferMediumDetails

	m.DumpTransferDetails = model.DumpTransferDetails

	m.DatapumpSettings = model.DatapumpSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.ExcludeObjects = make([]DatabaseObject, len(model.ExcludeObjects))
	copy(m.ExcludeObjects, model.ExcludeObjects)
	m.IncludeObjects = make([]DatabaseObject, len(model.IncludeObjects))
	copy(m.IncludeObjects, model.IncludeObjects)
	m.CsvText = model.CsvText

	m.GoldenGateDetails = model.GoldenGateDetails

	m.GoldenGateServiceDetails = model.GoldenGateServiceDetails

	m.VaultDetails = model.VaultDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Type = model.Type

	m.CompartmentId = model.CompartmentId

	m.SourceDatabaseConnectionId = model.SourceDatabaseConnectionId

	m.TargetDatabaseConnectionId = model.TargetDatabaseConnectionId

	return
}
