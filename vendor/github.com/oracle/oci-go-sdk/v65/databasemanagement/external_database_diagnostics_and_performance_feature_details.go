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

// ExternalDatabaseDiagnosticsAndPerformanceFeatureDetails The details required to enable diagnostics and performance feature.
type ExternalDatabaseDiagnosticsAndPerformanceFeatureDetails struct {
	ConnectorDetails ConnectorDetails `mandatory:"true" json:"connectorDetails"`

	// The Oracle license model that applies to the external database.
	LicenseModel ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

// GetConnectorDetails returns ConnectorDetails
func (m ExternalDatabaseDiagnosticsAndPerformanceFeatureDetails) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

func (m ExternalDatabaseDiagnosticsAndPerformanceFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDatabaseDiagnosticsAndPerformanceFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalDatabaseDiagnosticsAndPerformanceFeatureDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalDatabaseDiagnosticsAndPerformanceFeatureDetails ExternalDatabaseDiagnosticsAndPerformanceFeatureDetails
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeExternalDatabaseDiagnosticsAndPerformanceFeatureDetails
	}{
		"DIAGNOSTICS_AND_PERFORMANCE",
		(MarshalTypeExternalDatabaseDiagnosticsAndPerformanceFeatureDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExternalDatabaseDiagnosticsAndPerformanceFeatureDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectorDetails connectordetails                                                        `json:"connectorDetails"`
		LicenseModel     ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum `json:"licenseModel"`
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

// ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum Enum with underlying type: string
type ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum string

// Set of constants representing the allowable values for ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum
const (
	ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelLicenseIncluded     ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelBringYourOwnLicense ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum = map[string]ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelBringYourOwnLicense,
}

var mappingExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnumLowerCase = map[string]ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum{
	"license_included":       ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelBringYourOwnLicense,
}

// GetExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnumValues Enumerates the set of values for ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum
func GetExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnumValues() []ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum {
	values := make([]ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum, 0)
	for _, v := range mappingExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnumStringValues Enumerates the set of values in String for ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum
func GetExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum(val string) (ExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnum, bool) {
	enum, ok := mappingExternalDatabaseDiagnosticsAndPerformanceFeatureDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
