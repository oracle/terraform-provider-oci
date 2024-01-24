// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_streaming_connect_harness", StreamingConnectHarnessResource())
	tfresource.RegisterResource("oci_streaming_stream", StreamingStreamResource())
	tfresource.RegisterResource("oci_streaming_stream_pool", StreamingStreamPoolResource())
}
