// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"strings"
)

// MacsecEncryptionCipherEnum Enum with underlying type: string
type MacsecEncryptionCipherEnum string

// Set of constants representing the allowable values for MacsecEncryptionCipherEnum
const (
	MacsecEncryptionCipherAes128Gcm    MacsecEncryptionCipherEnum = "AES128_GCM"
	MacsecEncryptionCipherAes128GcmXpn MacsecEncryptionCipherEnum = "AES128_GCM_XPN"
	MacsecEncryptionCipherAes256Gcm    MacsecEncryptionCipherEnum = "AES256_GCM"
	MacsecEncryptionCipherAes256GcmXpn MacsecEncryptionCipherEnum = "AES256_GCM_XPN"
)

var mappingMacsecEncryptionCipherEnum = map[string]MacsecEncryptionCipherEnum{
	"AES128_GCM":     MacsecEncryptionCipherAes128Gcm,
	"AES128_GCM_XPN": MacsecEncryptionCipherAes128GcmXpn,
	"AES256_GCM":     MacsecEncryptionCipherAes256Gcm,
	"AES256_GCM_XPN": MacsecEncryptionCipherAes256GcmXpn,
}

var mappingMacsecEncryptionCipherEnumLowerCase = map[string]MacsecEncryptionCipherEnum{
	"aes128_gcm":     MacsecEncryptionCipherAes128Gcm,
	"aes128_gcm_xpn": MacsecEncryptionCipherAes128GcmXpn,
	"aes256_gcm":     MacsecEncryptionCipherAes256Gcm,
	"aes256_gcm_xpn": MacsecEncryptionCipherAes256GcmXpn,
}

// GetMacsecEncryptionCipherEnumValues Enumerates the set of values for MacsecEncryptionCipherEnum
func GetMacsecEncryptionCipherEnumValues() []MacsecEncryptionCipherEnum {
	values := make([]MacsecEncryptionCipherEnum, 0)
	for _, v := range mappingMacsecEncryptionCipherEnum {
		values = append(values, v)
	}
	return values
}

// GetMacsecEncryptionCipherEnumStringValues Enumerates the set of values in String for MacsecEncryptionCipherEnum
func GetMacsecEncryptionCipherEnumStringValues() []string {
	return []string{
		"AES128_GCM",
		"AES128_GCM_XPN",
		"AES256_GCM",
		"AES256_GCM_XPN",
	}
}

// GetMappingMacsecEncryptionCipherEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacsecEncryptionCipherEnum(val string) (MacsecEncryptionCipherEnum, bool) {
	enum, ok := mappingMacsecEncryptionCipherEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
