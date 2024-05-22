// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeTypeLevelDetails Details of node type level used to trigger the creation of a new node backup configuration and node replacement configuration.
type NodeTypeLevelDetails struct {

	// Type of the node or nodes of the node backup configuration or node replacement configuration which are going to be created.
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`
}

func (m NodeTypeLevelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeTypeLevelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNodeNodeTypeEnum(string(m.NodeType)); !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetNodeNodeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NodeTypeLevelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNodeTypeLevelDetails NodeTypeLevelDetails
	s := struct {
		DiscriminatorParam string `json:"levelType"`
		MarshalTypeNodeTypeLevelDetails
	}{
		"NODE_TYPE_LEVEL",
		(MarshalTypeNodeTypeLevelDetails)(m),
	}

	return json.Marshal(&s)
}
