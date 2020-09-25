// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// KerberosKeytabEntry Details of each keytab entry read from Keytab file.
type KerberosKeytabEntry struct {

	// Keytab principal.
	Principal *string `mandatory:"false" json:"principal"`

	// Encryption type with with keytab was generated.
	// Secure: aes128-cts-hmac-sha256-128
	// Secure: aes256-cts-hmac-sha384-192
	// Less Secure: aes128-cts-hmac-sha1-96
	// Less Secure: aes256-cts-hmac-sha1-96
	EncryptionType KerberosKeytabEntryEncryptionTypeEnum `mandatory:"false" json:"encryptionType,omitempty"`

	// Kerberos kvno (key version number) for key in keytab entry.
	KeyVersionNumber *int64 `mandatory:"false" json:"keyVersionNumber"`

	// The date and time the Keytab was rotated, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeRotated *common.SDKTime `mandatory:"false" json:"timeRotated"`
}

func (m KerberosKeytabEntry) String() string {
	return common.PointerString(m)
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

var mappingKerberosKeytabEntryEncryptionType = map[string]KerberosKeytabEntryEncryptionTypeEnum{
	"AES128_CTS_HMAC_SHA256_128": KerberosKeytabEntryEncryptionTypeAes128CtsHmacSha256128,
	"AES256_CTS_HMAC_SHA384_192": KerberosKeytabEntryEncryptionTypeAes256CtsHmacSha384192,
	"AES128_CTS_HMAC_SHA1_96":    KerberosKeytabEntryEncryptionTypeAes128CtsHmacSha196,
	"AES256_CTS_HMAC_SHA1_96":    KerberosKeytabEntryEncryptionTypeAes256CtsHmacSha196,
}

// GetKerberosKeytabEntryEncryptionTypeEnumValues Enumerates the set of values for KerberosKeytabEntryEncryptionTypeEnum
func GetKerberosKeytabEntryEncryptionTypeEnumValues() []KerberosKeytabEntryEncryptionTypeEnum {
	values := make([]KerberosKeytabEntryEncryptionTypeEnum, 0)
	for _, v := range mappingKerberosKeytabEntryEncryptionType {
		values = append(values, v)
	}
	return values
}
