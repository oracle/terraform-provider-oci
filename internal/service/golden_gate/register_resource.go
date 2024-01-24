// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_golden_gate_connection", GoldenGateConnectionResource())
	tfresource.RegisterResource("oci_golden_gate_connection_assignment", GoldenGateConnectionAssignmentResource())
	tfresource.RegisterResource("oci_golden_gate_database_registration", GoldenGateDatabaseRegistrationResource())
	tfresource.RegisterResource("oci_golden_gate_deployment", GoldenGateDeploymentResource())
	tfresource.RegisterResource("oci_golden_gate_deployment_backup", GoldenGateDeploymentBackupResource())
	tfresource.RegisterResource("oci_golden_gate_deployment_certificate", GoldenGateDeploymentCertificateResource())
}
