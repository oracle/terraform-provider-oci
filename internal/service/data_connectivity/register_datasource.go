// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_data_connectivity_registries", DataConnectivityRegistriesDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry", DataConnectivityRegistryDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry_connection", DataConnectivityRegistryConnectionDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry_connections", DataConnectivityRegistryConnectionsDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry_data_asset", DataConnectivityRegistryDataAssetDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry_data_assets", DataConnectivityRegistryDataAssetsDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry_folder", DataConnectivityRegistryFolderDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry_folders", DataConnectivityRegistryFoldersDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry_type", DataConnectivityRegistryTypeDataSource())
	tfresource.RegisterDatasource("oci_data_connectivity_registry_types", DataConnectivityRegistryTypesDataSource())
}
