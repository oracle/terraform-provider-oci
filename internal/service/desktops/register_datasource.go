// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package desktops

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_desktops_desktop", DesktopsDesktopDataSource())
	tfresource.RegisterDatasource("oci_desktops_desktop_pool", DesktopsDesktopPoolDataSource())
	tfresource.RegisterDatasource("oci_desktops_desktop_pool_desktops", DesktopsDesktopPoolDesktopsDataSource())
	tfresource.RegisterDatasource("oci_desktops_desktop_pool_volumes", DesktopsDesktopPoolVolumesDataSource())
	tfresource.RegisterDatasource("oci_desktops_desktop_pools", DesktopsDesktopPoolsDataSource())
	tfresource.RegisterDatasource("oci_desktops_desktops", DesktopsDesktopsDataSource())
}
