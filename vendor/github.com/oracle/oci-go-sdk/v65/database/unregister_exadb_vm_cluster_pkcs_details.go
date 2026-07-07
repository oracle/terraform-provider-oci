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

// UnregisterExadbVmClusterPkcsDetails Details of unregistering PKCS11 driver.
type UnregisterExadbVmClusterPkcsDetails struct {

	// TDE keystore type
	TdeKeyStoreType UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum `mandatory:"true" json:"tdeKeyStoreType"`
}

func (m UnregisterExadbVmClusterPkcsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnregisterExadbVmClusterPkcsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum(string(m.TdeKeyStoreType)); !ok && m.TdeKeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TdeKeyStoreType: %s. Supported values are: %s.", m.TdeKeyStoreType, strings.Join(GetUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum Enum with underlying type: string
type UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum string

// Set of constants representing the allowable values for UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum
const (
	UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAzure UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "AZURE"
	UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeOci   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "OCI"
	UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeGcp   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "GCP"
	UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAws   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "AWS"
)

var mappingUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum = map[string]UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"AZURE": UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"OCI":   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeOci,
	"GCP":   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeGcp,
	"AWS":   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAws,
}

var mappingUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase = map[string]UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"azure": UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"oci":   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeOci,
	"gcp":   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeGcp,
	"aws":   UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeAws,
}

// GetUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues Enumerates the set of values for UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues() []UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
	values := make([]UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum, 0)
	for _, v := range mappingUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues Enumerates the set of values in String for UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues() []string {
	return []string{
		"AZURE",
		"OCI",
		"GCP",
		"AWS",
	}
}

// GetMappingUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum(val string) (UnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnum, bool) {
	enum, ok := mappingUnregisterExadbVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
