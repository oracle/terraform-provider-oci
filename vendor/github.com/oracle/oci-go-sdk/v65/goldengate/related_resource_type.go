// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// RelatedResourceTypeEnum Enum with underlying type: string
type RelatedResourceTypeEnum string

// Set of constants representing the allowable values for RelatedResourceTypeEnum
const (
	RelatedResourceTypeGoldengatedeployment RelatedResourceTypeEnum = "GOLDENGATEDEPLOYMENT"
	RelatedResourceTypeGoldengateconnection RelatedResourceTypeEnum = "GOLDENGATECONNECTION"
	RelatedResourceTypeVaultsecret          RelatedResourceTypeEnum = "VAULTSECRET"
)

var mappingRelatedResourceTypeEnum = map[string]RelatedResourceTypeEnum{
	"GOLDENGATEDEPLOYMENT": RelatedResourceTypeGoldengatedeployment,
	"GOLDENGATECONNECTION": RelatedResourceTypeGoldengateconnection,
	"VAULTSECRET":          RelatedResourceTypeVaultsecret,
}

var mappingRelatedResourceTypeEnumLowerCase = map[string]RelatedResourceTypeEnum{
	"goldengatedeployment": RelatedResourceTypeGoldengatedeployment,
	"goldengateconnection": RelatedResourceTypeGoldengateconnection,
	"vaultsecret":          RelatedResourceTypeVaultsecret,
}

// GetRelatedResourceTypeEnumValues Enumerates the set of values for RelatedResourceTypeEnum
func GetRelatedResourceTypeEnumValues() []RelatedResourceTypeEnum {
	values := make([]RelatedResourceTypeEnum, 0)
	for _, v := range mappingRelatedResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRelatedResourceTypeEnumStringValues Enumerates the set of values in String for RelatedResourceTypeEnum
func GetRelatedResourceTypeEnumStringValues() []string {
	return []string{
		"GOLDENGATEDEPLOYMENT",
		"GOLDENGATECONNECTION",
		"VAULTSECRET",
	}
}

// GetMappingRelatedResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRelatedResourceTypeEnum(val string) (RelatedResourceTypeEnum, bool) {
	enum, ok := mappingRelatedResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
