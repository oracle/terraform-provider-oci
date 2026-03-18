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

// CreateDistributedAutonomousDatabaseDetails Details required for creation of the Globally distributed autonomous database.
type CreateDistributedAutonomousDatabaseDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Globally distributed autonomous database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Oracle Database version for the shards and catalog used in Globally distributed autonomous database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Unique name prefix for the Globally distributed autonomous databases. Only alpha-numeric values are allowed. First character
	// has to be a letter followed by any combination of letter and number.
	Prefix *string `mandatory:"true" json:"prefix"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
	PrivateEndpointIds []string `mandatory:"true" json:"privateEndpointIds"`

	// Sharding Methods for the Globally distributed autonomous database.
	ShardingMethod CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

	// Possible workload types. Currently only OLTP workload type is supported.
	DbWorkload CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum `mandatory:"true" json:"dbWorkload"`

	// The character set for the database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// The listener port number for the Globally distributed autonomous database. The listener port number
	// has to be unique for a customer tenancy across all distributed autonomous databases. Same port number
	// should not be re-used for any other distributed autonomous database.
	ListenerPort *int `mandatory:"true" json:"listenerPort"`

	// Ons local port number for Globally distributed autonomous database. The onsPortLocal has to be unique for
	// a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for
	// any other distributed autonomous database.
	OnsPortLocal *int `mandatory:"true" json:"onsPortLocal"`

	// Ons remote port number for Globally distributed autonomous database. The onsPortRemote has to be unique for
	// a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for
	// any other distributed autonomous database.
	OnsPortRemote *int `mandatory:"true" json:"onsPortRemote"`

	// The distributed autonomous database deployment type.
	DbDeploymentType CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`

	// Collection of shards for the Globally distributed autonomous database.
	ShardDetails []CreateDistributedAutonomousDatabaseShardDetails `mandatory:"true" json:"shardDetails"`

	// Collection of catalog for the Globally distributed autonomous database.
	CatalogDetails []CreateDistributedAutonomousDatabaseCatalogDetails `mandatory:"true" json:"catalogDetails"`

	// Number of chunks in a shardspace. The value of chunks must be
	// greater than 2 times the size of the largest shardgroup in any shardspace. Chunks is
	// required to be provided for distributed autonomous databases being created with
	// SYSTEM shardingMethod. For USER shardingMethod, chunks should not be set in create payload.
	Chunks *int `mandatory:"false" json:"chunks"`

	// The TLS listener port number for Globally distributed autonomous database. The TLS listener port number
	// has to be unique for a customer tenancy across all distributed autonomous databases. Same port number
	// should not be re-used for any other distributed autonomous database. The listenerPortTls is mandatory
	// for dedicated infrastructure based distributed autonomous databases.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// The Replication method for Globally distributed autonomous database. Use RAFT for Raft based replication.
	// With RAFT replication, shards cannot have peers details set on them. In case shards need to
	// have peers, please do not set RAFT replicationMethod. For all non RAFT replication cases (with or
	// without peers), please set replicationMethod as DG or do not set any value for replicationMethod.
	ReplicationMethod CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`

	// The Replication factor for RAFT replication based Globally distributed autonomous database. Currently supported values are 3, 5 and 7.
	ReplicationFactor *int `mandatory:"false" json:"replicationFactor"`

	// The replication unit count for RAFT based distributed autonomous database. For RAFT replication based
	// Globally distributed autonomous database, the value should be at least twice the number of shards.
	ReplicationUnit *int `mandatory:"false" json:"replicationUnit"`

	DbBackupConfig *DistributedAutonomousDbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDistributedAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedAutonomousDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDistributedAutonomousDatabaseDetailsShardingMethodEnum(string(m.ShardingMethod)); !ok && m.ShardingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShardingMethod: %s. Supported values are: %s.", m.ShardingMethod, strings.Join(GetCreateDistributedAutonomousDatabaseDetailsShardingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum(string(m.DbDeploymentType)); !ok && m.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", m.DbDeploymentType, strings.Join(GetCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum(string(m.ReplicationMethod)); !ok && m.ReplicationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationMethod: %s. Supported values are: %s.", m.ReplicationMethod, strings.Join(GetCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDistributedAutonomousDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Chunks             *int                                                            `json:"chunks"`
		ListenerPortTls    *int                                                            `json:"listenerPortTls"`
		ReplicationMethod  CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum `json:"replicationMethod"`
		ReplicationFactor  *int                                                            `json:"replicationFactor"`
		ReplicationUnit    *int                                                            `json:"replicationUnit"`
		DbBackupConfig     *DistributedAutonomousDbBackupConfig                            `json:"dbBackupConfig"`
		FreeformTags       map[string]string                                               `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{}                               `json:"definedTags"`
		CompartmentId      *string                                                         `json:"compartmentId"`
		DisplayName        *string                                                         `json:"displayName"`
		DatabaseVersion    *string                                                         `json:"databaseVersion"`
		Prefix             *string                                                         `json:"prefix"`
		PrivateEndpointIds []string                                                        `json:"privateEndpointIds"`
		ShardingMethod     CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum    `json:"shardingMethod"`
		DbWorkload         CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum        `json:"dbWorkload"`
		CharacterSet       *string                                                         `json:"characterSet"`
		NcharacterSet      *string                                                         `json:"ncharacterSet"`
		ListenerPort       *int                                                            `json:"listenerPort"`
		OnsPortLocal       *int                                                            `json:"onsPortLocal"`
		OnsPortRemote      *int                                                            `json:"onsPortRemote"`
		DbDeploymentType   CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum  `json:"dbDeploymentType"`
		ShardDetails       []createdistributedautonomousdatabasesharddetails               `json:"shardDetails"`
		CatalogDetails     []createdistributedautonomousdatabasecatalogdetails             `json:"catalogDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Chunks = model.Chunks

	m.ListenerPortTls = model.ListenerPortTls

	m.ReplicationMethod = model.ReplicationMethod

	m.ReplicationFactor = model.ReplicationFactor

	m.ReplicationUnit = model.ReplicationUnit

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

	m.DbWorkload = model.DbWorkload

	m.CharacterSet = model.CharacterSet

	m.NcharacterSet = model.NcharacterSet

	m.ListenerPort = model.ListenerPort

	m.OnsPortLocal = model.OnsPortLocal

	m.OnsPortRemote = model.OnsPortRemote

	m.DbDeploymentType = model.DbDeploymentType

	m.ShardDetails = make([]CreateDistributedAutonomousDatabaseShardDetails, len(model.ShardDetails))
	for i, n := range model.ShardDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ShardDetails[i] = nn.(CreateDistributedAutonomousDatabaseShardDetails)
		} else {
			m.ShardDetails[i] = nil
		}
	}
	m.CatalogDetails = make([]CreateDistributedAutonomousDatabaseCatalogDetails, len(model.CatalogDetails))
	for i, n := range model.CatalogDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.CatalogDetails[i] = nn.(CreateDistributedAutonomousDatabaseCatalogDetails)
		} else {
			m.CatalogDetails[i] = nil
		}
	}
	return
}

// CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum Enum with underlying type: string
type CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum string

// Set of constants representing the allowable values for CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum
const (
	CreateDistributedAutonomousDatabaseDetailsShardingMethodUser   CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum = "USER"
	CreateDistributedAutonomousDatabaseDetailsShardingMethodSystem CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum = "SYSTEM"
)

var mappingCreateDistributedAutonomousDatabaseDetailsShardingMethodEnum = map[string]CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum{
	"USER":   CreateDistributedAutonomousDatabaseDetailsShardingMethodUser,
	"SYSTEM": CreateDistributedAutonomousDatabaseDetailsShardingMethodSystem,
}

var mappingCreateDistributedAutonomousDatabaseDetailsShardingMethodEnumLowerCase = map[string]CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum{
	"user":   CreateDistributedAutonomousDatabaseDetailsShardingMethodUser,
	"system": CreateDistributedAutonomousDatabaseDetailsShardingMethodSystem,
}

// GetCreateDistributedAutonomousDatabaseDetailsShardingMethodEnumValues Enumerates the set of values for CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum
func GetCreateDistributedAutonomousDatabaseDetailsShardingMethodEnumValues() []CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum {
	values := make([]CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum, 0)
	for _, v := range mappingCreateDistributedAutonomousDatabaseDetailsShardingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedAutonomousDatabaseDetailsShardingMethodEnumStringValues Enumerates the set of values in String for CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum
func GetCreateDistributedAutonomousDatabaseDetailsShardingMethodEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingCreateDistributedAutonomousDatabaseDetailsShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedAutonomousDatabaseDetailsShardingMethodEnum(val string) (CreateDistributedAutonomousDatabaseDetailsShardingMethodEnum, bool) {
	enum, ok := mappingCreateDistributedAutonomousDatabaseDetailsShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum Enum with underlying type: string
type CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum string

// Set of constants representing the allowable values for CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum
const (
	CreateDistributedAutonomousDatabaseDetailsDbWorkloadOltp CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum = "OLTP"
	CreateDistributedAutonomousDatabaseDetailsDbWorkloadDw   CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum = "DW"
)

var mappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum = map[string]CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum{
	"OLTP": CreateDistributedAutonomousDatabaseDetailsDbWorkloadOltp,
	"DW":   CreateDistributedAutonomousDatabaseDetailsDbWorkloadDw,
}

var mappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnumLowerCase = map[string]CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum{
	"oltp": CreateDistributedAutonomousDatabaseDetailsDbWorkloadOltp,
	"dw":   CreateDistributedAutonomousDatabaseDetailsDbWorkloadDw,
}

// GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnumValues Enumerates the set of values for CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum
func GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnumValues() []CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum {
	values := make([]CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum, 0)
	for _, v := range mappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnumStringValues Enumerates the set of values in String for CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum
func GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
	}
}

// GetMappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum(val string) (CreateDistributedAutonomousDatabaseDetailsDbWorkloadEnum, bool) {
	enum, ok := mappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum Enum with underlying type: string
type CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum string

// Set of constants representing the allowable values for CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum
const (
	CreateDistributedAutonomousDatabaseDetailsReplicationMethodRaft CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum = "RAFT"
	CreateDistributedAutonomousDatabaseDetailsReplicationMethodDg   CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum = "DG"
)

var mappingCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum = map[string]CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum{
	"RAFT": CreateDistributedAutonomousDatabaseDetailsReplicationMethodRaft,
	"DG":   CreateDistributedAutonomousDatabaseDetailsReplicationMethodDg,
}

var mappingCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnumLowerCase = map[string]CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum{
	"raft": CreateDistributedAutonomousDatabaseDetailsReplicationMethodRaft,
	"dg":   CreateDistributedAutonomousDatabaseDetailsReplicationMethodDg,
}

// GetCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnumValues Enumerates the set of values for CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum
func GetCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnumValues() []CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum {
	values := make([]CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum, 0)
	for _, v := range mappingCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnumStringValues Enumerates the set of values in String for CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum
func GetCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnumStringValues() []string {
	return []string{
		"RAFT",
		"DG",
	}
}

// GetMappingCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum(val string) (CreateDistributedAutonomousDatabaseDetailsReplicationMethodEnum, bool) {
	enum, ok := mappingCreateDistributedAutonomousDatabaseDetailsReplicationMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum Enum with underlying type: string
type CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum string

// Set of constants representing the allowable values for CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum
const (
	CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeAdbD CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum = "ADB_D"
)

var mappingCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum = map[string]CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum{
	"ADB_D": CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeAdbD,
}

var mappingCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnumLowerCase = map[string]CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum{
	"adb_d": CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeAdbD,
}

// GetCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnumValues Enumerates the set of values for CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum
func GetCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnumValues() []CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum {
	values := make([]CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum, 0)
	for _, v := range mappingCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnumStringValues Enumerates the set of values in String for CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum
func GetCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnumStringValues() []string {
	return []string{
		"ADB_D",
	}
}

// GetMappingCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum(val string) (CreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnum, bool) {
	enum, ok := mappingCreateDistributedAutonomousDatabaseDetailsDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
