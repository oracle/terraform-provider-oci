// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BdsInstance Description of the cluster.
type BdsInstance struct {

	// The OCID of the Big Data Service resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the cluster.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the cluster.
	LifecycleState BdsInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Boolean flag specifying whether or not the cluster is highly available (HA)
	IsHighAvailability *bool `mandatory:"true" json:"isHighAvailability"`

	// Boolean flag specifying whether or not the cluster should be set up as secure.
	IsSecure *bool `mandatory:"true" json:"isSecure"`

	// Boolean flag specifying whether or not Cloud SQL should be configured.
	IsCloudSqlConfigured *bool `mandatory:"true" json:"isCloudSqlConfigured"`

	// Boolean flag specifying whether or not Kafka should be configured.
	IsKafkaConfigured *bool `mandatory:"true" json:"isKafkaConfigured"`

	// The list of nodes in the cluster.
	Nodes []Node `mandatory:"true" json:"nodes"`

	// Number of nodes that forming the cluster
	NumberOfNodes *int `mandatory:"true" json:"numberOfNodes"`

	// Version of the Hadoop distribution.
	ClusterVersion BdsInstanceClusterVersionEnum `mandatory:"false" json:"clusterVersion,omitempty"`

	NetworkConfig *NetworkConfig `mandatory:"false" json:"networkConfig"`

	ClusterDetails *ClusterDetails `mandatory:"false" json:"clusterDetails"`

	CloudSqlDetails *CloudSqlDetails `mandatory:"false" json:"cloudSqlDetails"`

	// The user who created the cluster.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the cluster was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Number of nodes that require a maintenance reboot
	NumberOfNodesRequiringMaintenanceReboot *int `mandatory:"false" json:"numberOfNodesRequiringMaintenanceReboot"`

	// pre-authenticated URL of the bootstrap script in Object Store that can be downloaded and executed.
	BootstrapScriptUrl *string `mandatory:"false" json:"bootstrapScriptUrl"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Exists for cross-compatibility only. For example, `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example, `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID of the Key Management master encryption key.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// Profile of the Big Data Service cluster.
	ClusterProfile BdsInstanceClusterProfileEnum `mandatory:"false" json:"clusterProfile,omitempty"`
}

func (m BdsInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBdsInstanceClusterVersionEnum(string(m.ClusterVersion)); !ok && m.ClusterVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterVersion: %s. Supported values are: %s.", m.ClusterVersion, strings.Join(GetBdsInstanceClusterVersionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBdsInstanceClusterProfileEnum(string(m.ClusterProfile)); !ok && m.ClusterProfile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterProfile: %s. Supported values are: %s.", m.ClusterProfile, strings.Join(GetBdsInstanceClusterProfileEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsInstanceLifecycleStateEnum Enum with underlying type: string
type BdsInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for BdsInstanceLifecycleStateEnum
const (
	BdsInstanceLifecycleStateCreating   BdsInstanceLifecycleStateEnum = "CREATING"
	BdsInstanceLifecycleStateActive     BdsInstanceLifecycleStateEnum = "ACTIVE"
	BdsInstanceLifecycleStateUpdating   BdsInstanceLifecycleStateEnum = "UPDATING"
	BdsInstanceLifecycleStateSuspending BdsInstanceLifecycleStateEnum = "SUSPENDING"
	BdsInstanceLifecycleStateSuspended  BdsInstanceLifecycleStateEnum = "SUSPENDED"
	BdsInstanceLifecycleStateResuming   BdsInstanceLifecycleStateEnum = "RESUMING"
	BdsInstanceLifecycleStateDeleting   BdsInstanceLifecycleStateEnum = "DELETING"
	BdsInstanceLifecycleStateDeleted    BdsInstanceLifecycleStateEnum = "DELETED"
	BdsInstanceLifecycleStateFailed     BdsInstanceLifecycleStateEnum = "FAILED"
	BdsInstanceLifecycleStateInactive   BdsInstanceLifecycleStateEnum = "INACTIVE"
)

var mappingBdsInstanceLifecycleStateEnum = map[string]BdsInstanceLifecycleStateEnum{
	"CREATING":   BdsInstanceLifecycleStateCreating,
	"ACTIVE":     BdsInstanceLifecycleStateActive,
	"UPDATING":   BdsInstanceLifecycleStateUpdating,
	"SUSPENDING": BdsInstanceLifecycleStateSuspending,
	"SUSPENDED":  BdsInstanceLifecycleStateSuspended,
	"RESUMING":   BdsInstanceLifecycleStateResuming,
	"DELETING":   BdsInstanceLifecycleStateDeleting,
	"DELETED":    BdsInstanceLifecycleStateDeleted,
	"FAILED":     BdsInstanceLifecycleStateFailed,
	"INACTIVE":   BdsInstanceLifecycleStateInactive,
}

var mappingBdsInstanceLifecycleStateEnumLowerCase = map[string]BdsInstanceLifecycleStateEnum{
	"creating":   BdsInstanceLifecycleStateCreating,
	"active":     BdsInstanceLifecycleStateActive,
	"updating":   BdsInstanceLifecycleStateUpdating,
	"suspending": BdsInstanceLifecycleStateSuspending,
	"suspended":  BdsInstanceLifecycleStateSuspended,
	"resuming":   BdsInstanceLifecycleStateResuming,
	"deleting":   BdsInstanceLifecycleStateDeleting,
	"deleted":    BdsInstanceLifecycleStateDeleted,
	"failed":     BdsInstanceLifecycleStateFailed,
	"inactive":   BdsInstanceLifecycleStateInactive,
}

// GetBdsInstanceLifecycleStateEnumValues Enumerates the set of values for BdsInstanceLifecycleStateEnum
func GetBdsInstanceLifecycleStateEnumValues() []BdsInstanceLifecycleStateEnum {
	values := make([]BdsInstanceLifecycleStateEnum, 0)
	for _, v := range mappingBdsInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for BdsInstanceLifecycleStateEnum
func GetBdsInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"SUSPENDING",
		"SUSPENDED",
		"RESUMING",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingBdsInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsInstanceLifecycleStateEnum(val string) (BdsInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingBdsInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BdsInstanceClusterVersionEnum Enum with underlying type: string
type BdsInstanceClusterVersionEnum string

// Set of constants representing the allowable values for BdsInstanceClusterVersionEnum
const (
	BdsInstanceClusterVersionCdh5  BdsInstanceClusterVersionEnum = "CDH5"
	BdsInstanceClusterVersionCdh6  BdsInstanceClusterVersionEnum = "CDH6"
	BdsInstanceClusterVersionOdh1  BdsInstanceClusterVersionEnum = "ODH1"
	BdsInstanceClusterVersionOdh09 BdsInstanceClusterVersionEnum = "ODH0_9"
	BdsInstanceClusterVersionOdh20 BdsInstanceClusterVersionEnum = "ODH2_0"
)

var mappingBdsInstanceClusterVersionEnum = map[string]BdsInstanceClusterVersionEnum{
	"CDH5":   BdsInstanceClusterVersionCdh5,
	"CDH6":   BdsInstanceClusterVersionCdh6,
	"ODH1":   BdsInstanceClusterVersionOdh1,
	"ODH0_9": BdsInstanceClusterVersionOdh09,
	"ODH2_0": BdsInstanceClusterVersionOdh20,
}

var mappingBdsInstanceClusterVersionEnumLowerCase = map[string]BdsInstanceClusterVersionEnum{
	"cdh5":   BdsInstanceClusterVersionCdh5,
	"cdh6":   BdsInstanceClusterVersionCdh6,
	"odh1":   BdsInstanceClusterVersionOdh1,
	"odh0_9": BdsInstanceClusterVersionOdh09,
	"odh2_0": BdsInstanceClusterVersionOdh20,
}

// GetBdsInstanceClusterVersionEnumValues Enumerates the set of values for BdsInstanceClusterVersionEnum
func GetBdsInstanceClusterVersionEnumValues() []BdsInstanceClusterVersionEnum {
	values := make([]BdsInstanceClusterVersionEnum, 0)
	for _, v := range mappingBdsInstanceClusterVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsInstanceClusterVersionEnumStringValues Enumerates the set of values in String for BdsInstanceClusterVersionEnum
func GetBdsInstanceClusterVersionEnumStringValues() []string {
	return []string{
		"CDH5",
		"CDH6",
		"ODH1",
		"ODH0_9",
		"ODH2_0",
	}
}

// GetMappingBdsInstanceClusterVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsInstanceClusterVersionEnum(val string) (BdsInstanceClusterVersionEnum, bool) {
	enum, ok := mappingBdsInstanceClusterVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BdsInstanceClusterProfileEnum Enum with underlying type: string
type BdsInstanceClusterProfileEnum string

// Set of constants representing the allowable values for BdsInstanceClusterProfileEnum
const (
	BdsInstanceClusterProfileHadoopExtended BdsInstanceClusterProfileEnum = "HADOOP_EXTENDED"
	BdsInstanceClusterProfileHadoop         BdsInstanceClusterProfileEnum = "HADOOP"
	BdsInstanceClusterProfileHive           BdsInstanceClusterProfileEnum = "HIVE"
	BdsInstanceClusterProfileSpark          BdsInstanceClusterProfileEnum = "SPARK"
	BdsInstanceClusterProfileHbase          BdsInstanceClusterProfileEnum = "HBASE"
	BdsInstanceClusterProfileTrino          BdsInstanceClusterProfileEnum = "TRINO"
	BdsInstanceClusterProfileKafka          BdsInstanceClusterProfileEnum = "KAFKA"
)

var mappingBdsInstanceClusterProfileEnum = map[string]BdsInstanceClusterProfileEnum{
	"HADOOP_EXTENDED": BdsInstanceClusterProfileHadoopExtended,
	"HADOOP":          BdsInstanceClusterProfileHadoop,
	"HIVE":            BdsInstanceClusterProfileHive,
	"SPARK":           BdsInstanceClusterProfileSpark,
	"HBASE":           BdsInstanceClusterProfileHbase,
	"TRINO":           BdsInstanceClusterProfileTrino,
	"KAFKA":           BdsInstanceClusterProfileKafka,
}

var mappingBdsInstanceClusterProfileEnumLowerCase = map[string]BdsInstanceClusterProfileEnum{
	"hadoop_extended": BdsInstanceClusterProfileHadoopExtended,
	"hadoop":          BdsInstanceClusterProfileHadoop,
	"hive":            BdsInstanceClusterProfileHive,
	"spark":           BdsInstanceClusterProfileSpark,
	"hbase":           BdsInstanceClusterProfileHbase,
	"trino":           BdsInstanceClusterProfileTrino,
	"kafka":           BdsInstanceClusterProfileKafka,
}

// GetBdsInstanceClusterProfileEnumValues Enumerates the set of values for BdsInstanceClusterProfileEnum
func GetBdsInstanceClusterProfileEnumValues() []BdsInstanceClusterProfileEnum {
	values := make([]BdsInstanceClusterProfileEnum, 0)
	for _, v := range mappingBdsInstanceClusterProfileEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsInstanceClusterProfileEnumStringValues Enumerates the set of values in String for BdsInstanceClusterProfileEnum
func GetBdsInstanceClusterProfileEnumStringValues() []string {
	return []string{
		"HADOOP_EXTENDED",
		"HADOOP",
		"HIVE",
		"SPARK",
		"HBASE",
		"TRINO",
		"KAFKA",
	}
}

// GetMappingBdsInstanceClusterProfileEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsInstanceClusterProfileEnum(val string) (BdsInstanceClusterProfileEnum, bool) {
	enum, ok := mappingBdsInstanceClusterProfileEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
