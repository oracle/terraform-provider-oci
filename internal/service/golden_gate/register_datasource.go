// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_golden_gate_connection", GoldenGateConnectionDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_connection_assignment", GoldenGateConnectionAssignmentDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_connection_assignments", GoldenGateConnectionAssignmentsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_connections", GoldenGateConnectionsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_database_registration", GoldenGateDatabaseRegistrationDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_database_registrations", GoldenGateDatabaseRegistrationsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment", GoldenGateDeploymentDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_backup", GoldenGateDeploymentBackupDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_backups", GoldenGateDeploymentBackupsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_certificate", GoldenGateDeploymentCertificateDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_certificates", GoldenGateDeploymentCertificatesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_environments", GoldenGateDeploymentEnvironmentsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_peers", GoldenGateDeploymentPeersDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_type", GoldenGateDeploymentTypeDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_types", GoldenGateDeploymentTypesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_upgrade", GoldenGateDeploymentUpgradeDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_upgrades", GoldenGateDeploymentUpgradesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_versions", GoldenGateDeploymentVersionsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployments", GoldenGateDeploymentsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_message", GoldenGateMessageDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_messages", GoldenGateMessagesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_pipeline", GoldenGatePipelineDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_pipeline_running_processes", GoldenGatePipelineRunningProcessesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_pipeline_schema_tables", GoldenGatePipelineSchemaTablesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_pipeline_schemas", GoldenGatePipelineSchemasDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_pipelines", GoldenGatePipelinesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_recipes", GoldenGateRecipesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_trail_file", GoldenGateTrailFileDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_trail_files", GoldenGateTrailFilesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_trail_sequence", GoldenGateTrailSequenceDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_trail_sequences", GoldenGateTrailSequencesDataSource())
}
