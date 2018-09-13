// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// LaunchDbSystemDetails The representation of LaunchDbSystemDetails
type LaunchDbSystemDetails struct {

	// The availability domain where the DB system is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment the DB system  belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of CPU cores to enable for a bare metal or Exadata DB system. The valid values depend on the specified shape:
	// - BM.DenseIO1.36 - Specify a multiple of 2, from 2 to 36.
	// - BM.DenseIO2.52 - Specify a multiple of 2, from 2 to 52.
	// - Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84.
	// - Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168.
	// - Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.
	// - Exadata.Quarter2.92 - Specify a multiple of 2, from 0 to 92.
	// - Exadata.Half2.184 - Specify a multiple of 4, from 0 to 184.
	// - Exadata.Full2.368 - Specify a multiple of 8, from 0 to 368.
	// This parameter is not used for virtual machine DB systems because virtual machine DB systems have a set number of cores for each shape.
	// For information about the number of cores for a virtual machine DB system shape, see Virtual Machine DB Systems (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Concepts/overview.htm#virtualmachine)
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The hostname for the DB system. The hostname must begin with an alphabetic character and
	// can contain a maximum of 30 alphanumeric characters, including hyphens (-).
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// the DB system will fail to provision.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The shape of the DB system. The shape determines resources allocated to the DB system.
	// - For virtual machine shapes, the number of CPU cores and memory
	// - For bare metal and Exadata shapes, the number of CPU cores, memory, and storage
	// To get a list of shapes, use the ListDbSystemShapes operation.
	Shape *string `mandatory:"true" json:"shape"`

	// The public key portion of the key pair to use for SSH access to the DB system. Multiple public keys can be provided. The length of the combined keys cannot exceed 10,000 characters.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	DbHome *CreateDbHomeDetails `mandatory:"true" json:"dbHome"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.
	// **Subnet Restrictions:** See the subnet restrictions information for **subnetId**.
	BackupSubnetId *string `mandatory:"false" json:"backupSubnetId"`

	// The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups).
	// Specify 80 or 40. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The user-friendly name for the DB system. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A domain name used for the DB system. If the Oracle-provided Internet and VCN
	// Resolver is enabled for the specified subnet, the domain name for the subnet is used
	// (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.
	Domain *string `mandatory:"false" json:"domain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Size (in GB) of the initial data volume that will be created and attached to a virtual machine DB system. You can scale up storage after provisioning, as needed. Note that the total storage size attached will be more than the amount you specify to allow for REDO/RECO space and software volume.
	InitialDataStorageSizeInGB *int `mandatory:"false" json:"initialDataStorageSizeInGB"`

	// The number of nodes to launch for a 2-node RAC virtual machine DB system.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The Oracle Database Edition that applies to all the databases on the DB system.
	// Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE.
	DatabaseEdition LaunchDbSystemDetailsDatabaseEditionEnum `mandatory:"true" json:"databaseEdition"`

	// The type of redundancy configured for the DB system.
	// Normal is 2-way redundancy, recommended for test and development systems.
	// High is 3-way redundancy, recommended for production systems.
	DiskRedundancy LaunchDbSystemDetailsDiskRedundancyEnum `mandatory:"false" json:"diskRedundancy,omitempty"`

	// The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED.
	LicenseModel LaunchDbSystemDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
}

//GetAvailabilityDomain returns AvailabilityDomain
func (m LaunchDbSystemDetails) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

//GetBackupSubnetId returns BackupSubnetId
func (m LaunchDbSystemDetails) GetBackupSubnetId() *string {
	return m.BackupSubnetId
}

//GetClusterName returns ClusterName
func (m LaunchDbSystemDetails) GetClusterName() *string {
	return m.ClusterName
}

//GetCompartmentId returns CompartmentId
func (m LaunchDbSystemDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetCpuCoreCount returns CpuCoreCount
func (m LaunchDbSystemDetails) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

//GetDataStoragePercentage returns DataStoragePercentage
func (m LaunchDbSystemDetails) GetDataStoragePercentage() *int {
	return m.DataStoragePercentage
}

//GetDefinedTags returns DefinedTags
func (m LaunchDbSystemDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetDisplayName returns DisplayName
func (m LaunchDbSystemDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDomain returns Domain
func (m LaunchDbSystemDetails) GetDomain() *string {
	return m.Domain
}

//GetFreeformTags returns FreeformTags
func (m LaunchDbSystemDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetHostname returns Hostname
func (m LaunchDbSystemDetails) GetHostname() *string {
	return m.Hostname
}

//GetInitialDataStorageSizeInGB returns InitialDataStorageSizeInGB
func (m LaunchDbSystemDetails) GetInitialDataStorageSizeInGB() *int {
	return m.InitialDataStorageSizeInGB
}

//GetNodeCount returns NodeCount
func (m LaunchDbSystemDetails) GetNodeCount() *int {
	return m.NodeCount
}

//GetShape returns Shape
func (m LaunchDbSystemDetails) GetShape() *string {
	return m.Shape
}

//GetSshPublicKeys returns SshPublicKeys
func (m LaunchDbSystemDetails) GetSshPublicKeys() []string {
	return m.SshPublicKeys
}

//GetSubnetId returns SubnetId
func (m LaunchDbSystemDetails) GetSubnetId() *string {
	return m.SubnetId
}

func (m LaunchDbSystemDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m LaunchDbSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLaunchDbSystemDetails LaunchDbSystemDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeLaunchDbSystemDetails
	}{
		"NONE",
		(MarshalTypeLaunchDbSystemDetails)(m),
	}

	return json.Marshal(&s)
}

// LaunchDbSystemDetailsDatabaseEditionEnum Enum with underlying type: string
type LaunchDbSystemDetailsDatabaseEditionEnum string

// Set of constants representing the allowable values for LaunchDbSystemDetailsDatabaseEdition
const (
	LaunchDbSystemDetailsDatabaseEditionStandardEdition                     LaunchDbSystemDetailsDatabaseEditionEnum = "STANDARD_EDITION"
	LaunchDbSystemDetailsDatabaseEditionEnterpriseEdition                   LaunchDbSystemDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION"
	LaunchDbSystemDetailsDatabaseEditionEnterpriseEditionHighPerformance    LaunchDbSystemDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	LaunchDbSystemDetailsDatabaseEditionEnterpriseEditionExtremePerformance LaunchDbSystemDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingLaunchDbSystemDetailsDatabaseEdition = map[string]LaunchDbSystemDetailsDatabaseEditionEnum{
	"STANDARD_EDITION":                       LaunchDbSystemDetailsDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     LaunchDbSystemDetailsDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    LaunchDbSystemDetailsDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": LaunchDbSystemDetailsDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetLaunchDbSystemDetailsDatabaseEditionEnumValues Enumerates the set of values for LaunchDbSystemDetailsDatabaseEdition
func GetLaunchDbSystemDetailsDatabaseEditionEnumValues() []LaunchDbSystemDetailsDatabaseEditionEnum {
	values := make([]LaunchDbSystemDetailsDatabaseEditionEnum, 0)
	for _, v := range mappingLaunchDbSystemDetailsDatabaseEdition {
		values = append(values, v)
	}
	return values
}

// LaunchDbSystemDetailsDiskRedundancyEnum Enum with underlying type: string
type LaunchDbSystemDetailsDiskRedundancyEnum string

// Set of constants representing the allowable values for LaunchDbSystemDetailsDiskRedundancy
const (
	LaunchDbSystemDetailsDiskRedundancyHigh   LaunchDbSystemDetailsDiskRedundancyEnum = "HIGH"
	LaunchDbSystemDetailsDiskRedundancyNormal LaunchDbSystemDetailsDiskRedundancyEnum = "NORMAL"
)

var mappingLaunchDbSystemDetailsDiskRedundancy = map[string]LaunchDbSystemDetailsDiskRedundancyEnum{
	"HIGH":   LaunchDbSystemDetailsDiskRedundancyHigh,
	"NORMAL": LaunchDbSystemDetailsDiskRedundancyNormal,
}

// GetLaunchDbSystemDetailsDiskRedundancyEnumValues Enumerates the set of values for LaunchDbSystemDetailsDiskRedundancy
func GetLaunchDbSystemDetailsDiskRedundancyEnumValues() []LaunchDbSystemDetailsDiskRedundancyEnum {
	values := make([]LaunchDbSystemDetailsDiskRedundancyEnum, 0)
	for _, v := range mappingLaunchDbSystemDetailsDiskRedundancy {
		values = append(values, v)
	}
	return values
}

// LaunchDbSystemDetailsLicenseModelEnum Enum with underlying type: string
type LaunchDbSystemDetailsLicenseModelEnum string

// Set of constants representing the allowable values for LaunchDbSystemDetailsLicenseModel
const (
	LaunchDbSystemDetailsLicenseModelLicenseIncluded     LaunchDbSystemDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	LaunchDbSystemDetailsLicenseModelBringYourOwnLicense LaunchDbSystemDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingLaunchDbSystemDetailsLicenseModel = map[string]LaunchDbSystemDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       LaunchDbSystemDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": LaunchDbSystemDetailsLicenseModelBringYourOwnLicense,
}

// GetLaunchDbSystemDetailsLicenseModelEnumValues Enumerates the set of values for LaunchDbSystemDetailsLicenseModel
func GetLaunchDbSystemDetailsLicenseModelEnumValues() []LaunchDbSystemDetailsLicenseModelEnum {
	values := make([]LaunchDbSystemDetailsLicenseModelEnum, 0)
	for _, v := range mappingLaunchDbSystemDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
