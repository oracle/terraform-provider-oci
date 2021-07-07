---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_enterprise_manager_bridge"
sidebar_current: "docs-oci-resource-opsi-enterprise_manager_bridge"
description: |-
  Provides the Enterprise Manager Bridge resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_enterprise_manager_bridge
This resource provides the Enterprise Manager Bridge resource in Oracle Cloud Infrastructure Opsi service.

Create a Enterprise Manager bridge in Operations Insights.


## Example Usage

```hcl
resource "oci_opsi_enterprise_manager_bridge" "test_enterprise_manager_bridge" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.enterprise_manager_bridge_display_name
	object_storage_bucket_name = oci_objectstorage_bucket.test_bucket.name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.enterprise_manager_bridge_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment identifier of the Enterprise Manager bridge
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of Enterprise Manager Bridge
* `display_name` - (Required) (Updatable) User-friedly name of Enterprise Manager Bridge that does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `object_storage_bucket_name` - (Required) Object Storage Bucket Name


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the Enterprise Manager bridge
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of Enterprise Manager Bridge
* `display_name` - User-friedly name of Enterprise Manager Bridge that does not have to be unique.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Enterprise Manager bridge identifier
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `object_storage_bucket_name` - Object Storage Bucket Name
* `object_storage_bucket_status_details` - A message describing status of the object storage bucket of this resource. For example, it can be used to provide actionable information about the permission and content validity of the bucket.
* `object_storage_namespace_name` - Object Storage Namespace Name
* `state` - The current state of the Enterprise Manager bridge.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the Enterprise Manager bridge was first created. An RFC3339 formatted datetime string
* `time_updated` - The time the Enterprise Manager bridge was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Enterprise Manager Bridge
	* `update` - (Defaults to 20 minutes), when updating the Enterprise Manager Bridge
	* `delete` - (Defaults to 20 minutes), when destroying the Enterprise Manager Bridge


## Import

EnterpriseManagerBridges can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge "id"
```

