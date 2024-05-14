// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"strings"
)

// DomainVerificationStatusTypeEnum Enum with underlying type: string
type DomainVerificationStatusTypeEnum string

// Set of constants representing the allowable values for DomainVerificationStatusTypeEnum
const (
	DomainVerificationStatusTypeNone     DomainVerificationStatusTypeEnum = "NONE"
	DomainVerificationStatusTypeDomainid DomainVerificationStatusTypeEnum = "DOMAINID"
	DomainVerificationStatusTypeOther    DomainVerificationStatusTypeEnum = "OTHER"
)

var mappingDomainVerificationStatusTypeEnum = map[string]DomainVerificationStatusTypeEnum{
	"NONE":     DomainVerificationStatusTypeNone,
	"DOMAINID": DomainVerificationStatusTypeDomainid,
	"OTHER":    DomainVerificationStatusTypeOther,
}

var mappingDomainVerificationStatusTypeEnumLowerCase = map[string]DomainVerificationStatusTypeEnum{
	"none":     DomainVerificationStatusTypeNone,
	"domainid": DomainVerificationStatusTypeDomainid,
	"other":    DomainVerificationStatusTypeOther,
}

// GetDomainVerificationStatusTypeEnumValues Enumerates the set of values for DomainVerificationStatusTypeEnum
func GetDomainVerificationStatusTypeEnumValues() []DomainVerificationStatusTypeEnum {
	values := make([]DomainVerificationStatusTypeEnum, 0)
	for _, v := range mappingDomainVerificationStatusTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDomainVerificationStatusTypeEnumStringValues Enumerates the set of values in String for DomainVerificationStatusTypeEnum
func GetDomainVerificationStatusTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"DOMAINID",
		"OTHER",
	}
}

// GetMappingDomainVerificationStatusTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDomainVerificationStatusTypeEnum(val string) (DomainVerificationStatusTypeEnum, bool) {
	enum, ok := mappingDomainVerificationStatusTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
