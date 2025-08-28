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

// UpdateKafkaClusterDetails The data to update a KafkaCluster.
type UpdateKafkaClusterDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// CA certificate bundle for mTLS broker authentication.
	ClientCertificateBundle *string `mandatory:"false" json:"clientCertificateBundle"`

	BrokerShape *BrokerShape `mandatory:"false" json:"brokerShape"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Kafka Cluster configuration object
	ClusterConfigId *string `mandatory:"false" json:"clusterConfigId"`

	// The version of configuration object
	ClusterConfigVersion *int `mandatory:"false" json:"clusterConfigVersion"`

	// Subnets where broker/coordinator VNICs will be created.
	AccessSubnets []SubnetSet `mandatory:"false" json:"accessSubnets"`

	// Kafka coordination type. Set of available types depends on Kafka version
	CoordinationType KafkaClusterCoordinationTypeEnum `mandatory:"false" json:"coordinationType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateKafkaClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateKafkaClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingKafkaClusterCoordinationTypeEnum(string(m.CoordinationType)); !ok && m.CoordinationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CoordinationType: %s. Supported values are: %s.", m.CoordinationType, strings.Join(GetKafkaClusterCoordinationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
