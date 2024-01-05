// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package nosql

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_nosql_index", NosqlIndexDataSource())
	tfresource.RegisterDatasource("oci_nosql_indexes", NosqlIndexesDataSource())
	tfresource.RegisterDatasource("oci_nosql_table", NosqlTableDataSource())
	tfresource.RegisterDatasource("oci_nosql_tables", NosqlTablesDataSource())
}
