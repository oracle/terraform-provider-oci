// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

// RestoreDeploymentTypeEnum Enum with underlying type: string
type RestoreDeploymentTypeEnum string

// Set of constants representing the allowable values for RestoreDeploymentTypeEnum
const (
	RestoreDeploymentTypeDefault RestoreDeploymentTypeEnum = "DEFAULT"
)

var mappingRestoreDeploymentType = map[string]RestoreDeploymentTypeEnum{
	"DEFAULT": RestoreDeploymentTypeDefault,
}

// GetRestoreDeploymentTypeEnumValues Enumerates the set of values for RestoreDeploymentTypeEnum
func GetRestoreDeploymentTypeEnumValues() []RestoreDeploymentTypeEnum {
	values := make([]RestoreDeploymentTypeEnum, 0)
	for _, v := range mappingRestoreDeploymentType {
		values = append(values, v)
	}
	return values
}
