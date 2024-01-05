// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_oda_oda_instance", OdaOdaInstanceDataSource())
	tfresource.RegisterDatasource("oci_oda_oda_instances", OdaOdaInstancesDataSource())
	tfresource.RegisterDatasource("oci_oda_oda_private_endpoint", OdaOdaPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_oda_oda_private_endpoint_attachment", OdaOdaPrivateEndpointAttachmentDataSource())
	tfresource.RegisterDatasource("oci_oda_oda_private_endpoint_attachments", OdaOdaPrivateEndpointAttachmentsDataSource())
	tfresource.RegisterDatasource("oci_oda_oda_private_endpoint_scan_proxies", OdaOdaPrivateEndpointScanProxiesDataSource())
	tfresource.RegisterDatasource("oci_oda_oda_private_endpoint_scan_proxy", OdaOdaPrivateEndpointScanProxyDataSource())
	tfresource.RegisterDatasource("oci_oda_oda_private_endpoints", OdaOdaPrivateEndpointsDataSource())
}
