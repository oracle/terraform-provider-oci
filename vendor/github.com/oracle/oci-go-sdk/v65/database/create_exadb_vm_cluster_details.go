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

// CreateExadbVmClusterDetails Details for the create Exadata VM cluster on Exascale Infrastructure operation. Applies to Exadata Cloud Service instances only.
type CreateExadbVmClusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the availability domain that VM cluster on Exascale Infrastructure is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure.
	BackupSubnetId *string `mandatory:"true" json:"backupSubnetId"`

	// The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The hostname for the Exadata VM cluster on Exascale Infrastructure. The hostname must begin with an alphabetic character, and
	// can contain alphanumeric characters and hyphens (-). The maximum length of the hostname is 12 characters for Exadata systems.
	//
	// The maximum length of the combined hostname and domain is 63 characters.
	//
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// the Exadata VM cluster on Exascale Infrastructure will fail to provision.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The number of nodes in the Exadata VM cluster on Exascale Infrastructure.
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// The number of CPU cores reserved for a Exadata VM cluster on Exascale Infrastructure.
	ReservedCpuCoreCount *int `mandatory:"true" json:"reservedCpuCoreCount"`

	// The number of CPU cores to enable for a Exadata VM cluster on Exascale Infrastructure.
	EnabledCpuCoreCount *int `mandatory:"true" json:"enabledCpuCoreCount"`

	VmFileSystemStorage *ExadbVmClusterStorageDetails `mandatory:"true" json:"vmFileSystemStorage"`

	HighCapacityDatabaseStorage *ExadbVmClusterStorageDetails `mandatory:"true" json:"highCapacityDatabaseStorage"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	GiVersion *string `mandatory:"true" json:"giVersion"`

	// The cluster name for Exadata VM cluster on Exascale Infrastructure. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 35, 40, 60 and 80. The default is 80 percent assigned to DATA storage. See Storage Configuration (https://docs.cloud.oracle.com/Content/Database/Concepts/exaoverview.htm#Exadata) in the Exadata documentation for details on the impact of the configuration settings on storage.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage"`

	// A domain name used for the Exadata VM cluster on Exascale Infrastructure. If the Oracle-provided internet and VCN
	// resolver is enabled for the specified subnet, the domain name for the subnet is used
	// (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.
	// Applies to Exadata Cloud Service instances only.
	Domain *string `mandatory:"false" json:"domain"`

	// The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel CreateExadbVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// If true, database backup on local Exadata storage is configured for the Exadata VM cluster on Exascale Infrastructure. If false, database backup on local Exadata storage is not available in the Exadata VM cluster on Exascale Infrastructure.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// If true, sparse disk group is configured for the Exadata VM cluster on Exascale Infrastructure. If false, sparse disk group is not created.
	IsSparseDiskgroupEnabled *bool `mandatory:"false" json:"isSparseDiskgroupEnabled"`

	// The time zone to use for the Exadata VM cluster on Exascale Infrastructure. For details, see Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The TCP Single Client Access Name (SCAN) port. The default port is 1521.
	ScanListenerPortTcp *int `mandatory:"false" json:"scanListenerPortTcp"`

	// The TCPS Single Client Access Name (SCAN) port. The default port is 2484.
	ScanListenerPortTcpSsl *int `mandatory:"false" json:"scanListenerPortTcpSsl"`

	// The private zone id in which DNS records needs to be created.
	PrivateZoneId *string `mandatory:"false" json:"privateZoneId"`

	ExtremeFlashDatabaseStorage *ExadbVmClusterStorageDetails `mandatory:"false" json:"extremeFlashDatabaseStorage"`

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

func (m CreateExadbVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExadbVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateExadbVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateExadbVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateExadbVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type CreateExadbVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateExadbVmClusterDetailsLicenseModelEnum
const (
	CreateExadbVmClusterDetailsLicenseModelLicenseIncluded     CreateExadbVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateExadbVmClusterDetailsLicenseModelBringYourOwnLicense CreateExadbVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateExadbVmClusterDetailsLicenseModelEnum = map[string]CreateExadbVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateExadbVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateExadbVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateExadbVmClusterDetailsLicenseModelEnumLowerCase = map[string]CreateExadbVmClusterDetailsLicenseModelEnum{
	"license_included":       CreateExadbVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateExadbVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateExadbVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for CreateExadbVmClusterDetailsLicenseModelEnum
func GetCreateExadbVmClusterDetailsLicenseModelEnumValues() []CreateExadbVmClusterDetailsLicenseModelEnum {
	values := make([]CreateExadbVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateExadbVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateExadbVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateExadbVmClusterDetailsLicenseModelEnum
func GetCreateExadbVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateExadbVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateExadbVmClusterDetailsLicenseModelEnum(val string) (CreateExadbVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateExadbVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
