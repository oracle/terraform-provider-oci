// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_capacity_management_internal_occ_availability_catalogs", CapacityManagementInternalOccAvailabilityCatalogsDataSource())
	tfresource.RegisterDatasource("oci_capacity_management_occ_availability_catalog", CapacityManagementOccAvailabilityCatalogDataSource())
	tfresource.RegisterDatasource("oci_capacity_management_occ_availability_catalog_content", CapacityManagementOccAvailabilityCatalogContentDataSource())
	tfresource.RegisterDatasource("oci_capacity_management_occ_availability_catalog_occ_availabilities", CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSource())
	tfresource.RegisterDatasource("oci_capacity_management_occ_availability_catalogs", CapacityManagementOccAvailabilityCatalogsDataSource())
	tfresource.RegisterDatasource("oci_capacity_management_occ_capacity_request", CapacityManagementOccCapacityRequestDataSource())
	tfresource.RegisterDatasource("oci_capacity_management_occ_capacity_requests", CapacityManagementOccCapacityRequestsDataSource())
	tfresource.RegisterDatasource("oci_capacity_management_occ_customer_group", CapacityManagementOccCustomerGroupDataSource())
	tfresource.RegisterDatasource("oci_capacity_management_occ_customer_groups", CapacityManagementOccCustomerGroupsDataSource())
}
