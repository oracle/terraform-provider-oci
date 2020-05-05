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

// UpdateDbSystemDetails Describes the parameters for updating the DB system.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateDbSystemDetails struct {

	// The new number of CPU cores to set for the DB system. Not applicable for virtual machine DB systems.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	Version *PatchDetails `mandatory:"false" json:"version"`

	// The public key portion of the key pair to use for SSH access to the DB system. Multiple public keys can be provided. The length of the combined keys cannot exceed 40,000 characters.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	// The size, in gigabytes, to scale the attached storage up to for this virtual machine DB system. This value must be greater than current storage size. Note that the resulting total storage size attached will be greater than the amount requested to allow for REDO/RECO space and software volume. Applies only to virtual machine DB systems.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The shape of the DB system. The shape determines resources allocated to the DB system.
	// - For virtual machine shapes, the number of CPU cores and memory
	// To get a list of shapes, use the ListDbSystemShapes operation.
	Shape *string `mandatory:"false" json:"shape"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata DB systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The Oracle Database license model that applies to all databases on the DB system. The default is LICENSE_INCLUDED.
	LicenseModel UpdateDbSystemDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`
}

func (m UpdateDbSystemDetails) String() string {
	return common.PointerString(m)
}

// UpdateDbSystemDetailsLicenseModelEnum Enum with underlying type: string
type UpdateDbSystemDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateDbSystemDetailsLicenseModelEnum
const (
	UpdateDbSystemDetailsLicenseModelLicenseIncluded     UpdateDbSystemDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateDbSystemDetailsLicenseModelBringYourOwnLicense UpdateDbSystemDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateDbSystemDetailsLicenseModel = map[string]UpdateDbSystemDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateDbSystemDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateDbSystemDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateDbSystemDetailsLicenseModelEnumValues Enumerates the set of values for UpdateDbSystemDetailsLicenseModelEnum
func GetUpdateDbSystemDetailsLicenseModelEnumValues() []UpdateDbSystemDetailsLicenseModelEnum {
	values := make([]UpdateDbSystemDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateDbSystemDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
