// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VnicShapeConfig Shape config of VNIC that will be used to allocate resource in the data plane once the VNIC is attached
type VnicShapeConfig struct {

	// It defines the percentage number of concurrent connections that can be tracked to the VNIC.
	PercentageOfConnTrack *int `mandatory:"false" json:"percentageOfConnTrack"`

	// It defines the bandwidthMbps for the shape.
	AggregateBandwidthBps *int64 `mandatory:"false" json:"aggregateBandwidthBps"`

	// VCNCP will use this flag to set the internet bandwidth for always free vnic.
	IsAlwaysFree *bool `mandatory:"false" json:"isAlwaysFree"`

	// An enum for different platforms (E3, A1, X5, etc).
	VnicShapeCategory VnicShapeConfigVnicShapeCategoryEnum `mandatory:"false" json:"vnicShapeCategory,omitempty"`
}

func (m VnicShapeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VnicShapeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVnicShapeConfigVnicShapeCategoryEnum(string(m.VnicShapeCategory)); !ok && m.VnicShapeCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VnicShapeCategory: %s. Supported values are: %s.", m.VnicShapeCategory, strings.Join(GetVnicShapeConfigVnicShapeCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VnicShapeConfigVnicShapeCategoryEnum Enum with underlying type: string
type VnicShapeConfigVnicShapeCategoryEnum string

// Set of constants representing the allowable values for VnicShapeConfigVnicShapeCategoryEnum
const (
	VnicShapeConfigVnicShapeCategoryX510gNicmodelBasic10g   VnicShapeConfigVnicShapeCategoryEnum = "X5_10G_NicModel.Basic10G"
	VnicShapeConfigVnicShapeCategoryX625gNicmodelBasic25g   VnicShapeConfigVnicShapeCategoryEnum = "X6_25G_NicModel.Basic25G"
	VnicShapeConfigVnicShapeCategoryX725gNicmodelBasic25g   VnicShapeConfigVnicShapeCategoryEnum = "X7_25G_NicModel.Basic25G"
	VnicShapeConfigVnicShapeCategoryX950gNicmodelBasic50g   VnicShapeConfigVnicShapeCategoryEnum = "X9_50G_NicModel.Basic50G"
	VnicShapeConfigVnicShapeCategoryX9100gNicmodelBasic100g VnicShapeConfigVnicShapeCategoryEnum = "X9_100G_NicModel.Basic100G"
	VnicShapeConfigVnicShapeCategoryE125gNicmodelBasic25g   VnicShapeConfigVnicShapeCategoryEnum = "E1_25G_NicModel.Basic25G"
	VnicShapeConfigVnicShapeCategoryE350gNicmodelBasic50g   VnicShapeConfigVnicShapeCategoryEnum = "E3_50G_NicModel.Basic50G"
	VnicShapeConfigVnicShapeCategoryE450gNicmodelBasic50g   VnicShapeConfigVnicShapeCategoryEnum = "E4_50G_NicModel.Basic50G"
	VnicShapeConfigVnicShapeCategoryE550gNicmodelBasic50g   VnicShapeConfigVnicShapeCategoryEnum = "E5_50G_NicModel.Basic50G"
	VnicShapeConfigVnicShapeCategoryE5100gNicmodelBasic100g VnicShapeConfigVnicShapeCategoryEnum = "E5_100G_NicModel.Basic100G"
	VnicShapeConfigVnicShapeCategoryA150gNicmodelBasic50g   VnicShapeConfigVnicShapeCategoryEnum = "A1_50G_NicModel.Basic50G"
)

var mappingVnicShapeConfigVnicShapeCategoryEnum = map[string]VnicShapeConfigVnicShapeCategoryEnum{
	"X5_10G_NicModel.Basic10G":   VnicShapeConfigVnicShapeCategoryX510gNicmodelBasic10g,
	"X6_25G_NicModel.Basic25G":   VnicShapeConfigVnicShapeCategoryX625gNicmodelBasic25g,
	"X7_25G_NicModel.Basic25G":   VnicShapeConfigVnicShapeCategoryX725gNicmodelBasic25g,
	"X9_50G_NicModel.Basic50G":   VnicShapeConfigVnicShapeCategoryX950gNicmodelBasic50g,
	"X9_100G_NicModel.Basic100G": VnicShapeConfigVnicShapeCategoryX9100gNicmodelBasic100g,
	"E1_25G_NicModel.Basic25G":   VnicShapeConfigVnicShapeCategoryE125gNicmodelBasic25g,
	"E3_50G_NicModel.Basic50G":   VnicShapeConfigVnicShapeCategoryE350gNicmodelBasic50g,
	"E4_50G_NicModel.Basic50G":   VnicShapeConfigVnicShapeCategoryE450gNicmodelBasic50g,
	"E5_50G_NicModel.Basic50G":   VnicShapeConfigVnicShapeCategoryE550gNicmodelBasic50g,
	"E5_100G_NicModel.Basic100G": VnicShapeConfigVnicShapeCategoryE5100gNicmodelBasic100g,
	"A1_50G_NicModel.Basic50G":   VnicShapeConfigVnicShapeCategoryA150gNicmodelBasic50g,
}

var mappingVnicShapeConfigVnicShapeCategoryEnumLowerCase = map[string]VnicShapeConfigVnicShapeCategoryEnum{
	"x5_10g_nicmodel.basic10g":   VnicShapeConfigVnicShapeCategoryX510gNicmodelBasic10g,
	"x6_25g_nicmodel.basic25g":   VnicShapeConfigVnicShapeCategoryX625gNicmodelBasic25g,
	"x7_25g_nicmodel.basic25g":   VnicShapeConfigVnicShapeCategoryX725gNicmodelBasic25g,
	"x9_50g_nicmodel.basic50g":   VnicShapeConfigVnicShapeCategoryX950gNicmodelBasic50g,
	"x9_100g_nicmodel.basic100g": VnicShapeConfigVnicShapeCategoryX9100gNicmodelBasic100g,
	"e1_25g_nicmodel.basic25g":   VnicShapeConfigVnicShapeCategoryE125gNicmodelBasic25g,
	"e3_50g_nicmodel.basic50g":   VnicShapeConfigVnicShapeCategoryE350gNicmodelBasic50g,
	"e4_50g_nicmodel.basic50g":   VnicShapeConfigVnicShapeCategoryE450gNicmodelBasic50g,
	"e5_50g_nicmodel.basic50g":   VnicShapeConfigVnicShapeCategoryE550gNicmodelBasic50g,
	"e5_100g_nicmodel.basic100g": VnicShapeConfigVnicShapeCategoryE5100gNicmodelBasic100g,
	"a1_50g_nicmodel.basic50g":   VnicShapeConfigVnicShapeCategoryA150gNicmodelBasic50g,
}

// GetVnicShapeConfigVnicShapeCategoryEnumValues Enumerates the set of values for VnicShapeConfigVnicShapeCategoryEnum
func GetVnicShapeConfigVnicShapeCategoryEnumValues() []VnicShapeConfigVnicShapeCategoryEnum {
	values := make([]VnicShapeConfigVnicShapeCategoryEnum, 0)
	for _, v := range mappingVnicShapeConfigVnicShapeCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetVnicShapeConfigVnicShapeCategoryEnumStringValues Enumerates the set of values in String for VnicShapeConfigVnicShapeCategoryEnum
func GetVnicShapeConfigVnicShapeCategoryEnumStringValues() []string {
	return []string{
		"X5_10G_NicModel.Basic10G",
		"X6_25G_NicModel.Basic25G",
		"X7_25G_NicModel.Basic25G",
		"X9_50G_NicModel.Basic50G",
		"X9_100G_NicModel.Basic100G",
		"E1_25G_NicModel.Basic25G",
		"E3_50G_NicModel.Basic50G",
		"E4_50G_NicModel.Basic50G",
		"E5_50G_NicModel.Basic50G",
		"E5_100G_NicModel.Basic100G",
		"A1_50G_NicModel.Basic50G",
	}
}

// GetMappingVnicShapeConfigVnicShapeCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVnicShapeConfigVnicShapeCategoryEnum(val string) (VnicShapeConfigVnicShapeCategoryEnum, bool) {
	enum, ok := mappingVnicShapeConfigVnicShapeCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
