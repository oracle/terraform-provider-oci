// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseDiagnosticsAndPerformanceFeatureDetails The details required to enable diagnostics and performance feature.
type DatabaseDiagnosticsAndPerformanceFeatureDetails struct {
	DatabaseConnectionDetails *DatabaseConnectionDetails `mandatory:"true" json:"databaseConnectionDetails"`

	ConnectorDetails ConnectorDetails `mandatory:"true" json:"connectorDetails"`

	// Indicates whether the pluggable database can be enabled automatically.
	IsAutoEnablePluggableDatabase *bool `mandatory:"false" json:"isAutoEnablePluggableDatabase"`

	// management type for the database.
	ManagementType DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum `mandatory:"true" json:"managementType"`
}

// GetDatabaseConnectionDetails returns DatabaseConnectionDetails
func (m DatabaseDiagnosticsAndPerformanceFeatureDetails) GetDatabaseConnectionDetails() *DatabaseConnectionDetails {
	return m.DatabaseConnectionDetails
}

// GetConnectorDetails returns ConnectorDetails
func (m DatabaseDiagnosticsAndPerformanceFeatureDetails) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

func (m DatabaseDiagnosticsAndPerformanceFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseDiagnosticsAndPerformanceFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum(string(m.ManagementType)); !ok && m.ManagementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementType: %s. Supported values are: %s.", m.ManagementType, strings.Join(GetDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseDiagnosticsAndPerformanceFeatureDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseDiagnosticsAndPerformanceFeatureDetails DatabaseDiagnosticsAndPerformanceFeatureDetails
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeDatabaseDiagnosticsAndPerformanceFeatureDetails
	}{
		"DIAGNOSTICS_AND_PERFORMANCE",
		(MarshalTypeDatabaseDiagnosticsAndPerformanceFeatureDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseDiagnosticsAndPerformanceFeatureDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsAutoEnablePluggableDatabase *bool                                                             `json:"isAutoEnablePluggableDatabase"`
		DatabaseConnectionDetails     *DatabaseConnectionDetails                                        `json:"databaseConnectionDetails"`
		ConnectorDetails              connectordetails                                                  `json:"connectorDetails"`
		ManagementType                DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum `json:"managementType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsAutoEnablePluggableDatabase = model.IsAutoEnablePluggableDatabase

	m.DatabaseConnectionDetails = model.DatabaseConnectionDetails

	nn, e = model.ConnectorDetails.UnmarshalPolymorphicJSON(model.ConnectorDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectorDetails = nn.(ConnectorDetails)
	} else {
		m.ConnectorDetails = nil
	}

	m.ManagementType = model.ManagementType

	return
}

// DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum Enum with underlying type: string
type DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum string

// Set of constants representing the allowable values for DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum
const (
	DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeBasic    DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum = "BASIC"
	DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeAdvanced DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum = "ADVANCED"
)

var mappingDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum = map[string]DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum{
	"BASIC":    DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeBasic,
	"ADVANCED": DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeAdvanced,
}

var mappingDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnumLowerCase = map[string]DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum{
	"basic":    DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeBasic,
	"advanced": DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeAdvanced,
}

// GetDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnumValues Enumerates the set of values for DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum
func GetDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnumValues() []DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum {
	values := make([]DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum, 0)
	for _, v := range mappingDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnumStringValues Enumerates the set of values in String for DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum
func GetDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
	}
}

// GetMappingDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum(val string) (DatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnum, bool) {
	enum, ok := mappingDatabaseDiagnosticsAndPerformanceFeatureDetailsManagementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
