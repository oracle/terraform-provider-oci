// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDataGuardAssociationWithNewDbSystemDetails The configuration details for creating a Data Guard association for a virtual machine DB system database. For this type of DB system database, the `creationType` should be `NewDbSystem`. A new DB system will be launched to create the standby database.
// To create a Data Guard association for a database in a bare metal or Exadata DB system, use the CreateDataGuardAssociationToExistingDbSystemDetails subtype instead.
type CreateDataGuardAssociationWithNewDbSystemDetails struct {

	// A strong password for the `SYS`, `SYSTEM`, and `PDB Admin` users to apply during standby creation.
	// The password must contain no fewer than nine characters and include:
	// * At least two uppercase characters.
	// * At least two lowercase characters.
	// * At least two numeric characters.
	// * At least two special characters. Valid special characters include "_", "#", and "-" only.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword"`

	// The database software image OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// True if active Data Guard is enabled.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`

	// Specifies the `DB_UNIQUE_NAME` of the peer database to be created.
	PeerDbUniqueName *string `mandatory:"false" json:"peerDbUniqueName"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	PeerSidPrefix *string `mandatory:"false" json:"peerSidPrefix"`

	// The user-friendly name of the DB system that will contain the the standby database. The display name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The name of the availability domain that the standby database DB system will be located in. For example- "Uocm:PHX-AD-1".
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The virtual machine DB system shape to launch for the standby database in the Data Guard association. The shape determines the number of CPU cores and the amount of memory available for the DB system.
	// Only virtual machine shapes are valid options. If you do not supply this parameter, the default shape is the shape of the primary DB system.
	// To get a list of all shapes, use the ListDbSystemShapes operation.
	Shape *string `mandatory:"false" json:"shape"`

	// The number of OCPU cores available for AMD-based virtual machine DB systems.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of nodes to launch for the DB system of the standby in the Data Guard association. For a 2-node RAC virtual machine DB system, specify either 1 or 2. If you do not supply this parameter, the default is the node count of the primary DB system.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The OCID of the subnet the DB system is associated with.
	// **Subnet Restrictions:**
	// - For 1- and 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.16.16/28
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The hostname for the DB node.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The time zone of the dataguard standby DB system. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

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

	// The IPv4 address from the provided OCI subnet which needs to be assigned to the VNIC. If not provided, it will
	// be auto-assigned with an available IPv4 address from the subnet.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	DbSystemFreeformTags map[string]string `mandatory:"false" json:"dbSystemFreeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DbSystemDefinedTags map[string]map[string]interface{} `mandatory:"false" json:"dbSystemDefinedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	DatabaseFreeformTags map[string]string `mandatory:"false" json:"databaseFreeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DatabaseDefinedTags map[string]map[string]interface{} `mandatory:"false" json:"databaseDefinedTags"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	// The block storage volume performance level. Valid values are `BALANCED` and `HIGH_PERFORMANCE`. See Block Volume Performance (https://docs.cloud.oracle.com/Content/Block/Concepts/blockvolumeperformance.htm) for more information.
	StorageVolumePerformanceMode CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum `mandatory:"false" json:"storageVolumePerformanceMode,omitempty"`

	// The Oracle license model that applies to all the databases on the dataguard standby DB system. The default is LICENSE_INCLUDED.
	LicenseModel CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The protection mode to set up between the primary and standby databases. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only protection mode currently supported by the Database service is MAXIMUM_PERFORMANCE.
	ProtectionMode CreateDataGuardAssociationDetailsProtectionModeEnum `mandatory:"true" json:"protectionMode"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database service is ASYNC.
	TransportType CreateDataGuardAssociationDetailsTransportTypeEnum `mandatory:"true" json:"transportType"`
}

//GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

//GetDatabaseAdminPassword returns DatabaseAdminPassword
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetDatabaseAdminPassword() *string {
	return m.DatabaseAdminPassword
}

//GetProtectionMode returns ProtectionMode
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetProtectionMode() CreateDataGuardAssociationDetailsProtectionModeEnum {
	return m.ProtectionMode
}

//GetTransportType returns TransportType
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetTransportType() CreateDataGuardAssociationDetailsTransportTypeEnum {
	return m.TransportType
}

//GetIsActiveDataGuardEnabled returns IsActiveDataGuardEnabled
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetIsActiveDataGuardEnabled() *bool {
	return m.IsActiveDataGuardEnabled
}

//GetPeerDbUniqueName returns PeerDbUniqueName
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetPeerDbUniqueName() *string {
	return m.PeerDbUniqueName
}

//GetPeerSidPrefix returns PeerSidPrefix
func (m CreateDataGuardAssociationWithNewDbSystemDetails) GetPeerSidPrefix() *string {
	return m.PeerSidPrefix
}

func (m CreateDataGuardAssociationWithNewDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataGuardAssociationWithNewDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum(string(m.StorageVolumePerformanceMode)); !ok && m.StorageVolumePerformanceMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageVolumePerformanceMode: %s. Supported values are: %s.", m.StorageVolumePerformanceMode, strings.Join(GetCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateDataGuardAssociationDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetCreateDataGuardAssociationDetailsProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDataGuardAssociationDetailsTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetCreateDataGuardAssociationDetailsTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDataGuardAssociationWithNewDbSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataGuardAssociationWithNewDbSystemDetails CreateDataGuardAssociationWithNewDbSystemDetails
	s := struct {
		DiscriminatorParam string `json:"creationType"`
		MarshalTypeCreateDataGuardAssociationWithNewDbSystemDetails
	}{
		"NewDbSystem",
		(MarshalTypeCreateDataGuardAssociationWithNewDbSystemDetails)(m),
	}

	return json.Marshal(&s)
}

// CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum Enum with underlying type: string
type CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum string

// Set of constants representing the allowable values for CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum
const (
	CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeBalanced        CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum = "BALANCED"
	CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeHighPerformance CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum = "HIGH_PERFORMANCE"
)

var mappingCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum = map[string]CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum{
	"BALANCED":         CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeBalanced,
	"HIGH_PERFORMANCE": CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeHighPerformance,
}

var mappingCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnumLowerCase = map[string]CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum{
	"balanced":         CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeBalanced,
	"high_performance": CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeHighPerformance,
}

// GetCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnumValues Enumerates the set of values for CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum
func GetCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnumValues() []CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum {
	values := make([]CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum, 0)
	for _, v := range mappingCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnumStringValues Enumerates the set of values in String for CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum
func GetCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnumStringValues() []string {
	return []string{
		"BALANCED",
		"HIGH_PERFORMANCE",
	}
}

// GetMappingCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum(val string) (CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum, bool) {
	enum, ok := mappingCreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum Enum with underlying type: string
type CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum
const (
	CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelLicenseIncluded     CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelBringYourOwnLicense CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum = map[string]CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnumLowerCase = map[string]CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum{
	"license_included":       CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnumValues Enumerates the set of values for CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum
func GetCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnumValues() []CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum {
	values := make([]CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum
func GetCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum(val string) (CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
