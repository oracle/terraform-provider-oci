// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package pss

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_pss_instance", PssInstanceDataSource())
	tfresource.RegisterDatasource("oci_pss_instances", PssInstancesDataSource())
}
