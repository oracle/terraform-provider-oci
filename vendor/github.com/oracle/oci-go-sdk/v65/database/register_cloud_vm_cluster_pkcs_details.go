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

// RegisterCloudVmClusterPkcsDetails Details of registering PKCS11 driver.
type RegisterCloudVmClusterPkcsDetails struct {

	// TDE keystore type
	TdeKeyStoreType RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum `mandatory:"true" json:"tdeKeyStoreType"`
}

func (m RegisterCloudVmClusterPkcsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegisterCloudVmClusterPkcsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum(string(m.TdeKeyStoreType)); !ok && m.TdeKeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TdeKeyStoreType: %s. Supported values are: %s.", m.TdeKeyStoreType, strings.Join(GetRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum Enum with underlying type: string
type RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum string

// Set of constants representing the allowable values for RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum
const (
	RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeAzure RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "AZURE"
	RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeOci   RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "OCI"
	RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeGcp   RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "GCP"
)

var mappingRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum = map[string]RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"AZURE": RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"OCI":   RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeOci,
	"GCP":   RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeGcp,
}

var mappingRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase = map[string]RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"azure": RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"oci":   RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeOci,
	"gcp":   RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeGcp,
}

// GetRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues Enumerates the set of values for RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues() []RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
	values := make([]RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum, 0)
	for _, v := range mappingRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues Enumerates the set of values in String for RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues() []string {
	return []string{
		"AZURE",
		"OCI",
		"GCP",
	}
}

// GetMappingRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum(val string) (RegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnum, bool) {
	enum, ok := mappingRegisterCloudVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
