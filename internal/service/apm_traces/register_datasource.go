// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_apm_traces_query_quick_picks", ApmTracesQueryQuickPicksDataSource())
	tfresource.RegisterDatasource("oci_apm_traces_trace", ApmTracesTraceDataSource())
	tfresource.RegisterDatasource("oci_apm_traces_trace_aggregated_snapshot_data", ApmTracesTraceAggregatedSnapshotDataDataSource())
	tfresource.RegisterDatasource("oci_apm_traces_trace_snapshot_data", ApmTracesTraceSnapshotDataDataSource())
}
