// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RestCredential The user credential information.
type RestCredential struct {

	// The name of the user.
	Username *string `mandatory:"true" json:"username"`

	// The password of the user.
	Password *string `mandatory:"true" json:"password"`

	// The SSL truststore type.
	SslTrustStoreType RestCredentialSslTrustStoreTypeEnum `mandatory:"false" json:"sslTrustStoreType,omitempty"`

	// The full path of the SSL truststore location in the agent.
	SslTrustStoreLocation *string `mandatory:"false" json:"sslTrustStoreLocation"`

	// The password of the SSL truststore location in the agent.
	SslTrustStorePassword *string `mandatory:"false" json:"sslTrustStorePassword"`
}

func (m RestCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRestCredentialSslTrustStoreTypeEnum(string(m.SslTrustStoreType)); !ok && m.SslTrustStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SslTrustStoreType: %s. Supported values are: %s.", m.SslTrustStoreType, strings.Join(GetRestCredentialSslTrustStoreTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RestCredentialSslTrustStoreTypeEnum Enum with underlying type: string
type RestCredentialSslTrustStoreTypeEnum string

// Set of constants representing the allowable values for RestCredentialSslTrustStoreTypeEnum
const (
	RestCredentialSslTrustStoreTypeJks   RestCredentialSslTrustStoreTypeEnum = "JKS"
	RestCredentialSslTrustStoreTypeBcfks RestCredentialSslTrustStoreTypeEnum = "BCFKS"
)

var mappingRestCredentialSslTrustStoreTypeEnum = map[string]RestCredentialSslTrustStoreTypeEnum{
	"JKS":   RestCredentialSslTrustStoreTypeJks,
	"BCFKS": RestCredentialSslTrustStoreTypeBcfks,
}

var mappingRestCredentialSslTrustStoreTypeEnumLowerCase = map[string]RestCredentialSslTrustStoreTypeEnum{
	"jks":   RestCredentialSslTrustStoreTypeJks,
	"bcfks": RestCredentialSslTrustStoreTypeBcfks,
}

// GetRestCredentialSslTrustStoreTypeEnumValues Enumerates the set of values for RestCredentialSslTrustStoreTypeEnum
func GetRestCredentialSslTrustStoreTypeEnumValues() []RestCredentialSslTrustStoreTypeEnum {
	values := make([]RestCredentialSslTrustStoreTypeEnum, 0)
	for _, v := range mappingRestCredentialSslTrustStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRestCredentialSslTrustStoreTypeEnumStringValues Enumerates the set of values in String for RestCredentialSslTrustStoreTypeEnum
func GetRestCredentialSslTrustStoreTypeEnumStringValues() []string {
	return []string{
		"JKS",
		"BCFKS",
	}
}

// GetMappingRestCredentialSslTrustStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRestCredentialSslTrustStoreTypeEnum(val string) (RestCredentialSslTrustStoreTypeEnum, bool) {
	enum, ok := mappingRestCredentialSslTrustStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
