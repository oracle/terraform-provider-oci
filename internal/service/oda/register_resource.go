// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_oda_oda_instance", OdaOdaInstanceResource())
	tfresource.RegisterResource("oci_oda_oda_private_endpoint", OdaOdaPrivateEndpointResource())
	tfresource.RegisterResource("oci_oda_oda_private_endpoint_attachment", OdaOdaPrivateEndpointAttachmentResource())
	tfresource.RegisterResource("oci_oda_oda_private_endpoint_scan_proxy", OdaOdaPrivateEndpointScanProxyResource())
}
