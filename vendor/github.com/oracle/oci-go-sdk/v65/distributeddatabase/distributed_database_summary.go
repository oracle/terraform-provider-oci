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

// DistributedDatabaseSummary Globally distributed database.
type DistributedDatabaseSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed database compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Globally distributed database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the Globally distributed database was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Globally distributed database was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Oracle Database version for the shards and catalog used in Globally distributed database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Lifecycle state of sharded database.
	LifecycleState DistributedDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

	// The listener port number for the Globally distributed database.
	ListenerPort *int `mandatory:"true" json:"listenerPort"`

	// Ons local port number for the Globally distributed database.
	OnsPortLocal *int `mandatory:"true" json:"onsPortLocal"`

	// Ons remote port number for the Globally distributed database.
	OnsPortRemote *int `mandatory:"true" json:"onsPortRemote"`

	// Sharding methods for the Globally distributed database.
	ShardingMethod DistributedDatabaseSummaryShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

	// The distributed database deployment type.
	DbDeploymentType DistributedDatabaseSummaryDbDeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed database.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The TLS listener port number for the Globally distributed database.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// Count of chunks associated with system raft clusters or system data guard databases.
	SystemChunkCount *int `mandatory:"false" json:"systemChunkCount"`

	AutoResourceManagementConfig *AutoResourceManagementConfigurationDetails `mandatory:"false" json:"autoResourceManagementConfig"`

	// Number of replication units associated with system raft clusters.
	SystemRaftReplicationUnitCount *int `mandatory:"false" json:"systemRaftReplicationUnitCount"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// The Replication method for Globally distributed database. Use RAFT for Raft based replication.
	// With RAFT replication, shards cannot have peers details set on them. In case shards need to
	// have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or
	// without peers), please set replicationMethod as DG or do not set any value for replicationMethod.
	ReplicationMethod DistributedDatabaseSummaryReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`

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

func (m DistributedDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDistributedDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDatabaseSummaryShardingMethodEnum(string(m.ShardingMethod)); !ok && m.ShardingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShardingMethod: %s. Supported values are: %s.", m.ShardingMethod, strings.Join(GetDistributedDatabaseSummaryShardingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDatabaseSummaryDbDeploymentTypeEnum(string(m.DbDeploymentType)); !ok && m.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", m.DbDeploymentType, strings.Join(GetDistributedDatabaseSummaryDbDeploymentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDistributedDatabaseSummaryReplicationMethodEnum(string(m.ReplicationMethod)); !ok && m.ReplicationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationMethod: %s. Supported values are: %s.", m.ReplicationMethod, strings.Join(GetDistributedDatabaseSummaryReplicationMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseSummaryShardingMethodEnum Enum with underlying type: string
type DistributedDatabaseSummaryShardingMethodEnum string

// Set of constants representing the allowable values for DistributedDatabaseSummaryShardingMethodEnum
const (
	DistributedDatabaseSummaryShardingMethodUser      DistributedDatabaseSummaryShardingMethodEnum = "USER"
	DistributedDatabaseSummaryShardingMethodSystem    DistributedDatabaseSummaryShardingMethodEnum = "SYSTEM"
	DistributedDatabaseSummaryShardingMethodComposite DistributedDatabaseSummaryShardingMethodEnum = "COMPOSITE"
)

var mappingDistributedDatabaseSummaryShardingMethodEnum = map[string]DistributedDatabaseSummaryShardingMethodEnum{
	"USER":      DistributedDatabaseSummaryShardingMethodUser,
	"SYSTEM":    DistributedDatabaseSummaryShardingMethodSystem,
	"COMPOSITE": DistributedDatabaseSummaryShardingMethodComposite,
}

var mappingDistributedDatabaseSummaryShardingMethodEnumLowerCase = map[string]DistributedDatabaseSummaryShardingMethodEnum{
	"user":      DistributedDatabaseSummaryShardingMethodUser,
	"system":    DistributedDatabaseSummaryShardingMethodSystem,
	"composite": DistributedDatabaseSummaryShardingMethodComposite,
}

// GetDistributedDatabaseSummaryShardingMethodEnumValues Enumerates the set of values for DistributedDatabaseSummaryShardingMethodEnum
func GetDistributedDatabaseSummaryShardingMethodEnumValues() []DistributedDatabaseSummaryShardingMethodEnum {
	values := make([]DistributedDatabaseSummaryShardingMethodEnum, 0)
	for _, v := range mappingDistributedDatabaseSummaryShardingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseSummaryShardingMethodEnumStringValues Enumerates the set of values in String for DistributedDatabaseSummaryShardingMethodEnum
func GetDistributedDatabaseSummaryShardingMethodEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
		"COMPOSITE",
	}
}

// GetMappingDistributedDatabaseSummaryShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseSummaryShardingMethodEnum(val string) (DistributedDatabaseSummaryShardingMethodEnum, bool) {
	enum, ok := mappingDistributedDatabaseSummaryShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDatabaseSummaryReplicationMethodEnum Enum with underlying type: string
type DistributedDatabaseSummaryReplicationMethodEnum string

// Set of constants representing the allowable values for DistributedDatabaseSummaryReplicationMethodEnum
const (
	DistributedDatabaseSummaryReplicationMethodRaft DistributedDatabaseSummaryReplicationMethodEnum = "RAFT"
	DistributedDatabaseSummaryReplicationMethodDg   DistributedDatabaseSummaryReplicationMethodEnum = "DG"
)

var mappingDistributedDatabaseSummaryReplicationMethodEnum = map[string]DistributedDatabaseSummaryReplicationMethodEnum{
	"RAFT": DistributedDatabaseSummaryReplicationMethodRaft,
	"DG":   DistributedDatabaseSummaryReplicationMethodDg,
}

var mappingDistributedDatabaseSummaryReplicationMethodEnumLowerCase = map[string]DistributedDatabaseSummaryReplicationMethodEnum{
	"raft": DistributedDatabaseSummaryReplicationMethodRaft,
	"dg":   DistributedDatabaseSummaryReplicationMethodDg,
}

// GetDistributedDatabaseSummaryReplicationMethodEnumValues Enumerates the set of values for DistributedDatabaseSummaryReplicationMethodEnum
func GetDistributedDatabaseSummaryReplicationMethodEnumValues() []DistributedDatabaseSummaryReplicationMethodEnum {
	values := make([]DistributedDatabaseSummaryReplicationMethodEnum, 0)
	for _, v := range mappingDistributedDatabaseSummaryReplicationMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseSummaryReplicationMethodEnumStringValues Enumerates the set of values in String for DistributedDatabaseSummaryReplicationMethodEnum
func GetDistributedDatabaseSummaryReplicationMethodEnumStringValues() []string {
	return []string{
		"RAFT",
		"DG",
	}
}

// GetMappingDistributedDatabaseSummaryReplicationMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseSummaryReplicationMethodEnum(val string) (DistributedDatabaseSummaryReplicationMethodEnum, bool) {
	enum, ok := mappingDistributedDatabaseSummaryReplicationMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDatabaseSummaryDbDeploymentTypeEnum Enum with underlying type: string
type DistributedDatabaseSummaryDbDeploymentTypeEnum string

// Set of constants representing the allowable values for DistributedDatabaseSummaryDbDeploymentTypeEnum
const (
	DistributedDatabaseSummaryDbDeploymentTypeExadb DistributedDatabaseSummaryDbDeploymentTypeEnum = "EXADB"
)

var mappingDistributedDatabaseSummaryDbDeploymentTypeEnum = map[string]DistributedDatabaseSummaryDbDeploymentTypeEnum{
	"EXADB": DistributedDatabaseSummaryDbDeploymentTypeExadb,
}

var mappingDistributedDatabaseSummaryDbDeploymentTypeEnumLowerCase = map[string]DistributedDatabaseSummaryDbDeploymentTypeEnum{
	"exadb": DistributedDatabaseSummaryDbDeploymentTypeExadb,
}

// GetDistributedDatabaseSummaryDbDeploymentTypeEnumValues Enumerates the set of values for DistributedDatabaseSummaryDbDeploymentTypeEnum
func GetDistributedDatabaseSummaryDbDeploymentTypeEnumValues() []DistributedDatabaseSummaryDbDeploymentTypeEnum {
	values := make([]DistributedDatabaseSummaryDbDeploymentTypeEnum, 0)
	for _, v := range mappingDistributedDatabaseSummaryDbDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseSummaryDbDeploymentTypeEnumStringValues Enumerates the set of values in String for DistributedDatabaseSummaryDbDeploymentTypeEnum
func GetDistributedDatabaseSummaryDbDeploymentTypeEnumStringValues() []string {
	return []string{
		"EXADB",
	}
}

// GetMappingDistributedDatabaseSummaryDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseSummaryDbDeploymentTypeEnum(val string) (DistributedDatabaseSummaryDbDeploymentTypeEnum, bool) {
	enum, ok := mappingDistributedDatabaseSummaryDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
