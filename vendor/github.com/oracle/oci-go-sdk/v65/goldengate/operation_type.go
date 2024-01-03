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

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeGoldengateDatabaseRegistrationCreate  OperationTypeEnum = "GOLDENGATE_DATABASE_REGISTRATION_CREATE"
	OperationTypeGoldengateDatabaseRegistrationUpdate  OperationTypeEnum = "GOLDENGATE_DATABASE_REGISTRATION_UPDATE"
	OperationTypeGoldengateDatabaseRegistrationDelete  OperationTypeEnum = "GOLDENGATE_DATABASE_REGISTRATION_DELETE"
	OperationTypeGoldengateDatabaseRegistrationMove    OperationTypeEnum = "GOLDENGATE_DATABASE_REGISTRATION_MOVE"
	OperationTypeGoldengateDeploymentCreate            OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_CREATE"
	OperationTypeGoldengateDeploymentUpdate            OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPDATE"
	OperationTypeGoldengateDeploymentDelete            OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_DELETE"
	OperationTypeGoldengateDeploymentMove              OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_MOVE"
	OperationTypeGoldengateDeploymentRestore           OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_RESTORE"
	OperationTypeGoldengateDeploymentStart             OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_START"
	OperationTypeGoldengateDeploymentStop              OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_STOP"
	OperationTypeGoldengateDeploymentUpgrade           OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPGRADE"
	OperationTypeGoldengateDeploymentBackupCreate      OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_BACKUP_CREATE"
	OperationTypeGoldengateDeploymentBackupDelete      OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_BACKUP_DELETE"
	OperationTypeGoldengateDeploymentBackupCancel      OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_BACKUP_CANCEL"
	OperationTypeGoldengateDeploymentBackupCopy        OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_BACKUP_COPY"
	OperationTypeGoldengateConnectionCreate            OperationTypeEnum = "GOLDENGATE_CONNECTION_CREATE"
	OperationTypeGoldengateConnectionUpdate            OperationTypeEnum = "GOLDENGATE_CONNECTION_UPDATE"
	OperationTypeGoldengateConnectionDelete            OperationTypeEnum = "GOLDENGATE_CONNECTION_DELETE"
	OperationTypeGoldengateConnectionMove              OperationTypeEnum = "GOLDENGATE_CONNECTION_MOVE"
	OperationTypeGoldengateConnectionAssignmentCreate  OperationTypeEnum = "GOLDENGATE_CONNECTION_ASSIGNMENT_CREATE"
	OperationTypeGoldengateConnectionAssigmnentDelete  OperationTypeEnum = "GOLDENGATE_CONNECTION_ASSIGMNENT_DELETE"
	OperationTypeGoldengateDeploymentDiagnosticCollect OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_DIAGNOSTIC_COLLECT"
	OperationTypeGoldengateDeploymentWalletExport      OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_WALLET_EXPORT"
	OperationTypeGoldengateDeploymentWalletImport      OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_WALLET_IMPORT"
	OperationTypeGoldengateDeploymentUpgradeUpgrade    OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPGRADE_UPGRADE"
	OperationTypeGoldengateDeploymentUpgradeRollback   OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPGRADE_ROLLBACK"
	OperationTypeGoldengateDeploymentUpgradeSnooze     OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPGRADE_SNOOZE"
	OperationTypeGoldengateDeploymentCertificateCreate OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_CERTIFICATE_CREATE"
	OperationTypeGoldengateDeploymentCertificateDelete OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_CERTIFICATE_DELETE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"GOLDENGATE_DATABASE_REGISTRATION_CREATE":  OperationTypeGoldengateDatabaseRegistrationCreate,
	"GOLDENGATE_DATABASE_REGISTRATION_UPDATE":  OperationTypeGoldengateDatabaseRegistrationUpdate,
	"GOLDENGATE_DATABASE_REGISTRATION_DELETE":  OperationTypeGoldengateDatabaseRegistrationDelete,
	"GOLDENGATE_DATABASE_REGISTRATION_MOVE":    OperationTypeGoldengateDatabaseRegistrationMove,
	"GOLDENGATE_DEPLOYMENT_CREATE":             OperationTypeGoldengateDeploymentCreate,
	"GOLDENGATE_DEPLOYMENT_UPDATE":             OperationTypeGoldengateDeploymentUpdate,
	"GOLDENGATE_DEPLOYMENT_DELETE":             OperationTypeGoldengateDeploymentDelete,
	"GOLDENGATE_DEPLOYMENT_MOVE":               OperationTypeGoldengateDeploymentMove,
	"GOLDENGATE_DEPLOYMENT_RESTORE":            OperationTypeGoldengateDeploymentRestore,
	"GOLDENGATE_DEPLOYMENT_START":              OperationTypeGoldengateDeploymentStart,
	"GOLDENGATE_DEPLOYMENT_STOP":               OperationTypeGoldengateDeploymentStop,
	"GOLDENGATE_DEPLOYMENT_UPGRADE":            OperationTypeGoldengateDeploymentUpgrade,
	"GOLDENGATE_DEPLOYMENT_BACKUP_CREATE":      OperationTypeGoldengateDeploymentBackupCreate,
	"GOLDENGATE_DEPLOYMENT_BACKUP_DELETE":      OperationTypeGoldengateDeploymentBackupDelete,
	"GOLDENGATE_DEPLOYMENT_BACKUP_CANCEL":      OperationTypeGoldengateDeploymentBackupCancel,
	"GOLDENGATE_DEPLOYMENT_BACKUP_COPY":        OperationTypeGoldengateDeploymentBackupCopy,
	"GOLDENGATE_CONNECTION_CREATE":             OperationTypeGoldengateConnectionCreate,
	"GOLDENGATE_CONNECTION_UPDATE":             OperationTypeGoldengateConnectionUpdate,
	"GOLDENGATE_CONNECTION_DELETE":             OperationTypeGoldengateConnectionDelete,
	"GOLDENGATE_CONNECTION_MOVE":               OperationTypeGoldengateConnectionMove,
	"GOLDENGATE_CONNECTION_ASSIGNMENT_CREATE":  OperationTypeGoldengateConnectionAssignmentCreate,
	"GOLDENGATE_CONNECTION_ASSIGMNENT_DELETE":  OperationTypeGoldengateConnectionAssigmnentDelete,
	"GOLDENGATE_DEPLOYMENT_DIAGNOSTIC_COLLECT": OperationTypeGoldengateDeploymentDiagnosticCollect,
	"GOLDENGATE_DEPLOYMENT_WALLET_EXPORT":      OperationTypeGoldengateDeploymentWalletExport,
	"GOLDENGATE_DEPLOYMENT_WALLET_IMPORT":      OperationTypeGoldengateDeploymentWalletImport,
	"GOLDENGATE_DEPLOYMENT_UPGRADE_UPGRADE":    OperationTypeGoldengateDeploymentUpgradeUpgrade,
	"GOLDENGATE_DEPLOYMENT_UPGRADE_ROLLBACK":   OperationTypeGoldengateDeploymentUpgradeRollback,
	"GOLDENGATE_DEPLOYMENT_UPGRADE_SNOOZE":     OperationTypeGoldengateDeploymentUpgradeSnooze,
	"GOLDENGATE_DEPLOYMENT_CERTIFICATE_CREATE": OperationTypeGoldengateDeploymentCertificateCreate,
	"GOLDENGATE_DEPLOYMENT_CERTIFICATE_DELETE": OperationTypeGoldengateDeploymentCertificateDelete,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"goldengate_database_registration_create":  OperationTypeGoldengateDatabaseRegistrationCreate,
	"goldengate_database_registration_update":  OperationTypeGoldengateDatabaseRegistrationUpdate,
	"goldengate_database_registration_delete":  OperationTypeGoldengateDatabaseRegistrationDelete,
	"goldengate_database_registration_move":    OperationTypeGoldengateDatabaseRegistrationMove,
	"goldengate_deployment_create":             OperationTypeGoldengateDeploymentCreate,
	"goldengate_deployment_update":             OperationTypeGoldengateDeploymentUpdate,
	"goldengate_deployment_delete":             OperationTypeGoldengateDeploymentDelete,
	"goldengate_deployment_move":               OperationTypeGoldengateDeploymentMove,
	"goldengate_deployment_restore":            OperationTypeGoldengateDeploymentRestore,
	"goldengate_deployment_start":              OperationTypeGoldengateDeploymentStart,
	"goldengate_deployment_stop":               OperationTypeGoldengateDeploymentStop,
	"goldengate_deployment_upgrade":            OperationTypeGoldengateDeploymentUpgrade,
	"goldengate_deployment_backup_create":      OperationTypeGoldengateDeploymentBackupCreate,
	"goldengate_deployment_backup_delete":      OperationTypeGoldengateDeploymentBackupDelete,
	"goldengate_deployment_backup_cancel":      OperationTypeGoldengateDeploymentBackupCancel,
	"goldengate_deployment_backup_copy":        OperationTypeGoldengateDeploymentBackupCopy,
	"goldengate_connection_create":             OperationTypeGoldengateConnectionCreate,
	"goldengate_connection_update":             OperationTypeGoldengateConnectionUpdate,
	"goldengate_connection_delete":             OperationTypeGoldengateConnectionDelete,
	"goldengate_connection_move":               OperationTypeGoldengateConnectionMove,
	"goldengate_connection_assignment_create":  OperationTypeGoldengateConnectionAssignmentCreate,
	"goldengate_connection_assigmnent_delete":  OperationTypeGoldengateConnectionAssigmnentDelete,
	"goldengate_deployment_diagnostic_collect": OperationTypeGoldengateDeploymentDiagnosticCollect,
	"goldengate_deployment_wallet_export":      OperationTypeGoldengateDeploymentWalletExport,
	"goldengate_deployment_wallet_import":      OperationTypeGoldengateDeploymentWalletImport,
	"goldengate_deployment_upgrade_upgrade":    OperationTypeGoldengateDeploymentUpgradeUpgrade,
	"goldengate_deployment_upgrade_rollback":   OperationTypeGoldengateDeploymentUpgradeRollback,
	"goldengate_deployment_upgrade_snooze":     OperationTypeGoldengateDeploymentUpgradeSnooze,
	"goldengate_deployment_certificate_create": OperationTypeGoldengateDeploymentCertificateCreate,
	"goldengate_deployment_certificate_delete": OperationTypeGoldengateDeploymentCertificateDelete,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"GOLDENGATE_DATABASE_REGISTRATION_CREATE",
		"GOLDENGATE_DATABASE_REGISTRATION_UPDATE",
		"GOLDENGATE_DATABASE_REGISTRATION_DELETE",
		"GOLDENGATE_DATABASE_REGISTRATION_MOVE",
		"GOLDENGATE_DEPLOYMENT_CREATE",
		"GOLDENGATE_DEPLOYMENT_UPDATE",
		"GOLDENGATE_DEPLOYMENT_DELETE",
		"GOLDENGATE_DEPLOYMENT_MOVE",
		"GOLDENGATE_DEPLOYMENT_RESTORE",
		"GOLDENGATE_DEPLOYMENT_START",
		"GOLDENGATE_DEPLOYMENT_STOP",
		"GOLDENGATE_DEPLOYMENT_UPGRADE",
		"GOLDENGATE_DEPLOYMENT_BACKUP_CREATE",
		"GOLDENGATE_DEPLOYMENT_BACKUP_DELETE",
		"GOLDENGATE_DEPLOYMENT_BACKUP_CANCEL",
		"GOLDENGATE_DEPLOYMENT_BACKUP_COPY",
		"GOLDENGATE_CONNECTION_CREATE",
		"GOLDENGATE_CONNECTION_UPDATE",
		"GOLDENGATE_CONNECTION_DELETE",
		"GOLDENGATE_CONNECTION_MOVE",
		"GOLDENGATE_CONNECTION_ASSIGNMENT_CREATE",
		"GOLDENGATE_CONNECTION_ASSIGMNENT_DELETE",
		"GOLDENGATE_DEPLOYMENT_DIAGNOSTIC_COLLECT",
		"GOLDENGATE_DEPLOYMENT_WALLET_EXPORT",
		"GOLDENGATE_DEPLOYMENT_WALLET_IMPORT",
		"GOLDENGATE_DEPLOYMENT_UPGRADE_UPGRADE",
		"GOLDENGATE_DEPLOYMENT_UPGRADE_ROLLBACK",
		"GOLDENGATE_DEPLOYMENT_UPGRADE_SNOOZE",
		"GOLDENGATE_DEPLOYMENT_CERTIFICATE_CREATE",
		"GOLDENGATE_DEPLOYMENT_CERTIFICATE_DELETE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
