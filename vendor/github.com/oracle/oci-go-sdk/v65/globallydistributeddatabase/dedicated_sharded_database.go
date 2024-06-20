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

// DedicatedShardedDatabase Details of ATP-D based sharded database.
type DedicatedShardedDatabase struct {

	// Sharded Database identifier
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment in which sharded database exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Oracle sharded database display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the the Sharded Database was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Sharded Database was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Detailed message for the lifecycle state.
	LifecycleStateDetails *string `mandatory:"true" json:"lifecycleStateDetails"`

	// The character set for the database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// Oracle Database version number.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// Unique prefix for the sharded database.
	Prefix *string `mandatory:"true" json:"prefix"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

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

	// The GSM listener port number.
	ListenerPort *int `mandatory:"false" json:"listenerPort"`

	// The TLS listener port number for sharded database.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// Ons local port number.
	OnsPortLocal *int `mandatory:"false" json:"onsPortLocal"`

	// Ons remote port number.
	OnsPortRemote *int `mandatory:"false" json:"onsPortRemote"`

	// The OCID of private endpoint being used by the sharded database.
	PrivateEndpoint *string `mandatory:"false" json:"privateEndpoint"`

	ConnectionStrings *ConnectionString `mandatory:"false" json:"connectionStrings"`

	// Timezone associated with the sharded database.
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// Details of GSM instances for the sharded database.
	Gsms []GsmDetails `mandatory:"false" json:"gsms"`

	// Details of ATP-D based shards.
	ShardDetails []DedicatedShardDetails `mandatory:"false" json:"shardDetails"`

	// Details of ATP-D based catalogs.
	CatalogDetails []DedicatedCatalogDetails `mandatory:"false" json:"catalogDetails"`

	// The Replication method for sharded database. Use RAFT for Raft replication, and DG for
	// DataGuard. If replicationMethod is not provided, it defaults to DG.
	ReplicationMethod DedicatedShardedDatabaseReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`

	// Possible workload types.
	DbWorkload DedicatedShardedDatabaseDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// Sharding Method.
	ShardingMethod DedicatedShardedDatabaseShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

	// Lifecycle states for sharded databases.
	LifecycleState ShardedDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m DedicatedShardedDatabase) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m DedicatedShardedDatabase) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m DedicatedShardedDatabase) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m DedicatedShardedDatabase) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DedicatedShardedDatabase) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m DedicatedShardedDatabase) GetLifecycleState() ShardedDatabaseLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleStateDetails returns LifecycleStateDetails
func (m DedicatedShardedDatabase) GetLifecycleStateDetails() *string {
	return m.LifecycleStateDetails
}

// GetFreeformTags returns FreeformTags
func (m DedicatedShardedDatabase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m DedicatedShardedDatabase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m DedicatedShardedDatabase) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DedicatedShardedDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedShardedDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedShardedDatabaseReplicationMethodEnum(string(m.ReplicationMethod)); !ok && m.ReplicationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationMethod: %s. Supported values are: %s.", m.ReplicationMethod, strings.Join(GetDedicatedShardedDatabaseReplicationMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDedicatedShardedDatabaseDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetDedicatedShardedDatabaseDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDedicatedShardedDatabaseShardingMethodEnum(string(m.ShardingMethod)); !ok && m.ShardingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShardingMethod: %s. Supported values are: %s.", m.ShardingMethod, strings.Join(GetDedicatedShardedDatabaseShardingMethodEnumStringValues(), ",")))
	}

	if _, ok := GetMappingShardedDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShardedDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DedicatedShardedDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDedicatedShardedDatabase DedicatedShardedDatabase
	s := struct {
		DiscriminatorParam string `json:"dbDeploymentType"`
		MarshalTypeDedicatedShardedDatabase
	}{
		"DEDICATED",
		(MarshalTypeDedicatedShardedDatabase)(m),
	}

	return json.Marshal(&s)
}

// DedicatedShardedDatabaseReplicationMethodEnum Enum with underlying type: string
type DedicatedShardedDatabaseReplicationMethodEnum string

// Set of constants representing the allowable values for DedicatedShardedDatabaseReplicationMethodEnum
const (
	DedicatedShardedDatabaseReplicationMethodRaft DedicatedShardedDatabaseReplicationMethodEnum = "RAFT"
	DedicatedShardedDatabaseReplicationMethodDg   DedicatedShardedDatabaseReplicationMethodEnum = "DG"
)

var mappingDedicatedShardedDatabaseReplicationMethodEnum = map[string]DedicatedShardedDatabaseReplicationMethodEnum{
	"RAFT": DedicatedShardedDatabaseReplicationMethodRaft,
	"DG":   DedicatedShardedDatabaseReplicationMethodDg,
}

var mappingDedicatedShardedDatabaseReplicationMethodEnumLowerCase = map[string]DedicatedShardedDatabaseReplicationMethodEnum{
	"raft": DedicatedShardedDatabaseReplicationMethodRaft,
	"dg":   DedicatedShardedDatabaseReplicationMethodDg,
}

// GetDedicatedShardedDatabaseReplicationMethodEnumValues Enumerates the set of values for DedicatedShardedDatabaseReplicationMethodEnum
func GetDedicatedShardedDatabaseReplicationMethodEnumValues() []DedicatedShardedDatabaseReplicationMethodEnum {
	values := make([]DedicatedShardedDatabaseReplicationMethodEnum, 0)
	for _, v := range mappingDedicatedShardedDatabaseReplicationMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedShardedDatabaseReplicationMethodEnumStringValues Enumerates the set of values in String for DedicatedShardedDatabaseReplicationMethodEnum
func GetDedicatedShardedDatabaseReplicationMethodEnumStringValues() []string {
	return []string{
		"RAFT",
		"DG",
	}
}

// GetMappingDedicatedShardedDatabaseReplicationMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedShardedDatabaseReplicationMethodEnum(val string) (DedicatedShardedDatabaseReplicationMethodEnum, bool) {
	enum, ok := mappingDedicatedShardedDatabaseReplicationMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DedicatedShardedDatabaseDbWorkloadEnum Enum with underlying type: string
type DedicatedShardedDatabaseDbWorkloadEnum string

// Set of constants representing the allowable values for DedicatedShardedDatabaseDbWorkloadEnum
const (
	DedicatedShardedDatabaseDbWorkloadOltp DedicatedShardedDatabaseDbWorkloadEnum = "OLTP"
	DedicatedShardedDatabaseDbWorkloadDw   DedicatedShardedDatabaseDbWorkloadEnum = "DW"
)

var mappingDedicatedShardedDatabaseDbWorkloadEnum = map[string]DedicatedShardedDatabaseDbWorkloadEnum{
	"OLTP": DedicatedShardedDatabaseDbWorkloadOltp,
	"DW":   DedicatedShardedDatabaseDbWorkloadDw,
}

var mappingDedicatedShardedDatabaseDbWorkloadEnumLowerCase = map[string]DedicatedShardedDatabaseDbWorkloadEnum{
	"oltp": DedicatedShardedDatabaseDbWorkloadOltp,
	"dw":   DedicatedShardedDatabaseDbWorkloadDw,
}

// GetDedicatedShardedDatabaseDbWorkloadEnumValues Enumerates the set of values for DedicatedShardedDatabaseDbWorkloadEnum
func GetDedicatedShardedDatabaseDbWorkloadEnumValues() []DedicatedShardedDatabaseDbWorkloadEnum {
	values := make([]DedicatedShardedDatabaseDbWorkloadEnum, 0)
	for _, v := range mappingDedicatedShardedDatabaseDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedShardedDatabaseDbWorkloadEnumStringValues Enumerates the set of values in String for DedicatedShardedDatabaseDbWorkloadEnum
func GetDedicatedShardedDatabaseDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
	}
}

// GetMappingDedicatedShardedDatabaseDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedShardedDatabaseDbWorkloadEnum(val string) (DedicatedShardedDatabaseDbWorkloadEnum, bool) {
	enum, ok := mappingDedicatedShardedDatabaseDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DedicatedShardedDatabaseShardingMethodEnum Enum with underlying type: string
type DedicatedShardedDatabaseShardingMethodEnum string

// Set of constants representing the allowable values for DedicatedShardedDatabaseShardingMethodEnum
const (
	DedicatedShardedDatabaseShardingMethodUser   DedicatedShardedDatabaseShardingMethodEnum = "USER"
	DedicatedShardedDatabaseShardingMethodSystem DedicatedShardedDatabaseShardingMethodEnum = "SYSTEM"
)

var mappingDedicatedShardedDatabaseShardingMethodEnum = map[string]DedicatedShardedDatabaseShardingMethodEnum{
	"USER":   DedicatedShardedDatabaseShardingMethodUser,
	"SYSTEM": DedicatedShardedDatabaseShardingMethodSystem,
}

var mappingDedicatedShardedDatabaseShardingMethodEnumLowerCase = map[string]DedicatedShardedDatabaseShardingMethodEnum{
	"user":   DedicatedShardedDatabaseShardingMethodUser,
	"system": DedicatedShardedDatabaseShardingMethodSystem,
}

// GetDedicatedShardedDatabaseShardingMethodEnumValues Enumerates the set of values for DedicatedShardedDatabaseShardingMethodEnum
func GetDedicatedShardedDatabaseShardingMethodEnumValues() []DedicatedShardedDatabaseShardingMethodEnum {
	values := make([]DedicatedShardedDatabaseShardingMethodEnum, 0)
	for _, v := range mappingDedicatedShardedDatabaseShardingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedShardedDatabaseShardingMethodEnumStringValues Enumerates the set of values in String for DedicatedShardedDatabaseShardingMethodEnum
func GetDedicatedShardedDatabaseShardingMethodEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingDedicatedShardedDatabaseShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedShardedDatabaseShardingMethodEnum(val string) (DedicatedShardedDatabaseShardingMethodEnum, bool) {
	enum, ok := mappingDedicatedShardedDatabaseShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
