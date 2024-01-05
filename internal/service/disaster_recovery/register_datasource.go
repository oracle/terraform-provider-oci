// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_disaster_recovery_dr_plan", DisasterRecoveryDrPlanDataSource())
	tfresource.RegisterDatasource("oci_disaster_recovery_dr_plan_execution", DisasterRecoveryDrPlanExecutionDataSource())
	tfresource.RegisterDatasource("oci_disaster_recovery_dr_plan_executions", DisasterRecoveryDrPlanExecutionsDataSource())
	tfresource.RegisterDatasource("oci_disaster_recovery_dr_plans", DisasterRecoveryDrPlansDataSource())
	tfresource.RegisterDatasource("oci_disaster_recovery_dr_protection_group", DisasterRecoveryDrProtectionGroupDataSource())
	tfresource.RegisterDatasource("oci_disaster_recovery_dr_protection_groups", DisasterRecoveryDrProtectionGroupsDataSource())
}
