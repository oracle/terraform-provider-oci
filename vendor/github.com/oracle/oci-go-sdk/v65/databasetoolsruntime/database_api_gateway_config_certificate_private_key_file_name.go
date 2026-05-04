// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseApiGatewayConfigCertificatePrivateKeyFileName Describes a certificate private key file to be used with SSL
type DatabaseApiGatewayConfigCertificatePrivateKeyFileName struct {

	// The format of the file
	Format DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum `mandatory:"false" json:"format,omitempty"`

	// The path to the file
	Path *string `mandatory:"false" json:"path"`
}

func (m DatabaseApiGatewayConfigCertificatePrivateKeyFileName) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseApiGatewayConfigCertificatePrivateKeyFileName) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum Enum with underlying type: string
type DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum string

// Set of constants representing the allowable values for DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum
const (
	DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatDer DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum = "DER"
	DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatPem DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum = "PEM"
)

var mappingDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum = map[string]DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum{
	"DER": DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatDer,
	"PEM": DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatPem,
}

var mappingDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnumLowerCase = map[string]DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum{
	"der": DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatDer,
	"pem": DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatPem,
}

// GetDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnumValues Enumerates the set of values for DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum
func GetDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnumValues() []DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum {
	values := make([]DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum, 0)
	for _, v := range mappingDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnumStringValues Enumerates the set of values in String for DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum
func GetDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnumStringValues() []string {
	return []string{
		"DER",
		"PEM",
	}
}

// GetMappingDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum(val string) (DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum, bool) {
	enum, ok := mappingDatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
