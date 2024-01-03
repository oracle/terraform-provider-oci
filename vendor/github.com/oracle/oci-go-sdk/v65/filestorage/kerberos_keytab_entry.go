// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KerberosKeytabEntry Details of each keytab entry read from the keytab file.
type KerberosKeytabEntry struct {

	// Keytab principal.
	Principal *string `mandatory:"true" json:"principal"`

	// Encryption type with with keytab was generated.
	// Secure: aes128-cts-hmac-sha256-128
	// Secure: aes256-cts-hmac-sha384-192
	// Less Secure: aes128-cts-hmac-sha1-96
	// Less Secure: aes256-cts-hmac-sha1-96
	EncryptionType KerberosKeytabEntryEncryptionTypeEnum `mandatory:"true" json:"encryptionType"`

	// Kerberos KVNO (key version number) for key in keytab entry.
	KeyVersionNumber *int64 `mandatory:"true" json:"keyVersionNumber"`
}

func (m KerberosKeytabEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KerberosKeytabEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKerberosKeytabEntryEncryptionTypeEnum(string(m.EncryptionType)); !ok && m.EncryptionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EncryptionType: %s. Supported values are: %s.", m.EncryptionType, strings.Join(GetKerberosKeytabEntryEncryptionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KerberosKeytabEntryEncryptionTypeEnum Enum with underlying type: string
type KerberosKeytabEntryEncryptionTypeEnum string

// Set of constants representing the allowable values for KerberosKeytabEntryEncryptionTypeEnum
const (
	KerberosKeytabEntryEncryptionTypeAes128CtsHmacSha256128 KerberosKeytabEntryEncryptionTypeEnum = "AES128_CTS_HMAC_SHA256_128"
	KerberosKeytabEntryEncryptionTypeAes256CtsHmacSha384192 KerberosKeytabEntryEncryptionTypeEnum = "AES256_CTS_HMAC_SHA384_192"
	KerberosKeytabEntryEncryptionTypeAes128CtsHmacSha196    KerberosKeytabEntryEncryptionTypeEnum = "AES128_CTS_HMAC_SHA1_96"
	KerberosKeytabEntryEncryptionTypeAes256CtsHmacSha196    KerberosKeytabEntryEncryptionTypeEnum = "AES256_CTS_HMAC_SHA1_96"
)

var mappingKerberosKeytabEntryEncryptionTypeEnum = map[string]KerberosKeytabEntryEncryptionTypeEnum{
	"AES128_CTS_HMAC_SHA256_128": KerberosKeytabEntryEncryptionTypeAes128CtsHmacSha256128,
	"AES256_CTS_HMAC_SHA384_192": KerberosKeytabEntryEncryptionTypeAes256CtsHmacSha384192,
	"AES128_CTS_HMAC_SHA1_96":    KerberosKeytabEntryEncryptionTypeAes128CtsHmacSha196,
	"AES256_CTS_HMAC_SHA1_96":    KerberosKeytabEntryEncryptionTypeAes256CtsHmacSha196,
}

var mappingKerberosKeytabEntryEncryptionTypeEnumLowerCase = map[string]KerberosKeytabEntryEncryptionTypeEnum{
	"aes128_cts_hmac_sha256_128": KerberosKeytabEntryEncryptionTypeAes128CtsHmacSha256128,
	"aes256_cts_hmac_sha384_192": KerberosKeytabEntryEncryptionTypeAes256CtsHmacSha384192,
	"aes128_cts_hmac_sha1_96":    KerberosKeytabEntryEncryptionTypeAes128CtsHmacSha196,
	"aes256_cts_hmac_sha1_96":    KerberosKeytabEntryEncryptionTypeAes256CtsHmacSha196,
}

// GetKerberosKeytabEntryEncryptionTypeEnumValues Enumerates the set of values for KerberosKeytabEntryEncryptionTypeEnum
func GetKerberosKeytabEntryEncryptionTypeEnumValues() []KerberosKeytabEntryEncryptionTypeEnum {
	values := make([]KerberosKeytabEntryEncryptionTypeEnum, 0)
	for _, v := range mappingKerberosKeytabEntryEncryptionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKerberosKeytabEntryEncryptionTypeEnumStringValues Enumerates the set of values in String for KerberosKeytabEntryEncryptionTypeEnum
func GetKerberosKeytabEntryEncryptionTypeEnumStringValues() []string {
	return []string{
		"AES128_CTS_HMAC_SHA256_128",
		"AES256_CTS_HMAC_SHA384_192",
		"AES128_CTS_HMAC_SHA1_96",
		"AES256_CTS_HMAC_SHA1_96",
	}
}

// GetMappingKerberosKeytabEntryEncryptionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKerberosKeytabEntryEncryptionTypeEnum(val string) (KerberosKeytabEntryEncryptionTypeEnum, bool) {
	enum, ok := mappingKerberosKeytabEntryEncryptionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
