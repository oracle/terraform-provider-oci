// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_data_connectivity_registry", DataConnectivityRegistryResource())
	tfresource.RegisterResource("oci_data_connectivity_registry_connection", DataConnectivityRegistryConnectionResource())
	tfresource.RegisterResource("oci_data_connectivity_registry_data_asset", DataConnectivityRegistryDataAssetResource())
	tfresource.RegisterResource("oci_data_connectivity_registry_folder", DataConnectivityRegistryFolderResource())
}
