// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeGoldengateDatabaseRegistrationCreate OperationTypeEnum = "GOLDENGATE_DATABASE_REGISTRATION_CREATE"
	OperationTypeGoldengateDatabaseRegistrationUpdate OperationTypeEnum = "GOLDENGATE_DATABASE_REGISTRATION_UPDATE"
	OperationTypeGoldengateDatabaseRegistrationDelete OperationTypeEnum = "GOLDENGATE_DATABASE_REGISTRATION_DELETE"
	OperationTypeGoldengateDatabaseRegistrationMove   OperationTypeEnum = "GOLDENGATE_DATABASE_REGISTRATION_MOVE"
	OperationTypeGoldengateDeploymentCreate           OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_CREATE"
	OperationTypeGoldengateDeploymentUpdate           OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPDATE"
	OperationTypeGoldengateDeploymentDelete           OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_DELETE"
	OperationTypeGoldengateDeploymentMove             OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_MOVE"
	OperationTypeGoldengateDeploymentRestore          OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_RESTORE"
	OperationTypeGoldengateDeploymentStart            OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_START"
	OperationTypeGoldengateDeploymentStop             OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_STOP"
	OperationTypeGoldengateDeploymentPatch            OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_PATCH"
	OperationTypeGoldengateDeploymentUpgrade          OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPGRADE"
	OperationTypeGoldengateDeploymentBackupCreate     OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_BACKUP_CREATE"
	OperationTypeGoldengateDeploymentBackupDelete     OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_BACKUP_DELETE"
	OperationTypeGoldengateDeploymentBackupCancel     OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_BACKUP_CANCEL"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"GOLDENGATE_DATABASE_REGISTRATION_CREATE": OperationTypeGoldengateDatabaseRegistrationCreate,
	"GOLDENGATE_DATABASE_REGISTRATION_UPDATE": OperationTypeGoldengateDatabaseRegistrationUpdate,
	"GOLDENGATE_DATABASE_REGISTRATION_DELETE": OperationTypeGoldengateDatabaseRegistrationDelete,
	"GOLDENGATE_DATABASE_REGISTRATION_MOVE":   OperationTypeGoldengateDatabaseRegistrationMove,
	"GOLDENGATE_DEPLOYMENT_CREATE":            OperationTypeGoldengateDeploymentCreate,
	"GOLDENGATE_DEPLOYMENT_UPDATE":            OperationTypeGoldengateDeploymentUpdate,
	"GOLDENGATE_DEPLOYMENT_DELETE":            OperationTypeGoldengateDeploymentDelete,
	"GOLDENGATE_DEPLOYMENT_MOVE":              OperationTypeGoldengateDeploymentMove,
	"GOLDENGATE_DEPLOYMENT_RESTORE":           OperationTypeGoldengateDeploymentRestore,
	"GOLDENGATE_DEPLOYMENT_START":             OperationTypeGoldengateDeploymentStart,
	"GOLDENGATE_DEPLOYMENT_STOP":              OperationTypeGoldengateDeploymentStop,
	"GOLDENGATE_DEPLOYMENT_PATCH":             OperationTypeGoldengateDeploymentPatch,
	"GOLDENGATE_DEPLOYMENT_UPGRADE":           OperationTypeGoldengateDeploymentUpgrade,
	"GOLDENGATE_DEPLOYMENT_BACKUP_CREATE":     OperationTypeGoldengateDeploymentBackupCreate,
	"GOLDENGATE_DEPLOYMENT_BACKUP_DELETE":     OperationTypeGoldengateDeploymentBackupDelete,
	"GOLDENGATE_DEPLOYMENT_BACKUP_CANCEL":     OperationTypeGoldengateDeploymentBackupCancel,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
