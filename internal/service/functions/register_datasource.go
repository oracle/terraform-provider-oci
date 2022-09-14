// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_functions_application", FunctionsApplicationDataSource())
	tfresource.RegisterDatasource("oci_functions_applications", FunctionsApplicationsDataSource())
	tfresource.RegisterDatasource("oci_functions_function", FunctionsFunctionDataSource())
	tfresource.RegisterDatasource("oci_functions_functions", FunctionsFunctionsDataSource())
}
