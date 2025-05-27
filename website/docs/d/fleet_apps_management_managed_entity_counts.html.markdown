---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_managed_entity_counts"
sidebar_current: "docs-oci-datasource-fleet_apps_management-managed_entity_counts"
description: |-
  Provides the list of Managed Entity Counts in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_managed_entity_counts
This data source provides the list of Managed Entity Counts in Oracle Cloud Infrastructure Fleet Apps Management service.

Retrieve  aggregated summary information of Managed entities within a Compartment.


## Example Usage

```hcl
data "oci_fleet_apps_management_managed_entity_counts" "test_managed_entity_counts" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.managed_entity_count_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `compartment_id_in_subtree` - (Optional) If set to true, resources will be returned for not only the provided compartment, but all compartments which descend from it. Which resources are returned and their field contents depends on the value of accessLevel. 


## Attributes Reference

The following attributes are exported:

* `managed_entity_aggregation_collection` - The list of managed_entity_aggregation_collection.

### ManagedEntityCount Reference

The following attributes are exported:

* `items` - List of ManagedEntityAggregation objects.
	* `dimensions` - Aggregated summary information for ComplianceRecord
		* `entity` - Level at which the compliance is calculated.
	* `managed_entity_count_count` - Count of managed entities in a compartment.

