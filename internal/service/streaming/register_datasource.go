// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_streaming_connect_harness", StreamingConnectHarnessDataSource())
	tfresource.RegisterDatasource("oci_streaming_connect_harnesses", StreamingConnectHarnessesDataSource())
	tfresource.RegisterDatasource("oci_streaming_stream", StreamingStreamDataSource())
	tfresource.RegisterDatasource("oci_streaming_stream_pool", StreamingStreamPoolDataSource())
	tfresource.RegisterDatasource("oci_streaming_stream_pools", StreamingStreamPoolsDataSource())
	tfresource.RegisterDatasource("oci_streaming_streams", StreamingStreamsDataSource())
}
