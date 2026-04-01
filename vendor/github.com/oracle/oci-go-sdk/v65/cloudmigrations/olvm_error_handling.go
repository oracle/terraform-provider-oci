// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmErrorHandling Error Handling in OLVM
type OlvmErrorHandling struct {

	// Migrate on error
	OnError OlvmErrorHandlingOnErrorEnum `mandatory:"false" json:"onError,omitempty"`
}

func (m OlvmErrorHandling) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmErrorHandling) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmErrorHandlingOnErrorEnum(string(m.OnError)); !ok && m.OnError != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OnError: %s. Supported values are: %s.", m.OnError, strings.Join(GetOlvmErrorHandlingOnErrorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmErrorHandlingOnErrorEnum Enum with underlying type: string
type OlvmErrorHandlingOnErrorEnum string

// Set of constants representing the allowable values for OlvmErrorHandlingOnErrorEnum
const (
	OlvmErrorHandlingOnErrorDoNotMigrate           OlvmErrorHandlingOnErrorEnum = "DO_NOT_MIGRATE"
	OlvmErrorHandlingOnErrorMigrate                OlvmErrorHandlingOnErrorEnum = "MIGRATE"
	OlvmErrorHandlingOnErrorMigrateHighlyAvailable OlvmErrorHandlingOnErrorEnum = "MIGRATE_HIGHLY_AVAILABLE"
)

var mappingOlvmErrorHandlingOnErrorEnum = map[string]OlvmErrorHandlingOnErrorEnum{
	"DO_NOT_MIGRATE":           OlvmErrorHandlingOnErrorDoNotMigrate,
	"MIGRATE":                  OlvmErrorHandlingOnErrorMigrate,
	"MIGRATE_HIGHLY_AVAILABLE": OlvmErrorHandlingOnErrorMigrateHighlyAvailable,
}

var mappingOlvmErrorHandlingOnErrorEnumLowerCase = map[string]OlvmErrorHandlingOnErrorEnum{
	"do_not_migrate":           OlvmErrorHandlingOnErrorDoNotMigrate,
	"migrate":                  OlvmErrorHandlingOnErrorMigrate,
	"migrate_highly_available": OlvmErrorHandlingOnErrorMigrateHighlyAvailable,
}

// GetOlvmErrorHandlingOnErrorEnumValues Enumerates the set of values for OlvmErrorHandlingOnErrorEnum
func GetOlvmErrorHandlingOnErrorEnumValues() []OlvmErrorHandlingOnErrorEnum {
	values := make([]OlvmErrorHandlingOnErrorEnum, 0)
	for _, v := range mappingOlvmErrorHandlingOnErrorEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmErrorHandlingOnErrorEnumStringValues Enumerates the set of values in String for OlvmErrorHandlingOnErrorEnum
func GetOlvmErrorHandlingOnErrorEnumStringValues() []string {
	return []string{
		"DO_NOT_MIGRATE",
		"MIGRATE",
		"MIGRATE_HIGHLY_AVAILABLE",
	}
}

// GetMappingOlvmErrorHandlingOnErrorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmErrorHandlingOnErrorEnum(val string) (OlvmErrorHandlingOnErrorEnum, bool) {
	enum, ok := mappingOlvmErrorHandlingOnErrorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
