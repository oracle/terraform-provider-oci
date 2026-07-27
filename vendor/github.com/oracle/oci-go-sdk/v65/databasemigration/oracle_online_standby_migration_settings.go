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

// OracleOnlineStandbyMigrationSettings Oracle online standby migration settings.
type OracleOnlineStandbyMigrationSettings struct {
	DataTransferMediumDetails OracleDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	InitialLoadSettings *OracleInitialLoadSettings `mandatory:"false" json:"initialLoadSettings"`

	AdvisorSettings *OracleAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	HubDetails *GoldenGateHubDetails `mandatory:"false" json:"hubDetails"`

	GgsDetails *OracleGgsDeploymentDetails `mandatory:"false" json:"ggsDetails"`

	// The OCID of the resource being referenced.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	// The OCID of the resource being referenced.
	SourceStandbyDatabaseConnectionId *string `mandatory:"false" json:"sourceStandbyDatabaseConnectionId"`

	// List of Migration Parameter objects.
	AdvancedParameters []MigrationParameterDetails `mandatory:"false" json:"advancedParameters"`
}

func (m OracleOnlineStandbyMigrationSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleOnlineStandbyMigrationSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleOnlineStandbyMigrationSettings) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleOnlineStandbyMigrationSettings OracleOnlineStandbyMigrationSettings
	s := struct {
		DiscriminatorParam string `json:"migrationMethod"`
		MarshalTypeOracleOnlineStandbyMigrationSettings
	}{
		"ONLINE_STANDBY",
		(MarshalTypeOracleOnlineStandbyMigrationSettings)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OracleOnlineStandbyMigrationSettings) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataTransferMediumDetails           oracledatatransfermediumdetails `json:"dataTransferMediumDetails"`
		InitialLoadSettings                 *OracleInitialLoadSettings      `json:"initialLoadSettings"`
		AdvisorSettings                     *OracleAdvisorSettings          `json:"advisorSettings"`
		HubDetails                          *GoldenGateHubDetails           `json:"hubDetails"`
		GgsDetails                          *OracleGgsDeploymentDetails     `json:"ggsDetails"`
		SourceContainerDatabaseConnectionId *string                         `json:"sourceContainerDatabaseConnectionId"`
		SourceStandbyDatabaseConnectionId   *string                         `json:"sourceStandbyDatabaseConnectionId"`
		AdvancedParameters                  []MigrationParameterDetails     `json:"advancedParameters"`
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
		m.DataTransferMediumDetails = nn.(OracleDataTransferMediumDetails)
	} else {
		m.DataTransferMediumDetails = nil
	}

	m.InitialLoadSettings = model.InitialLoadSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.HubDetails = model.HubDetails

	m.GgsDetails = model.GgsDetails

	m.SourceContainerDatabaseConnectionId = model.SourceContainerDatabaseConnectionId

	m.SourceStandbyDatabaseConnectionId = model.SourceStandbyDatabaseConnectionId

	m.AdvancedParameters = make([]MigrationParameterDetails, len(model.AdvancedParameters))
	copy(m.AdvancedParameters, model.AdvancedParameters)
	return
}
