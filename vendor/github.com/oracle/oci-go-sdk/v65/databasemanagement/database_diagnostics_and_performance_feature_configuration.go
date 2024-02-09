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

// DatabaseDiagnosticsAndPerformanceFeatureConfiguration The details required to enable diagnostics and performance feature.
type DatabaseDiagnosticsAndPerformanceFeatureConfiguration struct {
	ConnectorDetails ConnectorDetails `mandatory:"false" json:"connectorDetails"`

	DatabaseConnectionDetails *DatabaseConnectionDetails `mandatory:"false" json:"databaseConnectionDetails"`

	// The Oracle license model that applies to the external database.
	LicenseModel DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The list of the database management supported feature statuses:
	FeatureStatus DatabaseFeatureConfigurationFeatureStatusEnum `mandatory:"true" json:"featureStatus"`
}

// GetFeatureStatus returns FeatureStatus
func (m DatabaseDiagnosticsAndPerformanceFeatureConfiguration) GetFeatureStatus() DatabaseFeatureConfigurationFeatureStatusEnum {
	return m.FeatureStatus
}

// GetConnectorDetails returns ConnectorDetails
func (m DatabaseDiagnosticsAndPerformanceFeatureConfiguration) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

// GetDatabaseConnectionDetails returns DatabaseConnectionDetails
func (m DatabaseDiagnosticsAndPerformanceFeatureConfiguration) GetDatabaseConnectionDetails() *DatabaseConnectionDetails {
	return m.DatabaseConnectionDetails
}

func (m DatabaseDiagnosticsAndPerformanceFeatureConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseDiagnosticsAndPerformanceFeatureConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseFeatureConfigurationFeatureStatusEnum(string(m.FeatureStatus)); !ok && m.FeatureStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FeatureStatus: %s. Supported values are: %s.", m.FeatureStatus, strings.Join(GetDatabaseFeatureConfigurationFeatureStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseDiagnosticsAndPerformanceFeatureConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseDiagnosticsAndPerformanceFeatureConfiguration DatabaseDiagnosticsAndPerformanceFeatureConfiguration
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeDatabaseDiagnosticsAndPerformanceFeatureConfiguration
	}{
		"DIAGNOSTICS_AND_PERFORMANCE",
		(MarshalTypeDatabaseDiagnosticsAndPerformanceFeatureConfiguration)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseDiagnosticsAndPerformanceFeatureConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectorDetails          connectordetails                                                      `json:"connectorDetails"`
		DatabaseConnectionDetails *DatabaseConnectionDetails                                            `json:"databaseConnectionDetails"`
		LicenseModel              DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum `json:"licenseModel"`
		FeatureStatus             DatabaseFeatureConfigurationFeatureStatusEnum                         `json:"featureStatus"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ConnectorDetails.UnmarshalPolymorphicJSON(model.ConnectorDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectorDetails = nn.(ConnectorDetails)
	} else {
		m.ConnectorDetails = nil
	}

	m.DatabaseConnectionDetails = model.DatabaseConnectionDetails

	m.LicenseModel = model.LicenseModel

	m.FeatureStatus = model.FeatureStatus

	return
}

// DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum Enum with underlying type: string
type DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum string

// Set of constants representing the allowable values for DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum
const (
	DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelLicenseIncluded     DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum = "LICENSE_INCLUDED"
	DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelBringYourOwnLicense DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum = map[string]DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum{
	"LICENSE_INCLUDED":       DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelBringYourOwnLicense,
}

var mappingDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnumLowerCase = map[string]DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum{
	"license_included":       DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelLicenseIncluded,
	"bring_your_own_license": DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelBringYourOwnLicense,
}

// GetDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnumValues Enumerates the set of values for DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum
func GetDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnumValues() []DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum {
	values := make([]DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum, 0)
	for _, v := range mappingDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnumStringValues Enumerates the set of values in String for DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum
func GetDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum(val string) (DatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnum, bool) {
	enum, ok := mappingDatabaseDiagnosticsAndPerformanceFeatureConfigurationLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
