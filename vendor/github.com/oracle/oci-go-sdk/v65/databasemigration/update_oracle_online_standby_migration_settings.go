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

// UpdateOracleOnlineStandbyMigrationSettings Update settings for an online standby Oracle migration.
type UpdateOracleOnlineStandbyMigrationSettings struct {
	DataTransferMediumDetails UpdateOracleDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	InitialLoadSettings *UpdateOracleInitialLoadSettings `mandatory:"false" json:"initialLoadSettings"`

	AdvisorSettings *UpdateOracleAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	HubDetails *UpdateGoldenGateHubDetails `mandatory:"false" json:"hubDetails"`

	GgsDetails *UpdateOracleGgsDeploymentDetails `mandatory:"false" json:"ggsDetails"`

	// List of Migration Parameter objects.
	AdvancedParameters []MigrationParameterDetails `mandatory:"false" json:"advancedParameters"`

	// The OCID of the resource being updated.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	// The OCID of the resource being updated.
	SourceStandbyDatabaseConnectionId *string `mandatory:"false" json:"sourceStandbyDatabaseConnectionId"`
}

func (m UpdateOracleOnlineStandbyMigrationSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleOnlineStandbyMigrationSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOracleOnlineStandbyMigrationSettings) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOracleOnlineStandbyMigrationSettings UpdateOracleOnlineStandbyMigrationSettings
	s := struct {
		DiscriminatorParam string `json:"migrationMethod"`
		MarshalTypeUpdateOracleOnlineStandbyMigrationSettings
	}{
		"ONLINE_STANDBY",
		(MarshalTypeUpdateOracleOnlineStandbyMigrationSettings)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateOracleOnlineStandbyMigrationSettings) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataTransferMediumDetails           updateoracledatatransfermediumdetails `json:"dataTransferMediumDetails"`
		InitialLoadSettings                 *UpdateOracleInitialLoadSettings      `json:"initialLoadSettings"`
		AdvisorSettings                     *UpdateOracleAdvisorSettings          `json:"advisorSettings"`
		HubDetails                          *UpdateGoldenGateHubDetails           `json:"hubDetails"`
		GgsDetails                          *UpdateOracleGgsDeploymentDetails     `json:"ggsDetails"`
		AdvancedParameters                  []MigrationParameterDetails           `json:"advancedParameters"`
		SourceContainerDatabaseConnectionId *string                               `json:"sourceContainerDatabaseConnectionId"`
		SourceStandbyDatabaseConnectionId   *string                               `json:"sourceStandbyDatabaseConnectionId"`
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
		m.DataTransferMediumDetails = nn.(UpdateOracleDataTransferMediumDetails)
	} else {
		m.DataTransferMediumDetails = nil
	}

	m.InitialLoadSettings = model.InitialLoadSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.HubDetails = model.HubDetails

	m.GgsDetails = model.GgsDetails

	m.AdvancedParameters = make([]MigrationParameterDetails, len(model.AdvancedParameters))
	copy(m.AdvancedParameters, model.AdvancedParameters)
	m.SourceContainerDatabaseConnectionId = model.SourceContainerDatabaseConnectionId

	m.SourceStandbyDatabaseConnectionId = model.SourceStandbyDatabaseConnectionId

	return
}
