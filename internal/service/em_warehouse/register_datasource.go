// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package em_warehouse

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_em_warehouse_em_warehouse", EmWarehouseEmWarehouseDataSource())
	tfresource.RegisterDatasource("oci_em_warehouse_em_warehouse_etl_run", EmWarehouseEmWarehouseEtlRunDataSource())
	tfresource.RegisterDatasource("oci_em_warehouse_em_warehouse_etl_runs", EmWarehouseEmWarehouseEtlRunsDataSource())
	tfresource.RegisterDatasource("oci_em_warehouse_em_warehouse_resource_usage", EmWarehouseEmWarehouseResourceUsageDataSource())
	tfresource.RegisterDatasource("oci_em_warehouse_em_warehouses", EmWarehouseEmWarehousesDataSource())
}
