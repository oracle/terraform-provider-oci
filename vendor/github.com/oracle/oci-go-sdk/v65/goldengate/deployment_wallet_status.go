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

// DeploymentWalletStatusEnum Enum with underlying type: string
type DeploymentWalletStatusEnum string

// Set of constants representing the allowable values for DeploymentWalletStatusEnum
const (
	DeploymentWalletStatusExporting DeploymentWalletStatusEnum = "EXPORTING"
	DeploymentWalletStatusExported  DeploymentWalletStatusEnum = "EXPORTED"
	DeploymentWalletStatusImported  DeploymentWalletStatusEnum = "IMPORTED"
	DeploymentWalletStatusImporting DeploymentWalletStatusEnum = "IMPORTING"
	DeploymentWalletStatusFailed    DeploymentWalletStatusEnum = "FAILED"
)

var mappingDeploymentWalletStatusEnum = map[string]DeploymentWalletStatusEnum{
	"EXPORTING": DeploymentWalletStatusExporting,
	"EXPORTED":  DeploymentWalletStatusExported,
	"IMPORTED":  DeploymentWalletStatusImported,
	"IMPORTING": DeploymentWalletStatusImporting,
	"FAILED":    DeploymentWalletStatusFailed,
}

var mappingDeploymentWalletStatusEnumLowerCase = map[string]DeploymentWalletStatusEnum{
	"exporting": DeploymentWalletStatusExporting,
	"exported":  DeploymentWalletStatusExported,
	"imported":  DeploymentWalletStatusImported,
	"importing": DeploymentWalletStatusImporting,
	"failed":    DeploymentWalletStatusFailed,
}

// GetDeploymentWalletStatusEnumValues Enumerates the set of values for DeploymentWalletStatusEnum
func GetDeploymentWalletStatusEnumValues() []DeploymentWalletStatusEnum {
	values := make([]DeploymentWalletStatusEnum, 0)
	for _, v := range mappingDeploymentWalletStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentWalletStatusEnumStringValues Enumerates the set of values in String for DeploymentWalletStatusEnum
func GetDeploymentWalletStatusEnumStringValues() []string {
	return []string{
		"EXPORTING",
		"EXPORTED",
		"IMPORTED",
		"IMPORTING",
		"FAILED",
	}
}

// GetMappingDeploymentWalletStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentWalletStatusEnum(val string) (DeploymentWalletStatusEnum, bool) {
	enum, ok := mappingDeploymentWalletStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
