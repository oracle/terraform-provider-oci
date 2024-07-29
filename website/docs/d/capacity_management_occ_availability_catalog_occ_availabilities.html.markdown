---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_availability_catalog_occ_availabilities"
sidebar_current: "docs-oci-datasource-capacity_management-occ_availability_catalog_occ_availabilities"
description: |-
  Provides the list of Occ Availability Catalog Occ Availabilities in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occ_availability_catalog_occ_availabilities
This data source provides the list of Occ Availability Catalog Occ Availabilities in Oracle Cloud Infrastructure Capacity Management service.

Lists availabilities for a particular availability catalog.

## Example Usage

```hcl
data "oci_capacity_management_occ_availability_catalog_occ_availabilities" "test_occ_availability_catalog_occ_availabilities" {
	#Required
	occ_availability_catalog_id = oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id

	#Optional
	date_expected_capacity_handover = var.occ_availability_catalog_occ_availability_date_expected_capacity_handover
	resource_name = oci_usage_proxy_resource.test_resource.name
	resource_type = var.occ_availability_catalog_occ_availability_resource_type
	workload_type = var.occ_availability_catalog_occ_availability_workload_type
}
```

## Argument Reference

The following arguments are supported:

* `date_expected_capacity_handover` - (Optional) The capacity handover date of the capacity constraint to filter the list of capacity constraints.
* `occ_availability_catalog_id` - (Required) The OCID of the availability catalog.
* `resource_name` - (Optional) The name of the resource to filter the list of capacity constraints.
* `resource_type` - (Optional) Resource type using which the capacity constraints of an availability catalog can be filtered.
* `workload_type` - (Optional) Workload type using the resources in an availability catalog can be filtered.


## Attributes Reference

The following attributes are exported:

* `occ_availability_collection` - The list of occ_availability_collection.

### OccAvailabilityCatalogOccAvailability Reference

The following attributes are exported:

* `items` - An array of capacity constraints.
	* `available_quantity` - The quantity of resource currently available that the customer can request.
	* `catalog_id` - The OCID of the availability catalog.
	* `date_expected_capacity_handover` - The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
	* `date_final_customer_order` - The date by which the customer must place the order to have their capacity requirements met by the customer handover date.
	* `demanded_quantity` - The quantity of resource currently demanded by the customer.
	* `namespace` - The name of the Oracle Cloud Infrastructure service in consideration. For example, Compute, Exadata, and so on. 
	* `resource_name` - The name of the resource that the customer can request.
	* `resource_type` - The different types of resources against which customers can place capacity requests.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `total_available_quantity` - The total quantity of resource that the customer can request.
	* `unit` - The unit in which the resource available is measured.
	* `workload_type` - The type of workload (Generic/ROW).

