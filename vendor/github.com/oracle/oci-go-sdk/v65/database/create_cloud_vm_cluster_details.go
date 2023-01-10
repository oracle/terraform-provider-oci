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

// CreateCloudVmClusterDetails Details for the create cloud VM cluster operation. Applies to Exadata Cloud Service instances only.
type CreateCloudVmClusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet associated with the cloud VM cluster.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the cloud VM cluster.
	BackupSubnetId *string `mandatory:"true" json:"backupSubnetId"`

	// The number of CPU cores to enable for a cloud VM cluster. Valid values depend on the specified shape:
	// - Exadata.Base.48 - Specify a multiple of 2, from 0 to 48.
	// - Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84.
	// - Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168.
	// - Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.
	// - Exadata.Quarter2.92 - Specify a multiple of 2, from 0 to 92.
	// - Exadata.Half2.184 - Specify a multiple of 4, from 0 to 184.
	// - Exadata.Full2.368 - Specify a multiple of 8, from 0 to 368.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The user-friendly name for the cloud VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure resource.
	CloudExadataInfrastructureId *string `mandatory:"true" json:"cloudExadataInfrastructureId"`

	// The hostname for the cloud VM cluster. The hostname must begin with an alphabetic character, and
	// can contain alphanumeric characters and hyphens (-). The maximum length of the hostname is 16 characters for bare metal and virtual machine DB systems, and 12 characters for Exadata systems.
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// the cloud VM Cluster will fail to provision.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The public key portion of one or more key pairs used for SSH access to the cloud VM cluster.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	GiVersion *string `mandatory:"true" json:"giVersion"`

	// The number of OCPU cores to enable for a cloud VM cluster. Only 1 decimal place is allowed for the fractional part.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The memory to be allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The local node storage to be allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The data disk group size to be allocated in TBs.
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The list of Db servers.
	DbServers []string `mandatory:"false" json:"dbServers"`

	// The cluster name for cloud VM cluster. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 35, 40, 60 and 80. The default is 80 percent assigned to DATA storage. See Storage Configuration (https://docs.cloud.oracle.com/Content/Database/Concepts/exaoverview.htm#Exadata) in the Exadata documentation for details on the impact of the configuration settings on storage.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage"`

	// A domain name used for the cloud VM cluster. If the Oracle-provided internet and VCN
	// resolver is enabled for the specified subnet, the domain name for the subnet is used
	// (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.
	// Applies to Exadata Cloud Service instances only.
	Domain *string `mandatory:"false" json:"domain"`

	// The Oracle license model that applies to the cloud VM cluster. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel CreateCloudVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// If true, the sparse disk group is configured for the cloud VM cluster. If false, the sparse disk group is not created.
	IsSparseDiskgroupEnabled *bool `mandatory:"false" json:"isSparseDiskgroupEnabled"`

	// If true, database backup on local Exadata storage is configured for the cloud VM cluster. If false, database backup on local Exadata storage is not available in the cloud VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// The time zone to use for the cloud VM cluster. For details, see Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The TCP Single Client Access Name (SCAN) port. The default port is 1521.
	ScanListenerPortTcp *int `mandatory:"false" json:"scanListenerPortTcp"`

	// The TCPS Single Client Access Name (SCAN) port. The default port is 2484.
	ScanListenerPortTcpSsl *int `mandatory:"false" json:"scanListenerPortTcpSsl"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`
}

func (m CreateCloudVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCloudVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateCloudVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateCloudVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateCloudVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type CreateCloudVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateCloudVmClusterDetailsLicenseModelEnum
const (
	CreateCloudVmClusterDetailsLicenseModelLicenseIncluded     CreateCloudVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateCloudVmClusterDetailsLicenseModelBringYourOwnLicense CreateCloudVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateCloudVmClusterDetailsLicenseModelEnum = map[string]CreateCloudVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateCloudVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateCloudVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateCloudVmClusterDetailsLicenseModelEnumLowerCase = map[string]CreateCloudVmClusterDetailsLicenseModelEnum{
	"license_included":       CreateCloudVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateCloudVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateCloudVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for CreateCloudVmClusterDetailsLicenseModelEnum
func GetCreateCloudVmClusterDetailsLicenseModelEnumValues() []CreateCloudVmClusterDetailsLicenseModelEnum {
	values := make([]CreateCloudVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateCloudVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCloudVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateCloudVmClusterDetailsLicenseModelEnum
func GetCreateCloudVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateCloudVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCloudVmClusterDetailsLicenseModelEnum(val string) (CreateCloudVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateCloudVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
