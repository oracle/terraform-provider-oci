// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package api_platform

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_api_platform_api_platform_instance", ApiPlatformApiPlatformInstanceDataSource())
	tfresource.RegisterDatasource("oci_api_platform_api_platform_instances", ApiPlatformApiPlatformInstancesDataSource())
}
