// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_functions_application", FunctionsApplicationDataSource())
	tfresource.RegisterDatasource("oci_functions_applications", FunctionsApplicationsDataSource())
	tfresource.RegisterDatasource("oci_functions_function", FunctionsFunctionDataSource())
	tfresource.RegisterDatasource("oci_functions_functions", FunctionsFunctionsDataSource())
	tfresource.RegisterDatasource("oci_functions_pbf_listing", FunctionsPbfListingDataSource())
	tfresource.RegisterDatasource("oci_functions_pbf_listing_triggers", FunctionsPbfListingTriggersDataSource())
	tfresource.RegisterDatasource("oci_functions_pbf_listing_version", FunctionsPbfListingVersionDataSource())
	tfresource.RegisterDatasource("oci_functions_pbf_listing_versions", FunctionsPbfListingVersionsDataSource())
	tfresource.RegisterDatasource("oci_functions_pbf_listings", FunctionsPbfListingsDataSource())
}
