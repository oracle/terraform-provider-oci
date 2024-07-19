// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDedicatedShardedDatabase Request details for creation of ATP-Dedicated based sharded database.
type CreateDedicatedShardedDatabase struct {

	// Identifier of the compartment where sharded database is to be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Oracle sharded database display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Oracle Database version of the Autonomous Container Database.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The character set for the new shard database being created. Use database api ListAutonomousDatabaseCharacterSets to
	// get the list of allowed character set for autonomous dedicated database. See documentation:
	// https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/AutonomousDatabaseCharacterSets/ListAutonomousDatabaseCharacterSets
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the new shard database being created. Use database api ListAutonomousDatabaseCharacterSets to
	// get the list of allowed national character set for autonomous dedicated database. See documentation:
	// https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/AutonomousDatabaseCharacterSets/ListAutonomousDatabaseCharacterSets
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// The listener port number for sharded database.
	ListenerPort *int `mandatory:"true" json:"listenerPort"`

	// The TLS listener port number for sharded database.
	ListenerPortTls *int `mandatory:"true" json:"listenerPortTls"`

	// Ons port local for sharded database.
	OnsPortLocal *int `mandatory:"true" json:"onsPortLocal"`

	// Ons remote port for sharded database.
	OnsPortRemote *int `mandatory:"true" json:"onsPortRemote"`

	// Unique name prefix for the sharded databases. Only alpha-numeric values are allowed. First character
	// has to be a letter followed by any combination of letter and number.
	Prefix *string `mandatory:"true" json:"prefix"`

	// Collection of ATP-Dedicated shards that needs to be created.
	ShardDetails []CreateDedicatedShardDetail `mandatory:"true" json:"shardDetails"`

	// Collection of ATP-Dedicated catalogs that needs to be created.
	CatalogDetails []CreateDedicatedCatalogDetail `mandatory:"true" json:"catalogDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Replication factor for RAFT replication based sharded database. Currently supported values are 3, 5 and 7.
	ReplicationFactor *int `mandatory:"false" json:"replicationFactor"`

	// For RAFT replication based sharded database, the value should be atleast twice the number of shards.
	ReplicationUnit *int `mandatory:"false" json:"replicationUnit"`

	// The certificate common name used in all cloudAutonomousVmClusters for the sharded database topology. Eg. Production.
	// All the clusters used in one sharded database topology shall have same CABundle setup. Valid characterset for
	// clusterCertificateCommonName include uppercase or lowercase letters, numbers, hyphens, underscores, and period.
	ClusterCertificateCommonName *string `mandatory:"false" json:"clusterCertificateCommonName"`

	// The default number of unique chunks in a shardspace. The value of chunks must be
	// greater than 2 times the size of the largest shardgroup in any shardspace.
	Chunks *int `mandatory:"false" json:"chunks"`

	// Possible workload types.
	DbWorkload CreateDedicatedShardedDatabaseDbWorkloadEnum `mandatory:"true" json:"dbWorkload"`

	// Sharding Method.
	ShardingMethod CreateDedicatedShardedDatabaseShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

	// The Replication method for sharded database.
	ReplicationMethod DedicatedShardedDatabaseReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`
}

// GetCompartmentId returns CompartmentId
func (m CreateDedicatedShardedDatabase) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateDedicatedShardedDatabase) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m CreateDedicatedShardedDatabase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateDedicatedShardedDatabase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateDedicatedShardedDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDedicatedShardedDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDedicatedShardedDatabaseDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetCreateDedicatedShardedDatabaseDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDedicatedShardedDatabaseShardingMethodEnum(string(m.ShardingMethod)); !ok && m.ShardingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShardingMethod: %s. Supported values are: %s.", m.ShardingMethod, strings.Join(GetCreateDedicatedShardedDatabaseShardingMethodEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDedicatedShardedDatabaseReplicationMethodEnum(string(m.ReplicationMethod)); !ok && m.ReplicationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationMethod: %s. Supported values are: %s.", m.ReplicationMethod, strings.Join(GetDedicatedShardedDatabaseReplicationMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDedicatedShardedDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDedicatedShardedDatabase CreateDedicatedShardedDatabase
	s := struct {
		DiscriminatorParam string `json:"dbDeploymentType"`
		MarshalTypeCreateDedicatedShardedDatabase
	}{
		"DEDICATED",
		(MarshalTypeCreateDedicatedShardedDatabase)(m),
	}

	return json.Marshal(&s)
}

// CreateDedicatedShardedDatabaseDbWorkloadEnum Enum with underlying type: string
type CreateDedicatedShardedDatabaseDbWorkloadEnum string

// Set of constants representing the allowable values for CreateDedicatedShardedDatabaseDbWorkloadEnum
const (
	CreateDedicatedShardedDatabaseDbWorkloadOltp CreateDedicatedShardedDatabaseDbWorkloadEnum = "OLTP"
	CreateDedicatedShardedDatabaseDbWorkloadDw   CreateDedicatedShardedDatabaseDbWorkloadEnum = "DW"
)

var mappingCreateDedicatedShardedDatabaseDbWorkloadEnum = map[string]CreateDedicatedShardedDatabaseDbWorkloadEnum{
	"OLTP": CreateDedicatedShardedDatabaseDbWorkloadOltp,
	"DW":   CreateDedicatedShardedDatabaseDbWorkloadDw,
}

var mappingCreateDedicatedShardedDatabaseDbWorkloadEnumLowerCase = map[string]CreateDedicatedShardedDatabaseDbWorkloadEnum{
	"oltp": CreateDedicatedShardedDatabaseDbWorkloadOltp,
	"dw":   CreateDedicatedShardedDatabaseDbWorkloadDw,
}

// GetCreateDedicatedShardedDatabaseDbWorkloadEnumValues Enumerates the set of values for CreateDedicatedShardedDatabaseDbWorkloadEnum
func GetCreateDedicatedShardedDatabaseDbWorkloadEnumValues() []CreateDedicatedShardedDatabaseDbWorkloadEnum {
	values := make([]CreateDedicatedShardedDatabaseDbWorkloadEnum, 0)
	for _, v := range mappingCreateDedicatedShardedDatabaseDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDedicatedShardedDatabaseDbWorkloadEnumStringValues Enumerates the set of values in String for CreateDedicatedShardedDatabaseDbWorkloadEnum
func GetCreateDedicatedShardedDatabaseDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
	}
}

// GetMappingCreateDedicatedShardedDatabaseDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDedicatedShardedDatabaseDbWorkloadEnum(val string) (CreateDedicatedShardedDatabaseDbWorkloadEnum, bool) {
	enum, ok := mappingCreateDedicatedShardedDatabaseDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDedicatedShardedDatabaseShardingMethodEnum Enum with underlying type: string
type CreateDedicatedShardedDatabaseShardingMethodEnum string

// Set of constants representing the allowable values for CreateDedicatedShardedDatabaseShardingMethodEnum
const (
	CreateDedicatedShardedDatabaseShardingMethodUser   CreateDedicatedShardedDatabaseShardingMethodEnum = "USER"
	CreateDedicatedShardedDatabaseShardingMethodSystem CreateDedicatedShardedDatabaseShardingMethodEnum = "SYSTEM"
)

var mappingCreateDedicatedShardedDatabaseShardingMethodEnum = map[string]CreateDedicatedShardedDatabaseShardingMethodEnum{
	"USER":   CreateDedicatedShardedDatabaseShardingMethodUser,
	"SYSTEM": CreateDedicatedShardedDatabaseShardingMethodSystem,
}

var mappingCreateDedicatedShardedDatabaseShardingMethodEnumLowerCase = map[string]CreateDedicatedShardedDatabaseShardingMethodEnum{
	"user":   CreateDedicatedShardedDatabaseShardingMethodUser,
	"system": CreateDedicatedShardedDatabaseShardingMethodSystem,
}

// GetCreateDedicatedShardedDatabaseShardingMethodEnumValues Enumerates the set of values for CreateDedicatedShardedDatabaseShardingMethodEnum
func GetCreateDedicatedShardedDatabaseShardingMethodEnumValues() []CreateDedicatedShardedDatabaseShardingMethodEnum {
	values := make([]CreateDedicatedShardedDatabaseShardingMethodEnum, 0)
	for _, v := range mappingCreateDedicatedShardedDatabaseShardingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDedicatedShardedDatabaseShardingMethodEnumStringValues Enumerates the set of values in String for CreateDedicatedShardedDatabaseShardingMethodEnum
func GetCreateDedicatedShardedDatabaseShardingMethodEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingCreateDedicatedShardedDatabaseShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDedicatedShardedDatabaseShardingMethodEnum(val string) (CreateDedicatedShardedDatabaseShardingMethodEnum, bool) {
	enum, ok := mappingCreateDedicatedShardedDatabaseShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
