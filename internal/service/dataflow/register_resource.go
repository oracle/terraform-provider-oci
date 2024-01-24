// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_dataflow_application", DataflowApplicationResource())
	tfresource.RegisterResource("oci_dataflow_invoke_run", DataflowInvokeRunResource())
	tfresource.RegisterResource("oci_dataflow_pool", DataflowPoolResource())
	tfresource.RegisterResource("oci_dataflow_private_endpoint", DataflowPrivateEndpointResource())
	tfresource.RegisterResource("oci_dataflow_run_statement", DataflowRunStatementResource())
	tfresource.RegisterResource("oci_dataflow_sql_endpoint", DataflowSqlEndpointResource())
}
