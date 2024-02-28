// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package sch

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_sch_connector_plugin", SchConnectorPluginDataSource())
	tfresource.RegisterDatasource("oci_sch_connector_plugins", SchConnectorPluginsDataSource())
	tfresource.RegisterDatasource("oci_sch_service_connector", SchServiceConnectorDataSource())
	tfresource.RegisterDatasource("oci_sch_service_connectors", SchServiceConnectorsDataSource())
}
