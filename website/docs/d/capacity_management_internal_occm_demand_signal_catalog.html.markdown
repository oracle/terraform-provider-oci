---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal_catalog"
sidebar_current: "docs-oci-datasource-capacity_management-internal_occm_demand_signal_catalog"
description: |-
  Provides details about a specific Internal Occm Demand Signal Catalog in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_occm_demand_signal_catalog
This data source provides details about a specific Internal Occm Demand Signal Catalog resource in Oracle Cloud Infrastructure Capacity Management service.

This API helps in getting the details about a specific occm demand signal catalog.


## Example Usage

```hcl
data "oci_capacity_management_internal_occm_demand_signal_catalog" "test_internal_occm_demand_signal_catalog" {
	#Required
	occm_demand_signal_catalog_id = oci_datacatalog_catalog.test_catalog.id
}
```

## Argument Reference

The following arguments are supported:

* `occm_demand_signal_catalog_id` - (Required) The OCID of the demand signal catalog. 


## Attributes Reference

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

