// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"strings"
)

// InfluxVersionEnum Enum with underlying type: string
type InfluxVersionEnum string

// Set of constants representing the allowable values for InfluxVersionEnum
const (
	InfluxVersionV18 InfluxVersionEnum = "V_1_8"
	InfluxVersionV20 InfluxVersionEnum = "V_2_0"
)

var mappingInfluxVersionEnum = map[string]InfluxVersionEnum{
	"V_1_8": InfluxVersionV18,
	"V_2_0": InfluxVersionV20,
}

var mappingInfluxVersionEnumLowerCase = map[string]InfluxVersionEnum{
	"v_1_8": InfluxVersionV18,
	"v_2_0": InfluxVersionV20,
}

// GetInfluxVersionEnumValues Enumerates the set of values for InfluxVersionEnum
func GetInfluxVersionEnumValues() []InfluxVersionEnum {
	values := make([]InfluxVersionEnum, 0)
	for _, v := range mappingInfluxVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetInfluxVersionEnumStringValues Enumerates the set of values in String for InfluxVersionEnum
func GetInfluxVersionEnumStringValues() []string {
	return []string{
		"V_1_8",
		"V_2_0",
	}
}

// GetMappingInfluxVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfluxVersionEnum(val string) (InfluxVersionEnum, bool) {
	enum, ok := mappingInfluxVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
