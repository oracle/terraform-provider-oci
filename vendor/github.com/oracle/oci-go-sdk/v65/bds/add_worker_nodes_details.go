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

// AddWorkerNodesDetails The information about added nodes.
type AddWorkerNodesDetails struct {

	// Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// Number of additional worker nodes for the cluster.
	NumberOfWorkerNodes *int `mandatory:"true" json:"numberOfWorkerNodes"`

	// Worker node types, can either be Worker Data node or Compute only worker node.
	NodeType AddWorkerNodesDetailsNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// Shape of the node. This has to be specified when adding compute only worker node at the first time. Otherwise, it's a read-only property.
	Shape *string `mandatory:"false" json:"shape"`

	// The size of block volume in GB to be attached to the given node. This has to be specified when adding compute only worker node at the first time. Otherwise, it's a read-only property.
	BlockVolumeSizeInGBs *int64 `mandatory:"false" json:"blockVolumeSizeInGBs"`

	ShapeConfig *ShapeConfigDetails `mandatory:"false" json:"shapeConfig"`
}

func (m AddWorkerNodesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddWorkerNodesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddWorkerNodesDetailsNodeTypeEnum(string(m.NodeType)); !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetAddWorkerNodesDetailsNodeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddWorkerNodesDetailsNodeTypeEnum Enum with underlying type: string
type AddWorkerNodesDetailsNodeTypeEnum string

// Set of constants representing the allowable values for AddWorkerNodesDetailsNodeTypeEnum
const (
	AddWorkerNodesDetailsNodeTypeWorker            AddWorkerNodesDetailsNodeTypeEnum = "WORKER"
	AddWorkerNodesDetailsNodeTypeComputeOnlyWorker AddWorkerNodesDetailsNodeTypeEnum = "COMPUTE_ONLY_WORKER"
	AddWorkerNodesDetailsNodeTypeEdge              AddWorkerNodesDetailsNodeTypeEnum = "EDGE"
	AddWorkerNodesDetailsNodeTypeKafkaBroker       AddWorkerNodesDetailsNodeTypeEnum = "KAFKA_BROKER"
)

var mappingAddWorkerNodesDetailsNodeTypeEnum = map[string]AddWorkerNodesDetailsNodeTypeEnum{
	"WORKER":              AddWorkerNodesDetailsNodeTypeWorker,
	"COMPUTE_ONLY_WORKER": AddWorkerNodesDetailsNodeTypeComputeOnlyWorker,
	"EDGE":                AddWorkerNodesDetailsNodeTypeEdge,
	"KAFKA_BROKER":        AddWorkerNodesDetailsNodeTypeKafkaBroker,
}

var mappingAddWorkerNodesDetailsNodeTypeEnumLowerCase = map[string]AddWorkerNodesDetailsNodeTypeEnum{
	"worker":              AddWorkerNodesDetailsNodeTypeWorker,
	"compute_only_worker": AddWorkerNodesDetailsNodeTypeComputeOnlyWorker,
	"edge":                AddWorkerNodesDetailsNodeTypeEdge,
	"kafka_broker":        AddWorkerNodesDetailsNodeTypeKafkaBroker,
}

// GetAddWorkerNodesDetailsNodeTypeEnumValues Enumerates the set of values for AddWorkerNodesDetailsNodeTypeEnum
func GetAddWorkerNodesDetailsNodeTypeEnumValues() []AddWorkerNodesDetailsNodeTypeEnum {
	values := make([]AddWorkerNodesDetailsNodeTypeEnum, 0)
	for _, v := range mappingAddWorkerNodesDetailsNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddWorkerNodesDetailsNodeTypeEnumStringValues Enumerates the set of values in String for AddWorkerNodesDetailsNodeTypeEnum
func GetAddWorkerNodesDetailsNodeTypeEnumStringValues() []string {
	return []string{
		"WORKER",
		"COMPUTE_ONLY_WORKER",
		"EDGE",
		"KAFKA_BROKER",
	}
}

// GetMappingAddWorkerNodesDetailsNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddWorkerNodesDetailsNodeTypeEnum(val string) (AddWorkerNodesDetailsNodeTypeEnum, bool) {
	enum, ok := mappingAddWorkerNodesDetailsNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
