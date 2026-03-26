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

// UpdateCloudExadataInfrastructureDetails The details required to update the Exadata infrastructure.
type UpdateCloudExadataInfrastructureDetails struct {

	// The unique key of the discovery request.
	DiscoveryKey *string `mandatory:"false" json:"discoveryKey"`

	// The Oracle license model that applies to the database management resources.
	LicenseModel UpdateCloudExadataInfrastructureDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The name of the Exadata infrastructure.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The list of all the VM Cluster OCIDs.
	VmClusterIds []string `mandatory:"false" json:"vmClusterIds"`

	// The list of the names of Exadata storage servers to be monitored. If not specified, it includes all Exadata storage servers associated with the monitored VM Clusters.
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

func (m UpdateCloudExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCloudExadataInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCloudExadataInfrastructureDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateCloudExadataInfrastructureDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCloudExadataInfrastructureDetailsLicenseModelEnum Enum with underlying type: string
type UpdateCloudExadataInfrastructureDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateCloudExadataInfrastructureDetailsLicenseModelEnum
const (
	UpdateCloudExadataInfrastructureDetailsLicenseModelLicenseIncluded     UpdateCloudExadataInfrastructureDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateCloudExadataInfrastructureDetailsLicenseModelBringYourOwnLicense UpdateCloudExadataInfrastructureDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateCloudExadataInfrastructureDetailsLicenseModelEnum = map[string]UpdateCloudExadataInfrastructureDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateCloudExadataInfrastructureDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateCloudExadataInfrastructureDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateCloudExadataInfrastructureDetailsLicenseModelEnumLowerCase = map[string]UpdateCloudExadataInfrastructureDetailsLicenseModelEnum{
	"license_included":       UpdateCloudExadataInfrastructureDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateCloudExadataInfrastructureDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateCloudExadataInfrastructureDetailsLicenseModelEnumValues Enumerates the set of values for UpdateCloudExadataInfrastructureDetailsLicenseModelEnum
func GetUpdateCloudExadataInfrastructureDetailsLicenseModelEnumValues() []UpdateCloudExadataInfrastructureDetailsLicenseModelEnum {
	values := make([]UpdateCloudExadataInfrastructureDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateCloudExadataInfrastructureDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCloudExadataInfrastructureDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateCloudExadataInfrastructureDetailsLicenseModelEnum
func GetUpdateCloudExadataInfrastructureDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateCloudExadataInfrastructureDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCloudExadataInfrastructureDetailsLicenseModelEnum(val string) (UpdateCloudExadataInfrastructureDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateCloudExadataInfrastructureDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
