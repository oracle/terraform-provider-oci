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

// CreateKafkaClusterDetails The data to create a KafkaCluster.
type CreateKafkaClusterDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the KafkaCluster in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Subnets where broker/coordinator VNICs will be created.
	AccessSubnets []SubnetSet `mandatory:"true" json:"accessSubnets"`

	// Version of Kafka to use to spin up the cluster
	KafkaVersion *string `mandatory:"true" json:"kafkaVersion"`

	// Type of the cluster to spin up.
	// DEVELOPMENT - setting that allows to sacrifice HA and spin up cluster on single node
	// PRODUCTION - Minimum allowed broker count is 3
	ClusterType KafkaClusterClusterTypeEnum `mandatory:"true" json:"clusterType"`

	BrokerShape *BrokerShape `mandatory:"true" json:"brokerShape"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Kafka Cluster configuration object
	ClusterConfigId *string `mandatory:"true" json:"clusterConfigId"`

	// The version of configuration object
	ClusterConfigVersion *int `mandatory:"true" json:"clusterConfigVersion"`

	// Kafka coordination type. Set of available types depends on Kafka version
	CoordinationType KafkaClusterCoordinationTypeEnum `mandatory:"true" json:"coordinationType"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// CA certificate bundle for mTLS broker authentication.
	ClientCertificateBundle *string `mandatory:"false" json:"clientCertificateBundle"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateKafkaClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateKafkaClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaClusterClusterTypeEnum(string(m.ClusterType)); !ok && m.ClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterType: %s. Supported values are: %s.", m.ClusterType, strings.Join(GetKafkaClusterClusterTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKafkaClusterCoordinationTypeEnum(string(m.CoordinationType)); !ok && m.CoordinationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CoordinationType: %s. Supported values are: %s.", m.CoordinationType, strings.Join(GetKafkaClusterCoordinationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
