// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

// DeploymentBackupTypeEnum Enum with underlying type: string
type DeploymentBackupTypeEnum string

// Set of constants representing the allowable values for DeploymentBackupTypeEnum
const (
	DeploymentBackupTypeIncremental DeploymentBackupTypeEnum = "INCREMENTAL"
	DeploymentBackupTypeFull        DeploymentBackupTypeEnum = "FULL"
)

var mappingDeploymentBackupType = map[string]DeploymentBackupTypeEnum{
	"INCREMENTAL": DeploymentBackupTypeIncremental,
	"FULL":        DeploymentBackupTypeFull,
}

// GetDeploymentBackupTypeEnumValues Enumerates the set of values for DeploymentBackupTypeEnum
func GetDeploymentBackupTypeEnumValues() []DeploymentBackupTypeEnum {
	values := make([]DeploymentBackupTypeEnum, 0)
	for _, v := range mappingDeploymentBackupType {
		values = append(values, v)
	}
	return values
}
