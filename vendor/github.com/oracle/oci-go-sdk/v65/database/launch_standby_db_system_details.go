// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LaunchStandbyDbSystemDetails Used for creating a new standby database system in dataguard environment.
type LaunchStandbyDbSystemDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DB system  belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain where the DB system is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The shape of the DB system. The shape determines resources allocated to the DB system.
	// - For virtual machine shapes, the number of CPU cores and memory
	// - For bare metal and Exadata shapes, the number of CPU cores, memory, and storage
	// To get a list of shapes, use the ListDbSystemShapes operation.
	Shape *string `mandatory:"true" json:"shape"`

	// The public key portion of the key pair to use for SSH access to the DB system. Multiple public keys can be provided. The length of the combined keys cannot exceed 40,000 characters.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The hostname for the DB system. The hostname must begin with an alphabetic character, and
	// can contain alphanumeric characters and hyphens (-). The maximum length of the hostname is 16 characters for bare metal and virtual machine DB systems, and 12 characters for Exadata DB systems.
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// the DB system will fail to provision.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	PrimaryDbSystemId *string `mandatory:"true" json:"primaryDbSystemId"`

	DbHome *CreateDbHomeForStandbyDbSystem `mandatory:"true" json:"dbHome"`

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
	FaultDomains []string `mandatory:"false" json:"faultDomains"`

	// The user-friendly name for the DB system. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.
	// **Subnet Restrictions:** See the subnet restrictions information for **subnetId**.
	BackupSubnetId *string `mandatory:"false" json:"backupSubnetId"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous AI Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The time zone to use for the DB system. For details, see DB System Time Zones (https://docs.oracle.com/iaas/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	DbSystemOptions *DbSystemOptions `mandatory:"false" json:"dbSystemOptions"`

	// If true, Sparse Diskgroup is configured for Exadata dbsystem. If False, Sparse diskgroup is not configured.
	SparseDiskgroup *bool `mandatory:"false" json:"sparseDiskgroup"`

	// A domain name used for the DB system. If the Oracle-provided Internet and VCN
	// Resolver is enabled for the specified subnet, the domain name for the subnet is used
	// (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.
	Domain *string `mandatory:"false" json:"domain"`

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
	// For information about the number of cores for a virtual machine DB system shape, see Virtual Machine DB Systems (https://docs.oracle.com/iaas/Content/Database/Concepts/overview.htm#virtualmachine)
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups).
	// Specify 80 or 40. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage"`

	// Size (in GB) of the initial data volume that will be created and attached to a virtual machine DB system. You can scale up storage after provisioning, as needed. Note that the total storage size attached will be more than the amount you specify to allow for REDO/RECO space and software volume. By default this will be set to 256.
	InitialDataStorageSizeInGB *int `mandatory:"false" json:"initialDataStorageSizeInGB"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The number of nodes to launch for a virtual machine DB system. Specify either 1 or 2. By default this will be set to 1.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

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

	// A private IP address of your choice. Must be an available IP address within the subnet's CIDR.
	// If you don't specify a value, Oracle automatically assigns a private IP address from the subnet.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// A private IPv6 address of your choice. Must be an available IP address within the subnet's CIDR.
	// If you don't specify a value and the subnet is dual stack, Oracle automatically assigns a private IPv6 address from the subnet.
	PrivateIpV6 *string `mandatory:"false" json:"privateIpV6"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Exadata Infrastructure or Db System.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	// The number of compute servers for the DB system.
	ComputeCount *int `mandatory:"false" json:"computeCount"`

	// The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED.
	LicenseModel LaunchStandbyDbSystemDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The block storage volume performance level. Valid values are `BALANCED` and `HIGH_PERFORMANCE`. See Block Volume Performance (https://docs.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm) for more information.
	StorageVolumePerformanceMode LaunchDbSystemBaseStorageVolumePerformanceModeEnum `mandatory:"false" json:"storageVolumePerformanceMode,omitempty"`

	// The compute model for Base Database Service. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. The ECPU compute model is the recommended model, and the OCPU compute model is legacy.
	ComputeModel LaunchDbSystemBaseComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`
}

// GetCompartmentId returns CompartmentId
func (m LaunchStandbyDbSystemDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFaultDomains returns FaultDomains
func (m LaunchStandbyDbSystemDetails) GetFaultDomains() []string {
	return m.FaultDomains
}

// GetDisplayName returns DisplayName
func (m LaunchStandbyDbSystemDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m LaunchStandbyDbSystemDetails) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

// GetSubnetId returns SubnetId
func (m LaunchStandbyDbSystemDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetBackupSubnetId returns BackupSubnetId
func (m LaunchStandbyDbSystemDetails) GetBackupSubnetId() *string {
	return m.BackupSubnetId
}

// GetNsgIds returns NsgIds
func (m LaunchStandbyDbSystemDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetBackupNetworkNsgIds returns BackupNetworkNsgIds
func (m LaunchStandbyDbSystemDetails) GetBackupNetworkNsgIds() []string {
	return m.BackupNetworkNsgIds
}

// GetShape returns Shape
func (m LaunchStandbyDbSystemDetails) GetShape() *string {
	return m.Shape
}

// GetTimeZone returns TimeZone
func (m LaunchStandbyDbSystemDetails) GetTimeZone() *string {
	return m.TimeZone
}

// GetDbSystemOptions returns DbSystemOptions
func (m LaunchStandbyDbSystemDetails) GetDbSystemOptions() *DbSystemOptions {
	return m.DbSystemOptions
}

// GetStorageVolumePerformanceMode returns StorageVolumePerformanceMode
func (m LaunchStandbyDbSystemDetails) GetStorageVolumePerformanceMode() LaunchDbSystemBaseStorageVolumePerformanceModeEnum {
	return m.StorageVolumePerformanceMode
}

// GetSparseDiskgroup returns SparseDiskgroup
func (m LaunchStandbyDbSystemDetails) GetSparseDiskgroup() *bool {
	return m.SparseDiskgroup
}

// GetSshPublicKeys returns SshPublicKeys
func (m LaunchStandbyDbSystemDetails) GetSshPublicKeys() []string {
	return m.SshPublicKeys
}

// GetHostname returns Hostname
func (m LaunchStandbyDbSystemDetails) GetHostname() *string {
	return m.Hostname
}

// GetDomain returns Domain
func (m LaunchStandbyDbSystemDetails) GetDomain() *string {
	return m.Domain
}

// GetCpuCoreCount returns CpuCoreCount
func (m LaunchStandbyDbSystemDetails) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

// GetClusterName returns ClusterName
func (m LaunchStandbyDbSystemDetails) GetClusterName() *string {
	return m.ClusterName
}

// GetDataStoragePercentage returns DataStoragePercentage
func (m LaunchStandbyDbSystemDetails) GetDataStoragePercentage() *int {
	return m.DataStoragePercentage
}

// GetInitialDataStorageSizeInGB returns InitialDataStorageSizeInGB
func (m LaunchStandbyDbSystemDetails) GetInitialDataStorageSizeInGB() *int {
	return m.InitialDataStorageSizeInGB
}

// GetKmsKeyId returns KmsKeyId
func (m LaunchStandbyDbSystemDetails) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetKmsKeyVersionId returns KmsKeyVersionId
func (m LaunchStandbyDbSystemDetails) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

// GetNodeCount returns NodeCount
func (m LaunchStandbyDbSystemDetails) GetNodeCount() *int {
	return m.NodeCount
}

// GetFreeformTags returns FreeformTags
func (m LaunchStandbyDbSystemDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m LaunchStandbyDbSystemDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSecurityAttributes returns SecurityAttributes
func (m LaunchStandbyDbSystemDetails) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

// GetPrivateIp returns PrivateIp
func (m LaunchStandbyDbSystemDetails) GetPrivateIp() *string {
	return m.PrivateIp
}

// GetPrivateIpV6 returns PrivateIpV6
func (m LaunchStandbyDbSystemDetails) GetPrivateIpV6() *string {
	return m.PrivateIpV6
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m LaunchStandbyDbSystemDetails) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSubscriptionId returns SubscriptionId
func (m LaunchStandbyDbSystemDetails) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetDataCollectionOptions returns DataCollectionOptions
func (m LaunchStandbyDbSystemDetails) GetDataCollectionOptions() *DataCollectionOptions {
	return m.DataCollectionOptions
}

// GetComputeModel returns ComputeModel
func (m LaunchStandbyDbSystemDetails) GetComputeModel() LaunchDbSystemBaseComputeModelEnum {
	return m.ComputeModel
}

// GetComputeCount returns ComputeCount
func (m LaunchStandbyDbSystemDetails) GetComputeCount() *int {
	return m.ComputeCount
}

func (m LaunchStandbyDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LaunchStandbyDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLaunchStandbyDbSystemDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLaunchStandbyDbSystemDetailsLicenseModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLaunchDbSystemBaseStorageVolumePerformanceModeEnum(string(m.StorageVolumePerformanceMode)); !ok && m.StorageVolumePerformanceMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageVolumePerformanceMode: %s. Supported values are: %s.", m.StorageVolumePerformanceMode, strings.Join(GetLaunchDbSystemBaseStorageVolumePerformanceModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLaunchDbSystemBaseComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetLaunchDbSystemBaseComputeModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LaunchStandbyDbSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLaunchStandbyDbSystemDetails LaunchStandbyDbSystemDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeLaunchStandbyDbSystemDetails
	}{
		"DATAGUARD",
		(MarshalTypeLaunchStandbyDbSystemDetails)(m),
	}

	return json.Marshal(&s)
}

// LaunchStandbyDbSystemDetailsLicenseModelEnum Enum with underlying type: string
type LaunchStandbyDbSystemDetailsLicenseModelEnum string

// Set of constants representing the allowable values for LaunchStandbyDbSystemDetailsLicenseModelEnum
const (
	LaunchStandbyDbSystemDetailsLicenseModelLicenseIncluded     LaunchStandbyDbSystemDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	LaunchStandbyDbSystemDetailsLicenseModelBringYourOwnLicense LaunchStandbyDbSystemDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingLaunchStandbyDbSystemDetailsLicenseModelEnum = map[string]LaunchStandbyDbSystemDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       LaunchStandbyDbSystemDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": LaunchStandbyDbSystemDetailsLicenseModelBringYourOwnLicense,
}

var mappingLaunchStandbyDbSystemDetailsLicenseModelEnumLowerCase = map[string]LaunchStandbyDbSystemDetailsLicenseModelEnum{
	"license_included":       LaunchStandbyDbSystemDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": LaunchStandbyDbSystemDetailsLicenseModelBringYourOwnLicense,
}

// GetLaunchStandbyDbSystemDetailsLicenseModelEnumValues Enumerates the set of values for LaunchStandbyDbSystemDetailsLicenseModelEnum
func GetLaunchStandbyDbSystemDetailsLicenseModelEnumValues() []LaunchStandbyDbSystemDetailsLicenseModelEnum {
	values := make([]LaunchStandbyDbSystemDetailsLicenseModelEnum, 0)
	for _, v := range mappingLaunchStandbyDbSystemDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetLaunchStandbyDbSystemDetailsLicenseModelEnumStringValues Enumerates the set of values in String for LaunchStandbyDbSystemDetailsLicenseModelEnum
func GetLaunchStandbyDbSystemDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingLaunchStandbyDbSystemDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLaunchStandbyDbSystemDetailsLicenseModelEnum(val string) (LaunchStandbyDbSystemDetailsLicenseModelEnum, bool) {
	enum, ok := mappingLaunchStandbyDbSystemDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
