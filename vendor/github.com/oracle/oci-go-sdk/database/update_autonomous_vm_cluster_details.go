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

// UpdateAutonomousVmClusterDetails Details for updating the Autonomous VM cluster.
type UpdateAutonomousVmClusterDetails struct {

	// The Oracle license model that applies to the Autonomous VM cluster. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel UpdateAutonomousVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateAutonomousVmClusterDetails) String() string {
	return common.PointerString(m)
}

// UpdateAutonomousVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type UpdateAutonomousVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousVmClusterDetailsLicenseModelEnum
const (
	UpdateAutonomousVmClusterDetailsLicenseModelLicenseIncluded     UpdateAutonomousVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense UpdateAutonomousVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateAutonomousVmClusterDetailsLicenseModel = map[string]UpdateAutonomousVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateAutonomousVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for UpdateAutonomousVmClusterDetailsLicenseModelEnum
func GetUpdateAutonomousVmClusterDetailsLicenseModelEnumValues() []UpdateAutonomousVmClusterDetailsLicenseModelEnum {
	values := make([]UpdateAutonomousVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateAutonomousVmClusterDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
