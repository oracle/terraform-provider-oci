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

// DedicatedShardedDatabaseSummary Summary of ATP-D based sharded database.
type DedicatedShardedDatabaseSummary struct {

	// Sharded Database identifier
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment where sharded database exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Oracle sharded database display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the the Sharded Database was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Sharded Database was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Detailed message for the lifecycle state.
	LifecycleStateDetails *string `mandatory:"true" json:"lifecycleStateDetails"`

	// The character set for the sharded database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the sharded database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

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

	// Oracle Database version of the Autonomous Container Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The listener port number for the sharded database.
	ListenerPort *int `mandatory:"false" json:"listenerPort"`

	// The TLS listener port number for sharded database.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// Ons local port number.
	OnsPortLocal *int `mandatory:"false" json:"onsPortLocal"`

	// Ons remote port number.
	OnsPortRemote *int `mandatory:"false" json:"onsPortRemote"`

	// Name prefix for the sharded databases.
	Prefix *string `mandatory:"false" json:"prefix"`

	// Total cpu count usage for shards and catalogs of the sharded database.
	TotalCpuCount *int `mandatory:"false" json:"totalCpuCount"`

	// The aggregarted value of dataStorageSizeInGbs for all shards and catalogs.
	TotalDataStorageSizeInGbs *float64 `mandatory:"false" json:"totalDataStorageSizeInGbs"`

	// Possible workload types.
	DbWorkload DedicatedShardedDatabaseSummaryDbWorkloadEnum `mandatory:"true" json:"dbWorkload"`

	// Sharding Method.
	ShardingMethod DedicatedShardedDatabaseSummaryShardingMethodEnum `mandatory:"true" json:"shardingMethod"`

	// Lifecycle state of sharded database.
	LifecycleState ShardedDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Replication method for sharded database.
	ReplicationMethod DedicatedShardedDatabaseReplicationMethodEnum `mandatory:"false" json:"replicationMethod,omitempty"`
}

// GetId returns Id
func (m DedicatedShardedDatabaseSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m DedicatedShardedDatabaseSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m DedicatedShardedDatabaseSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m DedicatedShardedDatabaseSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DedicatedShardedDatabaseSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m DedicatedShardedDatabaseSummary) GetLifecycleState() ShardedDatabaseLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleStateDetails returns LifecycleStateDetails
func (m DedicatedShardedDatabaseSummary) GetLifecycleStateDetails() *string {
	return m.LifecycleStateDetails
}

// GetFreeformTags returns FreeformTags
func (m DedicatedShardedDatabaseSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m DedicatedShardedDatabaseSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m DedicatedShardedDatabaseSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DedicatedShardedDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedShardedDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedShardedDatabaseSummaryDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetDedicatedShardedDatabaseSummaryDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDedicatedShardedDatabaseSummaryShardingMethodEnum(string(m.ShardingMethod)); !ok && m.ShardingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShardingMethod: %s. Supported values are: %s.", m.ShardingMethod, strings.Join(GetDedicatedShardedDatabaseSummaryShardingMethodEnumStringValues(), ",")))
	}

	if _, ok := GetMappingShardedDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShardedDatabaseLifecycleStateEnumStringValues(), ",")))
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
func (m DedicatedShardedDatabaseSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDedicatedShardedDatabaseSummary DedicatedShardedDatabaseSummary
	s := struct {
		DiscriminatorParam string `json:"dbDeploymentType"`
		MarshalTypeDedicatedShardedDatabaseSummary
	}{
		"DEDICATED",
		(MarshalTypeDedicatedShardedDatabaseSummary)(m),
	}

	return json.Marshal(&s)
}

// DedicatedShardedDatabaseSummaryDbWorkloadEnum Enum with underlying type: string
type DedicatedShardedDatabaseSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for DedicatedShardedDatabaseSummaryDbWorkloadEnum
const (
	DedicatedShardedDatabaseSummaryDbWorkloadOltp DedicatedShardedDatabaseSummaryDbWorkloadEnum = "OLTP"
	DedicatedShardedDatabaseSummaryDbWorkloadDw   DedicatedShardedDatabaseSummaryDbWorkloadEnum = "DW"
)

var mappingDedicatedShardedDatabaseSummaryDbWorkloadEnum = map[string]DedicatedShardedDatabaseSummaryDbWorkloadEnum{
	"OLTP": DedicatedShardedDatabaseSummaryDbWorkloadOltp,
	"DW":   DedicatedShardedDatabaseSummaryDbWorkloadDw,
}

var mappingDedicatedShardedDatabaseSummaryDbWorkloadEnumLowerCase = map[string]DedicatedShardedDatabaseSummaryDbWorkloadEnum{
	"oltp": DedicatedShardedDatabaseSummaryDbWorkloadOltp,
	"dw":   DedicatedShardedDatabaseSummaryDbWorkloadDw,
}

// GetDedicatedShardedDatabaseSummaryDbWorkloadEnumValues Enumerates the set of values for DedicatedShardedDatabaseSummaryDbWorkloadEnum
func GetDedicatedShardedDatabaseSummaryDbWorkloadEnumValues() []DedicatedShardedDatabaseSummaryDbWorkloadEnum {
	values := make([]DedicatedShardedDatabaseSummaryDbWorkloadEnum, 0)
	for _, v := range mappingDedicatedShardedDatabaseSummaryDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedShardedDatabaseSummaryDbWorkloadEnumStringValues Enumerates the set of values in String for DedicatedShardedDatabaseSummaryDbWorkloadEnum
func GetDedicatedShardedDatabaseSummaryDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
	}
}

// GetMappingDedicatedShardedDatabaseSummaryDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedShardedDatabaseSummaryDbWorkloadEnum(val string) (DedicatedShardedDatabaseSummaryDbWorkloadEnum, bool) {
	enum, ok := mappingDedicatedShardedDatabaseSummaryDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DedicatedShardedDatabaseSummaryShardingMethodEnum Enum with underlying type: string
type DedicatedShardedDatabaseSummaryShardingMethodEnum string

// Set of constants representing the allowable values for DedicatedShardedDatabaseSummaryShardingMethodEnum
const (
	DedicatedShardedDatabaseSummaryShardingMethodUser   DedicatedShardedDatabaseSummaryShardingMethodEnum = "USER"
	DedicatedShardedDatabaseSummaryShardingMethodSystem DedicatedShardedDatabaseSummaryShardingMethodEnum = "SYSTEM"
)

var mappingDedicatedShardedDatabaseSummaryShardingMethodEnum = map[string]DedicatedShardedDatabaseSummaryShardingMethodEnum{
	"USER":   DedicatedShardedDatabaseSummaryShardingMethodUser,
	"SYSTEM": DedicatedShardedDatabaseSummaryShardingMethodSystem,
}

var mappingDedicatedShardedDatabaseSummaryShardingMethodEnumLowerCase = map[string]DedicatedShardedDatabaseSummaryShardingMethodEnum{
	"user":   DedicatedShardedDatabaseSummaryShardingMethodUser,
	"system": DedicatedShardedDatabaseSummaryShardingMethodSystem,
}

// GetDedicatedShardedDatabaseSummaryShardingMethodEnumValues Enumerates the set of values for DedicatedShardedDatabaseSummaryShardingMethodEnum
func GetDedicatedShardedDatabaseSummaryShardingMethodEnumValues() []DedicatedShardedDatabaseSummaryShardingMethodEnum {
	values := make([]DedicatedShardedDatabaseSummaryShardingMethodEnum, 0)
	for _, v := range mappingDedicatedShardedDatabaseSummaryShardingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedShardedDatabaseSummaryShardingMethodEnumStringValues Enumerates the set of values in String for DedicatedShardedDatabaseSummaryShardingMethodEnum
func GetDedicatedShardedDatabaseSummaryShardingMethodEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingDedicatedShardedDatabaseSummaryShardingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedShardedDatabaseSummaryShardingMethodEnum(val string) (DedicatedShardedDatabaseSummaryShardingMethodEnum, bool) {
	enum, ok := mappingDedicatedShardedDatabaseSummaryShardingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
