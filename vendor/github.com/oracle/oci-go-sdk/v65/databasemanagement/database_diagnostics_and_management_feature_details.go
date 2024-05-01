// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseDiagnosticsAndManagementFeatureDetails The details required to enable the Diagnostics and Management feature.
type DatabaseDiagnosticsAndManagementFeatureDetails struct {
	DatabaseConnectionDetails *DatabaseConnectionDetails `mandatory:"true" json:"databaseConnectionDetails"`

	ConnectorDetails ConnectorDetails `mandatory:"true" json:"connectorDetails"`

	// Indicates whether the pluggable database can be enabled automatically.
	IsAutoEnablePluggableDatabase *bool `mandatory:"false" json:"isAutoEnablePluggableDatabase"`

	// The management type for the database.
	ManagementType DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum `mandatory:"true" json:"managementType"`
}

// GetDatabaseConnectionDetails returns DatabaseConnectionDetails
func (m DatabaseDiagnosticsAndManagementFeatureDetails) GetDatabaseConnectionDetails() *DatabaseConnectionDetails {
	return m.DatabaseConnectionDetails
}

// GetConnectorDetails returns ConnectorDetails
func (m DatabaseDiagnosticsAndManagementFeatureDetails) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

func (m DatabaseDiagnosticsAndManagementFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseDiagnosticsAndManagementFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum(string(m.ManagementType)); !ok && m.ManagementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementType: %s. Supported values are: %s.", m.ManagementType, strings.Join(GetDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseDiagnosticsAndManagementFeatureDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseDiagnosticsAndManagementFeatureDetails DatabaseDiagnosticsAndManagementFeatureDetails
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeDatabaseDiagnosticsAndManagementFeatureDetails
	}{
		"DIAGNOSTICS_AND_MANAGEMENT",
		(MarshalTypeDatabaseDiagnosticsAndManagementFeatureDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseDiagnosticsAndManagementFeatureDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsAutoEnablePluggableDatabase *bool                                                            `json:"isAutoEnablePluggableDatabase"`
		DatabaseConnectionDetails     *DatabaseConnectionDetails                                       `json:"databaseConnectionDetails"`
		ConnectorDetails              connectordetails                                                 `json:"connectorDetails"`
		ManagementType                DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum `json:"managementType"`
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

// DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum Enum with underlying type: string
type DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum string

// Set of constants representing the allowable values for DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum
const (
	DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeBasic    DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum = "BASIC"
	DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeAdvanced DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum = "ADVANCED"
)

var mappingDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum = map[string]DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum{
	"BASIC":    DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeBasic,
	"ADVANCED": DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeAdvanced,
}

var mappingDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnumLowerCase = map[string]DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum{
	"basic":    DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeBasic,
	"advanced": DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeAdvanced,
}

// GetDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnumValues Enumerates the set of values for DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum
func GetDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnumValues() []DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum {
	values := make([]DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum, 0)
	for _, v := range mappingDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnumStringValues Enumerates the set of values in String for DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum
func GetDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
	}
}

// GetMappingDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum(val string) (DatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnum, bool) {
	enum, ok := mappingDatabaseDiagnosticsAndManagementFeatureDetailsManagementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
