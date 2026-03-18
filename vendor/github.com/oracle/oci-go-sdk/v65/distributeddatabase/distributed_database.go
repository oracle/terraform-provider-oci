// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"encoding/json"
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

	// The time the Globally distributed database was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Globally distributed database was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Oracle Database version for the shards and catalog used in Globally distributed database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Lifecycle states for the Globally distributed database.
	LifecycleState DistributedDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The lifecycleDetails for the Globally distributed database.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Unique name prefix for the Globally distributed databases. Only alpha-numeric values are allowed. First character
	// has to be a letter followed by any combination of letter and number.
	Prefix *string `mandatory:"true" json:"prefix"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
	PrivateEndpointIds []string `mandatory:"true" json:"privateEndpointIds"`

	// Sharding Methods for the Globally distributed database.
	ShardingMethod DistributedDatabaseShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

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

	// The distributed database deployment type.
	DbDeploymentType DistributedDatabaseDbDeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`

	ConnectionStrings *DistributedDbConnectionString `mandatory:"false" json:"connectionStrings"`

	LatestGsmImageDetails *DistributedDbGsmImage `mandatory:"false" json:"latestGsmImageDetails"`

	// The default number of unique chunks in a shardspace. The value of chunks must be
	// greater than 2 times the size of the largest shardgroup in any shardspace.
	Chunks *int `mandatory:"false" json:"chunks"`

	// The TLS listener port number for Globally distributed database.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// The TCP Single Client Access Name (SCAN) port for Globally distributed database clusters.
	ScanListenerPort *int `mandatory:"false" json:"scanListenerPort"`

	// The Replication method for Globally distributed database. Use RAFT for Raft replication, and DG for
	// DataGuard. If replicationMethod is not provided, it defaults to DG.
	ReplicationMethod DistributedDatabaseReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`

	// The Replication factor for RAFT replication based Globally distributed database. Currently supported values are 3, 5 and 7.
	ReplicationFactor *int `mandatory:"false" json:"replicationFactor"`

	// The replication unit count for RAFT based distributed database. For RAFT replication based
	// Globally distributed database, the value should be at least twice the number of shards.
	ReplicationUnit *int `mandatory:"false" json:"replicationUnit"`

	// Collection of shards associated with the Globally distributed database.
	ShardDetails []DistributedDatabaseShard `mandatory:"false" json:"shardDetails"`

	// Collection of catalogs associated with the Globally distributed database.
	CatalogDetails []DistributedDatabaseCatalog `mandatory:"false" json:"catalogDetails"`

	// Collection of catalogs associated with the Globally distributed database.
	GsmDetails []DistributedDatabaseGsm `mandatory:"false" json:"gsmDetails"`

	DbBackupConfig *DistributedDbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// The SSH public key for Global service manager instances.
	GsmSshPublicKey *string `mandatory:"false" json:"gsmSshPublicKey"`

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

// UnmarshalJSON unmarshals from json
func (m *DistributedDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionStrings     *DistributedDbConnectionString           `json:"connectionStrings"`
		LatestGsmImageDetails *DistributedDbGsmImage                   `json:"latestGsmImageDetails"`
		Chunks                *int                                     `json:"chunks"`
		ListenerPortTls       *int                                     `json:"listenerPortTls"`
		ScanListenerPort      *int                                     `json:"scanListenerPort"`
		ReplicationMethod     DistributedDatabaseReplicationMethodEnum `json:"replicationMethod"`
		ReplicationFactor     *int                                     `json:"replicationFactor"`
		ReplicationUnit       *int                                     `json:"replicationUnit"`
		ShardDetails          []distributeddatabaseshard               `json:"shardDetails"`
		CatalogDetails        []distributeddatabasecatalog             `json:"catalogDetails"`
		GsmDetails            []DistributedDatabaseGsm                 `json:"gsmDetails"`
		DbBackupConfig        *DistributedDbBackupConfig               `json:"dbBackupConfig"`
		GsmSshPublicKey       *string                                  `json:"gsmSshPublicKey"`
		Metadata              *DistributedDbMetadata                   `json:"metadata"`
		FreeformTags          map[string]string                        `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}        `json:"definedTags"`
		SystemTags            map[string]map[string]interface{}        `json:"systemTags"`
		Id                    *string                                  `json:"id"`
		CompartmentId         *string                                  `json:"compartmentId"`
		DisplayName           *string                                  `json:"displayName"`
		TimeCreated           *common.SDKTime                          `json:"timeCreated"`
		TimeUpdated           *common.SDKTime                          `json:"timeUpdated"`
		DatabaseVersion       *string                                  `json:"databaseVersion"`
		LifecycleState        DistributedDatabaseLifecycleStateEnum    `json:"lifecycleState"`
		LifecycleDetails      *string                                  `json:"lifecycleDetails"`
		Prefix                *string                                  `json:"prefix"`
		PrivateEndpointIds    []string                                 `json:"privateEndpointIds"`
		ShardingMethod        DistributedDatabaseShardingMethodEnum    `json:"shardingMethod"`
		CharacterSet          *string                                  `json:"characterSet"`
		NcharacterSet         *string                                  `json:"ncharacterSet"`
		ListenerPort          *int                                     `json:"listenerPort"`
		OnsPortLocal          *int                                     `json:"onsPortLocal"`
		OnsPortRemote         *int                                     `json:"onsPortRemote"`
		DbDeploymentType      DistributedDatabaseDbDeploymentTypeEnum  `json:"dbDeploymentType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConnectionStrings = model.ConnectionStrings

	m.LatestGsmImageDetails = model.LatestGsmImageDetails

	m.Chunks = model.Chunks

	m.ListenerPortTls = model.ListenerPortTls

	m.ScanListenerPort = model.ScanListenerPort

	m.ReplicationMethod = model.ReplicationMethod

	m.ReplicationFactor = model.ReplicationFactor

	m.ReplicationUnit = model.ReplicationUnit

	m.ShardDetails = make([]DistributedDatabaseShard, len(model.ShardDetails))
	for i, n := range model.ShardDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ShardDetails[i] = nn.(DistributedDatabaseShard)
		} else {
			m.ShardDetails[i] = nil
		}
	}
	m.CatalogDetails = make([]DistributedDatabaseCatalog, len(model.CatalogDetails))
	for i, n := range model.CatalogDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.CatalogDetails[i] = nn.(DistributedDatabaseCatalog)
		} else {
			m.CatalogDetails[i] = nil
		}
	}
	m.GsmDetails = make([]DistributedDatabaseGsm, len(model.GsmDetails))
	copy(m.GsmDetails, model.GsmDetails)
	m.DbBackupConfig = model.DbBackupConfig

	m.GsmSshPublicKey = model.GsmSshPublicKey

	m.Metadata = model.Metadata

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.DatabaseVersion = model.DatabaseVersion

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.Prefix = model.Prefix

	m.PrivateEndpointIds = make([]string, len(model.PrivateEndpointIds))
	copy(m.PrivateEndpointIds, model.PrivateEndpointIds)
	m.ShardingMethod = model.ShardingMethod

	m.CharacterSet = model.CharacterSet

	m.NcharacterSet = model.NcharacterSet

	m.ListenerPort = model.ListenerPort

	m.OnsPortLocal = model.OnsPortLocal

	m.OnsPortRemote = model.OnsPortRemote

	m.DbDeploymentType = model.DbDeploymentType

	return
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
	DistributedDatabaseShardingMethodUser   DistributedDatabaseShardingMethodEnum = "USER"
	DistributedDatabaseShardingMethodSystem DistributedDatabaseShardingMethodEnum = "SYSTEM"
)

var mappingDistributedDatabaseShardingMethodEnum = map[string]DistributedDatabaseShardingMethodEnum{
	"USER":   DistributedDatabaseShardingMethodUser,
	"SYSTEM": DistributedDatabaseShardingMethodSystem,
}

var mappingDistributedDatabaseShardingMethodEnumLowerCase = map[string]DistributedDatabaseShardingMethodEnum{
	"user":   DistributedDatabaseShardingMethodUser,
	"system": DistributedDatabaseShardingMethodSystem,
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
	DistributedDatabaseDbDeploymentTypeExadbXs DistributedDatabaseDbDeploymentTypeEnum = "EXADB_XS"
)

var mappingDistributedDatabaseDbDeploymentTypeEnum = map[string]DistributedDatabaseDbDeploymentTypeEnum{
	"EXADB_XS": DistributedDatabaseDbDeploymentTypeExadbXs,
}

var mappingDistributedDatabaseDbDeploymentTypeEnumLowerCase = map[string]DistributedDatabaseDbDeploymentTypeEnum{
	"exadb_xs": DistributedDatabaseDbDeploymentTypeExadbXs,
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
		"EXADB_XS",
	}
}

// GetMappingDistributedDatabaseDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseDbDeploymentTypeEnum(val string) (DistributedDatabaseDbDeploymentTypeEnum, bool) {
	enum, ok := mappingDistributedDatabaseDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
