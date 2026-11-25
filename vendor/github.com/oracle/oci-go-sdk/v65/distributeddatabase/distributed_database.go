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

// DistributedDatabase Globally distributed database.
type DistributedDatabase struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Globally distributed database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the Globally distributed database was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Globally distributed database was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Oracle Database version for the shards and catalog used in Globally distributed database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Lifecycle states for the Globally distributed database.
	LifecycleState DistributedDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Sharding methods for the Globally distributed database.
	ShardingMethod DistributedDatabaseShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

	// The distributed database deployment type.
	DbDeploymentType DistributedDatabaseDbDeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`

	// The lifecycleDetails for the Globally distributed database.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Unique name prefix for the Globally distributed databases. Only alpha-numeric values are allowed. First character
	// has to be a letter followed by any combination of letter and number.
	Prefix *string `mandatory:"true" json:"prefix"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
	PrivateEndpointIds []string `mandatory:"true" json:"privateEndpointIds"`

	// The character set for the database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// The Global service manager listener port number for the Globally distributed database.
	ListenerPort *int `mandatory:"true" json:"listenerPort"`

	// Ons local port number.
	OnsPortLocal *int `mandatory:"true" json:"onsPortLocal"`

	// Ons remote port number.
	OnsPortRemote *int `mandatory:"true" json:"onsPortRemote"`

	// The Replication method for Globally distributed database. Use RAFT for Raft based replication.
	// With RAFT replication, shards cannot have peers details set on them. In case shards need to
	// have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or
	// without peers), please set replicationMethod as DG or do not set any value for replicationMethod.
	ReplicationMethod DistributedDatabaseReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`

	ConnectionStrings *DistributedDbConnectionString `mandatory:"false" json:"connectionStrings"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed database.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	LatestGsmImageDetails *DistributedDbGsmImage `mandatory:"false" json:"latestGsmImageDetails"`

	// The TLS listener port number for Globally distributed database.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// The TCP SCAN listener port for database clusters with source type XS_NEW_VAULT_AND_CLUSTER or XS_NEW_CLUSTER.
	ScanListenerPort *int `mandatory:"false" json:"scanListenerPort"`

	// Count of chunks associated with system raft clusters or system data guard databases.
	SystemChunkCount *int `mandatory:"false" json:"systemChunkCount"`

	// Number of replication units associated with system raft clusters.
	SystemRaftReplicationUnitCount *int `mandatory:"false" json:"systemRaftReplicationUnitCount"`

	// Collection of composite raft shards.
	CompositeRaftShardSpaces []CompositeRaftShardSpace `mandatory:"false" json:"compositeRaftShardSpaces"`

	// Collection of composite data guard shard spaces.
	CompositeDataGuardShardSpaces []CompositeDataGuardShardSpace `mandatory:"false" json:"compositeDataGuardShardSpaces"`

	// Collection of system raft clusters.
	SystemRaftClusters []SystemRaftCluster `mandatory:"false" json:"systemRaftClusters"`

	SystemDataGuardDatabases *SystemDataGuardDatabase `mandatory:"false" json:"systemDataGuardDatabases"`

	// Collection of user defined shard spaces.
	UserShardSpaces []UserShardSpace `mandatory:"false" json:"userShardSpaces"`

	// Catalog details associated with the distributed database.
	CatalogDetails []DistributedDatabaseCatalog `mandatory:"false" json:"catalogDetails"`

	// Global Service Manager (GSM) instances associated with the distributed database.
	GsmDetails []DistributedDatabaseGsm `mandatory:"false" json:"gsmDetails"`

	// Global Database Services Control(GDS CTL) instances associated with the distributed database.
	GdsControlNodeDetails []DistributedDatabaseGdsControlNode `mandatory:"false" json:"gdsControlNodeDetails"`

	DbBackupConfig *DistributedDbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	AutoResourceManagementConfig *AutoResourceManagementConfigurationDetails `mandatory:"false" json:"autoResourceManagementConfig"`

	// The SSH public key for Global service manager instances.
	GsmSshPublicKey *string `mandatory:"false" json:"gsmSshPublicKey"`

	// The list of network security group (NSG) details associated with the distributed database.
	VcnNsgIds []VcnNsgIdsDetails `mandatory:"false" json:"vcnNsgIds"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DistributedDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDistributedDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDatabaseShardingMethodEnum(string(m.ShardingMethod)); !ok && m.ShardingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShardingMethod: %s. Supported values are: %s.", m.ShardingMethod, strings.Join(GetDistributedDatabaseShardingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDatabaseDbDeploymentTypeEnum(string(m.DbDeploymentType)); !ok && m.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", m.DbDeploymentType, strings.Join(GetDistributedDatabaseDbDeploymentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDistributedDatabaseReplicationMethodEnum(string(m.ReplicationMethod)); !ok && m.ReplicationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationMethod: %s. Supported values are: %s.", m.ReplicationMethod, strings.Join(GetDistributedDatabaseReplicationMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseLifecycleStateEnum Enum with underlying type: string
type DistributedDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for DistributedDatabaseLifecycleStateEnum
const (
	DistributedDatabaseLifecycleStateActive         DistributedDatabaseLifecycleStateEnum = "ACTIVE"
	DistributedDatabaseLifecycleStateFailed         DistributedDatabaseLifecycleStateEnum = "FAILED"
	DistributedDatabaseLifecycleStateNeedsAttention DistributedDatabaseLifecycleStateEnum = "NEEDS_ATTENTION"
	DistributedDatabaseLifecycleStateInactive       DistributedDatabaseLifecycleStateEnum = "INACTIVE"
	DistributedDatabaseLifecycleStateDeleting       DistributedDatabaseLifecycleStateEnum = "DELETING"
	DistributedDatabaseLifecycleStateDeleted        DistributedDatabaseLifecycleStateEnum = "DELETED"
	DistributedDatabaseLifecycleStateUpdating       DistributedDatabaseLifecycleStateEnum = "UPDATING"
	DistributedDatabaseLifecycleStateCreating       DistributedDatabaseLifecycleStateEnum = "CREATING"
)

var mappingDistributedDatabaseLifecycleStateEnum = map[string]DistributedDatabaseLifecycleStateEnum{
	"ACTIVE":          DistributedDatabaseLifecycleStateActive,
	"FAILED":          DistributedDatabaseLifecycleStateFailed,
	"NEEDS_ATTENTION": DistributedDatabaseLifecycleStateNeedsAttention,
	"INACTIVE":        DistributedDatabaseLifecycleStateInactive,
	"DELETING":        DistributedDatabaseLifecycleStateDeleting,
	"DELETED":         DistributedDatabaseLifecycleStateDeleted,
	"UPDATING":        DistributedDatabaseLifecycleStateUpdating,
	"CREATING":        DistributedDatabaseLifecycleStateCreating,
}

var mappingDistributedDatabaseLifecycleStateEnumLowerCase = map[string]DistributedDatabaseLifecycleStateEnum{
	"active":          DistributedDatabaseLifecycleStateActive,
	"failed":          DistributedDatabaseLifecycleStateFailed,
	"needs_attention": DistributedDatabaseLifecycleStateNeedsAttention,
	"inactive":        DistributedDatabaseLifecycleStateInactive,
	"deleting":        DistributedDatabaseLifecycleStateDeleting,
	"deleted":         DistributedDatabaseLifecycleStateDeleted,
	"updating":        DistributedDatabaseLifecycleStateUpdating,
	"creating":        DistributedDatabaseLifecycleStateCreating,
}

// GetDistributedDatabaseLifecycleStateEnumValues Enumerates the set of values for DistributedDatabaseLifecycleStateEnum
func GetDistributedDatabaseLifecycleStateEnumValues() []DistributedDatabaseLifecycleStateEnum {
	values := make([]DistributedDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingDistributedDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for DistributedDatabaseLifecycleStateEnum
func GetDistributedDatabaseLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"FAILED",
		"NEEDS_ATTENTION",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
	}
}

// GetMappingDistributedDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseLifecycleStateEnum(val string) (DistributedDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingDistributedDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDatabaseShardingMethodEnum Enum with underlying type: string
type DistributedDatabaseShardingMethodEnum string

// Set of constants representing the allowable values for DistributedDatabaseShardingMethodEnum
const (
	DistributedDatabaseShardingMethodUser      DistributedDatabaseShardingMethodEnum = "USER"
	DistributedDatabaseShardingMethodSystem    DistributedDatabaseShardingMethodEnum = "SYSTEM"
	DistributedDatabaseShardingMethodComposite DistributedDatabaseShardingMethodEnum = "COMPOSITE"
)

var mappingDistributedDatabaseShardingMethodEnum = map[string]DistributedDatabaseShardingMethodEnum{
	"USER":      DistributedDatabaseShardingMethodUser,
	"SYSTEM":    DistributedDatabaseShardingMethodSystem,
	"COMPOSITE": DistributedDatabaseShardingMethodComposite,
}

var mappingDistributedDatabaseShardingMethodEnumLowerCase = map[string]DistributedDatabaseShardingMethodEnum{
	"user":      DistributedDatabaseShardingMethodUser,
	"system":    DistributedDatabaseShardingMethodSystem,
	"composite": DistributedDatabaseShardingMethodComposite,
}

// GetDistributedDatabaseShardingMethodEnumValues Enumerates the set of values for DistributedDatabaseShardingMethodEnum
func GetDistributedDatabaseShardingMethodEnumValues() []DistributedDatabaseShardingMethodEnum {
	values := make([]DistributedDatabaseShardingMethodEnum, 0)
	for _, v := range mappingDistributedDatabaseShardingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseShardingMethodEnumStringValues Enumerates the set of values in String for DistributedDatabaseShardingMethodEnum
func GetDistributedDatabaseShardingMethodEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
		"COMPOSITE",
	}
}

// GetMappingDistributedDatabaseShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseShardingMethodEnum(val string) (DistributedDatabaseShardingMethodEnum, bool) {
	enum, ok := mappingDistributedDatabaseShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDatabaseReplicationMethodEnum Enum with underlying type: string
type DistributedDatabaseReplicationMethodEnum string

// Set of constants representing the allowable values for DistributedDatabaseReplicationMethodEnum
const (
	DistributedDatabaseReplicationMethodRaft DistributedDatabaseReplicationMethodEnum = "RAFT"
	DistributedDatabaseReplicationMethodDg   DistributedDatabaseReplicationMethodEnum = "DG"
)

var mappingDistributedDatabaseReplicationMethodEnum = map[string]DistributedDatabaseReplicationMethodEnum{
	"RAFT": DistributedDatabaseReplicationMethodRaft,
	"DG":   DistributedDatabaseReplicationMethodDg,
}

var mappingDistributedDatabaseReplicationMethodEnumLowerCase = map[string]DistributedDatabaseReplicationMethodEnum{
	"raft": DistributedDatabaseReplicationMethodRaft,
	"dg":   DistributedDatabaseReplicationMethodDg,
}

// GetDistributedDatabaseReplicationMethodEnumValues Enumerates the set of values for DistributedDatabaseReplicationMethodEnum
func GetDistributedDatabaseReplicationMethodEnumValues() []DistributedDatabaseReplicationMethodEnum {
	values := make([]DistributedDatabaseReplicationMethodEnum, 0)
	for _, v := range mappingDistributedDatabaseReplicationMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseReplicationMethodEnumStringValues Enumerates the set of values in String for DistributedDatabaseReplicationMethodEnum
func GetDistributedDatabaseReplicationMethodEnumStringValues() []string {
	return []string{
		"RAFT",
		"DG",
	}
}

// GetMappingDistributedDatabaseReplicationMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseReplicationMethodEnum(val string) (DistributedDatabaseReplicationMethodEnum, bool) {
	enum, ok := mappingDistributedDatabaseReplicationMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDatabaseDbDeploymentTypeEnum Enum with underlying type: string
type DistributedDatabaseDbDeploymentTypeEnum string

// Set of constants representing the allowable values for DistributedDatabaseDbDeploymentTypeEnum
const (
	DistributedDatabaseDbDeploymentTypeExadb DistributedDatabaseDbDeploymentTypeEnum = "EXADB"
)

var mappingDistributedDatabaseDbDeploymentTypeEnum = map[string]DistributedDatabaseDbDeploymentTypeEnum{
	"EXADB": DistributedDatabaseDbDeploymentTypeExadb,
}

var mappingDistributedDatabaseDbDeploymentTypeEnumLowerCase = map[string]DistributedDatabaseDbDeploymentTypeEnum{
	"exadb": DistributedDatabaseDbDeploymentTypeExadb,
}

// GetDistributedDatabaseDbDeploymentTypeEnumValues Enumerates the set of values for DistributedDatabaseDbDeploymentTypeEnum
func GetDistributedDatabaseDbDeploymentTypeEnumValues() []DistributedDatabaseDbDeploymentTypeEnum {
	values := make([]DistributedDatabaseDbDeploymentTypeEnum, 0)
	for _, v := range mappingDistributedDatabaseDbDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseDbDeploymentTypeEnumStringValues Enumerates the set of values in String for DistributedDatabaseDbDeploymentTypeEnum
func GetDistributedDatabaseDbDeploymentTypeEnumStringValues() []string {
	return []string{
		"EXADB",
	}
}

// GetMappingDistributedDatabaseDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseDbDeploymentTypeEnum(val string) (DistributedDatabaseDbDeploymentTypeEnum, bool) {
	enum, ok := mappingDistributedDatabaseDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
