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

// UnregisterCloudAutonomousVmClusterPkcsDetails Details of unregistering PKCS11 driver.
type UnregisterCloudAutonomousVmClusterPkcsDetails struct {

	// TDE keystore type
	TdeKeyStoreType UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum `mandatory:"true" json:"tdeKeyStoreType"`
}

func (m UnregisterCloudAutonomousVmClusterPkcsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnregisterCloudAutonomousVmClusterPkcsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum(string(m.TdeKeyStoreType)); !ok && m.TdeKeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TdeKeyStoreType: %s. Supported values are: %s.", m.TdeKeyStoreType, strings.Join(GetUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum Enum with underlying type: string
type UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum string

// Set of constants representing the allowable values for UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum
const (
	UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeAzure UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "AZURE"
	UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeOci   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "OCI"
	UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeGcp   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "GCP"
	UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeAws   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum = "AWS"
)

var mappingUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum = map[string]UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"AZURE": UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"OCI":   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeOci,
	"GCP":   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeGcp,
	"AWS":   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeAws,
}

var mappingUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase = map[string]UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum{
	"azure": UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeAzure,
	"oci":   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeOci,
	"gcp":   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeGcp,
	"aws":   UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeAws,
}

// GetUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues Enumerates the set of values for UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnumValues() []UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
	values := make([]UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum, 0)
	for _, v := range mappingUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues Enumerates the set of values in String for UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum
func GetUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnumStringValues() []string {
	return []string{
		"AZURE",
		"OCI",
		"GCP",
		"AWS",
	}
}

// GetMappingUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum(val string) (UnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnum, bool) {
	enum, ok := mappingUnregisterCloudAutonomousVmClusterPkcsDetailsTdeKeyStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
