// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_recovery_protected_database", RecoveryProtectedDatabaseDataSource())
	tfresource.RegisterDatasource("oci_recovery_protected_database_fetch_configuration", RecoveryProtectedDatabaseFetchConfigurationDataSource())
	tfresource.RegisterDatasource("oci_recovery_protected_databases", RecoveryProtectedDatabasesDataSource())
	tfresource.RegisterDatasource("oci_recovery_protection_policies", RecoveryProtectionPoliciesDataSource())
	tfresource.RegisterDatasource("oci_recovery_protection_policy", RecoveryProtectionPolicyDataSource())
	tfresource.RegisterDatasource("oci_recovery_recovery_service_subnet", RecoveryRecoveryServiceSubnetDataSource())
	tfresource.RegisterDatasource("oci_recovery_recovery_service_subnets", RecoveryRecoveryServiceSubnetsDataSource())
}
