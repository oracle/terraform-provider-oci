// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateCloudVmClusterDetails Details for updating the cloud VM cluster. Applies to Exadata Cloud Service instances only.
type UpdateCloudVmClusterDetails struct {

	// The user-friendly name for the cloud VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The number of CPU cores to enable for the cloud VM cluster.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of OCPU cores to enable for a cloud VM cluster. Only 1 decimal place is allowed for the fractional part.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The Oracle license model that applies to the cloud VM cluster. The default is BRING_YOUR_OWN_LICENSE. Applies to Exadata Cloud Service instances only.
	LicenseModel UpdateCloudVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The public key portion of one or more key pairs used for SSH access to the cloud VM cluster.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	UpdateDetails *UpdateDetails `mandatory:"false" json:"updateDetails"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The list of compute servers to be added to the cloud VM cluster.
	ComputeNodes []string `mandatory:"false" json:"computeNodes"`

	// The disk group size to be allocated in GBs.
	StorageSizeInGBs *int `mandatory:"false" json:"storageSizeInGBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCloudVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCloudVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCloudVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateCloudVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCloudVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type UpdateCloudVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateCloudVmClusterDetailsLicenseModelEnum
const (
	UpdateCloudVmClusterDetailsLicenseModelLicenseIncluded     UpdateCloudVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateCloudVmClusterDetailsLicenseModelBringYourOwnLicense UpdateCloudVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateCloudVmClusterDetailsLicenseModelEnum = map[string]UpdateCloudVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateCloudVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateCloudVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateCloudVmClusterDetailsLicenseModelEnumLowerCase = map[string]UpdateCloudVmClusterDetailsLicenseModelEnum{
	"license_included":       UpdateCloudVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateCloudVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateCloudVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for UpdateCloudVmClusterDetailsLicenseModelEnum
func GetUpdateCloudVmClusterDetailsLicenseModelEnumValues() []UpdateCloudVmClusterDetailsLicenseModelEnum {
	values := make([]UpdateCloudVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateCloudVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCloudVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateCloudVmClusterDetailsLicenseModelEnum
func GetUpdateCloudVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateCloudVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCloudVmClusterDetailsLicenseModelEnum(val string) (UpdateCloudVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateCloudVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
