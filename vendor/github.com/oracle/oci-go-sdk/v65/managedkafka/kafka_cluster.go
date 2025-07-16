// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KafkaCluster A KafkaCluster is a description of a KafkaCluster.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type KafkaCluster struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the KafkaCluster was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the KafkaCluster.
	LifecycleState KafkaClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Subnets where broker/coordinator VNICs will be created.
	AccessSubnets []SubnetSet `mandatory:"true" json:"accessSubnets"`

	// Version of Kafka to use to spin up the cluster
	KafkaVersion *string `mandatory:"true" json:"kafkaVersion"`

	// Type of the cluster to spin up.
	// DEVELOPMENT - setting that allows to sacrifice HA and spin up cluster on a single node
	// PRODUCTION - Minimum allowed broker count is 3
	ClusterType KafkaClusterClusterTypeEnum `mandatory:"true" json:"clusterType"`

	BrokerShape *BrokerShape `mandatory:"true" json:"brokerShape"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Kafka Cluster configuration object
	ClusterConfigId *string `mandatory:"true" json:"clusterConfigId"`

	// The version of configuration object
	ClusterConfigVersion *int `mandatory:"true" json:"clusterConfigVersion"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the KafkaCluster was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the KafkaCluster in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// CA certificate bundle for mTLS broker authentication.
	ClientCertificateBundle *string `mandatory:"false" json:"clientCertificateBundle"`

	// Bootstrap URL that can be used to connect to Kafka
	KafkaBootstrapUrls []BootstrapUrl `mandatory:"false" json:"kafkaBootstrapUrls"`

	// Kafka coordination type. Set of available types depends on Kafka version
	CoordinationType KafkaClusterCoordinationTypeEnum `mandatory:"false" json:"coordinationType,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret that contains superuser password.
	SecretId *string `mandatory:"false" json:"secretId"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m KafkaCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KafkaCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetKafkaClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKafkaClusterClusterTypeEnum(string(m.ClusterType)); !ok && m.ClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterType: %s. Supported values are: %s.", m.ClusterType, strings.Join(GetKafkaClusterClusterTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingKafkaClusterCoordinationTypeEnum(string(m.CoordinationType)); !ok && m.CoordinationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CoordinationType: %s. Supported values are: %s.", m.CoordinationType, strings.Join(GetKafkaClusterCoordinationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KafkaClusterLifecycleStateEnum Enum with underlying type: string
type KafkaClusterLifecycleStateEnum string

// Set of constants representing the allowable values for KafkaClusterLifecycleStateEnum
const (
	KafkaClusterLifecycleStateCreating KafkaClusterLifecycleStateEnum = "CREATING"
	KafkaClusterLifecycleStateUpdating KafkaClusterLifecycleStateEnum = "UPDATING"
	KafkaClusterLifecycleStateActive   KafkaClusterLifecycleStateEnum = "ACTIVE"
	KafkaClusterLifecycleStateDeleting KafkaClusterLifecycleStateEnum = "DELETING"
	KafkaClusterLifecycleStateDeleted  KafkaClusterLifecycleStateEnum = "DELETED"
	KafkaClusterLifecycleStateFailed   KafkaClusterLifecycleStateEnum = "FAILED"
)

var mappingKafkaClusterLifecycleStateEnum = map[string]KafkaClusterLifecycleStateEnum{
	"CREATING": KafkaClusterLifecycleStateCreating,
	"UPDATING": KafkaClusterLifecycleStateUpdating,
	"ACTIVE":   KafkaClusterLifecycleStateActive,
	"DELETING": KafkaClusterLifecycleStateDeleting,
	"DELETED":  KafkaClusterLifecycleStateDeleted,
	"FAILED":   KafkaClusterLifecycleStateFailed,
}

var mappingKafkaClusterLifecycleStateEnumLowerCase = map[string]KafkaClusterLifecycleStateEnum{
	"creating": KafkaClusterLifecycleStateCreating,
	"updating": KafkaClusterLifecycleStateUpdating,
	"active":   KafkaClusterLifecycleStateActive,
	"deleting": KafkaClusterLifecycleStateDeleting,
	"deleted":  KafkaClusterLifecycleStateDeleted,
	"failed":   KafkaClusterLifecycleStateFailed,
}

// GetKafkaClusterLifecycleStateEnumValues Enumerates the set of values for KafkaClusterLifecycleStateEnum
func GetKafkaClusterLifecycleStateEnumValues() []KafkaClusterLifecycleStateEnum {
	values := make([]KafkaClusterLifecycleStateEnum, 0)
	for _, v := range mappingKafkaClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaClusterLifecycleStateEnumStringValues Enumerates the set of values in String for KafkaClusterLifecycleStateEnum
func GetKafkaClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingKafkaClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaClusterLifecycleStateEnum(val string) (KafkaClusterLifecycleStateEnum, bool) {
	enum, ok := mappingKafkaClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// KafkaClusterClusterTypeEnum Enum with underlying type: string
type KafkaClusterClusterTypeEnum string

// Set of constants representing the allowable values for KafkaClusterClusterTypeEnum
const (
	KafkaClusterClusterTypeDevelopment KafkaClusterClusterTypeEnum = "DEVELOPMENT"
	KafkaClusterClusterTypeProduction  KafkaClusterClusterTypeEnum = "PRODUCTION"
)

var mappingKafkaClusterClusterTypeEnum = map[string]KafkaClusterClusterTypeEnum{
	"DEVELOPMENT": KafkaClusterClusterTypeDevelopment,
	"PRODUCTION":  KafkaClusterClusterTypeProduction,
}

var mappingKafkaClusterClusterTypeEnumLowerCase = map[string]KafkaClusterClusterTypeEnum{
	"development": KafkaClusterClusterTypeDevelopment,
	"production":  KafkaClusterClusterTypeProduction,
}

// GetKafkaClusterClusterTypeEnumValues Enumerates the set of values for KafkaClusterClusterTypeEnum
func GetKafkaClusterClusterTypeEnumValues() []KafkaClusterClusterTypeEnum {
	values := make([]KafkaClusterClusterTypeEnum, 0)
	for _, v := range mappingKafkaClusterClusterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaClusterClusterTypeEnumStringValues Enumerates the set of values in String for KafkaClusterClusterTypeEnum
func GetKafkaClusterClusterTypeEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"PRODUCTION",
	}
}

// GetMappingKafkaClusterClusterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaClusterClusterTypeEnum(val string) (KafkaClusterClusterTypeEnum, bool) {
	enum, ok := mappingKafkaClusterClusterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// KafkaClusterCoordinationTypeEnum Enum with underlying type: string
type KafkaClusterCoordinationTypeEnum string

// Set of constants representing the allowable values for KafkaClusterCoordinationTypeEnum
const (
	KafkaClusterCoordinationTypeZookeeper KafkaClusterCoordinationTypeEnum = "ZOOKEEPER"
	KafkaClusterCoordinationTypeKraft     KafkaClusterCoordinationTypeEnum = "KRAFT"
)

var mappingKafkaClusterCoordinationTypeEnum = map[string]KafkaClusterCoordinationTypeEnum{
	"ZOOKEEPER": KafkaClusterCoordinationTypeZookeeper,
	"KRAFT":     KafkaClusterCoordinationTypeKraft,
}

var mappingKafkaClusterCoordinationTypeEnumLowerCase = map[string]KafkaClusterCoordinationTypeEnum{
	"zookeeper": KafkaClusterCoordinationTypeZookeeper,
	"kraft":     KafkaClusterCoordinationTypeKraft,
}

// GetKafkaClusterCoordinationTypeEnumValues Enumerates the set of values for KafkaClusterCoordinationTypeEnum
func GetKafkaClusterCoordinationTypeEnumValues() []KafkaClusterCoordinationTypeEnum {
	values := make([]KafkaClusterCoordinationTypeEnum, 0)
	for _, v := range mappingKafkaClusterCoordinationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaClusterCoordinationTypeEnumStringValues Enumerates the set of values in String for KafkaClusterCoordinationTypeEnum
func GetKafkaClusterCoordinationTypeEnumStringValues() []string {
	return []string{
		"ZOOKEEPER",
		"KRAFT",
	}
}

// GetMappingKafkaClusterCoordinationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaClusterCoordinationTypeEnum(val string) (KafkaClusterCoordinationTypeEnum, bool) {
	enum, ok := mappingKafkaClusterCoordinationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
