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

// AddBlockStorageDetails The information about added block volumes.
type AddBlockStorageDetails struct {

	// Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// The size of block volume in GB to be added to each worker node. All the
	// details needed for attaching the block volume are managed by service itself.
	BlockVolumeSizeInGBs *int64 `mandatory:"true" json:"blockVolumeSizeInGBs"`

	// Worker node types.
	NodeType AddBlockStorageDetailsNodeTypeEnum `mandatory:"true" json:"nodeType"`
}

func (m AddBlockStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddBlockStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddBlockStorageDetailsNodeTypeEnum(string(m.NodeType)); !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetAddBlockStorageDetailsNodeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddBlockStorageDetailsNodeTypeEnum Enum with underlying type: string
type AddBlockStorageDetailsNodeTypeEnum string

// Set of constants representing the allowable values for AddBlockStorageDetailsNodeTypeEnum
const (
	AddBlockStorageDetailsNodeTypeWorker            AddBlockStorageDetailsNodeTypeEnum = "WORKER"
	AddBlockStorageDetailsNodeTypeComputeOnlyWorker AddBlockStorageDetailsNodeTypeEnum = "COMPUTE_ONLY_WORKER"
	AddBlockStorageDetailsNodeTypeKafkaBroker       AddBlockStorageDetailsNodeTypeEnum = "KAFKA_BROKER"
)

var mappingAddBlockStorageDetailsNodeTypeEnum = map[string]AddBlockStorageDetailsNodeTypeEnum{
	"WORKER":              AddBlockStorageDetailsNodeTypeWorker,
	"COMPUTE_ONLY_WORKER": AddBlockStorageDetailsNodeTypeComputeOnlyWorker,
	"KAFKA_BROKER":        AddBlockStorageDetailsNodeTypeKafkaBroker,
}

var mappingAddBlockStorageDetailsNodeTypeEnumLowerCase = map[string]AddBlockStorageDetailsNodeTypeEnum{
	"worker":              AddBlockStorageDetailsNodeTypeWorker,
	"compute_only_worker": AddBlockStorageDetailsNodeTypeComputeOnlyWorker,
	"kafka_broker":        AddBlockStorageDetailsNodeTypeKafkaBroker,
}

// GetAddBlockStorageDetailsNodeTypeEnumValues Enumerates the set of values for AddBlockStorageDetailsNodeTypeEnum
func GetAddBlockStorageDetailsNodeTypeEnumValues() []AddBlockStorageDetailsNodeTypeEnum {
	values := make([]AddBlockStorageDetailsNodeTypeEnum, 0)
	for _, v := range mappingAddBlockStorageDetailsNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddBlockStorageDetailsNodeTypeEnumStringValues Enumerates the set of values in String for AddBlockStorageDetailsNodeTypeEnum
func GetAddBlockStorageDetailsNodeTypeEnumStringValues() []string {
	return []string{
		"WORKER",
		"COMPUTE_ONLY_WORKER",
		"KAFKA_BROKER",
	}
}

// GetMappingAddBlockStorageDetailsNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddBlockStorageDetailsNodeTypeEnum(val string) (AddBlockStorageDetailsNodeTypeEnum, bool) {
	enum, ok := mappingAddBlockStorageDetailsNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
