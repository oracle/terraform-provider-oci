// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package batch

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_batch_batch_context", BatchBatchContextDataSource())
	tfresource.RegisterDatasource("oci_batch_batch_context_shapes", BatchBatchContextShapesDataSource())
	tfresource.RegisterDatasource("oci_batch_batch_contexts", BatchBatchContextsDataSource())
	tfresource.RegisterDatasource("oci_batch_batch_job_pool", BatchBatchJobPoolDataSource())
	tfresource.RegisterDatasource("oci_batch_batch_job_pools", BatchBatchJobPoolsDataSource())
	tfresource.RegisterDatasource("oci_batch_batch_task_environment", BatchBatchTaskEnvironmentDataSource())
	tfresource.RegisterDatasource("oci_batch_batch_task_environments", BatchBatchTaskEnvironmentsDataSource())
	tfresource.RegisterDatasource("oci_batch_batch_task_profile", BatchBatchTaskProfileDataSource())
	tfresource.RegisterDatasource("oci_batch_batch_task_profiles", BatchBatchTaskProfilesDataSource())
}
