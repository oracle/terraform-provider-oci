// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FetchProtectedDatabaseConfigurationDetails Provides which configuration details to get.
type FetchProtectedDatabaseConfigurationDetails struct {

	// Currently has four config options ALL, TNSNAMES, HOSTS and CABUNDLE. All will return a zipped folder containing the contents of both tnsnames and the certificateChainPem.
	ConfigurationType FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum `mandatory:"false" json:"configurationType,omitempty"`
}

func (m FetchProtectedDatabaseConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FetchProtectedDatabaseConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum(string(m.ConfigurationType)); !ok && m.ConfigurationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigurationType: %s. Supported values are: %s.", m.ConfigurationType, strings.Join(GetFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum Enum with underlying type: string
type FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum string

// Set of constants representing the allowable values for FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum
const (
	FetchProtectedDatabaseConfigurationDetailsConfigurationTypeCabundle FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum = "CABUNDLE"
	FetchProtectedDatabaseConfigurationDetailsConfigurationTypeTnsnames FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum = "TNSNAMES"
	FetchProtectedDatabaseConfigurationDetailsConfigurationTypeHosts    FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum = "HOSTS"
	FetchProtectedDatabaseConfigurationDetailsConfigurationTypeAll      FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum = "ALL"
)

var mappingFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum = map[string]FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum{
	"CABUNDLE": FetchProtectedDatabaseConfigurationDetailsConfigurationTypeCabundle,
	"TNSNAMES": FetchProtectedDatabaseConfigurationDetailsConfigurationTypeTnsnames,
	"HOSTS":    FetchProtectedDatabaseConfigurationDetailsConfigurationTypeHosts,
	"ALL":      FetchProtectedDatabaseConfigurationDetailsConfigurationTypeAll,
}

var mappingFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnumLowerCase = map[string]FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum{
	"cabundle": FetchProtectedDatabaseConfigurationDetailsConfigurationTypeCabundle,
	"tnsnames": FetchProtectedDatabaseConfigurationDetailsConfigurationTypeTnsnames,
	"hosts":    FetchProtectedDatabaseConfigurationDetailsConfigurationTypeHosts,
	"all":      FetchProtectedDatabaseConfigurationDetailsConfigurationTypeAll,
}

// GetFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnumValues Enumerates the set of values for FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum
func GetFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnumValues() []FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum {
	values := make([]FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum, 0)
	for _, v := range mappingFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnumStringValues Enumerates the set of values in String for FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum
func GetFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnumStringValues() []string {
	return []string{
		"CABUNDLE",
		"TNSNAMES",
		"HOSTS",
		"ALL",
	}
}

// GetMappingFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum(val string) (FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum, bool) {
	enum, ok := mappingFetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
