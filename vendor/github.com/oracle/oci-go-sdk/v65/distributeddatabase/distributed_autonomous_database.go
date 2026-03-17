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

// DistributedAutonomousDatabase Globally distributed autonomous database.
type DistributedAutonomousDatabase struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Globally distributed autonomous database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the Globally distributed autonomous database was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Globally distributed autonomous database was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Oracle Database version for the shards and catalog used in Globally distributed autonomous database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Lifecycle states for the Globally distributed autonomous database.
	LifecycleState DistributedAutonomousDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The lifecycleDetails for the Globally distributed autonomous database.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Unique name prefix for the Globally distributed autonomous databases. Only alpha-numeric values are allowed. First character
	// has to be a letter followed by any combination of letter and number.
	Prefix *string `mandatory:"true" json:"prefix"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
	PrivateEndpointIds []string `mandatory:"true" json:"privateEndpointIds"`

	// Sharding Methods for the Globally distributed autonomous database.
	ShardingMethod DistributedAutonomousDatabaseShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

	// Possible workload types. Currently only OLTP workload type is supported.
	DbWorkload DistributedAutonomousDatabaseDbWorkloadEnum `mandatory:"true" json:"dbWorkload"`

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

	// The distributed autonomous database deployment type.
	DbDeploymentType DistributedAutonomousDatabaseDbDeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`

	ConnectionStrings *DistributedAutonomousDatabaseConnectionString `mandatory:"false" json:"connectionStrings"`

	// The default number of unique chunks in a shardspace. The value of chunks must be
	// greater than 2 times the size of the largest shardgroup in any shardspace.
	Chunks *int `mandatory:"false" json:"chunks"`

	// The TLS listener port number for Globally distributed autonomous database.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// The Replication method for Globally distributed autonomous database. Use RAFT for Raft replication, and DG for
	// DataGuard. If replicationMethod is not provided, it defaults to DG.
	ReplicationMethod DistributedAutonomousDatabaseReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`

	// The Replication factor for RAFT replication based Globally distributed autonomous database. Currently supported values are 3, 5 and 7.
	ReplicationFactor *int `mandatory:"false" json:"replicationFactor"`

	// The replication unit count for RAFT based distributed autonomous database. For RAFT replication based
	// Globally distributed autonomous database, the value should be at least twice the number of shards.
	ReplicationUnit *int `mandatory:"false" json:"replicationUnit"`

	LatestGsmImage *DistributedAutonomousDatabaseGsmImage `mandatory:"false" json:"latestGsmImage"`

	// Collection of shards associated with the Globally distributed autonomous database.
	ShardDetails []DistributedAutonomousDatabaseShard `mandatory:"false" json:"shardDetails"`

	// Collection of catalogs associated with the Globally distributed autonomous database.
	CatalogDetails []DistributedAutonomousDatabaseCatalog `mandatory:"false" json:"catalogDetails"`

	// Collection of catalogs associated with the Globally distributed autonomous database.
	GsmDetails []DistributedAutonomousDatabaseGsm `mandatory:"false" json:"gsmDetails"`

	DbBackupConfig *DistributedAutonomousDbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

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
	if _, ok := GetMappingDistributedAutonomousDatabaseDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetDistributedAutonomousDatabaseDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedAutonomousDatabaseDbDeploymentTypeEnum(string(m.DbDeploymentType)); !ok && m.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", m.DbDeploymentType, strings.Join(GetDistributedAutonomousDatabaseDbDeploymentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDistributedAutonomousDatabaseReplicationMethodEnum(string(m.ReplicationMethod)); !ok && m.ReplicationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationMethod: %s. Supported values are: %s.", m.ReplicationMethod, strings.Join(GetDistributedAutonomousDatabaseReplicationMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DistributedAutonomousDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionStrings  *DistributedAutonomousDatabaseConnectionString     `json:"connectionStrings"`
		Chunks             *int                                               `json:"chunks"`
		ListenerPortTls    *int                                               `json:"listenerPortTls"`
		ReplicationMethod  DistributedAutonomousDatabaseReplicationMethodEnum `json:"replicationMethod"`
		ReplicationFactor  *int                                               `json:"replicationFactor"`
		ReplicationUnit    *int                                               `json:"replicationUnit"`
		LatestGsmImage     *DistributedAutonomousDatabaseGsmImage             `json:"latestGsmImage"`
		ShardDetails       []distributedautonomousdatabaseshard               `json:"shardDetails"`
		CatalogDetails     []distributedautonomousdatabasecatalog             `json:"catalogDetails"`
		GsmDetails         []DistributedAutonomousDatabaseGsm                 `json:"gsmDetails"`
		DbBackupConfig     *DistributedAutonomousDbBackupConfig               `json:"dbBackupConfig"`
		Metadata           *DistributedAutonomousDbMetadata                   `json:"metadata"`
		FreeformTags       map[string]string                                  `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{}                  `json:"definedTags"`
		SystemTags         map[string]map[string]interface{}                  `json:"systemTags"`
		Id                 *string                                            `json:"id"`
		CompartmentId      *string                                            `json:"compartmentId"`
		DisplayName        *string                                            `json:"displayName"`
		TimeCreated        *common.SDKTime                                    `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                                    `json:"timeUpdated"`
		DatabaseVersion    *string                                            `json:"databaseVersion"`
		LifecycleState     DistributedAutonomousDatabaseLifecycleStateEnum    `json:"lifecycleState"`
		LifecycleDetails   *string                                            `json:"lifecycleDetails"`
		Prefix             *string                                            `json:"prefix"`
		PrivateEndpointIds []string                                           `json:"privateEndpointIds"`
		ShardingMethod     DistributedAutonomousDatabaseShardingMethodEnum    `json:"shardingMethod"`
		DbWorkload         DistributedAutonomousDatabaseDbWorkloadEnum        `json:"dbWorkload"`
		CharacterSet       *string                                            `json:"characterSet"`
		NcharacterSet      *string                                            `json:"ncharacterSet"`
		ListenerPort       *int                                               `json:"listenerPort"`
		OnsPortLocal       *int                                               `json:"onsPortLocal"`
		OnsPortRemote      *int                                               `json:"onsPortRemote"`
		DbDeploymentType   DistributedAutonomousDatabaseDbDeploymentTypeEnum  `json:"dbDeploymentType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConnectionStrings = model.ConnectionStrings

	m.Chunks = model.Chunks

	m.ListenerPortTls = model.ListenerPortTls

	m.ReplicationMethod = model.ReplicationMethod

	m.ReplicationFactor = model.ReplicationFactor

	m.ReplicationUnit = model.ReplicationUnit

	m.LatestGsmImage = model.LatestGsmImage

	m.ShardDetails = make([]DistributedAutonomousDatabaseShard, len(model.ShardDetails))
	for i, n := range model.ShardDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ShardDetails[i] = nn.(DistributedAutonomousDatabaseShard)
		} else {
			m.ShardDetails[i] = nil
		}
	}
	m.CatalogDetails = make([]DistributedAutonomousDatabaseCatalog, len(model.CatalogDetails))
	for i, n := range model.CatalogDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.CatalogDetails[i] = nn.(DistributedAutonomousDatabaseCatalog)
		} else {
			m.CatalogDetails[i] = nil
		}
	}
	m.GsmDetails = make([]DistributedAutonomousDatabaseGsm, len(model.GsmDetails))
	copy(m.GsmDetails, model.GsmDetails)
	m.DbBackupConfig = model.DbBackupConfig

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

	m.DbWorkload = model.DbWorkload

	m.CharacterSet = model.CharacterSet

	m.NcharacterSet = model.NcharacterSet

	m.ListenerPort = model.ListenerPort

	m.OnsPortLocal = model.OnsPortLocal

	m.OnsPortRemote = model.OnsPortRemote

	m.DbDeploymentType = model.DbDeploymentType

	return
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
	DistributedAutonomousDatabaseShardingMethodUser   DistributedAutonomousDatabaseShardingMethodEnum = "USER"
	DistributedAutonomousDatabaseShardingMethodSystem DistributedAutonomousDatabaseShardingMethodEnum = "SYSTEM"
)

var mappingDistributedAutonomousDatabaseShardingMethodEnum = map[string]DistributedAutonomousDatabaseShardingMethodEnum{
	"USER":   DistributedAutonomousDatabaseShardingMethodUser,
	"SYSTEM": DistributedAutonomousDatabaseShardingMethodSystem,
}

var mappingDistributedAutonomousDatabaseShardingMethodEnumLowerCase = map[string]DistributedAutonomousDatabaseShardingMethodEnum{
	"user":   DistributedAutonomousDatabaseShardingMethodUser,
	"system": DistributedAutonomousDatabaseShardingMethodSystem,
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
	}
}

// GetMappingDistributedAutonomousDatabaseShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseShardingMethodEnum(val string) (DistributedAutonomousDatabaseShardingMethodEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedAutonomousDatabaseDbWorkloadEnum Enum with underlying type: string
type DistributedAutonomousDatabaseDbWorkloadEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseDbWorkloadEnum
const (
	DistributedAutonomousDatabaseDbWorkloadOltp DistributedAutonomousDatabaseDbWorkloadEnum = "OLTP"
	DistributedAutonomousDatabaseDbWorkloadDw   DistributedAutonomousDatabaseDbWorkloadEnum = "DW"
)

var mappingDistributedAutonomousDatabaseDbWorkloadEnum = map[string]DistributedAutonomousDatabaseDbWorkloadEnum{
	"OLTP": DistributedAutonomousDatabaseDbWorkloadOltp,
	"DW":   DistributedAutonomousDatabaseDbWorkloadDw,
}

var mappingDistributedAutonomousDatabaseDbWorkloadEnumLowerCase = map[string]DistributedAutonomousDatabaseDbWorkloadEnum{
	"oltp": DistributedAutonomousDatabaseDbWorkloadOltp,
	"dw":   DistributedAutonomousDatabaseDbWorkloadDw,
}

// GetDistributedAutonomousDatabaseDbWorkloadEnumValues Enumerates the set of values for DistributedAutonomousDatabaseDbWorkloadEnum
func GetDistributedAutonomousDatabaseDbWorkloadEnumValues() []DistributedAutonomousDatabaseDbWorkloadEnum {
	values := make([]DistributedAutonomousDatabaseDbWorkloadEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseDbWorkloadEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseDbWorkloadEnum
func GetDistributedAutonomousDatabaseDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
	}
}

// GetMappingDistributedAutonomousDatabaseDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseDbWorkloadEnum(val string) (DistributedAutonomousDatabaseDbWorkloadEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseDbWorkloadEnumLowerCase[strings.ToLower(val)]
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
