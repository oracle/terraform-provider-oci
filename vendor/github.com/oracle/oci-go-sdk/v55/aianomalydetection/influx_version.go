// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

// InfluxVersionEnum Enum with underlying type: string
type InfluxVersionEnum string

// Set of constants representing the allowable values for InfluxVersionEnum
const (
	InfluxVersionV18 InfluxVersionEnum = "V_1_8"
	InfluxVersionV20 InfluxVersionEnum = "V_2_0"
)

var mappingInfluxVersion = map[string]InfluxVersionEnum{
	"V_1_8": InfluxVersionV18,
	"V_2_0": InfluxVersionV20,
}

// GetInfluxVersionEnumValues Enumerates the set of values for InfluxVersionEnum
func GetInfluxVersionEnumValues() []InfluxVersionEnum {
	values := make([]InfluxVersionEnum, 0)
	for _, v := range mappingInfluxVersion {
		values = append(values, v)
	}
	return values
}
