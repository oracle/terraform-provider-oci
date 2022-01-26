// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TlsConfig The details required to establish a TLS enabled connection.
type TlsConfig struct {

	// Status to represent whether the database connection is TLS enabled or not.
	Status TlsConfigStatusEnum `mandatory:"true" json:"status"`

	// The format of the certificate store.
	CertificateStoreType TlsConfigCertificateStoreTypeEnum `mandatory:"false" json:"certificateStoreType,omitempty"`

	// The password to read the trust store and key store files, if they are password protected.
	StorePassword *string `mandatory:"false" json:"storePassword"`

	// Base64 encoded string of trust store file content.
	TrustStoreContent *string `mandatory:"false" json:"trustStoreContent"`

	// Base64 encoded string of key store file content.
	KeyStoreContent *string `mandatory:"false" json:"keyStoreContent"`
}

func (m TlsConfig) String() string {
	return common.PointerString(m)
}

// TlsConfigStatusEnum Enum with underlying type: string
type TlsConfigStatusEnum string

// Set of constants representing the allowable values for TlsConfigStatusEnum
const (
	TlsConfigStatusEnabled  TlsConfigStatusEnum = "ENABLED"
	TlsConfigStatusDisabled TlsConfigStatusEnum = "DISABLED"
)

var mappingTlsConfigStatus = map[string]TlsConfigStatusEnum{
	"ENABLED":  TlsConfigStatusEnabled,
	"DISABLED": TlsConfigStatusDisabled,
}

// GetTlsConfigStatusEnumValues Enumerates the set of values for TlsConfigStatusEnum
func GetTlsConfigStatusEnumValues() []TlsConfigStatusEnum {
	values := make([]TlsConfigStatusEnum, 0)
	for _, v := range mappingTlsConfigStatus {
		values = append(values, v)
	}
	return values
}

// TlsConfigCertificateStoreTypeEnum Enum with underlying type: string
type TlsConfigCertificateStoreTypeEnum string

// Set of constants representing the allowable values for TlsConfigCertificateStoreTypeEnum
const (
	TlsConfigCertificateStoreTypeJks TlsConfigCertificateStoreTypeEnum = "JKS"
)

var mappingTlsConfigCertificateStoreType = map[string]TlsConfigCertificateStoreTypeEnum{
	"JKS": TlsConfigCertificateStoreTypeJks,
}

// GetTlsConfigCertificateStoreTypeEnumValues Enumerates the set of values for TlsConfigCertificateStoreTypeEnum
func GetTlsConfigCertificateStoreTypeEnumValues() []TlsConfigCertificateStoreTypeEnum {
	values := make([]TlsConfigCertificateStoreTypeEnum, 0)
	for _, v := range mappingTlsConfigCertificateStoreType {
		values = append(values, v)
	}
	return values
}
