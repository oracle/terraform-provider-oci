// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

// DeploymentUpgradeTypeEnum Enum with underlying type: string
type DeploymentUpgradeTypeEnum string

// Set of constants representing the allowable values for DeploymentUpgradeTypeEnum
const (
	DeploymentUpgradeTypeManual    DeploymentUpgradeTypeEnum = "MANUAL"
	DeploymentUpgradeTypeAutomatic DeploymentUpgradeTypeEnum = "AUTOMATIC"
)

var mappingDeploymentUpgradeType = map[string]DeploymentUpgradeTypeEnum{
	"MANUAL":    DeploymentUpgradeTypeManual,
	"AUTOMATIC": DeploymentUpgradeTypeAutomatic,
}

// GetDeploymentUpgradeTypeEnumValues Enumerates the set of values for DeploymentUpgradeTypeEnum
func GetDeploymentUpgradeTypeEnumValues() []DeploymentUpgradeTypeEnum {
	values := make([]DeploymentUpgradeTypeEnum, 0)
	for _, v := range mappingDeploymentUpgradeType {
		values = append(values, v)
	}
	return values
}
