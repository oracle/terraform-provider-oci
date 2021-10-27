// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ShapeNumaNodesPerSocketPlatformOptions Configuration options for NUMA nodes per socket.
type ShapeNumaNodesPerSocketPlatformOptions struct {

	// The supported values for this platform configuration property.
	AllowedValues []ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum `mandatory:"false" json:"allowedValues,omitempty"`

	// The default NUMA nodes per socket configuration.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`
}

func (m ShapeNumaNodesPerSocketPlatformOptions) String() string {
	return common.PointerString(m)
}

// ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum Enum with underlying type: string
type ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum string

// Set of constants representing the allowable values for ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum
const (
	ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesNps0 ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum = "NPS0"
	ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesNps1 ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum = "NPS1"
	ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesNps2 ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum = "NPS2"
	ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesNps4 ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum = "NPS4"
)

var mappingShapeNumaNodesPerSocketPlatformOptionsAllowedValues = map[string]ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum{
	"NPS0": ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesNps0,
	"NPS1": ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesNps1,
	"NPS2": ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesNps2,
	"NPS4": ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesNps4,
}

// GetShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnumValues Enumerates the set of values for ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum
func GetShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnumValues() []ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum {
	values := make([]ShapeNumaNodesPerSocketPlatformOptionsAllowedValuesEnum, 0)
	for _, v := range mappingShapeNumaNodesPerSocketPlatformOptionsAllowedValues {
		values = append(values, v)
	}
	return values
}
