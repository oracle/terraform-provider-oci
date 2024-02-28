---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_availability_catalog"
sidebar_current: "docs-oci-resource-capacity_management-occ_availability_catalog"
description: |-
  Provides the Occ Availability Catalog resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_occ_availability_catalog
This resource provides the Occ Availability Catalog resource in Oracle Cloud Infrastructure Capacity Management service.

Create availability catalog

## Example Usage

```hcl
resource "oci_capacity_management_occ_availability_catalog" "test_occ_availability_catalog" {
	#Required
	base64encoded_catalog_details = var.occ_availability_catalog_base64encoded_catalog_details
	compartment_id = var.compartment_id
	display_name = var.occ_availability_catalog_display_name
	namespace = var.occ_availability_catalog_namespace
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.occ_availability_catalog_description
	freeform_tags = {"bar-key"= "value"}
	metadata_details {
		#Required
		format_version = var.occ_availability_catalog_metadata_details_format_version
	}
}
```

## Argument Reference

The following arguments are supported:

* `base64encoded_catalog_details` - (Required) The base 64 encoded string corresponding to the catalog file contents.
* `compartment_id` - (Required) Since all resources are at tenancy level hence this will be the ocid of the tenancy where operation is to be performed.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Additional information about the availability catalog.
* `display_name` - (Required) (Updatable) The display name of the availability catalog.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `metadata_details` - (Optional) Used for representing the metadata of the catalog. This denotes the version and format of the CSV file for parsing.
	* `format_version` - (Required) The version for the format of the catalog file being uploaded.
* `namespace` - (Required) The name of the Oracle Cloud Infrastructure service in consideration. For example, Compute, Exadata, and so on.
* `occ_customer_group_id` - (Required) The OCID of the customer group.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `catalog_state` - The different states associated with the availability catalog.
* `compartment_id` - The OCID of the tenancy where the availability catalog resides.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Text information about the availability catalog.
* `details` - Details about capacity available for  different resources in catalog.
	* `available_quantity` - The quantity of available resource that the customer can request.
	* `catalog_id` - The OCID of the availability catalog.
	* `date_expected_capacity_handover` - The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
	* `date_final_customer_order` - The date by which the customer must place the order to have their capacity requirements met by the customer handover date.
	* `namespace` - The name of the Oracle Cloud Infrastructure service in consideration. For example, Compute, Exadata, and so on. 
	* `resource_name` - The name of the resource that the customer can request.
	* `resource_type` - The different types of resources against which customers can place capacity requests.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Occ Availability Catalog
	* `update` - (Defaults to 20 minutes), when updating the Occ Availability Catalog
	* `delete` - (Defaults to 20 minutes), when destroying the Occ Availability Catalog


## Import

OccAvailabilityCatalogs can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog "id"
```

