// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the <a href="https://docs.oracle.com/en-us/iaas/jms/doc/java-download.html">Java Download</a> feature of Java Management Service.
//

package jmsjavadownloads

import (
	"strings"
)

// ChecksumTypeEnum Enum with underlying type: string
type ChecksumTypeEnum string

// Set of constants representing the allowable values for ChecksumTypeEnum
const (
	ChecksumTypeSha256 ChecksumTypeEnum = "SHA256"
)

var mappingChecksumTypeEnum = map[string]ChecksumTypeEnum{
	"SHA256": ChecksumTypeSha256,
}

var mappingChecksumTypeEnumLowerCase = map[string]ChecksumTypeEnum{
	"sha256": ChecksumTypeSha256,
}

// GetChecksumTypeEnumValues Enumerates the set of values for ChecksumTypeEnum
func GetChecksumTypeEnumValues() []ChecksumTypeEnum {
	values := make([]ChecksumTypeEnum, 0)
	for _, v := range mappingChecksumTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChecksumTypeEnumStringValues Enumerates the set of values in String for ChecksumTypeEnum
func GetChecksumTypeEnumStringValues() []string {
	return []string{
		"SHA256",
	}
}

// GetMappingChecksumTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChecksumTypeEnum(val string) (ChecksumTypeEnum, bool) {
	enum, ok := mappingChecksumTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
