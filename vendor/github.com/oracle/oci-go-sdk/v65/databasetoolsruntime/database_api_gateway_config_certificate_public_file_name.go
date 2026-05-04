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

// DatabaseApiGatewayConfigCertificatePublicFileName Describes a certificate file to be used with SSL. Ignored if the httpsPort is 0.
type DatabaseApiGatewayConfigCertificatePublicFileName struct {

	// The format of the file
	Format DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum `mandatory:"false" json:"format,omitempty"`

	// The path to the file
	Path *string `mandatory:"false" json:"path"`
}

func (m DatabaseApiGatewayConfigCertificatePublicFileName) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseApiGatewayConfigCertificatePublicFileName) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum Enum with underlying type: string
type DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum string

// Set of constants representing the allowable values for DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum
const (
	DatabaseApiGatewayConfigCertificatePublicFileNameFormatPem DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum = "PEM"
)

var mappingDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum = map[string]DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum{
	"PEM": DatabaseApiGatewayConfigCertificatePublicFileNameFormatPem,
}

var mappingDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnumLowerCase = map[string]DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum{
	"pem": DatabaseApiGatewayConfigCertificatePublicFileNameFormatPem,
}

// GetDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnumValues Enumerates the set of values for DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum
func GetDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnumValues() []DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum {
	values := make([]DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum, 0)
	for _, v := range mappingDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnumStringValues Enumerates the set of values in String for DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum
func GetDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnumStringValues() []string {
	return []string{
		"PEM",
	}
}

// GetMappingDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum(val string) (DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum, bool) {
	enum, ok := mappingDatabaseApiGatewayConfigCertificatePublicFileNameFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
