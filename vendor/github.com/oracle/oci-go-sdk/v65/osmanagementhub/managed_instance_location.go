// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// ManagedInstanceLocationEnum Enum with underlying type: string
type ManagedInstanceLocationEnum string

// Set of constants representing the allowable values for ManagedInstanceLocationEnum
const (
	ManagedInstanceLocationOnPremise  ManagedInstanceLocationEnum = "ON_PREMISE"
	ManagedInstanceLocationOciCompute ManagedInstanceLocationEnum = "OCI_COMPUTE"
	ManagedInstanceLocationAzure      ManagedInstanceLocationEnum = "AZURE"
	ManagedInstanceLocationEc2        ManagedInstanceLocationEnum = "EC2"
	ManagedInstanceLocationGcp        ManagedInstanceLocationEnum = "GCP"
)

var mappingManagedInstanceLocationEnum = map[string]ManagedInstanceLocationEnum{
	"ON_PREMISE":  ManagedInstanceLocationOnPremise,
	"OCI_COMPUTE": ManagedInstanceLocationOciCompute,
	"AZURE":       ManagedInstanceLocationAzure,
	"EC2":         ManagedInstanceLocationEc2,
	"GCP":         ManagedInstanceLocationGcp,
}

var mappingManagedInstanceLocationEnumLowerCase = map[string]ManagedInstanceLocationEnum{
	"on_premise":  ManagedInstanceLocationOnPremise,
	"oci_compute": ManagedInstanceLocationOciCompute,
	"azure":       ManagedInstanceLocationAzure,
	"ec2":         ManagedInstanceLocationEc2,
	"gcp":         ManagedInstanceLocationGcp,
}

// GetManagedInstanceLocationEnumValues Enumerates the set of values for ManagedInstanceLocationEnum
func GetManagedInstanceLocationEnumValues() []ManagedInstanceLocationEnum {
	values := make([]ManagedInstanceLocationEnum, 0)
	for _, v := range mappingManagedInstanceLocationEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedInstanceLocationEnumStringValues Enumerates the set of values in String for ManagedInstanceLocationEnum
func GetManagedInstanceLocationEnumStringValues() []string {
	return []string{
		"ON_PREMISE",
		"OCI_COMPUTE",
		"AZURE",
		"EC2",
		"GCP",
	}
}

// GetMappingManagedInstanceLocationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedInstanceLocationEnum(val string) (ManagedInstanceLocationEnum, bool) {
	enum, ok := mappingManagedInstanceLocationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
