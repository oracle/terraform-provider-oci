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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateExternalExadataInfrastructureDetails The details required to update the external Exadata infrastructure.
type UpdateExternalExadataInfrastructureDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique key of the discovery request.
	DiscoveryKey *string `mandatory:"false" json:"discoveryKey"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel UpdateExternalExadataInfrastructureDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The name of the Exadata infrastructure.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The list of all the DB systems OCIDs.
	DbSystemIds []string `mandatory:"false" json:"dbSystemIds"`

	// The list of the names of Exadata storage servers to be monitored. If not specified, it includes all Exadata storage servers associated with the monitored DB systems.
	StorageServerNames []string `mandatory:"false" json:"storageServerNames"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateExternalExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateExternalExadataInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateExternalExadataInfrastructureDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateExternalExadataInfrastructureDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateExternalExadataInfrastructureDetailsLicenseModelEnum Enum with underlying type: string
type UpdateExternalExadataInfrastructureDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateExternalExadataInfrastructureDetailsLicenseModelEnum
const (
	UpdateExternalExadataInfrastructureDetailsLicenseModelLicenseIncluded     UpdateExternalExadataInfrastructureDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateExternalExadataInfrastructureDetailsLicenseModelBringYourOwnLicense UpdateExternalExadataInfrastructureDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateExternalExadataInfrastructureDetailsLicenseModelEnum = map[string]UpdateExternalExadataInfrastructureDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateExternalExadataInfrastructureDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateExternalExadataInfrastructureDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateExternalExadataInfrastructureDetailsLicenseModelEnumLowerCase = map[string]UpdateExternalExadataInfrastructureDetailsLicenseModelEnum{
	"license_included":       UpdateExternalExadataInfrastructureDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateExternalExadataInfrastructureDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateExternalExadataInfrastructureDetailsLicenseModelEnumValues Enumerates the set of values for UpdateExternalExadataInfrastructureDetailsLicenseModelEnum
func GetUpdateExternalExadataInfrastructureDetailsLicenseModelEnumValues() []UpdateExternalExadataInfrastructureDetailsLicenseModelEnum {
	values := make([]UpdateExternalExadataInfrastructureDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateExternalExadataInfrastructureDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateExternalExadataInfrastructureDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateExternalExadataInfrastructureDetailsLicenseModelEnum
func GetUpdateExternalExadataInfrastructureDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateExternalExadataInfrastructureDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateExternalExadataInfrastructureDetailsLicenseModelEnum(val string) (UpdateExternalExadataInfrastructureDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateExternalExadataInfrastructureDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
