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

// CreateVmClusterDetails Details for the create Exadata VM cluster operation. Applies to Exadata Cloud@Customer instances only.
// For details on the create cloud Exadata VM cluster operation used with Exadata Cloud Service instances, see CreateCloudVmClusterDetails
type CreateVmClusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"true" json:"exadataInfrastructureId"`

	// The number of CPU cores to enable for the VM cluster.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	VmClusterNetworkId *string `mandatory:"true" json:"vmClusterNetworkId"`

	// The Oracle Grid Infrastructure software version for the VM cluster.
	GiVersion *string `mandatory:"true" json:"giVersion"`

	// The number of OCPU cores to enable for the VM cluster. Only one decimal place is allowed for the fractional part.
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
	LicenseModel CreateVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// If true, the sparse disk group is configured for the VM cluster. If false, the sparse disk group is not created.
	IsSparseDiskgroupEnabled *bool `mandatory:"false" json:"isSparseDiskgroupEnabled"`

	// If true, database backup on local Exadata storage is configured for the VM cluster. If false, database backup on local Exadata storage is not available in the VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// The time zone to use for the VM cluster. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The list of Db server.
	DbServers []string `mandatory:"false" json:"dbServers"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`
}

func (m CreateVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type CreateVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateVmClusterDetailsLicenseModelEnum
const (
	CreateVmClusterDetailsLicenseModelLicenseIncluded     CreateVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateVmClusterDetailsLicenseModelBringYourOwnLicense CreateVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateVmClusterDetailsLicenseModelEnum = map[string]CreateVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateVmClusterDetailsLicenseModelEnumLowerCase = map[string]CreateVmClusterDetailsLicenseModelEnum{
	"license_included":       CreateVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for CreateVmClusterDetailsLicenseModelEnum
func GetCreateVmClusterDetailsLicenseModelEnumValues() []CreateVmClusterDetailsLicenseModelEnum {
	values := make([]CreateVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateVmClusterDetailsLicenseModelEnum
func GetCreateVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateVmClusterDetailsLicenseModelEnum(val string) (CreateVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
