// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateVmClusterDetails Details for updating the VM cluster. Applies to Exadata Cloud@Customer instances only.
// For details on updating a cloud VM cluster in an Exadata Cloud Service instance, see UpdateCloudVmClusterDetails
type UpdateVmClusterDetails struct {

	// The number of CPU cores to enable for the VM cluster.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of OCPU cores to enable for the VM cluster. Only 1 decimal place is allowed for the fractional part.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The memory to be allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The local node storage to be allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The data disk group size to be allocated in TBs.
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The data disk group size to be allocated in GBs.
	DataStorageSizeInGBs *float64 `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The Oracle license model that applies to the VM cluster. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel UpdateVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	Version *PatchDetails `mandatory:"false" json:"version"`

	UpdateDetails *VmClusterUpdateDetails `mandatory:"false" json:"updateDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`
}

func (m UpdateVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type UpdateVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateVmClusterDetailsLicenseModelEnum
const (
	UpdateVmClusterDetailsLicenseModelLicenseIncluded     UpdateVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateVmClusterDetailsLicenseModelBringYourOwnLicense UpdateVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateVmClusterDetailsLicenseModelEnum = map[string]UpdateVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateVmClusterDetailsLicenseModelEnumLowerCase = map[string]UpdateVmClusterDetailsLicenseModelEnum{
	"license_included":       UpdateVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for UpdateVmClusterDetailsLicenseModelEnum
func GetUpdateVmClusterDetailsLicenseModelEnumValues() []UpdateVmClusterDetailsLicenseModelEnum {
	values := make([]UpdateVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateVmClusterDetailsLicenseModelEnum
func GetUpdateVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateVmClusterDetailsLicenseModelEnum(val string) (UpdateVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
