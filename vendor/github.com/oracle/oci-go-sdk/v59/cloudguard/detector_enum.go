// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// DetectorEnumEnum Enum with underlying type: string
type DetectorEnumEnum string

// Set of constants representing the allowable values for DetectorEnumEnum
const (
	DetectorEnumIaasActivityDetector      DetectorEnumEnum = "IAAS_ACTIVITY_DETECTOR"
	DetectorEnumIaasConfigurationDetector DetectorEnumEnum = "IAAS_CONFIGURATION_DETECTOR"
	DetectorEnumIaasThreatDetector        DetectorEnumEnum = "IAAS_THREAT_DETECTOR"
)

var mappingDetectorEnumEnum = map[string]DetectorEnumEnum{
	"IAAS_ACTIVITY_DETECTOR":      DetectorEnumIaasActivityDetector,
	"IAAS_CONFIGURATION_DETECTOR": DetectorEnumIaasConfigurationDetector,
	"IAAS_THREAT_DETECTOR":        DetectorEnumIaasThreatDetector,
}

var mappingDetectorEnumEnumLowerCase = map[string]DetectorEnumEnum{
	"iaas_activity_detector":      DetectorEnumIaasActivityDetector,
	"iaas_configuration_detector": DetectorEnumIaasConfigurationDetector,
	"iaas_threat_detector":        DetectorEnumIaasThreatDetector,
}

// GetDetectorEnumEnumValues Enumerates the set of values for DetectorEnumEnum
func GetDetectorEnumEnumValues() []DetectorEnumEnum {
	values := make([]DetectorEnumEnum, 0)
	for _, v := range mappingDetectorEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectorEnumEnumStringValues Enumerates the set of values in String for DetectorEnumEnum
func GetDetectorEnumEnumStringValues() []string {
	return []string{
		"IAAS_ACTIVITY_DETECTOR",
		"IAAS_CONFIGURATION_DETECTOR",
		"IAAS_THREAT_DETECTOR",
	}
}

// GetMappingDetectorEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectorEnumEnum(val string) (DetectorEnumEnum, bool) {
	enum, ok := mappingDetectorEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
