---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal_catalogs"
sidebar_current: "docs-oci-datasource-capacity_management-internal_occm_demand_signal_catalogs"
description: |-
  Provides the list of Internal Occm Demand Signal Catalogs in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_occm_demand_signal_catalogs
This data source provides the list of Internal Occm Demand Signal Catalogs in Oracle Cloud Infrastructure Capacity Management service.

This API will list demand signal catalogs for a given customer group.


## Example Usage

```hcl
data "oci_capacity_management_internal_occm_demand_signal_catalogs" "test_internal_occm_demand_signal_catalogs" {
	#Required
	compartment_id = var.compartment_id
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id

	#Optional
	display_name = var.internal_occm_demand_signal_catalog_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
* `display_name` - (Optional) A filter to return only the resources that match the entire display name. The match is not case sensitive.
* `occ_customer_group_id` - (Required) The customer group ocid by which we would filter the list.


## Attributes Reference

The following attributes are exported:

* `occm_demand_signal_catalog_collection` - The list of occm_demand_signal_catalog_collection.

### InternalOccmDemandSignalCatalog Reference

The following attributes are exported:

* `compartment_id` - compartment id from where demand signal catalog is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - description of demand signal catalog. 
* `display_name` - displayName of demand signal catalog. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The ocid of demand signal catalog. 
* `occ_customer_group_id` - The customer group OCID to which the availability catalog belongs.
* `state` - The current lifecycle state of the resource. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the demand signal catalog was created. 
* `time_updated` - The time when the demand signal catalog was last updated. 

