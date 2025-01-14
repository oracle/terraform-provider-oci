// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

	// The Big Data Service cluster node type.
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// Shape of the node. This has to be specified when starting the cluster. Defaults to wn0 for homogeneous clusters and remains empty for heterogeneous clusters.
	// If provided, all nodes in the node type will adopt the specified shape; otherwise, nodes retain their original shapes.
	Shape *string `mandatory:"true" json:"shape"`
}

func (m NodeTypeShapeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeTypeShapeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNodeNodeTypeEnum(string(m.NodeType)); !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetNodeNodeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
