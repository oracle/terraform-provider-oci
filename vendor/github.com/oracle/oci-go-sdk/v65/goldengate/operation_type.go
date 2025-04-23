// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	OperationTypeGoldengateConnectionRefresh           OperationTypeEnum = "GOLDENGATE_CONNECTION_REFRESH"
	OperationTypeGoldengateConnectionAssignmentCreate  OperationTypeEnum = "GOLDENGATE_CONNECTION_ASSIGNMENT_CREATE"
	OperationTypeGoldengateConnectionAssignmentDelete  OperationTypeEnum = "GOLDENGATE_CONNECTION_ASSIGNMENT_DELETE"
	OperationTypeGoldengateConnectionAssigmnentDelete  OperationTypeEnum = "GOLDENGATE_CONNECTION_ASSIGMNENT_DELETE"
	OperationTypeGoldengateDeploymentDiagnosticCollect OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_DIAGNOSTIC_COLLECT"
	OperationTypeGoldengateDeploymentWalletExport      OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_WALLET_EXPORT"
	OperationTypeGoldengateDeploymentWalletImport      OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_WALLET_IMPORT"
	OperationTypeGoldengateDeploymentUpgradeUpgrade    OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPGRADE_UPGRADE"
	OperationTypeGoldengateDeploymentUpgradeRollback   OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPGRADE_ROLLBACK"
	OperationTypeGoldengateDeploymentUpgradeSnooze     OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_UPGRADE_SNOOZE"
	OperationTypeGoldengateDeploymentCertificateCreate OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_CERTIFICATE_CREATE"
	OperationTypeGoldengateDeploymentCertificateDelete OperationTypeEnum = "GOLDENGATE_DEPLOYMENT_CERTIFICATE_DELETE"
	OperationTypeGoldengatePipelineCreate              OperationTypeEnum = "GOLDENGATE_PIPELINE_CREATE"
	OperationTypeGoldengatePipelineStart               OperationTypeEnum = "GOLDENGATE_PIPELINE_START"
	OperationTypeGoldengatePipelineStop                OperationTypeEnum = "GOLDENGATE_PIPELINE_STOP"
	OperationTypeGoldengatePipelineUpdate              OperationTypeEnum = "GOLDENGATE_PIPELINE_UPDATE"
	OperationTypeGoldengatePipelineDelete              OperationTypeEnum = "GOLDENGATE_PIPELINE_DELETE"
	OperationTypeGoldengatePipelineMove                OperationTypeEnum = "GOLDENGATE_PIPELINE_MOVE"
	OperationTypeGoldengatePipelineDiagnosticsCollect  OperationTypeEnum = "GOLDENGATE_PIPELINE_DIAGNOSTICS_COLLECT"
	OperationTypeGoldengateSwitchoverDeploymentPeer    OperationTypeEnum = "GOLDENGATE_SWITCHOVER_DEPLOYMENT_PEER"
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
	"GOLDENGATE_CONNECTION_REFRESH":            OperationTypeGoldengateConnectionRefresh,
	"GOLDENGATE_CONNECTION_ASSIGNMENT_CREATE":  OperationTypeGoldengateConnectionAssignmentCreate,
	"GOLDENGATE_CONNECTION_ASSIGNMENT_DELETE":  OperationTypeGoldengateConnectionAssignmentDelete,
	"GOLDENGATE_CONNECTION_ASSIGMNENT_DELETE":  OperationTypeGoldengateConnectionAssigmnentDelete,
	"GOLDENGATE_DEPLOYMENT_DIAGNOSTIC_COLLECT": OperationTypeGoldengateDeploymentDiagnosticCollect,
	"GOLDENGATE_DEPLOYMENT_WALLET_EXPORT":      OperationTypeGoldengateDeploymentWalletExport,
	"GOLDENGATE_DEPLOYMENT_WALLET_IMPORT":      OperationTypeGoldengateDeploymentWalletImport,
	"GOLDENGATE_DEPLOYMENT_UPGRADE_UPGRADE":    OperationTypeGoldengateDeploymentUpgradeUpgrade,
	"GOLDENGATE_DEPLOYMENT_UPGRADE_ROLLBACK":   OperationTypeGoldengateDeploymentUpgradeRollback,
	"GOLDENGATE_DEPLOYMENT_UPGRADE_SNOOZE":     OperationTypeGoldengateDeploymentUpgradeSnooze,
	"GOLDENGATE_DEPLOYMENT_CERTIFICATE_CREATE": OperationTypeGoldengateDeploymentCertificateCreate,
	"GOLDENGATE_DEPLOYMENT_CERTIFICATE_DELETE": OperationTypeGoldengateDeploymentCertificateDelete,
	"GOLDENGATE_PIPELINE_CREATE":               OperationTypeGoldengatePipelineCreate,
	"GOLDENGATE_PIPELINE_START":                OperationTypeGoldengatePipelineStart,
	"GOLDENGATE_PIPELINE_STOP":                 OperationTypeGoldengatePipelineStop,
	"GOLDENGATE_PIPELINE_UPDATE":               OperationTypeGoldengatePipelineUpdate,
	"GOLDENGATE_PIPELINE_DELETE":               OperationTypeGoldengatePipelineDelete,
	"GOLDENGATE_PIPELINE_MOVE":                 OperationTypeGoldengatePipelineMove,
	"GOLDENGATE_PIPELINE_DIAGNOSTICS_COLLECT":  OperationTypeGoldengatePipelineDiagnosticsCollect,
	"GOLDENGATE_SWITCHOVER_DEPLOYMENT_PEER":    OperationTypeGoldengateSwitchoverDeploymentPeer,
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
	"goldengate_connection_refresh":            OperationTypeGoldengateConnectionRefresh,
	"goldengate_connection_assignment_create":  OperationTypeGoldengateConnectionAssignmentCreate,
	"goldengate_connection_assignment_delete":  OperationTypeGoldengateConnectionAssignmentDelete,
	"goldengate_connection_assigmnent_delete":  OperationTypeGoldengateConnectionAssigmnentDelete,
	"goldengate_deployment_diagnostic_collect": OperationTypeGoldengateDeploymentDiagnosticCollect,
	"goldengate_deployment_wallet_export":      OperationTypeGoldengateDeploymentWalletExport,
	"goldengate_deployment_wallet_import":      OperationTypeGoldengateDeploymentWalletImport,
	"goldengate_deployment_upgrade_upgrade":    OperationTypeGoldengateDeploymentUpgradeUpgrade,
	"goldengate_deployment_upgrade_rollback":   OperationTypeGoldengateDeploymentUpgradeRollback,
	"goldengate_deployment_upgrade_snooze":     OperationTypeGoldengateDeploymentUpgradeSnooze,
	"goldengate_deployment_certificate_create": OperationTypeGoldengateDeploymentCertificateCreate,
	"goldengate_deployment_certificate_delete": OperationTypeGoldengateDeploymentCertificateDelete,
	"goldengate_pipeline_create":               OperationTypeGoldengatePipelineCreate,
	"goldengate_pipeline_start":                OperationTypeGoldengatePipelineStart,
	"goldengate_pipeline_stop":                 OperationTypeGoldengatePipelineStop,
	"goldengate_pipeline_update":               OperationTypeGoldengatePipelineUpdate,
	"goldengate_pipeline_delete":               OperationTypeGoldengatePipelineDelete,
	"goldengate_pipeline_move":                 OperationTypeGoldengatePipelineMove,
	"goldengate_pipeline_diagnostics_collect":  OperationTypeGoldengatePipelineDiagnosticsCollect,
	"goldengate_switchover_deployment_peer":    OperationTypeGoldengateSwitchoverDeploymentPeer,
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
		"GOLDENGATE_CONNECTION_REFRESH",
		"GOLDENGATE_CONNECTION_ASSIGNMENT_CREATE",
		"GOLDENGATE_CONNECTION_ASSIGNMENT_DELETE",
		"GOLDENGATE_CONNECTION_ASSIGMNENT_DELETE",
		"GOLDENGATE_DEPLOYMENT_DIAGNOSTIC_COLLECT",
		"GOLDENGATE_DEPLOYMENT_WALLET_EXPORT",
		"GOLDENGATE_DEPLOYMENT_WALLET_IMPORT",
		"GOLDENGATE_DEPLOYMENT_UPGRADE_UPGRADE",
		"GOLDENGATE_DEPLOYMENT_UPGRADE_ROLLBACK",
		"GOLDENGATE_DEPLOYMENT_UPGRADE_SNOOZE",
		"GOLDENGATE_DEPLOYMENT_CERTIFICATE_CREATE",
		"GOLDENGATE_DEPLOYMENT_CERTIFICATE_DELETE",
		"GOLDENGATE_PIPELINE_CREATE",
		"GOLDENGATE_PIPELINE_START",
		"GOLDENGATE_PIPELINE_STOP",
		"GOLDENGATE_PIPELINE_UPDATE",
		"GOLDENGATE_PIPELINE_DELETE",
		"GOLDENGATE_PIPELINE_MOVE",
		"GOLDENGATE_PIPELINE_DIAGNOSTICS_COLLECT",
		"GOLDENGATE_SWITCHOVER_DEPLOYMENT_PEER",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
