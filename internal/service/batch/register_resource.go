// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package batch

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_batch_batch_context", BatchBatchContextResource())
	tfresource.RegisterResource("oci_batch_batch_job_pool", BatchBatchJobPoolResource())
	tfresource.RegisterResource("oci_batch_batch_task_environment", BatchBatchTaskEnvironmentResource())
	tfresource.RegisterResource("oci_batch_batch_task_profile", BatchBatchTaskProfileResource())
}
