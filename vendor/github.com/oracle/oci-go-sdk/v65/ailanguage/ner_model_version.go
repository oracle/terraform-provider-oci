// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"strings"
)

// NerModelVersionEnum Enum with underlying type: string
type NerModelVersionEnum string

// Set of constants representing the allowable values for NerModelVersionEnum
const (
	NerModelVersionV21 NerModelVersionEnum = "V2_1"
	NerModelVersionV11 NerModelVersionEnum = "V1_1"
)

var mappingNerModelVersionEnum = map[string]NerModelVersionEnum{
	"V2_1": NerModelVersionV21,
	"V1_1": NerModelVersionV11,
}

var mappingNerModelVersionEnumLowerCase = map[string]NerModelVersionEnum{
	"v2_1": NerModelVersionV21,
	"v1_1": NerModelVersionV11,
}

// GetNerModelVersionEnumValues Enumerates the set of values for NerModelVersionEnum
func GetNerModelVersionEnumValues() []NerModelVersionEnum {
	values := make([]NerModelVersionEnum, 0)
	for _, v := range mappingNerModelVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetNerModelVersionEnumStringValues Enumerates the set of values in String for NerModelVersionEnum
func GetNerModelVersionEnumStringValues() []string {
	return []string{
		"V2_1",
		"V1_1",
	}
}

// GetMappingNerModelVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNerModelVersionEnum(val string) (NerModelVersionEnum, bool) {
	enum, ok := mappingNerModelVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
