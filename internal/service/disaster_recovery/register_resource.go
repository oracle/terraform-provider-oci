// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_disaster_recovery_dr_plan", DisasterRecoveryDrPlanResource())
	tfresource.RegisterResource("oci_disaster_recovery_dr_plan_execution", DisasterRecoveryDrPlanExecutionResource())
	tfresource.RegisterResource("oci_disaster_recovery_dr_protection_group", DisasterRecoveryDrProtectionGroupResource())
}
