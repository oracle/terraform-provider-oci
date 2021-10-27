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

// ShapePlatformConfigOptions The list of supported platform configuration options for this shape.
type ShapePlatformConfigOptions struct {

	// The type of platform being configured.
	Type ShapePlatformConfigOptionsTypeEnum `mandatory:"false" json:"type,omitempty"`

	SecureBootOptions *ShapeSecureBootOptions `mandatory:"false" json:"secureBootOptions"`

	MeasuredBootOptions *ShapeMeasuredBootOptions `mandatory:"false" json:"measuredBootOptions"`

	TrustedPlatformModuleOptions *ShapeTrustedPlatformModuleOptions `mandatory:"false" json:"trustedPlatformModuleOptions"`

	NumaNodesPerSocketPlatformOptions *ShapeNumaNodesPerSocketPlatformOptions `mandatory:"false" json:"numaNodesPerSocketPlatformOptions"`
}

func (m ShapePlatformConfigOptions) String() string {
	return common.PointerString(m)
}

// ShapePlatformConfigOptionsTypeEnum Enum with underlying type: string
type ShapePlatformConfigOptionsTypeEnum string

// Set of constants representing the allowable values for ShapePlatformConfigOptionsTypeEnum
const (
	ShapePlatformConfigOptionsTypeAmdMilanBm     ShapePlatformConfigOptionsTypeEnum = "AMD_MILAN_BM"
	ShapePlatformConfigOptionsTypeAmdRomeBm      ShapePlatformConfigOptionsTypeEnum = "AMD_ROME_BM"
	ShapePlatformConfigOptionsTypeIntelSkylakeBm ShapePlatformConfigOptionsTypeEnum = "INTEL_SKYLAKE_BM"
	ShapePlatformConfigOptionsTypeAmdVm          ShapePlatformConfigOptionsTypeEnum = "AMD_VM"
	ShapePlatformConfigOptionsTypeIntelVm        ShapePlatformConfigOptionsTypeEnum = "INTEL_VM"
)

var mappingShapePlatformConfigOptionsType = map[string]ShapePlatformConfigOptionsTypeEnum{
	"AMD_MILAN_BM":     ShapePlatformConfigOptionsTypeAmdMilanBm,
	"AMD_ROME_BM":      ShapePlatformConfigOptionsTypeAmdRomeBm,
	"INTEL_SKYLAKE_BM": ShapePlatformConfigOptionsTypeIntelSkylakeBm,
	"AMD_VM":           ShapePlatformConfigOptionsTypeAmdVm,
	"INTEL_VM":         ShapePlatformConfigOptionsTypeIntelVm,
}

// GetShapePlatformConfigOptionsTypeEnumValues Enumerates the set of values for ShapePlatformConfigOptionsTypeEnum
func GetShapePlatformConfigOptionsTypeEnumValues() []ShapePlatformConfigOptionsTypeEnum {
	values := make([]ShapePlatformConfigOptionsTypeEnum, 0)
	for _, v := range mappingShapePlatformConfigOptionsType {
		values = append(values, v)
	}
	return values
}
