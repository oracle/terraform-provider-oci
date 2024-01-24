// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_dataflow_application", DataflowApplicationDataSource())
	tfresource.RegisterDatasource("oci_dataflow_applications", DataflowApplicationsDataSource())
	tfresource.RegisterDatasource("oci_dataflow_invoke_run", DataflowInvokeRunDataSource())
	tfresource.RegisterDatasource("oci_dataflow_invoke_runs", DataflowInvokeRunsDataSource())
	tfresource.RegisterDatasource("oci_dataflow_pool", DataflowPoolDataSource())
	tfresource.RegisterDatasource("oci_dataflow_pools", DataflowPoolsDataSource())
	tfresource.RegisterDatasource("oci_dataflow_private_endpoint", DataflowPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_dataflow_private_endpoints", DataflowPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_dataflow_run_log", DataflowRunLogDataSource())
	tfresource.RegisterDatasource("oci_dataflow_run_logs", DataflowRunLogsDataSource())
	tfresource.RegisterDatasource("oci_dataflow_run_statement", DataflowRunStatementDataSource())
	tfresource.RegisterDatasource("oci_dataflow_run_statements", DataflowRunStatementsDataSource())
	tfresource.RegisterDatasource("oci_dataflow_sql_endpoint", DataflowSqlEndpointDataSource())
	tfresource.RegisterDatasource("oci_dataflow_sql_endpoints", DataflowSqlEndpointsDataSource())
}
