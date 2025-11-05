// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dif

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_dif_stack", DifStackDataSource())
	tfresource.RegisterDatasource("oci_dif_stacks", DifStacksDataSource())
}
