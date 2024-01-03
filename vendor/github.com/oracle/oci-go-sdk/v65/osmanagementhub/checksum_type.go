// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// ChecksumTypeEnum Enum with underlying type: string
type ChecksumTypeEnum string

// Set of constants representing the allowable values for ChecksumTypeEnum
const (
	ChecksumTypeSha1   ChecksumTypeEnum = "SHA1"
	ChecksumTypeSha256 ChecksumTypeEnum = "SHA256"
	ChecksumTypeSha384 ChecksumTypeEnum = "SHA384"
	ChecksumTypeSha512 ChecksumTypeEnum = "SHA512"
)

var mappingChecksumTypeEnum = map[string]ChecksumTypeEnum{
	"SHA1":   ChecksumTypeSha1,
	"SHA256": ChecksumTypeSha256,
	"SHA384": ChecksumTypeSha384,
	"SHA512": ChecksumTypeSha512,
}

var mappingChecksumTypeEnumLowerCase = map[string]ChecksumTypeEnum{
	"sha1":   ChecksumTypeSha1,
	"sha256": ChecksumTypeSha256,
	"sha384": ChecksumTypeSha384,
	"sha512": ChecksumTypeSha512,
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
		"SHA1",
		"SHA256",
		"SHA384",
		"SHA512",
	}
}

// GetMappingChecksumTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChecksumTypeEnum(val string) (ChecksumTypeEnum, bool) {
	enum, ok := mappingChecksumTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
