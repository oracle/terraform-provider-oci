// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_datacatalog_catalog", DatacatalogCatalogResource())
	tfresource.RegisterResource("oci_datacatalog_catalog_private_endpoint", DatacatalogCatalogPrivateEndpointResource())
	tfresource.RegisterResource("oci_datacatalog_connection", DatacatalogConnectionResource())
	tfresource.RegisterResource("oci_datacatalog_data_asset", DatacatalogDataAssetResource())
	tfresource.RegisterResource("oci_datacatalog_metastore", DatacatalogMetastoreResource())
}
