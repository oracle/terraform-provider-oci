// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"strings"
)

// PropertySetKeyEnum Enum with underlying type: string
type PropertySetKeyEnum string

// Set of constants representing the allowable values for PropertySetKeyEnum
const (
	PropertySetKeyApexDocumentGenerator                PropertySetKeyEnum = "APEX_DOCUMENT_GENERATOR"
	PropertySetKeyApex                                 PropertySetKeyEnum = "APEX"
	PropertySetKeyApexFaIntegration                    PropertySetKeyEnum = "APEX_FA_INTEGRATION"
	PropertySetKeyOracleDatabaseExternalAuthentication PropertySetKeyEnum = "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"
)

var mappingPropertySetKeyEnum = map[string]PropertySetKeyEnum{
	"APEX_DOCUMENT_GENERATOR": PropertySetKeyApexDocumentGenerator,
	"APEX":                    PropertySetKeyApex,
	"APEX_FA_INTEGRATION":     PropertySetKeyApexFaIntegration,
	"ORACLE_DATABASE_EXTERNAL_AUTHENTICATION": PropertySetKeyOracleDatabaseExternalAuthentication,
}

var mappingPropertySetKeyEnumLowerCase = map[string]PropertySetKeyEnum{
	"apex_document_generator": PropertySetKeyApexDocumentGenerator,
	"apex":                    PropertySetKeyApex,
	"apex_fa_integration":     PropertySetKeyApexFaIntegration,
	"oracle_database_external_authentication": PropertySetKeyOracleDatabaseExternalAuthentication,
}

// GetPropertySetKeyEnumValues Enumerates the set of values for PropertySetKeyEnum
func GetPropertySetKeyEnumValues() []PropertySetKeyEnum {
	values := make([]PropertySetKeyEnum, 0)
	for _, v := range mappingPropertySetKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetPropertySetKeyEnumStringValues Enumerates the set of values in String for PropertySetKeyEnum
func GetPropertySetKeyEnumStringValues() []string {
	return []string{
		"APEX_DOCUMENT_GENERATOR",
		"APEX",
		"APEX_FA_INTEGRATION",
		"ORACLE_DATABASE_EXTERNAL_AUTHENTICATION",
	}
}

// GetMappingPropertySetKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPropertySetKeyEnum(val string) (PropertySetKeyEnum, bool) {
	enum, ok := mappingPropertySetKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
