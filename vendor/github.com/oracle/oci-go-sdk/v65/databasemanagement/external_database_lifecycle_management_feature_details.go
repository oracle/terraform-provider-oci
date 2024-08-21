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

// ExternalDatabaseLifecycleManagementFeatureDetails The details required to enable the Database Lifecycle Management feature.
type ExternalDatabaseLifecycleManagementFeatureDetails struct {
	ConnectorDetails ConnectorDetails `mandatory:"true" json:"connectorDetails"`

	// The Oracle license model that applies to the external database.
	LicenseModel ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

// GetConnectorDetails returns ConnectorDetails
func (m ExternalDatabaseLifecycleManagementFeatureDetails) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

func (m ExternalDatabaseLifecycleManagementFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDatabaseLifecycleManagementFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalDatabaseLifecycleManagementFeatureDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalDatabaseLifecycleManagementFeatureDetails ExternalDatabaseLifecycleManagementFeatureDetails
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeExternalDatabaseLifecycleManagementFeatureDetails
	}{
		"DB_LIFECYCLE_MANAGEMENT",
		(MarshalTypeExternalDatabaseLifecycleManagementFeatureDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExternalDatabaseLifecycleManagementFeatureDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectorDetails connectordetails                                                  `json:"connectorDetails"`
		LicenseModel     ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum `json:"licenseModel"`
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

// ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum Enum with underlying type: string
type ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum string

// Set of constants representing the allowable values for ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum
const (
	ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelLicenseIncluded     ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelBringYourOwnLicense ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum = map[string]ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelBringYourOwnLicense,
}

var mappingExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnumLowerCase = map[string]ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum{
	"license_included":       ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelBringYourOwnLicense,
}

// GetExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnumValues Enumerates the set of values for ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum
func GetExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnumValues() []ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum {
	values := make([]ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum, 0)
	for _, v := range mappingExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnumStringValues Enumerates the set of values in String for ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum
func GetExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum(val string) (ExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnum, bool) {
	enum, ok := mappingExternalDatabaseLifecycleManagementFeatureDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
