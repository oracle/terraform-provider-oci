// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_golden_gate_database_registration", GoldenGateDatabaseRegistrationDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_database_registrations", GoldenGateDatabaseRegistrationsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment", GoldenGateDeploymentDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_backup", GoldenGateDeploymentBackupDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_backups", GoldenGateDeploymentBackupsDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_upgrade", GoldenGateDeploymentUpgradeDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployment_upgrades", GoldenGateDeploymentUpgradesDataSource())
	tfresource.RegisterDatasource("oci_golden_gate_deployments", GoldenGateDeploymentsDataSource())
}
