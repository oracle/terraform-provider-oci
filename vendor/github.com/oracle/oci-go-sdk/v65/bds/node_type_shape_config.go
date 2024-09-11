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

// NodeTypeShapeConfig Shape configuration at node type level. Start cluster will start all nodes as is if no config is specified.
type NodeTypeShapeConfig struct {

	// All node types that exist in the cluster
	NodeType NodeTypeShapeConfigNodeTypeEnum `mandatory:"false" json:"nodeType,omitempty"`

	// Shape of the node. This has to be specified when starting the cluster. Defaults to wn0 for homogeneous clusters and remains empty for heterogeneous clusters.
	// If provided, all nodes in the node type will adopt the specified shape; otherwise, nodes retain their original shapes.
	Shape *string `mandatory:"false" json:"shape"`
}

func (m NodeTypeShapeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeTypeShapeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNodeTypeShapeConfigNodeTypeEnum(string(m.NodeType)); !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetNodeTypeShapeConfigNodeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NodeTypeShapeConfigNodeTypeEnum Enum with underlying type: string
type NodeTypeShapeConfigNodeTypeEnum string

// Set of constants representing the allowable values for NodeTypeShapeConfigNodeTypeEnum
const (
	NodeTypeShapeConfigNodeTypeWorker            NodeTypeShapeConfigNodeTypeEnum = "WORKER"
	NodeTypeShapeConfigNodeTypeComputeOnlyWorker NodeTypeShapeConfigNodeTypeEnum = "COMPUTE_ONLY_WORKER"
	NodeTypeShapeConfigNodeTypeEdge              NodeTypeShapeConfigNodeTypeEnum = "EDGE"
	NodeTypeShapeConfigNodeTypeKafkaBroker       NodeTypeShapeConfigNodeTypeEnum = "KAFKA_BROKER"
	NodeTypeShapeConfigNodeTypeMaster            NodeTypeShapeConfigNodeTypeEnum = "MASTER"
	NodeTypeShapeConfigNodeTypeUtility           NodeTypeShapeConfigNodeTypeEnum = "UTILITY"
)

var mappingNodeTypeShapeConfigNodeTypeEnum = map[string]NodeTypeShapeConfigNodeTypeEnum{
	"WORKER":              NodeTypeShapeConfigNodeTypeWorker,
	"COMPUTE_ONLY_WORKER": NodeTypeShapeConfigNodeTypeComputeOnlyWorker,
	"EDGE":                NodeTypeShapeConfigNodeTypeEdge,
	"KAFKA_BROKER":        NodeTypeShapeConfigNodeTypeKafkaBroker,
	"MASTER":              NodeTypeShapeConfigNodeTypeMaster,
	"UTILITY":             NodeTypeShapeConfigNodeTypeUtility,
}

var mappingNodeTypeShapeConfigNodeTypeEnumLowerCase = map[string]NodeTypeShapeConfigNodeTypeEnum{
	"worker":              NodeTypeShapeConfigNodeTypeWorker,
	"compute_only_worker": NodeTypeShapeConfigNodeTypeComputeOnlyWorker,
	"edge":                NodeTypeShapeConfigNodeTypeEdge,
	"kafka_broker":        NodeTypeShapeConfigNodeTypeKafkaBroker,
	"master":              NodeTypeShapeConfigNodeTypeMaster,
	"utility":             NodeTypeShapeConfigNodeTypeUtility,
}

// GetNodeTypeShapeConfigNodeTypeEnumValues Enumerates the set of values for NodeTypeShapeConfigNodeTypeEnum
func GetNodeTypeShapeConfigNodeTypeEnumValues() []NodeTypeShapeConfigNodeTypeEnum {
	values := make([]NodeTypeShapeConfigNodeTypeEnum, 0)
	for _, v := range mappingNodeTypeShapeConfigNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeTypeShapeConfigNodeTypeEnumStringValues Enumerates the set of values in String for NodeTypeShapeConfigNodeTypeEnum
func GetNodeTypeShapeConfigNodeTypeEnumStringValues() []string {
	return []string{
		"WORKER",
		"COMPUTE_ONLY_WORKER",
		"EDGE",
		"KAFKA_BROKER",
		"MASTER",
		"UTILITY",
	}
}

// GetMappingNodeTypeShapeConfigNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeTypeShapeConfigNodeTypeEnum(val string) (NodeTypeShapeConfigNodeTypeEnum, bool) {
	enum, ok := mappingNodeTypeShapeConfigNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
