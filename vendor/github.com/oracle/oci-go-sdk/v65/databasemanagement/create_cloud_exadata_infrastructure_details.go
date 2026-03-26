// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateCloudExadataInfrastructureDetails The details required to create the Exadata infrastructure.
type CreateCloudExadataInfrastructureDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of VM Clusters in the Exadata infrastructure.
	VmClusterIds []string `mandatory:"true" json:"vmClusterIds"`

	// The unique key of the discovery request.
	DiscoveryKey *string `mandatory:"false" json:"discoveryKey"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel CreateCloudExadataInfrastructureDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The name of the Exadata infrastructure.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The list of all the Exadata storage server names to be included for monitoring purposes. If not specified, all the Exadata storage servers associated with the VM Clusters are included.
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

func (m CreateCloudExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCloudExadataInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateCloudExadataInfrastructureDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateCloudExadataInfrastructureDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateCloudExadataInfrastructureDetailsLicenseModelEnum Enum with underlying type: string
type CreateCloudExadataInfrastructureDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateCloudExadataInfrastructureDetailsLicenseModelEnum
const (
	CreateCloudExadataInfrastructureDetailsLicenseModelLicenseIncluded     CreateCloudExadataInfrastructureDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateCloudExadataInfrastructureDetailsLicenseModelBringYourOwnLicense CreateCloudExadataInfrastructureDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateCloudExadataInfrastructureDetailsLicenseModelEnum = map[string]CreateCloudExadataInfrastructureDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateCloudExadataInfrastructureDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateCloudExadataInfrastructureDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateCloudExadataInfrastructureDetailsLicenseModelEnumLowerCase = map[string]CreateCloudExadataInfrastructureDetailsLicenseModelEnum{
	"license_included":       CreateCloudExadataInfrastructureDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateCloudExadataInfrastructureDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateCloudExadataInfrastructureDetailsLicenseModelEnumValues Enumerates the set of values for CreateCloudExadataInfrastructureDetailsLicenseModelEnum
func GetCreateCloudExadataInfrastructureDetailsLicenseModelEnumValues() []CreateCloudExadataInfrastructureDetailsLicenseModelEnum {
	values := make([]CreateCloudExadataInfrastructureDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateCloudExadataInfrastructureDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCloudExadataInfrastructureDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateCloudExadataInfrastructureDetailsLicenseModelEnum
func GetCreateCloudExadataInfrastructureDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateCloudExadataInfrastructureDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCloudExadataInfrastructureDetailsLicenseModelEnum(val string) (CreateCloudExadataInfrastructureDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateCloudExadataInfrastructureDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
