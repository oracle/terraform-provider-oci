---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_availability_catalog"
sidebar_current: "docs-oci-datasource-capacity_management-occ_availability_catalog"
description: |-
  Provides details about a specific Occ Availability Catalog in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occ_availability_catalog
This data source provides details about a specific Occ Availability Catalog resource in Oracle Cloud Infrastructure Capacity Management service.

Get details about availability catalog.

## Example Usage

```hcl
data "oci_capacity_management_occ_availability_catalog" "test_occ_availability_catalog" {
	#Required
	occ_availability_catalog_id = oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id
}
```

## Argument Reference

The following arguments are supported:

* `occ_availability_catalog_id` - (Required) The OCID of the availability catalog.


## Attributes Reference

The following attributes are exported:

* `catalog_state` - The different states associated with the availability catalog.
* `compartment_id` - The OCID of the tenancy where the availability catalog resides.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Text information about the availability catalog.
* `details` - Details about capacity available for different resources in catalog.
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
* `display_name` - A user-friendly name for the availability catalog.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the availability catalog.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
* `metadata_details` - Used for representing the metadata of the catalog. This denotes the version and format of the CSV file for parsing.
	* `format_version` - The version for the format of the catalog file being uploaded.
* `namespace` - The name of the Oracle Cloud Infrastructure service in consideration. For example, Compute, Exadata, and so on. 
* `occ_customer_group_id` - The customer group OCID to which the availability catalog belongs.
* `state` - The current lifecycle state of the resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the availability catalog was created.
* `time_updated` - The time when the availability catalog was last updated.

