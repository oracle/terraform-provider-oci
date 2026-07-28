// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RegisterExadbVmClusterPkcsDetails Details of registering PKCS11 driver.
type RegisterExadbVmClusterPkcsDetails struct {

	// TDE keystore type
	TdeKeyStoreType RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum `mandatory:"true" json:"tdeKeyStoreType"`
}

func (m RegisterExadbVmClusterPkcsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegisterExadbVmClusterPkcsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum(string(m.TdeKeyStoreType)); !ok && m.TdeKeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TdeKeyStoreType: %s. Supported values are: %s.", m.TdeKeyStoreType, strings.Join(GetRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum Enum with underlying type: string
type RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum string

// Set of constants representing the allowable values for RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum
const (
	RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAzure RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "AZURE"
	RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeOci   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "OCI"
	RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeGcp   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "GCP"
	RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAws   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "AWS"
)

var mappingRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = map[string]RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"AZURE": RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"OCI":   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeOci,
	"GCP":   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeGcp,
	"AWS":   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAws,
}

var mappingRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase = map[string]RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"azure": RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"oci":   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeOci,
	"gcp":   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeGcp,
	"aws":   RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAws,
}

// GetRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues Enumerates the set of values for RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues() []RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
	values := make([]RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum, 0)
	for _, v := range mappingRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues Enumerates the set of values in String for RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues() []string {
	return []string{
		"AZURE",
		"OCI",
		"GCP",
		"AWS",
	}
}

// GetMappingRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum(val string) (RegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum, bool) {
	enum, ok := mappingRegisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
