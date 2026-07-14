// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"strings"
)

// ComputeClusterLogicalPlacementConstraintEnum Enum with underlying type: string
type ComputeClusterLogicalPlacementConstraintEnum string

// Set of constants representing the allowable values for ComputeClusterLogicalPlacementConstraintEnum
const (
	ComputeClusterLogicalPlacementConstraintSingleTier                   ComputeClusterLogicalPlacementConstraintEnum = "SINGLE_TIER"
	ComputeClusterLogicalPlacementConstraintSingleBlock                  ComputeClusterLogicalPlacementConstraintEnum = "SINGLE_BLOCK"
	ComputeClusterLogicalPlacementConstraintPackedDistributionMultiBlock ComputeClusterLogicalPlacementConstraintEnum = "PACKED_DISTRIBUTION_MULTI_BLOCK"
)

var mappingComputeClusterLogicalPlacementConstraintEnum = map[string]ComputeClusterLogicalPlacementConstraintEnum{
	"SINGLE_TIER":                     ComputeClusterLogicalPlacementConstraintSingleTier,
	"SINGLE_BLOCK":                    ComputeClusterLogicalPlacementConstraintSingleBlock,
	"PACKED_DISTRIBUTION_MULTI_BLOCK": ComputeClusterLogicalPlacementConstraintPackedDistributionMultiBlock,
}

var mappingComputeClusterLogicalPlacementConstraintEnumLowerCase = map[string]ComputeClusterLogicalPlacementConstraintEnum{
	"single_tier":                     ComputeClusterLogicalPlacementConstraintSingleTier,
	"single_block":                    ComputeClusterLogicalPlacementConstraintSingleBlock,
	"packed_distribution_multi_block": ComputeClusterLogicalPlacementConstraintPackedDistributionMultiBlock,
}

// GetComputeClusterLogicalPlacementConstraintEnumValues Enumerates the set of values for ComputeClusterLogicalPlacementConstraintEnum
func GetComputeClusterLogicalPlacementConstraintEnumValues() []ComputeClusterLogicalPlacementConstraintEnum {
	values := make([]ComputeClusterLogicalPlacementConstraintEnum, 0)
	for _, v := range mappingComputeClusterLogicalPlacementConstraintEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeClusterLogicalPlacementConstraintEnumStringValues Enumerates the set of values in String for ComputeClusterLogicalPlacementConstraintEnum
func GetComputeClusterLogicalPlacementConstraintEnumStringValues() []string {
	return []string{
		"SINGLE_TIER",
		"SINGLE_BLOCK",
		"PACKED_DISTRIBUTION_MULTI_BLOCK",
	}
}

// GetMappingComputeClusterLogicalPlacementConstraintEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeClusterLogicalPlacementConstraintEnum(val string) (ComputeClusterLogicalPlacementConstraintEnum, bool) {
	enum, ok := mappingComputeClusterLogicalPlacementConstraintEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
