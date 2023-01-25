// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_oda_oda_instance", OdaOdaInstanceDataSource())
	tfresource.RegisterDatasource("oci_oda_oda_instances", OdaOdaInstancesDataSource())
}
