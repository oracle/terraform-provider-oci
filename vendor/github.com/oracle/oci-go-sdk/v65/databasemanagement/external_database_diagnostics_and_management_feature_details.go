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

// ExternalDatabaseDiagnosticsAndManagementFeatureDetails The details required to enable the Diagnostics and Management feature.
type ExternalDatabaseDiagnosticsAndManagementFeatureDetails struct {
	ConnectorDetails ConnectorDetails `mandatory:"true" json:"connectorDetails"`

	// The Oracle license model that applies to the external database.
	LicenseModel ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

// GetConnectorDetails returns ConnectorDetails
func (m ExternalDatabaseDiagnosticsAndManagementFeatureDetails) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

func (m ExternalDatabaseDiagnosticsAndManagementFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDatabaseDiagnosticsAndManagementFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalDatabaseDiagnosticsAndManagementFeatureDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalDatabaseDiagnosticsAndManagementFeatureDetails ExternalDatabaseDiagnosticsAndManagementFeatureDetails
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeExternalDatabaseDiagnosticsAndManagementFeatureDetails
	}{
		"DIAGNOSTICS_AND_MANAGEMENT",
		(MarshalTypeExternalDatabaseDiagnosticsAndManagementFeatureDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExternalDatabaseDiagnosticsAndManagementFeatureDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectorDetails connectordetails                                                       `json:"connectorDetails"`
		LicenseModel     ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum `json:"licenseModel"`
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

	m.LicenseModel = model.LicenseModel

	return
}

// ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum Enum with underlying type: string
type ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum string

// Set of constants representing the allowable values for ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum
const (
	ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelLicenseIncluded     ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelBringYourOwnLicense ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum = map[string]ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelBringYourOwnLicense,
}

var mappingExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnumLowerCase = map[string]ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum{
	"license_included":       ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelBringYourOwnLicense,
}

// GetExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnumValues Enumerates the set of values for ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum
func GetExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnumValues() []ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum {
	values := make([]ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum, 0)
	for _, v := range mappingExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnumStringValues Enumerates the set of values in String for ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum
func GetExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum(val string) (ExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnum, bool) {
	enum, ok := mappingExternalDatabaseDiagnosticsAndManagementFeatureDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
