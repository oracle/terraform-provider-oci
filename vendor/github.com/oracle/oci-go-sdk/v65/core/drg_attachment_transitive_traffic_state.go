// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DrgAttachmentTransitiveTrafficStateEnum Enum with underlying type: string
type DrgAttachmentTransitiveTrafficStateEnum string

// Set of constants representing the allowable values for DrgAttachmentTransitiveTrafficStateEnum
const (
	DrgAttachmentTransitiveTrafficStateEnabled  DrgAttachmentTransitiveTrafficStateEnum = "ENABLED"
	DrgAttachmentTransitiveTrafficStateDisabled DrgAttachmentTransitiveTrafficStateEnum = "DISABLED"
)

var mappingDrgAttachmentTransitiveTrafficStateEnum = map[string]DrgAttachmentTransitiveTrafficStateEnum{
	"ENABLED":  DrgAttachmentTransitiveTrafficStateEnabled,
	"DISABLED": DrgAttachmentTransitiveTrafficStateDisabled,
}

var mappingDrgAttachmentTransitiveTrafficStateEnumLowerCase = map[string]DrgAttachmentTransitiveTrafficStateEnum{
	"enabled":  DrgAttachmentTransitiveTrafficStateEnabled,
	"disabled": DrgAttachmentTransitiveTrafficStateDisabled,
}

// GetDrgAttachmentTransitiveTrafficStateEnumValues Enumerates the set of values for DrgAttachmentTransitiveTrafficStateEnum
func GetDrgAttachmentTransitiveTrafficStateEnumValues() []DrgAttachmentTransitiveTrafficStateEnum {
	values := make([]DrgAttachmentTransitiveTrafficStateEnum, 0)
	for _, v := range mappingDrgAttachmentTransitiveTrafficStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentTransitiveTrafficStateEnumStringValues Enumerates the set of values in String for DrgAttachmentTransitiveTrafficStateEnum
func GetDrgAttachmentTransitiveTrafficStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingDrgAttachmentTransitiveTrafficStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentTransitiveTrafficStateEnum(val string) (DrgAttachmentTransitiveTrafficStateEnum, bool) {
	enum, ok := mappingDrgAttachmentTransitiveTrafficStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
