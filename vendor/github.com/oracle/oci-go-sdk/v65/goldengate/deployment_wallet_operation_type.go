// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DeploymentWalletOperationTypeEnum Enum with underlying type: string
type DeploymentWalletOperationTypeEnum string

// Set of constants representing the allowable values for DeploymentWalletOperationTypeEnum
const (
	DeploymentWalletOperationTypeExport DeploymentWalletOperationTypeEnum = "EXPORT"
	DeploymentWalletOperationTypeImport DeploymentWalletOperationTypeEnum = "IMPORT"
)

var mappingDeploymentWalletOperationTypeEnum = map[string]DeploymentWalletOperationTypeEnum{
	"EXPORT": DeploymentWalletOperationTypeExport,
	"IMPORT": DeploymentWalletOperationTypeImport,
}

var mappingDeploymentWalletOperationTypeEnumLowerCase = map[string]DeploymentWalletOperationTypeEnum{
	"export": DeploymentWalletOperationTypeExport,
	"import": DeploymentWalletOperationTypeImport,
}

// GetDeploymentWalletOperationTypeEnumValues Enumerates the set of values for DeploymentWalletOperationTypeEnum
func GetDeploymentWalletOperationTypeEnumValues() []DeploymentWalletOperationTypeEnum {
	values := make([]DeploymentWalletOperationTypeEnum, 0)
	for _, v := range mappingDeploymentWalletOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentWalletOperationTypeEnumStringValues Enumerates the set of values in String for DeploymentWalletOperationTypeEnum
func GetDeploymentWalletOperationTypeEnumStringValues() []string {
	return []string{
		"EXPORT",
		"IMPORT",
	}
}

// GetMappingDeploymentWalletOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentWalletOperationTypeEnum(val string) (DeploymentWalletOperationTypeEnum, bool) {
	enum, ok := mappingDeploymentWalletOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
