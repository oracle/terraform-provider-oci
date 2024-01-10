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

// WalletExistsDeploymentTypeEnum Enum with underlying type: string
type WalletExistsDeploymentTypeEnum string

// Set of constants representing the allowable values for WalletExistsDeploymentTypeEnum
const (
	WalletExistsDeploymentTypeDefault WalletExistsDeploymentTypeEnum = "DEFAULT"
)

var mappingWalletExistsDeploymentTypeEnum = map[string]WalletExistsDeploymentTypeEnum{
	"DEFAULT": WalletExistsDeploymentTypeDefault,
}

var mappingWalletExistsDeploymentTypeEnumLowerCase = map[string]WalletExistsDeploymentTypeEnum{
	"default": WalletExistsDeploymentTypeDefault,
}

// GetWalletExistsDeploymentTypeEnumValues Enumerates the set of values for WalletExistsDeploymentTypeEnum
func GetWalletExistsDeploymentTypeEnumValues() []WalletExistsDeploymentTypeEnum {
	values := make([]WalletExistsDeploymentTypeEnum, 0)
	for _, v := range mappingWalletExistsDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWalletExistsDeploymentTypeEnumStringValues Enumerates the set of values in String for WalletExistsDeploymentTypeEnum
func GetWalletExistsDeploymentTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingWalletExistsDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWalletExistsDeploymentTypeEnum(val string) (WalletExistsDeploymentTypeEnum, bool) {
	enum, ok := mappingWalletExistsDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
