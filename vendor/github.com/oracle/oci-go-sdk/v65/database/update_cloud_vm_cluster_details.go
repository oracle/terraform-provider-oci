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

// UpdateCloudVmClusterDetails Details for updating the cloud VM cluster. Applies to Exadata Cloud Service instances only.
type UpdateCloudVmClusterDetails struct {

	// The user-friendly name for the cloud VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// For fixed shapes, this is the total number of OCPUs to enable across the VM cluster.
	//  - Exadata.Base.48 - Specify a multiple of 2, from 0 to 48.
	//  - Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84.
	//  - Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168.
	//  - Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.
	//  - Exadata.Quarter2.92 - Specify a multiple of 2, from 0 to 92.
	//  - Exadata.Half2.184 - Specify a multiple of 4, from 0 to 184.
	//  - Exadata.Full2.368 - Specify a multiple of 8, from 0 to 368.
	//  - Exadata.Quarter3.100 - Specify a multiple of 2, from 0 to 100.
	//  - Exadata.Half3.200 - Specify a multiple of 4, from 0 to 200.
	//  - Exadata.Full3.400 - Specify a multiple of 8, from 0 to 400.
	// The API specification for fixed shape values is https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/DbSystemShapeSummary
	//
	// For flexible shapes X8M and X9M, this is the total number of OCPUs to enable across the VM cluster. The number available for the VM cluster will be based on the number of database servers selected for provisioning the VM cluster on the Exadata Infrastructure.
	//  - Exadata.X8M - Specify a multiple of 2, from 2 to 50 per X8M database server.
	//  - Exadata.X9M - Specify a multiple of 2, from 2 to 126 per X9M database server.
	// For flexible shapes X11M and higher, this is the total number of ECPUs to enable across the VM cluster. The number available for the VM cluster will be based on the number of database servers selected for provisioning the VM cluster on the Exadata Infrastructure.
	//  - Exadata.X11M - Specify a multiple of 8, from 8 to 760 per X11M database server.
	// The API specification for flexible shape values is https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/datatypes/FlexComponentSummary
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of OCPU cores to enable for a cloud VM cluster. Only 1 decimal place is allowed for the fractional part.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The memory to be allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The local node storage to be allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The data disk group size to be allocated in TBs.
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The Oracle license model that applies to the cloud VM cluster. The default is BRING_YOUR_OWN_LICENSE. Applies to Exadata Cloud Service instances only.
	LicenseModel UpdateCloudVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The public key portion of one or more key pairs used for SSH access to the cloud VM cluster.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	UpdateDetails *UpdateDetails `mandatory:"false" json:"updateDetails"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The list of compute servers to be added to the cloud VM cluster.
	ComputeNodes []string `mandatory:"false" json:"computeNodes"`

	// The disk group size to be allocated in GBs.
	StorageSizeInGBs *int `mandatory:"false" json:"storageSizeInGBs"`

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

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	// Details of the file system configuration of the VM cluster.
	FileSystemConfigurationDetails []FileSystemConfigurationDetail `mandatory:"false" json:"fileSystemConfigurationDetails"`

	CloudAutomationUpdateDetails *CloudAutomationUpdateDetails `mandatory:"false" json:"cloudAutomationUpdateDetails"`
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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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
