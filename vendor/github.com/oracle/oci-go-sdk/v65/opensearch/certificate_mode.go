// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// CertificateModeEnum Enum with underlying type: string
type CertificateModeEnum string

// Set of constants representing the allowable values for CertificateModeEnum
const (
	CertificateModeOciCertificatesService CertificateModeEnum = "OCI_CERTIFICATES_SERVICE"
	CertificateModeOpensearchService      CertificateModeEnum = "OPENSEARCH_SERVICE"
)

var mappingCertificateModeEnum = map[string]CertificateModeEnum{
	"OCI_CERTIFICATES_SERVICE": CertificateModeOciCertificatesService,
	"OPENSEARCH_SERVICE":       CertificateModeOpensearchService,
}

var mappingCertificateModeEnumLowerCase = map[string]CertificateModeEnum{
	"oci_certificates_service": CertificateModeOciCertificatesService,
	"opensearch_service":       CertificateModeOpensearchService,
}

// GetCertificateModeEnumValues Enumerates the set of values for CertificateModeEnum
func GetCertificateModeEnumValues() []CertificateModeEnum {
	values := make([]CertificateModeEnum, 0)
	for _, v := range mappingCertificateModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateModeEnumStringValues Enumerates the set of values in String for CertificateModeEnum
func GetCertificateModeEnumStringValues() []string {
	return []string{
		"OCI_CERTIFICATES_SERVICE",
		"OPENSEARCH_SERVICE",
	}
}

// GetMappingCertificateModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateModeEnum(val string) (CertificateModeEnum, bool) {
	enum, ok := mappingCertificateModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
