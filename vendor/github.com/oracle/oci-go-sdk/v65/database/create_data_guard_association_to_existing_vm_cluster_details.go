// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDataGuardAssociationToExistingVmClusterDetails The configuration details for creating a Data Guard association for a ExaCC Vmcluster database. For these types of vm cluster databases, the `creationType` should be `ExistingVmCluster`. A standby database will be created in the VM cluster you specify.
type CreateDataGuardAssociationToExistingVmClusterDetails struct {

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM Cluster in which to create the standby database.
	// You must supply this value if creationType is `ExistingVmCluster`.
	PeerVmClusterId *string `mandatory:"false" json:"peerVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB home in which to create the standby database.
	// You must supply this value to create standby database with an existing DB home
	PeerDbHomeId *string `mandatory:"false" json:"peerDbHomeId"`

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

// GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m CreateDataGuardAssociationToExistingVmClusterDetails) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

// GetDatabaseAdminPassword returns DatabaseAdminPassword
func (m CreateDataGuardAssociationToExistingVmClusterDetails) GetDatabaseAdminPassword() *string {
	return m.DatabaseAdminPassword
}

// GetProtectionMode returns ProtectionMode
func (m CreateDataGuardAssociationToExistingVmClusterDetails) GetProtectionMode() CreateDataGuardAssociationDetailsProtectionModeEnum {
	return m.ProtectionMode
}

// GetTransportType returns TransportType
func (m CreateDataGuardAssociationToExistingVmClusterDetails) GetTransportType() CreateDataGuardAssociationDetailsTransportTypeEnum {
	return m.TransportType
}

// GetIsActiveDataGuardEnabled returns IsActiveDataGuardEnabled
func (m CreateDataGuardAssociationToExistingVmClusterDetails) GetIsActiveDataGuardEnabled() *bool {
	return m.IsActiveDataGuardEnabled
}

// GetPeerDbUniqueName returns PeerDbUniqueName
func (m CreateDataGuardAssociationToExistingVmClusterDetails) GetPeerDbUniqueName() *string {
	return m.PeerDbUniqueName
}

// GetPeerSidPrefix returns PeerSidPrefix
func (m CreateDataGuardAssociationToExistingVmClusterDetails) GetPeerSidPrefix() *string {
	return m.PeerSidPrefix
}

func (m CreateDataGuardAssociationToExistingVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataGuardAssociationToExistingVmClusterDetails) ValidateEnumValue() (bool, error) {
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
func (m CreateDataGuardAssociationToExistingVmClusterDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataGuardAssociationToExistingVmClusterDetails CreateDataGuardAssociationToExistingVmClusterDetails
	s := struct {
		DiscriminatorParam string `json:"creationType"`
		MarshalTypeCreateDataGuardAssociationToExistingVmClusterDetails
	}{
		"ExistingVmCluster",
		(MarshalTypeCreateDataGuardAssociationToExistingVmClusterDetails)(m),
	}

	return json.Marshal(&s)
}
