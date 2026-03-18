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

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
	PrivateEndpointIds []string `mandatory:"true" json:"privateEndpointIds"`

	// Sharding Methods for the Globally distributed database.
	ShardingMethod CreateDistributedDatabaseDetailsShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

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

	// The distributed database deployment type.
	DbDeploymentType CreateDistributedDatabaseDetailsDbDeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`

	// Collection of shards for the Globally distributed database.
	ShardDetails []CreateDistributedDatabaseShardDetails `mandatory:"true" json:"shardDetails"`

	// Collection of catalog for the Globally distributed database.
	CatalogDetails []CreateDistributedDatabaseCatalogDetails `mandatory:"true" json:"catalogDetails"`

	// Number of chunks in a shardspace. The value of chunks must be
	// greater than 2 times the size of the largest shardgroup in any shardspace. Chunks is
	// required to be provided for distributed databases being created with
	// SYSTEM shardingMethod. For USER shardingMethod, chunks should not be set in create payload.
	Chunks *int `mandatory:"false" json:"chunks"`

	// The TLS listener port number for the Globally distributed database. The TLS listener port number
	// has to be unique for a customer tenancy across all distributed databases. Same port number should
	// not be re-used for any other distributed database. For BASE_DB and EXADB_XS based distributed databases,
	// tls is not supported hence the listenerPortTls is not needed to be provided in create payload.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// The TCP Single Client Access Name (SCAN) port for clusters created for Globally distributed database.
	// The scanListenerPort number should only be provided if shard and catalog have source type NEW_VAULT_AND_CLUSTER.
	// If shard and catalog have source type NEW_VAULT_AND_CLUSTER and scanListenerPort is not provided then the
	// scanListenerPort will default to value 1521.
	ScanListenerPort *int `mandatory:"false" json:"scanListenerPort"`

	// The Replication method for Globally distributed database. Use RAFT for Raft based replication.
	// With RAFT replication, shards cannot have peers details set on them. In case shards need to
	// have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or
	// without peers), please set replicationMethod as DG or do not set any value for replicationMethod.
	ReplicationMethod CreateDistributedDatabaseDetailsReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`

	// The Replication factor for RAFT replication based Globally distributed database. Currently supported values are 3, 5 and 7.
	ReplicationFactor *int `mandatory:"false" json:"replicationFactor"`

	// The replication unit count for RAFT based distributed database. For RAFT replication based
	// Globally distributed database, the value should be at least twice the number of shards.
	ReplicationUnit *int `mandatory:"false" json:"replicationUnit"`

	// The SSH public key for Global service manager instances.
	GsmSshPublicKey *string `mandatory:"false" json:"gsmSshPublicKey"`

	DbBackupConfig *DistributedDbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

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
	if _, ok := GetMappingCreateDistributedDatabaseDetailsShardingMethodEnum(string(m.ShardingMethod)); !ok && m.ShardingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShardingMethod: %s. Supported values are: %s.", m.ShardingMethod, strings.Join(GetCreateDistributedDatabaseDetailsShardingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDistributedDatabaseDetailsDbDeploymentTypeEnum(string(m.DbDeploymentType)); !ok && m.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", m.DbDeploymentType, strings.Join(GetCreateDistributedDatabaseDetailsDbDeploymentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateDistributedDatabaseDetailsReplicationMethodEnum(string(m.ReplicationMethod)); !ok && m.ReplicationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationMethod: %s. Supported values are: %s.", m.ReplicationMethod, strings.Join(GetCreateDistributedDatabaseDetailsReplicationMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDistributedDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Chunks             *int                                                  `json:"chunks"`
		ListenerPortTls    *int                                                  `json:"listenerPortTls"`
		ScanListenerPort   *int                                                  `json:"scanListenerPort"`
		ReplicationMethod  CreateDistributedDatabaseDetailsReplicationMethodEnum `json:"replicationMethod"`
		ReplicationFactor  *int                                                  `json:"replicationFactor"`
		ReplicationUnit    *int                                                  `json:"replicationUnit"`
		GsmSshPublicKey    *string                                               `json:"gsmSshPublicKey"`
		DbBackupConfig     *DistributedDbBackupConfig                            `json:"dbBackupConfig"`
		FreeformTags       map[string]string                                     `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{}                     `json:"definedTags"`
		CompartmentId      *string                                               `json:"compartmentId"`
		DisplayName        *string                                               `json:"displayName"`
		DatabaseVersion    *string                                               `json:"databaseVersion"`
		Prefix             *string                                               `json:"prefix"`
		PrivateEndpointIds []string                                              `json:"privateEndpointIds"`
		ShardingMethod     CreateDistributedDatabaseDetailsShardingMethodEnum    `json:"shardingMethod"`
		CharacterSet       *string                                               `json:"characterSet"`
		NcharacterSet      *string                                               `json:"ncharacterSet"`
		ListenerPort       *int                                                  `json:"listenerPort"`
		OnsPortLocal       *int                                                  `json:"onsPortLocal"`
		OnsPortRemote      *int                                                  `json:"onsPortRemote"`
		DbDeploymentType   CreateDistributedDatabaseDetailsDbDeploymentTypeEnum  `json:"dbDeploymentType"`
		ShardDetails       []createdistributeddatabasesharddetails               `json:"shardDetails"`
		CatalogDetails     []createdistributeddatabasecatalogdetails             `json:"catalogDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Chunks = model.Chunks

	m.ListenerPortTls = model.ListenerPortTls

	m.ScanListenerPort = model.ScanListenerPort

	m.ReplicationMethod = model.ReplicationMethod

	m.ReplicationFactor = model.ReplicationFactor

	m.ReplicationUnit = model.ReplicationUnit

	m.GsmSshPublicKey = model.GsmSshPublicKey

	m.DbBackupConfig = model.DbBackupConfig

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.DatabaseVersion = model.DatabaseVersion

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

	m.ShardDetails = make([]CreateDistributedDatabaseShardDetails, len(model.ShardDetails))
	for i, n := range model.ShardDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ShardDetails[i] = nn.(CreateDistributedDatabaseShardDetails)
		} else {
			m.ShardDetails[i] = nil
		}
	}
	m.CatalogDetails = make([]CreateDistributedDatabaseCatalogDetails, len(model.CatalogDetails))
	for i, n := range model.CatalogDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.CatalogDetails[i] = nn.(CreateDistributedDatabaseCatalogDetails)
		} else {
			m.CatalogDetails[i] = nil
		}
	}
	return
}

// CreateDistributedDatabaseDetailsShardingMethodEnum Enum with underlying type: string
type CreateDistributedDatabaseDetailsShardingMethodEnum string

// Set of constants representing the allowable values for CreateDistributedDatabaseDetailsShardingMethodEnum
const (
	CreateDistributedDatabaseDetailsShardingMethodUser   CreateDistributedDatabaseDetailsShardingMethodEnum = "USER"
	CreateDistributedDatabaseDetailsShardingMethodSystem CreateDistributedDatabaseDetailsShardingMethodEnum = "SYSTEM"
)

var mappingCreateDistributedDatabaseDetailsShardingMethodEnum = map[string]CreateDistributedDatabaseDetailsShardingMethodEnum{
	"USER":   CreateDistributedDatabaseDetailsShardingMethodUser,
	"SYSTEM": CreateDistributedDatabaseDetailsShardingMethodSystem,
}

var mappingCreateDistributedDatabaseDetailsShardingMethodEnumLowerCase = map[string]CreateDistributedDatabaseDetailsShardingMethodEnum{
	"user":   CreateDistributedDatabaseDetailsShardingMethodUser,
	"system": CreateDistributedDatabaseDetailsShardingMethodSystem,
}

// GetCreateDistributedDatabaseDetailsShardingMethodEnumValues Enumerates the set of values for CreateDistributedDatabaseDetailsShardingMethodEnum
func GetCreateDistributedDatabaseDetailsShardingMethodEnumValues() []CreateDistributedDatabaseDetailsShardingMethodEnum {
	values := make([]CreateDistributedDatabaseDetailsShardingMethodEnum, 0)
	for _, v := range mappingCreateDistributedDatabaseDetailsShardingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedDatabaseDetailsShardingMethodEnumStringValues Enumerates the set of values in String for CreateDistributedDatabaseDetailsShardingMethodEnum
func GetCreateDistributedDatabaseDetailsShardingMethodEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingCreateDistributedDatabaseDetailsShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedDatabaseDetailsShardingMethodEnum(val string) (CreateDistributedDatabaseDetailsShardingMethodEnum, bool) {
	enum, ok := mappingCreateDistributedDatabaseDetailsShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDistributedDatabaseDetailsReplicationMethodEnum Enum with underlying type: string
type CreateDistributedDatabaseDetailsReplicationMethodEnum string

// Set of constants representing the allowable values for CreateDistributedDatabaseDetailsReplicationMethodEnum
const (
	CreateDistributedDatabaseDetailsReplicationMethodRaft CreateDistributedDatabaseDetailsReplicationMethodEnum = "RAFT"
	CreateDistributedDatabaseDetailsReplicationMethodDg   CreateDistributedDatabaseDetailsReplicationMethodEnum = "DG"
)

var mappingCreateDistributedDatabaseDetailsReplicationMethodEnum = map[string]CreateDistributedDatabaseDetailsReplicationMethodEnum{
	"RAFT": CreateDistributedDatabaseDetailsReplicationMethodRaft,
	"DG":   CreateDistributedDatabaseDetailsReplicationMethodDg,
}

var mappingCreateDistributedDatabaseDetailsReplicationMethodEnumLowerCase = map[string]CreateDistributedDatabaseDetailsReplicationMethodEnum{
	"raft": CreateDistributedDatabaseDetailsReplicationMethodRaft,
	"dg":   CreateDistributedDatabaseDetailsReplicationMethodDg,
}

// GetCreateDistributedDatabaseDetailsReplicationMethodEnumValues Enumerates the set of values for CreateDistributedDatabaseDetailsReplicationMethodEnum
func GetCreateDistributedDatabaseDetailsReplicationMethodEnumValues() []CreateDistributedDatabaseDetailsReplicationMethodEnum {
	values := make([]CreateDistributedDatabaseDetailsReplicationMethodEnum, 0)
	for _, v := range mappingCreateDistributedDatabaseDetailsReplicationMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedDatabaseDetailsReplicationMethodEnumStringValues Enumerates the set of values in String for CreateDistributedDatabaseDetailsReplicationMethodEnum
func GetCreateDistributedDatabaseDetailsReplicationMethodEnumStringValues() []string {
	return []string{
		"RAFT",
		"DG",
	}
}

// GetMappingCreateDistributedDatabaseDetailsReplicationMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedDatabaseDetailsReplicationMethodEnum(val string) (CreateDistributedDatabaseDetailsReplicationMethodEnum, bool) {
	enum, ok := mappingCreateDistributedDatabaseDetailsReplicationMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDistributedDatabaseDetailsDbDeploymentTypeEnum Enum with underlying type: string
type CreateDistributedDatabaseDetailsDbDeploymentTypeEnum string

// Set of constants representing the allowable values for CreateDistributedDatabaseDetailsDbDeploymentTypeEnum
const (
	CreateDistributedDatabaseDetailsDbDeploymentTypeExadbXs CreateDistributedDatabaseDetailsDbDeploymentTypeEnum = "EXADB_XS"
)

var mappingCreateDistributedDatabaseDetailsDbDeploymentTypeEnum = map[string]CreateDistributedDatabaseDetailsDbDeploymentTypeEnum{
	"EXADB_XS": CreateDistributedDatabaseDetailsDbDeploymentTypeExadbXs,
}

var mappingCreateDistributedDatabaseDetailsDbDeploymentTypeEnumLowerCase = map[string]CreateDistributedDatabaseDetailsDbDeploymentTypeEnum{
	"exadb_xs": CreateDistributedDatabaseDetailsDbDeploymentTypeExadbXs,
}

// GetCreateDistributedDatabaseDetailsDbDeploymentTypeEnumValues Enumerates the set of values for CreateDistributedDatabaseDetailsDbDeploymentTypeEnum
func GetCreateDistributedDatabaseDetailsDbDeploymentTypeEnumValues() []CreateDistributedDatabaseDetailsDbDeploymentTypeEnum {
	values := make([]CreateDistributedDatabaseDetailsDbDeploymentTypeEnum, 0)
	for _, v := range mappingCreateDistributedDatabaseDetailsDbDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedDatabaseDetailsDbDeploymentTypeEnumStringValues Enumerates the set of values in String for CreateDistributedDatabaseDetailsDbDeploymentTypeEnum
func GetCreateDistributedDatabaseDetailsDbDeploymentTypeEnumStringValues() []string {
	return []string{
		"EXADB_XS",
	}
}

// GetMappingCreateDistributedDatabaseDetailsDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedDatabaseDetailsDbDeploymentTypeEnum(val string) (CreateDistributedDatabaseDetailsDbDeploymentTypeEnum, bool) {
	enum, ok := mappingCreateDistributedDatabaseDetailsDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
