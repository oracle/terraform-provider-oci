// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TlsConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTlsConfigStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTlsConfigStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTlsConfigCertificateStoreTypeEnum(string(m.CertificateStoreType)); !ok && m.CertificateStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateStoreType: %s. Supported values are: %s.", m.CertificateStoreType, strings.Join(GetTlsConfigCertificateStoreTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TlsConfigStatusEnum Enum with underlying type: string
type TlsConfigStatusEnum string

// Set of constants representing the allowable values for TlsConfigStatusEnum
const (
	TlsConfigStatusEnabled  TlsConfigStatusEnum = "ENABLED"
	TlsConfigStatusDisabled TlsConfigStatusEnum = "DISABLED"
)

var mappingTlsConfigStatusEnum = map[string]TlsConfigStatusEnum{
	"ENABLED":  TlsConfigStatusEnabled,
	"DISABLED": TlsConfigStatusDisabled,
}

// GetTlsConfigStatusEnumValues Enumerates the set of values for TlsConfigStatusEnum
func GetTlsConfigStatusEnumValues() []TlsConfigStatusEnum {
	values := make([]TlsConfigStatusEnum, 0)
	for _, v := range mappingTlsConfigStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTlsConfigStatusEnumStringValues Enumerates the set of values in String for TlsConfigStatusEnum
func GetTlsConfigStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingTlsConfigStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTlsConfigStatusEnum(val string) (TlsConfigStatusEnum, bool) {
	mappingTlsConfigStatusEnumIgnoreCase := make(map[string]TlsConfigStatusEnum)
	for k, v := range mappingTlsConfigStatusEnum {
		mappingTlsConfigStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTlsConfigStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TlsConfigCertificateStoreTypeEnum Enum with underlying type: string
type TlsConfigCertificateStoreTypeEnum string

// Set of constants representing the allowable values for TlsConfigCertificateStoreTypeEnum
const (
	TlsConfigCertificateStoreTypeJks TlsConfigCertificateStoreTypeEnum = "JKS"
)

var mappingTlsConfigCertificateStoreTypeEnum = map[string]TlsConfigCertificateStoreTypeEnum{
	"JKS": TlsConfigCertificateStoreTypeJks,
}

// GetTlsConfigCertificateStoreTypeEnumValues Enumerates the set of values for TlsConfigCertificateStoreTypeEnum
func GetTlsConfigCertificateStoreTypeEnumValues() []TlsConfigCertificateStoreTypeEnum {
	values := make([]TlsConfigCertificateStoreTypeEnum, 0)
	for _, v := range mappingTlsConfigCertificateStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTlsConfigCertificateStoreTypeEnumStringValues Enumerates the set of values in String for TlsConfigCertificateStoreTypeEnum
func GetTlsConfigCertificateStoreTypeEnumStringValues() []string {
	return []string{
		"JKS",
	}
}

// GetMappingTlsConfigCertificateStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTlsConfigCertificateStoreTypeEnum(val string) (TlsConfigCertificateStoreTypeEnum, bool) {
	mappingTlsConfigCertificateStoreTypeEnumIgnoreCase := make(map[string]TlsConfigCertificateStoreTypeEnum)
	for k, v := range mappingTlsConfigCertificateStoreTypeEnum {
		mappingTlsConfigCertificateStoreTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTlsConfigCertificateStoreTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
