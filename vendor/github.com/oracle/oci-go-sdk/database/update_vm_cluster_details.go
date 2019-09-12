// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateVmClusterDetails Details for updating the VM cluster.
type UpdateVmClusterDetails struct {

	// The number of CPU cores to enable for the VM cluster.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The Oracle license model that applies to the VM cluster. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel UpdateVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateVmClusterDetails) String() string {
	return common.PointerString(m)
}

// UpdateVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type UpdateVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateVmClusterDetailsLicenseModelEnum
const (
	UpdateVmClusterDetailsLicenseModelLicenseIncluded     UpdateVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateVmClusterDetailsLicenseModelBringYourOwnLicense UpdateVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateVmClusterDetailsLicenseModel = map[string]UpdateVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for UpdateVmClusterDetailsLicenseModelEnum
func GetUpdateVmClusterDetailsLicenseModelEnumValues() []UpdateVmClusterDetailsLicenseModelEnum {
	values := make([]UpdateVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateVmClusterDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
