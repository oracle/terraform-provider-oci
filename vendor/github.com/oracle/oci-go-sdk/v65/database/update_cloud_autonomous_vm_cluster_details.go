// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCloudAutonomousVmClusterDetails Details for updating the cloud Autonomous VM cluster.
type UpdateCloudAutonomousVmClusterDetails struct {

	// User defined description of the cloud Autonomous VM cluster.
	Description *string `mandatory:"false" json:"description"`

	// The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCloudAutonomousVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCloudAutonomousVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCloudAutonomousVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateCloudAutonomousVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum
const (
	UpdateCloudAutonomousVmClusterDetailsLicenseModelLicenseIncluded     UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateCloudAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateCloudAutonomousVmClusterDetailsLicenseModelEnum = map[string]UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateCloudAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateCloudAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateCloudAutonomousVmClusterDetailsLicenseModelEnumLowerCase = map[string]UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum{
	"license_included":       UpdateCloudAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateCloudAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateCloudAutonomousVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum
func GetUpdateCloudAutonomousVmClusterDetailsLicenseModelEnumValues() []UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum {
	values := make([]UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateCloudAutonomousVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCloudAutonomousVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum
func GetUpdateCloudAutonomousVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateCloudAutonomousVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCloudAutonomousVmClusterDetailsLicenseModelEnum(val string) (UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateCloudAutonomousVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
