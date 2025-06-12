package capacity_management

import (
	"fmt"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportCapacityManagementInternalOccmDemandSignalDeliveryHints.GetIdFn = getCapacityManagementInternalOccmDemandSignalDeliveryId
	exportCapacityManagementInternalOccmDemandSignalHints.GetIdFn = getCapacityManagementInternalOccmDemandSignalId
	tf_export.RegisterCompartmentGraphs("capacity_management", capacityManagementResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getCapacityManagementInternalOccmDemandSignalDeliveryId(resource *tf_export.OCIResource) (string, error) {

	occmDemandSignalDeliveryId, ok := resource.SourceAttributes["occm_demand_signal_delivery_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find occmDemandSignalDeliveryId for CapacityManagement InternalOccmDemandSignalDelivery")
	}
	return GetInternalOccmDemandSignalDeliveryCompositeId(occmDemandSignalDeliveryId), nil
}

func getCapacityManagementInternalOccmDemandSignalId(resource *tf_export.OCIResource) (string, error) {

	occmDemandSignalId, ok := resource.SourceAttributes["occm_demand_signal_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find occmDemandSignalId for CapacityManagement InternalOccmDemandSignal")
	}
	return GetInternalOccmDemandSignalCompositeId(occmDemandSignalId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportCapacityManagementOccCustomerGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_occ_customer_group",
	DatasourceClass:        "oci_capacity_management_occ_customer_groups",
	DatasourceItemsAttr:    "occ_customer_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occ_customer_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccCustomerGroupLifecycleStateActive),
	},
}

var exportCapacityManagementOccAvailabilityCatalogHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_occ_availability_catalog",
	DatasourceClass:        "oci_capacity_management_occ_availability_catalogs",
	DatasourceItemsAttr:    "occ_availability_catalog_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occ_availability_catalog",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccAvailabilityCatalogLifecycleStateActive),
	},
}

var exportCapacityManagementOccCapacityRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_occ_capacity_request",
	DatasourceClass:        "oci_capacity_management_occ_capacity_requests",
	DatasourceItemsAttr:    "occ_capacity_request_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occ_capacity_request",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccCapacityRequestLifecycleStateActive),
	},
}

var exportCapacityManagementInternalOccCapacityRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_internal_occ_capacity_request",
	DatasourceClass:        "oci_capacity_management_internal_occ_capacity_requests",
	DatasourceItemsAttr:    "occ_capacity_request_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "internal_occ_capacity_request",
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccCapacityRequestLifecycleStateActive),
	},
}

var exportCapacityManagementInternalOccmDemandSignalDeliveryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_internal_occm_demand_signal_delivery",
	DatasourceClass:        "oci_capacity_management_internal_occm_demand_signal_deliveries",
	DatasourceItemsAttr:    "internal_occm_demand_signal_delivery_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "internal_occm_demand_signal_delivery",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.InternalOccmDemandSignalDeliveryLifecycleStateActive),
	},
}

var exportCapacityManagementInternalOccmDemandSignalHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_internal_occm_demand_signal",
	DatasourceClass:        "oci_capacity_management_internal_occm_demand_signals",
	DatasourceItemsAttr:    "internal_occm_demand_signal_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "internal_occm_demand_signal",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.InternalOccmDemandSignalCatalogResourceLifecycleStateActive),
	},
}

var exportCapacityManagementOccmDemandSignalItemHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_occm_demand_signal_item",
	DatasourceClass:        "oci_capacity_management_occm_demand_signal_items",
	DatasourceItemsAttr:    "occm_demand_signal_item_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occm_demand_signal_item",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateActive),
	},
}

var exportCapacityManagementOccmDemandSignalHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_capacity_management_occm_demand_signal",
	DatasourceClass:        "oci_capacity_management_occm_demand_signals",
	DatasourceItemsAttr:    "occm_demand_signal_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occm_demand_signal",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_capacity_management.OccmDemandSignalLifecycleStateActive),
	},
}

var capacityManagementResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCapacityManagementOccCustomerGroupHints},
		{TerraformResourceHints: exportCapacityManagementInternalOccCapacityRequestHints},
		{TerraformResourceHints: exportCapacityManagementOccAvailabilityCatalogHints},
		{TerraformResourceHints: exportCapacityManagementOccCapacityRequestHints},
		{TerraformResourceHints: exportCapacityManagementOccmDemandSignalItemHints},
		{TerraformResourceHints: exportCapacityManagementOccmDemandSignalHints},
	},
	"oci_capacity_management_occ_customer_group": {
		{
			TerraformResourceHints: exportCapacityManagementInternalOccmDemandSignalHints,
			DatasourceQueryParams: map[string]string{
				"occ_customer_group_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCapacityManagementInternalOccmDemandSignalDeliveryHints,
			DatasourceQueryParams: map[string]string{
				"occ_customer_group_id": "id",
			},
		},
	},
}
