// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExternalExadataInfrastructureDetails The details required to create the external Exadata infrastructure.
type CreateExternalExadataInfrastructureDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the Exadata infrastructure.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The list of DB systems in the Exadata infrastructure.
	DbSystemIds []string `mandatory:"true" json:"dbSystemIds"`

	// The unique key of the discovery request.
	DiscoveryKey *string `mandatory:"false" json:"discoveryKey"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel CreateExternalExadataInfrastructureDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The list of all the Exadata storage server names to be included for monitoring purposes. If not specified, all the Exadata storage servers associated with the DB systems are included.
	StorageServerNames []string `mandatory:"false" json:"storageServerNames"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateExternalExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExternalExadataInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateExternalExadataInfrastructureDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateExternalExadataInfrastructureDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateExternalExadataInfrastructureDetailsLicenseModelEnum Enum with underlying type: string
type CreateExternalExadataInfrastructureDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateExternalExadataInfrastructureDetailsLicenseModelEnum
const (
	CreateExternalExadataInfrastructureDetailsLicenseModelLicenseIncluded     CreateExternalExadataInfrastructureDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateExternalExadataInfrastructureDetailsLicenseModelBringYourOwnLicense CreateExternalExadataInfrastructureDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateExternalExadataInfrastructureDetailsLicenseModelEnum = map[string]CreateExternalExadataInfrastructureDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateExternalExadataInfrastructureDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateExternalExadataInfrastructureDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateExternalExadataInfrastructureDetailsLicenseModelEnumLowerCase = map[string]CreateExternalExadataInfrastructureDetailsLicenseModelEnum{
	"license_included":       CreateExternalExadataInfrastructureDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateExternalExadataInfrastructureDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateExternalExadataInfrastructureDetailsLicenseModelEnumValues Enumerates the set of values for CreateExternalExadataInfrastructureDetailsLicenseModelEnum
func GetCreateExternalExadataInfrastructureDetailsLicenseModelEnumValues() []CreateExternalExadataInfrastructureDetailsLicenseModelEnum {
	values := make([]CreateExternalExadataInfrastructureDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateExternalExadataInfrastructureDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateExternalExadataInfrastructureDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateExternalExadataInfrastructureDetailsLicenseModelEnum
func GetCreateExternalExadataInfrastructureDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateExternalExadataInfrastructureDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateExternalExadataInfrastructureDetailsLicenseModelEnum(val string) (CreateExternalExadataInfrastructureDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateExternalExadataInfrastructureDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
