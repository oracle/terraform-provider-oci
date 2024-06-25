// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_scheduler

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_resource_scheduler_schedule", ResourceSchedulerScheduleDataSource())
	tfresource.RegisterDatasource("oci_resource_scheduler_schedules", ResourceSchedulerSchedulesDataSource())
}
