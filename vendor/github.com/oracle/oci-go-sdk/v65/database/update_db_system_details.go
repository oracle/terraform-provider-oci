// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDbSystemDetails Describes the parameters for updating the DB system.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateDbSystemDetails struct {

	// The new number of CPU cores to set for the DB system. Not applicable for INTEL based virtual machine DB systems.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	Version *PatchDetails `mandatory:"false" json:"version"`

	// The public key portion of the key pair to use for SSH access to the DB system. Multiple public keys can be provided. The length of the combined keys cannot exceed 40,000 characters.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	// The size, in gigabytes, to scale the attached storage up to for this virtual machine DB system. This value must be greater than current storage size. Note that the resulting total storage size attached will be greater than the amount requested to allow for REDO/RECO space and software volume. Applies only to virtual machine DB systems.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The size, in gigabytes, to scale the attached RECO storage up to for this virtual machine DB system. This value must be greater than current storage size. Note that the resulting total storage size attached will be greater than the amount requested to allow for the software volume. Applies only to virtual machine DB systems.
	RecoStorageSizeInGBs *int `mandatory:"false" json:"recoStorageSizeInGBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The shape of the DB system. The shape determines resources allocated to the DB system.
	// - For virtual machine shapes, the number of CPU cores and memory
	// To get a list of shapes, use the ListDbSystemShapes operation.
	Shape *string `mandatory:"false" json:"shape"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous AI Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The Oracle Database license model that applies to all databases on the DB system. The default is LICENSE_INCLUDED.
	LicenseModel UpdateDbSystemDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	// The compute model for Base Database Service. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. The ECPU compute model is the recommended model, and the OCPU compute model is legacy.
	ComputeModel UpdateDbSystemDetailsComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// The number of compute servers for the DB system.
	ComputeCount *int `mandatory:"false" json:"computeCount"`
}

func (m UpdateDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDbSystemDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateDbSystemDetailsLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateDbSystemDetailsComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetUpdateDbSystemDetailsComputeModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDbSystemDetailsLicenseModelEnum Enum with underlying type: string
type UpdateDbSystemDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateDbSystemDetailsLicenseModelEnum
const (
	UpdateDbSystemDetailsLicenseModelLicenseIncluded     UpdateDbSystemDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateDbSystemDetailsLicenseModelBringYourOwnLicense UpdateDbSystemDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateDbSystemDetailsLicenseModelEnum = map[string]UpdateDbSystemDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateDbSystemDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateDbSystemDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateDbSystemDetailsLicenseModelEnumLowerCase = map[string]UpdateDbSystemDetailsLicenseModelEnum{
	"license_included":       UpdateDbSystemDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateDbSystemDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateDbSystemDetailsLicenseModelEnumValues Enumerates the set of values for UpdateDbSystemDetailsLicenseModelEnum
func GetUpdateDbSystemDetailsLicenseModelEnumValues() []UpdateDbSystemDetailsLicenseModelEnum {
	values := make([]UpdateDbSystemDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateDbSystemDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDbSystemDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateDbSystemDetailsLicenseModelEnum
func GetUpdateDbSystemDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateDbSystemDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDbSystemDetailsLicenseModelEnum(val string) (UpdateDbSystemDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateDbSystemDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDbSystemDetailsComputeModelEnum Enum with underlying type: string
type UpdateDbSystemDetailsComputeModelEnum string

// Set of constants representing the allowable values for UpdateDbSystemDetailsComputeModelEnum
const (
	UpdateDbSystemDetailsComputeModelEcpu UpdateDbSystemDetailsComputeModelEnum = "ECPU"
	UpdateDbSystemDetailsComputeModelOcpu UpdateDbSystemDetailsComputeModelEnum = "OCPU"
)

var mappingUpdateDbSystemDetailsComputeModelEnum = map[string]UpdateDbSystemDetailsComputeModelEnum{
	"ECPU": UpdateDbSystemDetailsComputeModelEcpu,
	"OCPU": UpdateDbSystemDetailsComputeModelOcpu,
}

var mappingUpdateDbSystemDetailsComputeModelEnumLowerCase = map[string]UpdateDbSystemDetailsComputeModelEnum{
	"ecpu": UpdateDbSystemDetailsComputeModelEcpu,
	"ocpu": UpdateDbSystemDetailsComputeModelOcpu,
}

// GetUpdateDbSystemDetailsComputeModelEnumValues Enumerates the set of values for UpdateDbSystemDetailsComputeModelEnum
func GetUpdateDbSystemDetailsComputeModelEnumValues() []UpdateDbSystemDetailsComputeModelEnum {
	values := make([]UpdateDbSystemDetailsComputeModelEnum, 0)
	for _, v := range mappingUpdateDbSystemDetailsComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDbSystemDetailsComputeModelEnumStringValues Enumerates the set of values in String for UpdateDbSystemDetailsComputeModelEnum
func GetUpdateDbSystemDetailsComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingUpdateDbSystemDetailsComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDbSystemDetailsComputeModelEnum(val string) (UpdateDbSystemDetailsComputeModelEnum, bool) {
	enum, ok := mappingUpdateDbSystemDetailsComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
