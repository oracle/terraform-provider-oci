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

// DistributedAutonomousDatabase Globally distributed autonomous database.
type DistributedAutonomousDatabase struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Globally distributed autonomous database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the Globally distributed autonomous database was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Globally distributed autonomous database was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Oracle Database version for the shards and catalog used in Globally distributed autonomous database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Lifecycle states for the Globally distributed autonomous database.
	LifecycleState DistributedAutonomousDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Sharding methods for the Globally distributed autonomous database.
	ShardingMethod DistributedAutonomousDatabaseShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

	// The distributed autonomous database deployment type.
	DbDeploymentType DistributedAutonomousDatabaseDbDeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`

	// The lifecycleDetails for the Globally distributed autonomous database.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Unique name prefix for the Globally distributed autonomous databases. Only alpha-numeric values are allowed. First character
	// has to be a letter followed by any combination of letter and number.
	Prefix *string `mandatory:"true" json:"prefix"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
	PrivateEndpointIds []string `mandatory:"true" json:"privateEndpointIds"`

	// Possible workload types. Currently only OLTP workload type is supported.
	DbWorkloadType DistributedAutonomousDatabaseDbWorkloadTypeEnum `mandatory:"true" json:"dbWorkloadType"`

	// The character set for the database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// The listener port number for the Globally distributed autonomous database.
	ListenerPort *int `mandatory:"true" json:"listenerPort"`

	// Ons local port number for Globally distributed autonomous database.
	OnsPortLocal *int `mandatory:"true" json:"onsPortLocal"`

	// Ons remote port number for Globally distributed autonomous database.
	OnsPortRemote *int `mandatory:"true" json:"onsPortRemote"`

	// The Replication method for Globally distributed Autonomous database. Use RAFT for Raft based replication.
	// With RAFT replication, shards cannot have peers details set on them. In case shards need to
	// have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or
	// without peers), please set replicationMethod as DG or do not set any value for replicationMethod.
	ReplicationMethod DistributedAutonomousDatabaseReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`

	ConnectionStrings *DistributedAutonomousDatabaseConnectionString `mandatory:"false" json:"connectionStrings"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed autonomous database.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The TLS listener port number for Globally distributed autonomous database.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// Count of chunks associated with system raft clusters or system data guard databases.
	SystemChunkCount *int `mandatory:"false" json:"systemChunkCount"`

	// Number of replication units associated with system raft clusters.
	SystemRaftReplicationUnitCount *int `mandatory:"false" json:"systemRaftReplicationUnitCount"`

	LatestGsmImage *DistributedAutonomousDatabaseGsmImage `mandatory:"false" json:"latestGsmImage"`

	// Collection of composite raft shards.
	CompositeRaftShardSpaces []AutonomousCompositeRaftShardSpace `mandatory:"false" json:"compositeRaftShardSpaces"`

	// Collection of composite data guard shard spaces.
	CompositeDataGuardShardSpaces []AutonomousCompositeDataGuardShardSpace `mandatory:"false" json:"compositeDataGuardShardSpaces"`

	// Collection of system raft clusters.
	SystemRaftClusters []AutonomousSystemRaftCluster `mandatory:"false" json:"systemRaftClusters"`

	SystemDataGuardDatabases *AutonomousSystemDataGuardDatabase `mandatory:"false" json:"systemDataGuardDatabases"`

	// Collection of user defined shard spaces.
	UserShardSpaces []AutonomousUserShardSpace `mandatory:"false" json:"userShardSpaces"`

	// Catalog details associated with the distributed autonomous database.
	CatalogDetails []DistributedAutonomousDatabaseCatalog `mandatory:"false" json:"catalogDetails"`

	// Global Service Manager (GSM) instances associated with the distributed autonomous database.
	GsmDetails []DistributedAutonomousDatabaseGsm `mandatory:"false" json:"gsmDetails"`

	// Global Database Services Control(GDS CTL) instances associated with the distributed autonomous database.
	GdsControlNodeDetails []DistributedAutonomousDatabaseGdsControlNode `mandatory:"false" json:"gdsControlNodeDetails"`

	DbBackupConfig *DistributedAutonomousDbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	AutoResourceManagementConfig *AutoResourceManagementConfigurationDetails `mandatory:"false" json:"autoResourceManagementConfig"`

	// The list of network security group (NSG) details associated with the distributed autonomous database.
	VcnNsgIds []VcnNsgIdsDetails `mandatory:"false" json:"vcnNsgIds"`

	Metadata *DistributedAutonomousDbMetadata `mandatory:"false" json:"metadata"`

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

func (m DistributedAutonomousDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedAutonomousDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedAutonomousDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDistributedAutonomousDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedAutonomousDatabaseShardingMethodEnum(string(m.ShardingMethod)); !ok && m.ShardingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShardingMethod: %s. Supported values are: %s.", m.ShardingMethod, strings.Join(GetDistributedAutonomousDatabaseShardingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedAutonomousDatabaseDbDeploymentTypeEnum(string(m.DbDeploymentType)); !ok && m.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", m.DbDeploymentType, strings.Join(GetDistributedAutonomousDatabaseDbDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedAutonomousDatabaseDbWorkloadTypeEnum(string(m.DbWorkloadType)); !ok && m.DbWorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkloadType: %s. Supported values are: %s.", m.DbWorkloadType, strings.Join(GetDistributedAutonomousDatabaseDbWorkloadTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDistributedAutonomousDatabaseReplicationMethodEnum(string(m.ReplicationMethod)); !ok && m.ReplicationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationMethod: %s. Supported values are: %s.", m.ReplicationMethod, strings.Join(GetDistributedAutonomousDatabaseReplicationMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDatabaseLifecycleStateEnum Enum with underlying type: string
type DistributedAutonomousDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseLifecycleStateEnum
const (
	DistributedAutonomousDatabaseLifecycleStateActive         DistributedAutonomousDatabaseLifecycleStateEnum = "ACTIVE"
	DistributedAutonomousDatabaseLifecycleStateFailed         DistributedAutonomousDatabaseLifecycleStateEnum = "FAILED"
	DistributedAutonomousDatabaseLifecycleStateNeedsAttention DistributedAutonomousDatabaseLifecycleStateEnum = "NEEDS_ATTENTION"
	DistributedAutonomousDatabaseLifecycleStateInactive       DistributedAutonomousDatabaseLifecycleStateEnum = "INACTIVE"
	DistributedAutonomousDatabaseLifecycleStateDeleting       DistributedAutonomousDatabaseLifecycleStateEnum = "DELETING"
	DistributedAutonomousDatabaseLifecycleStateDeleted        DistributedAutonomousDatabaseLifecycleStateEnum = "DELETED"
	DistributedAutonomousDatabaseLifecycleStateUpdating       DistributedAutonomousDatabaseLifecycleStateEnum = "UPDATING"
	DistributedAutonomousDatabaseLifecycleStateCreating       DistributedAutonomousDatabaseLifecycleStateEnum = "CREATING"
)

var mappingDistributedAutonomousDatabaseLifecycleStateEnum = map[string]DistributedAutonomousDatabaseLifecycleStateEnum{
	"ACTIVE":          DistributedAutonomousDatabaseLifecycleStateActive,
	"FAILED":          DistributedAutonomousDatabaseLifecycleStateFailed,
	"NEEDS_ATTENTION": DistributedAutonomousDatabaseLifecycleStateNeedsAttention,
	"INACTIVE":        DistributedAutonomousDatabaseLifecycleStateInactive,
	"DELETING":        DistributedAutonomousDatabaseLifecycleStateDeleting,
	"DELETED":         DistributedAutonomousDatabaseLifecycleStateDeleted,
	"UPDATING":        DistributedAutonomousDatabaseLifecycleStateUpdating,
	"CREATING":        DistributedAutonomousDatabaseLifecycleStateCreating,
}

var mappingDistributedAutonomousDatabaseLifecycleStateEnumLowerCase = map[string]DistributedAutonomousDatabaseLifecycleStateEnum{
	"active":          DistributedAutonomousDatabaseLifecycleStateActive,
	"failed":          DistributedAutonomousDatabaseLifecycleStateFailed,
	"needs_attention": DistributedAutonomousDatabaseLifecycleStateNeedsAttention,
	"inactive":        DistributedAutonomousDatabaseLifecycleStateInactive,
	"deleting":        DistributedAutonomousDatabaseLifecycleStateDeleting,
	"deleted":         DistributedAutonomousDatabaseLifecycleStateDeleted,
	"updating":        DistributedAutonomousDatabaseLifecycleStateUpdating,
	"creating":        DistributedAutonomousDatabaseLifecycleStateCreating,
}

// GetDistributedAutonomousDatabaseLifecycleStateEnumValues Enumerates the set of values for DistributedAutonomousDatabaseLifecycleStateEnum
func GetDistributedAutonomousDatabaseLifecycleStateEnumValues() []DistributedAutonomousDatabaseLifecycleStateEnum {
	values := make([]DistributedAutonomousDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseLifecycleStateEnum
func GetDistributedAutonomousDatabaseLifecycleStateEnumStringValues() []string {
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

// GetMappingDistributedAutonomousDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseLifecycleStateEnum(val string) (DistributedAutonomousDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedAutonomousDatabaseShardingMethodEnum Enum with underlying type: string
type DistributedAutonomousDatabaseShardingMethodEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseShardingMethodEnum
const (
	DistributedAutonomousDatabaseShardingMethodUser      DistributedAutonomousDatabaseShardingMethodEnum = "USER"
	DistributedAutonomousDatabaseShardingMethodSystem    DistributedAutonomousDatabaseShardingMethodEnum = "SYSTEM"
	DistributedAutonomousDatabaseShardingMethodComposite DistributedAutonomousDatabaseShardingMethodEnum = "COMPOSITE"
)

var mappingDistributedAutonomousDatabaseShardingMethodEnum = map[string]DistributedAutonomousDatabaseShardingMethodEnum{
	"USER":      DistributedAutonomousDatabaseShardingMethodUser,
	"SYSTEM":    DistributedAutonomousDatabaseShardingMethodSystem,
	"COMPOSITE": DistributedAutonomousDatabaseShardingMethodComposite,
}

var mappingDistributedAutonomousDatabaseShardingMethodEnumLowerCase = map[string]DistributedAutonomousDatabaseShardingMethodEnum{
	"user":      DistributedAutonomousDatabaseShardingMethodUser,
	"system":    DistributedAutonomousDatabaseShardingMethodSystem,
	"composite": DistributedAutonomousDatabaseShardingMethodComposite,
}

// GetDistributedAutonomousDatabaseShardingMethodEnumValues Enumerates the set of values for DistributedAutonomousDatabaseShardingMethodEnum
func GetDistributedAutonomousDatabaseShardingMethodEnumValues() []DistributedAutonomousDatabaseShardingMethodEnum {
	values := make([]DistributedAutonomousDatabaseShardingMethodEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseShardingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseShardingMethodEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseShardingMethodEnum
func GetDistributedAutonomousDatabaseShardingMethodEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
		"COMPOSITE",
	}
}

// GetMappingDistributedAutonomousDatabaseShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseShardingMethodEnum(val string) (DistributedAutonomousDatabaseShardingMethodEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedAutonomousDatabaseReplicationMethodEnum Enum with underlying type: string
type DistributedAutonomousDatabaseReplicationMethodEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseReplicationMethodEnum
const (
	DistributedAutonomousDatabaseReplicationMethodRaft DistributedAutonomousDatabaseReplicationMethodEnum = "RAFT"
	DistributedAutonomousDatabaseReplicationMethodDg   DistributedAutonomousDatabaseReplicationMethodEnum = "DG"
)

var mappingDistributedAutonomousDatabaseReplicationMethodEnum = map[string]DistributedAutonomousDatabaseReplicationMethodEnum{
	"RAFT": DistributedAutonomousDatabaseReplicationMethodRaft,
	"DG":   DistributedAutonomousDatabaseReplicationMethodDg,
}

var mappingDistributedAutonomousDatabaseReplicationMethodEnumLowerCase = map[string]DistributedAutonomousDatabaseReplicationMethodEnum{
	"raft": DistributedAutonomousDatabaseReplicationMethodRaft,
	"dg":   DistributedAutonomousDatabaseReplicationMethodDg,
}

// GetDistributedAutonomousDatabaseReplicationMethodEnumValues Enumerates the set of values for DistributedAutonomousDatabaseReplicationMethodEnum
func GetDistributedAutonomousDatabaseReplicationMethodEnumValues() []DistributedAutonomousDatabaseReplicationMethodEnum {
	values := make([]DistributedAutonomousDatabaseReplicationMethodEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseReplicationMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseReplicationMethodEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseReplicationMethodEnum
func GetDistributedAutonomousDatabaseReplicationMethodEnumStringValues() []string {
	return []string{
		"RAFT",
		"DG",
	}
}

// GetMappingDistributedAutonomousDatabaseReplicationMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseReplicationMethodEnum(val string) (DistributedAutonomousDatabaseReplicationMethodEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseReplicationMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedAutonomousDatabaseDbDeploymentTypeEnum Enum with underlying type: string
type DistributedAutonomousDatabaseDbDeploymentTypeEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseDbDeploymentTypeEnum
const (
	DistributedAutonomousDatabaseDbDeploymentTypeAdbD DistributedAutonomousDatabaseDbDeploymentTypeEnum = "ADB_D"
)

var mappingDistributedAutonomousDatabaseDbDeploymentTypeEnum = map[string]DistributedAutonomousDatabaseDbDeploymentTypeEnum{
	"ADB_D": DistributedAutonomousDatabaseDbDeploymentTypeAdbD,
}

var mappingDistributedAutonomousDatabaseDbDeploymentTypeEnumLowerCase = map[string]DistributedAutonomousDatabaseDbDeploymentTypeEnum{
	"adb_d": DistributedAutonomousDatabaseDbDeploymentTypeAdbD,
}

// GetDistributedAutonomousDatabaseDbDeploymentTypeEnumValues Enumerates the set of values for DistributedAutonomousDatabaseDbDeploymentTypeEnum
func GetDistributedAutonomousDatabaseDbDeploymentTypeEnumValues() []DistributedAutonomousDatabaseDbDeploymentTypeEnum {
	values := make([]DistributedAutonomousDatabaseDbDeploymentTypeEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseDbDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseDbDeploymentTypeEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseDbDeploymentTypeEnum
func GetDistributedAutonomousDatabaseDbDeploymentTypeEnumStringValues() []string {
	return []string{
		"ADB_D",
	}
}

// GetMappingDistributedAutonomousDatabaseDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseDbDeploymentTypeEnum(val string) (DistributedAutonomousDatabaseDbDeploymentTypeEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedAutonomousDatabaseDbWorkloadTypeEnum Enum with underlying type: string
type DistributedAutonomousDatabaseDbWorkloadTypeEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseDbWorkloadTypeEnum
const (
	DistributedAutonomousDatabaseDbWorkloadTypeOltp DistributedAutonomousDatabaseDbWorkloadTypeEnum = "OLTP"
	DistributedAutonomousDatabaseDbWorkloadTypeDw   DistributedAutonomousDatabaseDbWorkloadTypeEnum = "DW"
)

var mappingDistributedAutonomousDatabaseDbWorkloadTypeEnum = map[string]DistributedAutonomousDatabaseDbWorkloadTypeEnum{
	"OLTP": DistributedAutonomousDatabaseDbWorkloadTypeOltp,
	"DW":   DistributedAutonomousDatabaseDbWorkloadTypeDw,
}

var mappingDistributedAutonomousDatabaseDbWorkloadTypeEnumLowerCase = map[string]DistributedAutonomousDatabaseDbWorkloadTypeEnum{
	"oltp": DistributedAutonomousDatabaseDbWorkloadTypeOltp,
	"dw":   DistributedAutonomousDatabaseDbWorkloadTypeDw,
}

// GetDistributedAutonomousDatabaseDbWorkloadTypeEnumValues Enumerates the set of values for DistributedAutonomousDatabaseDbWorkloadTypeEnum
func GetDistributedAutonomousDatabaseDbWorkloadTypeEnumValues() []DistributedAutonomousDatabaseDbWorkloadTypeEnum {
	values := make([]DistributedAutonomousDatabaseDbWorkloadTypeEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseDbWorkloadTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseDbWorkloadTypeEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseDbWorkloadTypeEnum
func GetDistributedAutonomousDatabaseDbWorkloadTypeEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
	}
}

// GetMappingDistributedAutonomousDatabaseDbWorkloadTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseDbWorkloadTypeEnum(val string) (DistributedAutonomousDatabaseDbWorkloadTypeEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseDbWorkloadTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
