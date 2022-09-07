---
subcategory: "Data Connectivity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_connectivity_registry"
sidebar_current: "docs-oci-resource-data_connectivity-registry"
description: |-
  Provides the Registry resource in Oracle Cloud Infrastructure Data Connectivity service
---

# oci_data_connectivity_registry
This resource provides the Registry resource in Oracle Cloud Infrastructure Data Connectivity service.

Creates a new Data Connectivity Management registry ready to perform data connectivity management.


## Example Usage

```hcl
resource "oci_data_connectivity_registry" "test_registry" {
	#Required
	display_name = var.registry_display_name

	#Optional
	compartment_id = var.compartment_id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.registry_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Data Connectivity Management Registry description
* `display_name` - (Required) (Updatable) The Data Connectivity Management Registry display name; registries can be renamed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists only for cross-compatibility. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Registry description
* `display_name` - Data Connectivity Management registry display name; registries can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
* `id` - A unique identifier that is immutable on creation.
* `state` - Lifecycle states for registries in the Data Connectivity Management Service CREATING - The resource is being created and may not be usable until the entire metadata is defined. UPDATING - The resource is being updated and may not be usable until all changes are commited. DELETING - The resource is being deleted and might require deep cleanup of children. ACTIVE   - The resource is valid and available for access. INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for administrative reasons. DELETED  - The resource has been deleted and isn't available. FAILED   - The resource is in a failed state due to validation or other errors. 
* `state_message` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `time_created` - Time when the Data Connectivity Management registry was created. An RFC3339 formatted datetime string.
* `time_updated` - Time when the Data Connectivity Management registry was updated. An RFC3339 formatted datetime string.
* `updated_by` - Name of the user who updated the DCMS registry.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Registry
	* `update` - (Defaults to 20 minutes), when updating the Registry
	* `delete` - (Defaults to 20 minutes), when destroying the Registry


## Import

Registries can be imported using the `id`, e.g.

```
$ terraform import oci_data_connectivity_registry.test_registry "id"
```

