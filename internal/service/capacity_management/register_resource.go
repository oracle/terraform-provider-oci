// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_capacity_management_internal_occm_demand_signal", CapacityManagementInternalOccmDemandSignalResource())
	tfresource.RegisterResource("oci_capacity_management_internal_occm_demand_signal_delivery", CapacityManagementInternalOccmDemandSignalDeliveryResource())
	tfresource.RegisterResource("oci_capacity_management_occ_availability_catalog", CapacityManagementOccAvailabilityCatalogResource())
	tfresource.RegisterResource("oci_capacity_management_occ_capacity_request", CapacityManagementOccCapacityRequestResource())
	tfresource.RegisterResource("oci_capacity_management_occ_customer_group", CapacityManagementOccCustomerGroupResource())
	tfresource.RegisterResource("oci_capacity_management_occ_customer_group_occ_customer", CapacityManagementOccCustomerGroupOccCustomerResource())
	tfresource.RegisterResource("oci_capacity_management_occm_demand_signal", CapacityManagementOccmDemandSignalResource())
	tfresource.RegisterResource("oci_capacity_management_occm_demand_signal_item", CapacityManagementOccmDemandSignalItemResource())
}
