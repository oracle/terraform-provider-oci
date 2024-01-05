// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_datacatalog_catalog", DatacatalogCatalogDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_catalog_private_endpoint", DatacatalogCatalogPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_catalog_private_endpoints", DatacatalogCatalogPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_catalog_type", DatacatalogCatalogTypeDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_catalog_types", DatacatalogCatalogTypesDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_catalogs", DatacatalogCatalogsDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_connection", DatacatalogConnectionDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_connections", DatacatalogConnectionsDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_data_asset", DatacatalogDataAssetDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_data_assets", DatacatalogDataAssetsDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_metastore", DatacatalogMetastoreDataSource())
	tfresource.RegisterDatasource("oci_datacatalog_metastores", DatacatalogMetastoresDataSource())
}
