// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UnregisterCloudVmClusterPkcsDetails Details of unregistering PKCS11 driver.
type UnregisterCloudVmClusterPkcsDetails struct {

	// TDE keystore type
	TdeKeyStoreType UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum `mandatory:"true" json:"tdeKeyStoreType"`
}

func (m UnregisterCloudVmClusterPkcsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnregisterCloudVmClusterPkcsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum(string(m.TdeKeyStoreType)); !ok && m.TdeKeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TdeKeyStoreType: %s. Supported values are: %s.", m.TdeKeyStoreType, strings.Join(GetUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum Enum with underlying type: string
type UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum string

// Set of constants representing the allowable values for UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum
const (
	UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeAzure UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "AZURE"
	UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeOci   UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "OCI"
)

var mappingUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum = map[string]UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"AZURE": UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"OCI":   UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeOci,
}

var mappingUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase = map[string]UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"azure": UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"oci":   UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeOci,
}

// GetUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues Enumerates the set of values for UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues() []UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
	values := make([]UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum, 0)
	for _, v := range mappingUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues Enumerates the set of values in String for UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues() []string {
	return []string{
		"AZURE",
		"OCI",
	}
}

// GetMappingUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum(val string) (UnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum, bool) {
	enum, ok := mappingUnregisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
