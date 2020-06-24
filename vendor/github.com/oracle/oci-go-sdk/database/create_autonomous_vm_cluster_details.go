// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateAutonomousVmClusterDetails Details for the create Autonomous VM cluster operation.
type CreateAutonomousVmClusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"true" json:"exadataInfrastructureId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	VmClusterNetworkId *string `mandatory:"true" json:"vmClusterNetworkId"`

	// The time zone to use for the Autonomous VM cluster. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// The Oracle license model that applies to the Autonomous VM cluster. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel CreateAutonomousVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAutonomousVmClusterDetails) String() string {
	return common.PointerString(m)
}

// CreateAutonomousVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type CreateAutonomousVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateAutonomousVmClusterDetailsLicenseModelEnum
const (
	CreateAutonomousVmClusterDetailsLicenseModelLicenseIncluded     CreateAutonomousVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense CreateAutonomousVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateAutonomousVmClusterDetailsLicenseModel = map[string]CreateAutonomousVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateAutonomousVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for CreateAutonomousVmClusterDetailsLicenseModelEnum
func GetCreateAutonomousVmClusterDetailsLicenseModelEnumValues() []CreateAutonomousVmClusterDetailsLicenseModelEnum {
	values := make([]CreateAutonomousVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateAutonomousVmClusterDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
