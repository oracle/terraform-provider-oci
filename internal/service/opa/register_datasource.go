// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opa

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_opa_opa_instance", OpaOpaInstanceDataSource())
	tfresource.RegisterDatasource("oci_opa_opa_instances", OpaOpaInstancesDataSource())
}
