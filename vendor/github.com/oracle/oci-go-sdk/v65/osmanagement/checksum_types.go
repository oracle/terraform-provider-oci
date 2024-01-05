// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"strings"
)

// ChecksumTypesEnum Enum with underlying type: string
type ChecksumTypesEnum string

// Set of constants representing the allowable values for ChecksumTypesEnum
const (
	ChecksumTypesSha1   ChecksumTypesEnum = "SHA1"
	ChecksumTypesSha256 ChecksumTypesEnum = "SHA256"
	ChecksumTypesSha384 ChecksumTypesEnum = "SHA384"
	ChecksumTypesSha512 ChecksumTypesEnum = "SHA512"
)

var mappingChecksumTypesEnum = map[string]ChecksumTypesEnum{
	"SHA1":   ChecksumTypesSha1,
	"SHA256": ChecksumTypesSha256,
	"SHA384": ChecksumTypesSha384,
	"SHA512": ChecksumTypesSha512,
}

var mappingChecksumTypesEnumLowerCase = map[string]ChecksumTypesEnum{
	"sha1":   ChecksumTypesSha1,
	"sha256": ChecksumTypesSha256,
	"sha384": ChecksumTypesSha384,
	"sha512": ChecksumTypesSha512,
}

// GetChecksumTypesEnumValues Enumerates the set of values for ChecksumTypesEnum
func GetChecksumTypesEnumValues() []ChecksumTypesEnum {
	values := make([]ChecksumTypesEnum, 0)
	for _, v := range mappingChecksumTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetChecksumTypesEnumStringValues Enumerates the set of values in String for ChecksumTypesEnum
func GetChecksumTypesEnumStringValues() []string {
	return []string{
		"SHA1",
		"SHA256",
		"SHA384",
		"SHA512",
	}
}

// GetMappingChecksumTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChecksumTypesEnum(val string) (ChecksumTypesEnum, bool) {
	enum, ok := mappingChecksumTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
