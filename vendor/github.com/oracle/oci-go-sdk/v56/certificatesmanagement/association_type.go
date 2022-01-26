// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

// AssociationTypeEnum Enum with underlying type: string
type AssociationTypeEnum string

// Set of constants representing the allowable values for AssociationTypeEnum
const (
	AssociationTypeCertificate          AssociationTypeEnum = "CERTIFICATE"
	AssociationTypeCertificateAuthority AssociationTypeEnum = "CERTIFICATE_AUTHORITY"
	AssociationTypeCaBundle             AssociationTypeEnum = "CA_BUNDLE"
)

var mappingAssociationType = map[string]AssociationTypeEnum{
	"CERTIFICATE":           AssociationTypeCertificate,
	"CERTIFICATE_AUTHORITY": AssociationTypeCertificateAuthority,
	"CA_BUNDLE":             AssociationTypeCaBundle,
}

// GetAssociationTypeEnumValues Enumerates the set of values for AssociationTypeEnum
func GetAssociationTypeEnumValues() []AssociationTypeEnum {
	values := make([]AssociationTypeEnum, 0)
	for _, v := range mappingAssociationType {
		values = append(values, v)
	}
	return values
}
