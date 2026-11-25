// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDistributedDatabaseDetails Details required for creation of the Globally distributed database.
type CreateDistributedDatabaseDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Globally distributed database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Oracle Database version for the shards and catalog used in Globally distributed database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Unique name prefix for the Globally distributed databases. Only alpha-numeric values are allowed. First character
	// has to be a letter followed by any combination of letter and number.
	Prefix *string `mandatory:"true" json:"prefix"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoints associated with the Globally distributed database.
	PrivateEndpointIds []string `mandatory:"true" json:"privateEndpointIds"`

	// The character set for the database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// The listener port number for the Globally distributed database. The listener port number
	// has to be unique for a customer tenancy across all distributed databases. Same port number should
	// not be re-used for any other distributed database.
	ListenerPort *int `mandatory:"true" json:"listenerPort"`

	// The ons local port number for the Globally distributed database. The onsPortLocal has to be
	// unique for a customer tenancy across all distributed databases. Same port number should not be
	// re-used for any other distributed database.
	OnsPortLocal *int `mandatory:"true" json:"onsPortLocal"`

	// The ons remote port number for the Globally distributed database. The onsPortRemote has to be
	// unique for a customer tenancy across all distributed databases. Same port number should not be
	// re-used for any other distributed database.
	OnsPortRemote *int `mandatory:"true" json:"onsPortRemote"`

	// Collection of catalog details for the Globally distributed database.
	CatalogDetails []CreateDistributedDatabaseCatalogDetails `mandatory:"true" json:"catalogDetails"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed database.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The TLS listener port number for the Globally distributed database. The TLS listener port number
	// has to be unique for a customer tenancy across all distributed databases. Same port number should
	// not be re-used for any other distributed database. For EXADB based distributed databases,
	// tls is not supported hence the listenerPortTls is not needed to be provided in create payload.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// TCP SCAN listener port for new clusters to be created.
	// Applicable only when shard and catalog source types are XS_NEW_VAULT_AND_CLUSTER or XS_NEW_CLUSTER.
	// Defaults to 1521 when not provided.
	ScanListenerPort *int `mandatory:"false" json:"scanListenerPort"`

	// Count of chunks associated with system raft clusters or system data guard databases.
	SystemChunkCount *int `mandatory:"false" json:"systemChunkCount"`

	// Number of replication units associated with system raft clusters.
	SystemRaftReplicationUnitCount *int `mandatory:"false" json:"systemRaftReplicationUnitCount"`

	// Collection of composite raft shards.
	CompositeRaftShardSpaces []CreateCompositeRaftShardSpaceDetails `mandatory:"false" json:"compositeRaftShardSpaces"`

	// Collection of composite data guard shard spaces.
	CompositeDataGuardShardSpaces []CreateCompositeDataGuardShardSpaceDetails `mandatory:"false" json:"compositeDataGuardShardSpaces"`

	// Collection of system raft clusters.
	SystemRaftClusters []CreateSystemRaftClusterDetails `mandatory:"false" json:"systemRaftClusters"`

	SystemDataGuardDatabases *CreateSystemDataGuardDatabaseDetails `mandatory:"false" json:"systemDataGuardDatabases"`

	// Collection of user defined shard spaces.
	UserShardSpaces []CreateUserShardSpaceDetails `mandatory:"false" json:"userShardSpaces"`

	// The SSH public key for Global service manager instances.
	GsmSshPublicKey *string `mandatory:"false" json:"gsmSshPublicKey"`

	DbBackupConfig *DistributedDbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// The list of network security group (NSG) details to be associated with the distributed database.
	VcnNsgIds []VcnNsgIdsDetails `mandatory:"false" json:"vcnNsgIds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDistributedDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
