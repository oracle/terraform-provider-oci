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

	// The OCID of the subnet the DB system is associated with.
	// **Subnet Restrictions:**
	// - For 1- and 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.16.16/28
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// The hostname for the DB node.
	Hostname *string `mandatory:"false" json:"hostname"`

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
