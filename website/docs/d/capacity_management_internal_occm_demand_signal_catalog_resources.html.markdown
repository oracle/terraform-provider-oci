---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal_catalog_resources"
sidebar_current: "docs-oci-datasource-capacity_management-internal_occm_demand_signal_catalog_resources"
description: |-
  Provides the list of Internal Occm Demand Signal Catalog Resources in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_occm_demand_signal_catalog_resources
This data source provides the list of Internal Occm Demand Signal Catalog Resources in Oracle Cloud Infrastructure Capacity Management service.

This API will list all the  resources across all demand signal catalogs for a given namespace and customer group.


## Example Usage

```hcl
data "oci_capacity_management_internal_occm_demand_signal_catalog_resources" "test_internal_occm_demand_signal_catalog_resources" {
	#Required
	compartment_id = var.compartment_id
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id
	occm_demand_signal_catalog_id = oci_datacatalog_catalog.test_catalog.id

	#Optional
	demand_signal_namespace = var.internal_occm_demand_signal_catalog_resource_demand_signal_namespace
	name = var.internal_occm_demand_signal_catalog_resource_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
* `demand_signal_namespace` - (Optional) A query parameter to filter the list of demand signal catalog resources based on the namespace. 
* `name` - (Optional) A query parameter to filter the list of demand signal catalog resource based on the resource name. 
* `occ_customer_group_id` - (Required) The customer group ocid by which we would filter the list.
* `occm_demand_signal_catalog_id` - (Required) The ocid of demand signal catalog id.


## Attributes Reference

The following attributes are exported:

* `internal_occm_demand_signal_catalog_resource_collection` - The list of internal_occm_demand_signal_catalog_resource_collection.

### InternalOccmDemandSignalCatalogResource Reference

The following attributes are exported:

* `items` - An array of items containing detailed information about different resources. 
	* `availability_domain` - The name of the availability domain for which you want to request the Oracle Cloud Infrastructure resource. This is an optional parameter. 
	* `compartment_id` - The OCID of the tenancy from which the request to create the demand signal catalog was made. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The OCID of the demand signal catalog resource. 
	* `name` - The name of the Oracle Cloud Infrastructure resource that you want to request. 
	* `namespace` - The name of the Oracle Cloud Infrastructure service in consideration for demand signal submission. For example: COMPUTE, NETWORK, GPU etc. 
	* `occ_customer_group_id` - The OCID of the customerGroup. 
	* `occm_demand_signal_catalog_id` - This OCID of the demand signal catalog 
	* `region` - The name of region for which you want to request the Oracle Cloud Infrastructure resource. This is an optional parameter. 
	* `resource_properties` - A list containing detailed information about a resource's properties.
		* `items` - An array of items containing detailed information about a resource's properties. 
			* `is_editable` - This will indicate if demand signal resource's property is editable. 
			* `property_max_value` - The maximum value of demand signal resource's property. This is an optional parameter. 
			* `property_min_value` - The minimum value of demand signal resource's property. This is an optional parameter. 
			* `property_name` - The name of demand signal resource's property. 
			* `property_options` - Predefined options for demand signal resource's property. This is an optional parameter. 
				* `option_key` - key of a property option like memoryValue, ocpuValue. 
				* `option_value` - value of a property option like 64, 2 fastconnect etc. 
			* `property_unit` - Unit for demand signal resource's property. 
			* `property_value` - Default value of demand signal resource's property. 
	* `resource_property_constraints` - A list containing detailed information about a resource's property constraints.
		* `items` - An array of items containing detailed information about a resource's property dependecies. 
			* `constraint_name` - The name of demand signal resource's property constraint. 
			* `constraint_value` - The value of demand signal resource's property constraint. 
	* `state` - The current lifecycle state of the demand signal catalog resource. 
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `target_compartment_id` - The OCID of the customer tenancy for which this resource will be available for the customer to order against. 
	* `time_created` - The time when the demand signal catalog resource was created. 
	* `time_updated` - The time when the demand signal catalog resource was last updated. 

