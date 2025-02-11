// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LaunchDbSystemBase Parameters for provisioning a bare metal, virtual machine, or Exadata DB system.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type LaunchDbSystemBase interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment the DB system  belongs in.
	GetCompartmentId() *string

	// The availability domain where the DB system is located.
	GetAvailabilityDomain() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	GetSubnetId() *string

	// The shape of the DB system. The shape determines resources allocated to the DB system.
	// - For virtual machine shapes, the number of CPU cores and memory
	// - For bare metal and Exadata shapes, the number of CPU cores, memory, and storage
	// To get a list of shapes, use the ListDbSystemShapes operation.
	GetShape() *string

	// The public key portion of the key pair to use for SSH access to the DB system. Multiple public keys can be provided. The length of the combined keys cannot exceed 40,000 characters.
	GetSshPublicKeys() []string

	// The hostname for the DB system. The hostname must begin with an alphabetic character, and
	// can contain alphanumeric characters and hyphens (-). The maximum length of the hostname is 16 characters for bare metal and virtual machine DB systems, and 12 characters for Exadata DB systems.
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// the DB system will fail to provision.
	GetHostname() *string

	// The number of CPU cores to enable for a bare metal or Exadata DB system or AMD VMDB Systems. The valid values depend on the specified shape:
	// - BM.DenseIO1.36 - Specify a multiple of 2, from 2 to 36.
	// - BM.DenseIO2.52 - Specify a multiple of 2, from 2 to 52.
	// - Exadata.Base.48 - Specify a multiple of 2, from 0 to 48.
	// - Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84.
	// - Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168.
	// - Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.
	// - Exadata.Quarter2.92 - Specify a multiple of 2, from 0 to 92.
	// - Exadata.Half2.184 - Specify a multiple of 4, from 0 to 184.
	// - Exadata.Full2.368 - Specify a multiple of 8, from 0 to 368.
	// - VM.Standard.E4.Flex - Specify any thing from 1 to 64.
	// This parameter is not used for INTEL virtual machine DB systems because virtual machine DB systems have a set number of cores for each shape.
	// For information about the number of cores for a virtual machine DB system shape, see Virtual Machine DB Systems (https://docs.cloud.oracle.com/Content/Database/Concepts/overview.htm#virtualmachine)
	GetCpuCoreCount() *int

	// A Fault Domain is a grouping of hardware and infrastructure within an availability domain.
	// Fault Domains let you distribute your instances so that they are not on the same physical
	// hardware within a single availability domain. A hardware failure or maintenance
	// that affects one Fault Domain does not affect DB systems in other Fault Domains.
	// If you do not specify the Fault Domain, the system selects one for you. To change the Fault
	// Domain for a DB system, terminate it and launch a new DB system in the preferred Fault Domain.
	// If the node count is greater than 1, you can specify which Fault Domains these nodes will be distributed into.
	// The system assigns your nodes automatically to the Fault Domains you specify so that
	// no Fault Domain contains more than one node.
	// To get a list of Fault Domains, use the
	// ListFaultDomains operation in the
	// Identity and Access Management Service API.
	// Example: `FAULT-DOMAIN-1`
	GetFaultDomains() []string

	// The user-friendly name for the DB system. The name does not have to be unique.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.
	// **Subnet Restrictions:** See the subnet restrictions information for **subnetId**.
	GetBackupSubnetId() *string

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	GetNsgIds() []string

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	GetBackupNetworkNsgIds() []string

	// The time zone to use for the DB system. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	GetTimeZone() *string

	GetDbSystemOptions() *DbSystemOptions

	// The block storage volume performance level. Valid values are `BALANCED` and `HIGH_PERFORMANCE`. See Block Volume Performance (https://docs.cloud.oracle.com/Content/Block/Concepts/blockvolumeperformance.htm) for more information.
	GetStorageVolumePerformanceMode() LaunchDbSystemBaseStorageVolumePerformanceModeEnum

	// If true, Sparse Diskgroup is configured for Exadata dbsystem. If False, Sparse diskgroup is not configured.
	GetSparseDiskgroup() *bool

	// A domain name used for the DB system. If the Oracle-provided Internet and VCN
	// Resolver is enabled for the specified subnet, the domain name for the subnet is used
	// (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.
	GetDomain() *string

	// The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	GetClusterName() *string

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups).
	// Specify 80 or 40. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems.
	GetDataStoragePercentage() *int

	// Size (in GB) of the initial data volume that will be created and attached to a virtual machine DB system. You can scale up storage after provisioning, as needed. Note that the total storage size attached will be more than the amount you specify to allow for REDO/RECO space and software volume.
	GetInitialDataStorageSizeInGB() *int

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	GetKmsKeyId() *string

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	GetKmsKeyVersionId() *string

	// The number of nodes to launch for a 2-node RAC virtual machine DB system. Specify either 1 or 2.
	GetNodeCount() *int

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	GetSecurityAttributes() map[string]map[string]interface{}

	// A private IP address of your choice. Must be an available IP address within the subnet's CIDR.
	// If you don't specify a value, Oracle automatically assigns a private IP address from the subnet.
	GetPrivateIp() *string

	// A private IPv6 address of your choice. Must be an available IP address within the subnet's CIDR.
	// If you don't specify a value and the subnet is dual stack, Oracle automatically assigns a private IPv6 address from the subnet.
	GetPrivateIpV6() *string

	GetDataCollectionOptions() *DataCollectionOptions
}

type launchdbsystembase struct {
	JsonData                     []byte
	FaultDomains                 []string                                           `mandatory:"false" json:"faultDomains"`
	DisplayName                  *string                                            `mandatory:"false" json:"displayName"`
	BackupSubnetId               *string                                            `mandatory:"false" json:"backupSubnetId"`
	NsgIds                       []string                                           `mandatory:"false" json:"nsgIds"`
	BackupNetworkNsgIds          []string                                           `mandatory:"false" json:"backupNetworkNsgIds"`
	TimeZone                     *string                                            `mandatory:"false" json:"timeZone"`
	DbSystemOptions              *DbSystemOptions                                   `mandatory:"false" json:"dbSystemOptions"`
	StorageVolumePerformanceMode LaunchDbSystemBaseStorageVolumePerformanceModeEnum `mandatory:"false" json:"storageVolumePerformanceMode,omitempty"`
	SparseDiskgroup              *bool                                              `mandatory:"false" json:"sparseDiskgroup"`
	Domain                       *string                                            `mandatory:"false" json:"domain"`
	ClusterName                  *string                                            `mandatory:"false" json:"clusterName"`
	DataStoragePercentage        *int                                               `mandatory:"false" json:"dataStoragePercentage"`
	InitialDataStorageSizeInGB   *int                                               `mandatory:"false" json:"initialDataStorageSizeInGB"`
	KmsKeyId                     *string                                            `mandatory:"false" json:"kmsKeyId"`
	KmsKeyVersionId              *string                                            `mandatory:"false" json:"kmsKeyVersionId"`
	NodeCount                    *int                                               `mandatory:"false" json:"nodeCount"`
	FreeformTags                 map[string]string                                  `mandatory:"false" json:"freeformTags"`
	DefinedTags                  map[string]map[string]interface{}                  `mandatory:"false" json:"definedTags"`
	SecurityAttributes           map[string]map[string]interface{}                  `mandatory:"false" json:"securityAttributes"`
	PrivateIp                    *string                                            `mandatory:"false" json:"privateIp"`
	PrivateIpV6                  *string                                            `mandatory:"false" json:"privateIpV6"`
	DataCollectionOptions        *DataCollectionOptions                             `mandatory:"false" json:"dataCollectionOptions"`
	CompartmentId                *string                                            `mandatory:"true" json:"compartmentId"`
	AvailabilityDomain           *string                                            `mandatory:"true" json:"availabilityDomain"`
	SubnetId                     *string                                            `mandatory:"true" json:"subnetId"`
	Shape                        *string                                            `mandatory:"true" json:"shape"`
	SshPublicKeys                []string                                           `mandatory:"true" json:"sshPublicKeys"`
	Hostname                     *string                                            `mandatory:"true" json:"hostname"`
	CpuCoreCount                 *int                                               `mandatory:"true" json:"cpuCoreCount"`
	Source                       string                                             `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *launchdbsystembase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlaunchdbsystembase launchdbsystembase
	s := struct {
		Model Unmarshalerlaunchdbsystembase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.AvailabilityDomain = s.Model.AvailabilityDomain
	m.SubnetId = s.Model.SubnetId
	m.Shape = s.Model.Shape
	m.SshPublicKeys = s.Model.SshPublicKeys
	m.Hostname = s.Model.Hostname
	m.CpuCoreCount = s.Model.CpuCoreCount
	m.FaultDomains = s.Model.FaultDomains
	m.DisplayName = s.Model.DisplayName
	m.BackupSubnetId = s.Model.BackupSubnetId
	m.NsgIds = s.Model.NsgIds
	m.BackupNetworkNsgIds = s.Model.BackupNetworkNsgIds
	m.TimeZone = s.Model.TimeZone
	m.DbSystemOptions = s.Model.DbSystemOptions
	m.StorageVolumePerformanceMode = s.Model.StorageVolumePerformanceMode
	m.SparseDiskgroup = s.Model.SparseDiskgroup
	m.Domain = s.Model.Domain
	m.ClusterName = s.Model.ClusterName
	m.DataStoragePercentage = s.Model.DataStoragePercentage
	m.InitialDataStorageSizeInGB = s.Model.InitialDataStorageSizeInGB
	m.KmsKeyId = s.Model.KmsKeyId
	m.KmsKeyVersionId = s.Model.KmsKeyVersionId
	m.NodeCount = s.Model.NodeCount
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SecurityAttributes = s.Model.SecurityAttributes
	m.PrivateIp = s.Model.PrivateIp
	m.PrivateIpV6 = s.Model.PrivateIpV6
	m.DataCollectionOptions = s.Model.DataCollectionOptions
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *launchdbsystembase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "NONE":
		mm := LaunchDbSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_SYSTEM":
		mm := LaunchDbSystemFromDbSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := LaunchDbSystemFromDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_BACKUP":
		mm := LaunchDbSystemFromBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for LaunchDbSystemBase: %s.", m.Source)
		return *m, nil
	}
}

// GetFaultDomains returns FaultDomains
func (m launchdbsystembase) GetFaultDomains() []string {
	return m.FaultDomains
}

// GetDisplayName returns DisplayName
func (m launchdbsystembase) GetDisplayName() *string {
	return m.DisplayName
}

// GetBackupSubnetId returns BackupSubnetId
func (m launchdbsystembase) GetBackupSubnetId() *string {
	return m.BackupSubnetId
}

// GetNsgIds returns NsgIds
func (m launchdbsystembase) GetNsgIds() []string {
	return m.NsgIds
}

// GetBackupNetworkNsgIds returns BackupNetworkNsgIds
func (m launchdbsystembase) GetBackupNetworkNsgIds() []string {
	return m.BackupNetworkNsgIds
}

// GetTimeZone returns TimeZone
func (m launchdbsystembase) GetTimeZone() *string {
	return m.TimeZone
}

// GetDbSystemOptions returns DbSystemOptions
func (m launchdbsystembase) GetDbSystemOptions() *DbSystemOptions {
	return m.DbSystemOptions
}

// GetStorageVolumePerformanceMode returns StorageVolumePerformanceMode
func (m launchdbsystembase) GetStorageVolumePerformanceMode() LaunchDbSystemBaseStorageVolumePerformanceModeEnum {
	return m.StorageVolumePerformanceMode
}

// GetSparseDiskgroup returns SparseDiskgroup
func (m launchdbsystembase) GetSparseDiskgroup() *bool {
	return m.SparseDiskgroup
}

// GetDomain returns Domain
func (m launchdbsystembase) GetDomain() *string {
	return m.Domain
}

// GetClusterName returns ClusterName
func (m launchdbsystembase) GetClusterName() *string {
	return m.ClusterName
}

// GetDataStoragePercentage returns DataStoragePercentage
func (m launchdbsystembase) GetDataStoragePercentage() *int {
	return m.DataStoragePercentage
}

// GetInitialDataStorageSizeInGB returns InitialDataStorageSizeInGB
func (m launchdbsystembase) GetInitialDataStorageSizeInGB() *int {
	return m.InitialDataStorageSizeInGB
}

// GetKmsKeyId returns KmsKeyId
func (m launchdbsystembase) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetKmsKeyVersionId returns KmsKeyVersionId
func (m launchdbsystembase) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

// GetNodeCount returns NodeCount
func (m launchdbsystembase) GetNodeCount() *int {
	return m.NodeCount
}

// GetFreeformTags returns FreeformTags
func (m launchdbsystembase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m launchdbsystembase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSecurityAttributes returns SecurityAttributes
func (m launchdbsystembase) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

// GetPrivateIp returns PrivateIp
func (m launchdbsystembase) GetPrivateIp() *string {
	return m.PrivateIp
}

// GetPrivateIpV6 returns PrivateIpV6
func (m launchdbsystembase) GetPrivateIpV6() *string {
	return m.PrivateIpV6
}

// GetDataCollectionOptions returns DataCollectionOptions
func (m launchdbsystembase) GetDataCollectionOptions() *DataCollectionOptions {
	return m.DataCollectionOptions
}

// GetCompartmentId returns CompartmentId
func (m launchdbsystembase) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m launchdbsystembase) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

// GetSubnetId returns SubnetId
func (m launchdbsystembase) GetSubnetId() *string {
	return m.SubnetId
}

// GetShape returns Shape
func (m launchdbsystembase) GetShape() *string {
	return m.Shape
}

// GetSshPublicKeys returns SshPublicKeys
func (m launchdbsystembase) GetSshPublicKeys() []string {
	return m.SshPublicKeys
}

// GetHostname returns Hostname
func (m launchdbsystembase) GetHostname() *string {
	return m.Hostname
}

// GetCpuCoreCount returns CpuCoreCount
func (m launchdbsystembase) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

func (m launchdbsystembase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m launchdbsystembase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLaunchDbSystemBaseStorageVolumePerformanceModeEnum(string(m.StorageVolumePerformanceMode)); !ok && m.StorageVolumePerformanceMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageVolumePerformanceMode: %s. Supported values are: %s.", m.StorageVolumePerformanceMode, strings.Join(GetLaunchDbSystemBaseStorageVolumePerformanceModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LaunchDbSystemBaseStorageVolumePerformanceModeEnum Enum with underlying type: string
type LaunchDbSystemBaseStorageVolumePerformanceModeEnum string

// Set of constants representing the allowable values for LaunchDbSystemBaseStorageVolumePerformanceModeEnum
const (
	LaunchDbSystemBaseStorageVolumePerformanceModeBalanced        LaunchDbSystemBaseStorageVolumePerformanceModeEnum = "BALANCED"
	LaunchDbSystemBaseStorageVolumePerformanceModeHighPerformance LaunchDbSystemBaseStorageVolumePerformanceModeEnum = "HIGH_PERFORMANCE"
)

var mappingLaunchDbSystemBaseStorageVolumePerformanceModeEnum = map[string]LaunchDbSystemBaseStorageVolumePerformanceModeEnum{
	"BALANCED":         LaunchDbSystemBaseStorageVolumePerformanceModeBalanced,
	"HIGH_PERFORMANCE": LaunchDbSystemBaseStorageVolumePerformanceModeHighPerformance,
}

var mappingLaunchDbSystemBaseStorageVolumePerformanceModeEnumLowerCase = map[string]LaunchDbSystemBaseStorageVolumePerformanceModeEnum{
	"balanced":         LaunchDbSystemBaseStorageVolumePerformanceModeBalanced,
	"high_performance": LaunchDbSystemBaseStorageVolumePerformanceModeHighPerformance,
}

// GetLaunchDbSystemBaseStorageVolumePerformanceModeEnumValues Enumerates the set of values for LaunchDbSystemBaseStorageVolumePerformanceModeEnum
func GetLaunchDbSystemBaseStorageVolumePerformanceModeEnumValues() []LaunchDbSystemBaseStorageVolumePerformanceModeEnum {
	values := make([]LaunchDbSystemBaseStorageVolumePerformanceModeEnum, 0)
	for _, v := range mappingLaunchDbSystemBaseStorageVolumePerformanceModeEnum {
		values = append(values, v)
	}
	return values
}

// GetLaunchDbSystemBaseStorageVolumePerformanceModeEnumStringValues Enumerates the set of values in String for LaunchDbSystemBaseStorageVolumePerformanceModeEnum
func GetLaunchDbSystemBaseStorageVolumePerformanceModeEnumStringValues() []string {
	return []string{
		"BALANCED",
		"HIGH_PERFORMANCE",
	}
}

// GetMappingLaunchDbSystemBaseStorageVolumePerformanceModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLaunchDbSystemBaseStorageVolumePerformanceModeEnum(val string) (LaunchDbSystemBaseStorageVolumePerformanceModeEnum, bool) {
	enum, ok := mappingLaunchDbSystemBaseStorageVolumePerformanceModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LaunchDbSystemBaseSourceEnum Enum with underlying type: string
type LaunchDbSystemBaseSourceEnum string

// Set of constants representing the allowable values for LaunchDbSystemBaseSourceEnum
const (
	LaunchDbSystemBaseSourceNone     LaunchDbSystemBaseSourceEnum = "NONE"
	LaunchDbSystemBaseSourceDbBackup LaunchDbSystemBaseSourceEnum = "DB_BACKUP"
	LaunchDbSystemBaseSourceDatabase LaunchDbSystemBaseSourceEnum = "DATABASE"
	LaunchDbSystemBaseSourceDbSystem LaunchDbSystemBaseSourceEnum = "DB_SYSTEM"
)

var mappingLaunchDbSystemBaseSourceEnum = map[string]LaunchDbSystemBaseSourceEnum{
	"NONE":      LaunchDbSystemBaseSourceNone,
	"DB_BACKUP": LaunchDbSystemBaseSourceDbBackup,
	"DATABASE":  LaunchDbSystemBaseSourceDatabase,
	"DB_SYSTEM": LaunchDbSystemBaseSourceDbSystem,
}

var mappingLaunchDbSystemBaseSourceEnumLowerCase = map[string]LaunchDbSystemBaseSourceEnum{
	"none":      LaunchDbSystemBaseSourceNone,
	"db_backup": LaunchDbSystemBaseSourceDbBackup,
	"database":  LaunchDbSystemBaseSourceDatabase,
	"db_system": LaunchDbSystemBaseSourceDbSystem,
}

// GetLaunchDbSystemBaseSourceEnumValues Enumerates the set of values for LaunchDbSystemBaseSourceEnum
func GetLaunchDbSystemBaseSourceEnumValues() []LaunchDbSystemBaseSourceEnum {
	values := make([]LaunchDbSystemBaseSourceEnum, 0)
	for _, v := range mappingLaunchDbSystemBaseSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetLaunchDbSystemBaseSourceEnumStringValues Enumerates the set of values in String for LaunchDbSystemBaseSourceEnum
func GetLaunchDbSystemBaseSourceEnumStringValues() []string {
	return []string{
		"NONE",
		"DB_BACKUP",
		"DATABASE",
		"DB_SYSTEM",
	}
}

// GetMappingLaunchDbSystemBaseSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLaunchDbSystemBaseSourceEnum(val string) (LaunchDbSystemBaseSourceEnum, bool) {
	enum, ok := mappingLaunchDbSystemBaseSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
