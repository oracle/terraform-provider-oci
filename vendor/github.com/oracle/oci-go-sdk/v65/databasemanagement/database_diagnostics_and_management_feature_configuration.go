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

// DatabaseDiagnosticsAndManagementFeatureConfiguration The details required to enable the Diagnostics and Management feature.
type DatabaseDiagnosticsAndManagementFeatureConfiguration struct {
	ConnectorDetails ConnectorDetails `mandatory:"false" json:"connectorDetails"`

	DatabaseConnectionDetails *DatabaseConnectionDetails `mandatory:"false" json:"databaseConnectionDetails"`

	// The Oracle license model that applies to the external database.
	LicenseModel DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The list of statuses for Database Management features.
	FeatureStatus DatabaseFeatureConfigurationFeatureStatusEnum `mandatory:"true" json:"featureStatus"`
}

// GetFeatureStatus returns FeatureStatus
func (m DatabaseDiagnosticsAndManagementFeatureConfiguration) GetFeatureStatus() DatabaseFeatureConfigurationFeatureStatusEnum {
	return m.FeatureStatus
}

// GetConnectorDetails returns ConnectorDetails
func (m DatabaseDiagnosticsAndManagementFeatureConfiguration) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

// GetDatabaseConnectionDetails returns DatabaseConnectionDetails
func (m DatabaseDiagnosticsAndManagementFeatureConfiguration) GetDatabaseConnectionDetails() *DatabaseConnectionDetails {
	return m.DatabaseConnectionDetails
}

func (m DatabaseDiagnosticsAndManagementFeatureConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseDiagnosticsAndManagementFeatureConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnumStringValues(), ",")))
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
func (m DatabaseDiagnosticsAndManagementFeatureConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseDiagnosticsAndManagementFeatureConfiguration DatabaseDiagnosticsAndManagementFeatureConfiguration
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeDatabaseDiagnosticsAndManagementFeatureConfiguration
	}{
		"DIAGNOSTICS_AND_MANAGEMENT",
		(MarshalTypeDatabaseDiagnosticsAndManagementFeatureConfiguration)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseDiagnosticsAndManagementFeatureConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectorDetails          connectordetails                                                     `json:"connectorDetails"`
		DatabaseConnectionDetails *DatabaseConnectionDetails                                           `json:"databaseConnectionDetails"`
		LicenseModel              DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum `json:"licenseModel"`
		FeatureStatus             DatabaseFeatureConfigurationFeatureStatusEnum                        `json:"featureStatus"`
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

// DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum Enum with underlying type: string
type DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum string

// Set of constants representing the allowable values for DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum
const (
	DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelLicenseIncluded     DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum = "LICENSE_INCLUDED"
	DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelBringYourOwnLicense DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum = map[string]DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum{
	"LICENSE_INCLUDED":       DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelBringYourOwnLicense,
}

var mappingDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnumLowerCase = map[string]DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum{
	"license_included":       DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelLicenseIncluded,
	"bring_your_own_license": DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelBringYourOwnLicense,
}

// GetDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnumValues Enumerates the set of values for DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum
func GetDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnumValues() []DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum {
	values := make([]DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum, 0)
	for _, v := range mappingDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnumStringValues Enumerates the set of values in String for DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum
func GetDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum(val string) (DatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnum, bool) {
	enum, ok := mappingDatabaseDiagnosticsAndManagementFeatureConfigurationLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
