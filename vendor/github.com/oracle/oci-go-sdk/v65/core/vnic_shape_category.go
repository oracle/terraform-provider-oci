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
	"strings"
)

// VnicShapeCategoryEnum Enum with underlying type: string
type VnicShapeCategoryEnum string

// Set of constants representing the allowable values for VnicShapeCategoryEnum
const (
	VnicShapeCategoryX510GNicModelBasic10G   VnicShapeCategoryEnum = "X5_10G_NicModel.Basic10G"
	VnicShapeCategoryX625GNicModelBasic25G   VnicShapeCategoryEnum = "X6_25G_NicModel.Basic25G"
	VnicShapeCategoryX725GNicModelBasic25G   VnicShapeCategoryEnum = "X7_25G_NicModel.Basic25G"
	VnicShapeCategoryX950GNicModelBasic50G   VnicShapeCategoryEnum = "X9_50G_NicModel.Basic50G"
	VnicShapeCategoryX9100GNicModelBasic100G VnicShapeCategoryEnum = "X9_100G_NicModel.Basic100G"
	VnicShapeCategoryE125GNicModelBasic25G   VnicShapeCategoryEnum = "E1_25G_NicModel.Basic25G"
	VnicShapeCategoryE350GNicModelBasic50G   VnicShapeCategoryEnum = "E3_50G_NicModel.Basic50G"
	VnicShapeCategoryE450GNicModelBasic50G   VnicShapeCategoryEnum = "E4_50G_NicModel.Basic50G"
	VnicShapeCategoryE550GNicModelBasic50G   VnicShapeCategoryEnum = "E5_50G_NicModel.Basic50G"
	VnicShapeCategoryE5100GNicModelBasic100G VnicShapeCategoryEnum = "E5_100G_NicModel.Basic100G"
	VnicShapeCategoryA150GNicModelBasic50G   VnicShapeCategoryEnum = "A1_50G_NicModel.Basic50G"
)

var mappingVnicShapeCategoryEnum = map[string]VnicShapeCategoryEnum{
	"X5_10G_NicModel.Basic10G":   VnicShapeCategoryX510GNicModelBasic10G,
	"X6_25G_NicModel.Basic25G":   VnicShapeCategoryX625GNicModelBasic25G,
	"X7_25G_NicModel.Basic25G":   VnicShapeCategoryX725GNicModelBasic25G,
	"X9_50G_NicModel.Basic50G":   VnicShapeCategoryX950GNicModelBasic50G,
	"X9_100G_NicModel.Basic100G": VnicShapeCategoryX9100GNicModelBasic100G,
	"E1_25G_NicModel.Basic25G":   VnicShapeCategoryE125GNicModelBasic25G,
	"E3_50G_NicModel.Basic50G":   VnicShapeCategoryE350GNicModelBasic50G,
	"E4_50G_NicModel.Basic50G":   VnicShapeCategoryE450GNicModelBasic50G,
	"E5_50G_NicModel.Basic50G":   VnicShapeCategoryE550GNicModelBasic50G,
	"E5_100G_NicModel.Basic100G": VnicShapeCategoryE5100GNicModelBasic100G,
	"A1_50G_NicModel.Basic50G":   VnicShapeCategoryA150GNicModelBasic50G,
}

var mappingVnicShapeCategoryEnumLowerCase = map[string]VnicShapeCategoryEnum{
	"x5_10g_nicmodel.basic10g":   VnicShapeCategoryX510GNicModelBasic10G,
	"x6_25g_nicmodel.basic25g":   VnicShapeCategoryX625GNicModelBasic25G,
	"x7_25g_nicmodel.basic25g":   VnicShapeCategoryX725GNicModelBasic25G,
	"x9_50g_nicmodel.basic50g":   VnicShapeCategoryX950GNicModelBasic50G,
	"x9_100g_nicmodel.basic100g": VnicShapeCategoryX9100GNicModelBasic100G,
	"e1_25g_nicmodel.basic25g":   VnicShapeCategoryE125GNicModelBasic25G,
	"e3_50g_nicmodel.basic50g":   VnicShapeCategoryE350GNicModelBasic50G,
	"e4_50g_nicmodel.basic50g":   VnicShapeCategoryE450GNicModelBasic50G,
	"e5_50g_nicmodel.basic50g":   VnicShapeCategoryE550GNicModelBasic50G,
	"e5_100g_nicmodel.basic100g": VnicShapeCategoryE5100GNicModelBasic100G,
	"a1_50g_nicmodel.basic50g":   VnicShapeCategoryA150GNicModelBasic50G,
}

// GetVnicShapeCategoryEnumValues Enumerates the set of values for VnicShapeCategoryEnum
func GetVnicShapeCategoryEnumValues() []VnicShapeCategoryEnum {
	values := make([]VnicShapeCategoryEnum, 0)
	for _, v := range mappingVnicShapeCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetVnicShapeCategoryEnumStringValues Enumerates the set of values in String for VnicShapeCategoryEnum
func GetVnicShapeCategoryEnumStringValues() []string {
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

// GetMappingVnicShapeCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVnicShapeCategoryEnum(val string) (VnicShapeCategoryEnum, bool) {
	enum, ok := mappingVnicShapeCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
